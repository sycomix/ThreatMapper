package controls

import (
	"context"
	"encoding/json"
	"time"

	"github.com/deepfence/ThreatMapper/deepfence_utils/controls"
	"github.com/deepfence/ThreatMapper/deepfence_utils/directory"
	"github.com/deepfence/ThreatMapper/deepfence_utils/log"
	"github.com/deepfence/ThreatMapper/deepfence_utils/utils"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetKubernetesClusterActions(ctx context.Context, nodeID string, workNumToExtract int) ([]controls.Action, []error) {
	// Append more actions here
	var actions []controls.Action

	// Diagnostic logs not part of workNumToExtract
	diagnosticLogActions, diagnosticLogErr := ExtractAgentDiagnosticLogRequests(ctx, nodeID, controls.KubernetesCluster, maxWork)
	if diagnosticLogErr == nil {
		actions = append(actions, diagnosticLogActions...)
	}

	if workNumToExtract == 0 {
		return actions, []error{diagnosticLogErr}
	}

	upgradeActions, upgradeErr := ExtractPendingKubernetesClusterUpgrade(ctx, nodeID, workNumToExtract)
	workNumToExtract -= len(upgradeActions)
	if upgradeErr == nil {
		actions = append(actions, upgradeActions...)
	}

	scanActions, scanErr := ExtractStartingKubernetesClusterScans(ctx, nodeID, workNumToExtract)
	workNumToExtract -= len(scanActions)
	if scanErr == nil {
		actions = append(actions, scanActions...)
	}

	return actions, []error{scanErr, upgradeErr, diagnosticLogErr}
}

func ExtractStartingKubernetesClusterScans(ctx context.Context, nodeID string, maxWork int) ([]controls.Action, error) {
	var res []controls.Action
	if len(nodeID) == 0 {
		return res, ErrMissingNodeID
	}

	client, err := directory.Neo4jClient(ctx)
	if err != nil {
		return res, err
	}

	session := client.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return res, err
	}
	defer session.Close()

	tx, err := session.BeginTransaction(neo4j.WithTxTimeout(30 * time.Second))
	if err != nil {
		return res, err
	}
	defer tx.Close()

	r, err := tx.Run(`MATCH (s) -[:SCHEDULED]-> (n:KubernetesCluster{node_id:$id})
		WHERE s.status = '`+utils.ScanStatusStarting+`'
		AND s.retries < 3
		WITH s LIMIT $max_work
		SET s.status = '`+utils.ScanStatusInProgress+`'
		WITH s
		RETURN s.trigger_action`,
		map[string]interface{}{"id": nodeID, "max_work": maxWork})

	if err != nil {
		return res, err
	}

	records, err := r.Collect()

	if err != nil {
		return res, err
	}

	for _, record := range records {
		var action controls.Action
		if record.Values[0] == nil {
			log.Error().Msgf("Invalid neo4j trigger_action result, skipping")
			continue
		}
		err := json.Unmarshal([]byte(record.Values[0].(string)), &action)
		if err != nil {
			log.Error().Msgf("Unmarshal of action failed: %v", err)
			continue
		}
		res = append(res, action)
	}

	return res, tx.Commit()

}

func ExtractPendingKubernetesClusterUpgrade(ctx context.Context, nodeID string, maxWork int) ([]controls.Action, error) {
	var res []controls.Action
	if len(nodeID) == 0 {
		return res, ErrMissingNodeID
	}

	client, err := directory.Neo4jClient(ctx)
	if err != nil {
		return res, err
	}

	session := client.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	tx, err := session.BeginTransaction(neo4j.WithTxTimeout(30 * time.Second))
	if err != nil {
		return res, err
	}
	defer tx.Close()

	r, err := tx.Run(`MATCH (s:AgentVersion) -[r:SCHEDULED]-> (n:KubernetesCluster{node_id:$id})
		WHERE r.status = '`+utils.ScanStatusStarting+`'
		AND r.retries < 3
		WITH r LIMIT $max_work
		SET r.status = '`+utils.ScanStatusInProgress+`'
		WITH r
		RETURN r.trigger_action`,
		map[string]interface{}{"id": nodeID, "max_work": maxWork})

	if err != nil {
		return res, err
	}

	records, err := r.Collect()

	if err != nil {
		return res, err
	}

	for _, record := range records {
		var action controls.Action
		if record.Values[0] == nil {
			log.Error().Msgf("Invalid neo4j trigger_action result, skipping")
			continue
		}
		err := json.Unmarshal([]byte(record.Values[0].(string)), &action)
		if err != nil {
			log.Error().Msgf("Unmarshal of action failed: %v", err)
			continue
		}
		res = append(res, action)
	}

	return res, tx.Commit()

}
