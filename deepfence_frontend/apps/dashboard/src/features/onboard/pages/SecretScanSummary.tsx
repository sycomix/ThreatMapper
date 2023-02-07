import cx from 'classnames';
import { Suspense } from 'react';
import { Await, Link, LoaderFunctionArgs, useLoaderData } from 'react-router-dom';
import { Card, CircleSpinner, Typography } from 'ui-components';

import { secretScanApiClient } from '@/api/api';
import { ModelSecretScanResult } from '@/api/generated/models/ModelSecretScanResult';
import LogoLinux from '@/assets/logo-linux.svg';
import { ConnectorHeader } from '@/features/onboard/components/ConnectorHeader';
import { ApiError, makeRequest } from '@/utils/api';
import { typedDefer, TypedDeferredData } from '@/utils/router';

const color: { [key: string]: string } = {
  alarm: 'bg-red-400 dark:bg-red-500',
  info: 'bg-blue-400 dark:bg-blue-500',
  ok: 'bg-green-400 dark:bg-green-500',
  skip: 'bg-gray-400 dark:bg-gray-500',
  critical: 'bg-red-400 dark:bg-red-500',
  high: 'bg-pink-400 dark:bg-pink-500',
  low: 'bg-yellow-300 dark:bg-yellow-500',
  medium: 'bg-blue-400 dark:bg-blue-500',
  unknown: 'bg-gray-400 dark:bg-gray-500',
};

type ScanType = {
  total: number;
  counts: SeverityType[] | null;
};
type SeverityType = {
  name: string;
  value: number;
};

type ScanData = {
  accountId: ModelSecretScanResult['kubernetes_cluster_name'];
  data: {
    total: number;
    counts: SeverityType[] | null;
  }[];
} | null;

export type LoaderDataType = {
  error?: string;
  message?: string;
  data?: ScanData[];
};

async function getScanSummary(scanIds: string): Promise<LoaderDataType> {
  const bulkRequest = scanIds.split(',').map((scanId) => {
    return makeRequest({
      apiFunction: secretScanApiClient().resultSecretScan,
      apiArgs: [
        {
          modelScanResultsReq: {
            scan_id: scanId,
            window: {
              offset: 0,
              size: 1000000,
            },
          },
        },
      ],
    });
  });
  const responses = await Promise.all(bulkRequest);
  const resultData = responses.map((response: ModelSecretScanResult | ApiError<void>) => {
    if (ApiError.isApiError(response)) {
      // TODO: handle any one request has an error on this bulk request
      return null;
    } else {
      const resp = response as ModelSecretScanResult;
      return {
        accountId: resp.kubernetes_cluster_name,
        data: [
          {
            total: Object.keys(resp.severity_counts ?? {}).reduce((acc, severity) => {
              acc = acc + (resp.severity_counts?.[severity] ?? 0);
              return acc;
            }, 0),
            counts: Object.keys(resp.severity_counts ?? {}).map((severity) => {
              return {
                name: severity,
                value: resp.severity_counts![severity],
              };
            }),
          },
        ],
      };
    }
  });

  return {
    data: resultData,
  };
}

const loader = ({
  params = {
    scanIds: '',
  },
}: LoaderFunctionArgs): TypedDeferredData<LoaderDataType> => {
  return typedDefer({
    data: getScanSummary(params.scanIds ?? ''),
  });
};

const AccountComponent = ({ accountId }: { accountId: string }) => {
  return (
    <div
      className={cx(
        'h-full flex flex-col items-center justify-center gap-y-3',
        'border-r dark:border-gray-700',
        'bg-gray-100 dark:bg-gray-700',
      )}
    >
      <img src={LogoLinux} alt="logo" height={40} width={40} />
      <data
        className={`${Typography.size.base} ${Typography.weight.normal} text-gray-700 dark:text-gray-300`}
      >
        {accountId}
      </data>
    </div>
  );
};

const TypeAndCountComponent = ({ total }: { total: number }) => {
  return (
    <div className="flex flex-col gap-y-1 items-center justify-center">
      <data className="text-sm text-gray-400 dark:text-gray-500">Total Results</data>
      <data className={'text-[2rem] text-gray-900 dark:text-gray-200 font-light'}>
        {total}
      </data>
    </div>
  );
};

const ChartComponent = ({ counts }: { counts: SeverityType[] }) => {
  const maxValue = Math.max(...counts.map((v) => v.value));

  return (
    <div>
      {counts.map(({ name, value }) => {
        return (
          <div className="flex items-center w-full" key={name}>
            <data
              className="pr-2 text-sm min-w-[100px] text-gray-500 text-end dark:text-gray-400"
              value={value}
            >
              {name.charAt(0).toUpperCase() + name.slice(1)}
            </data>
            <div
              className={cx(
                'w-[80%] overflow-hidden flex items-center',
                'cursor-pointer transition duration-100 hover:scale-y-125',
              )}
            >
              <div
                className={cx('rounded h-2 relative', color[name.toLowerCase()])}
                style={{
                  width: `${(100 / maxValue) * value}%`,
                }}
              ></div>
              <data className="ml-2 right-0 top-0 text-xs text-gray-500 dark:text-gray-400">
                {value}
              </data>
            </div>
          </div>
        );
      })}
    </div>
  );
};

const Scan = ({ scanData }: { scanData: ScanData }) => {
  if (!scanData) {
    return null;
  }
  const { accountId, data = [] } = scanData;

  return (
    <Card>
      <div className="grid grid-cols-[450px_1fr] items-center">
        <AccountComponent accountId={accountId} />
        <div className="flex flex-col">
          {data.map((severityData: ScanType | null, index: number) => {
            const { counts = [], total = 0 } = severityData ?? {};
            return (
              <div className="flex flex-col p-4" key={index}>
                <div className="grid grid-cols-[40%_60%]">
                  <TypeAndCountComponent total={total} />
                  {counts && <ChartComponent counts={counts} />}
                </div>
              </div>
            );
          })}
        </div>
      </div>
    </Card>
  );
};

const SecretScanSummary = () => {
  const loaderData = useLoaderData() as LoaderDataType;

  return (
    <div className="flex flex-col">
      <ConnectorHeader
        title={'Secret Scan Results Summary'}
        description={'Summary of secret scan result'}
      />

      <Link
        to="/dashboard"
        className={cx(
          `${Typography.size.sm} `,
          'underline underline-offset-2 ml-auto bg-transparent text-blue-600 dark:text-blue-500',
        )}
      >
        Go to Secret Dashboard to view details scan result
      </Link>

      <div className="flex flex-col gap-4 mt-4">
        <Suspense fallback={<CircleSpinner />}>
          <Await resolve={loaderData.data ?? []}>
            {(resolvedData) => {
              return resolvedData.data?.map((accountScanData: ScanData) => (
                <Scan key={accountScanData?.accountId} scanData={accountScanData} />
              ));
            }}
          </Await>
        </Suspense>
      </div>
    </div>
  );
};

export const module = {
  loader,
  element: <SecretScanSummary />,
};
