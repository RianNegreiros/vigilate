export default function ServerLoadingSkeleton() {
  const LoadingElement = ({ width, height }: { width: number | string, height: number | string }) => (
    <div className={`h-${height} bg-gray-200 rounded-full dark:bg-gray-700 w-${width} mb-2.5`}></div>
  );

  return (
    <div className="grid-card w-64 h-80 overflow-hidden max-w-md p-4 space-y-4 border border-gray-200 divide-y divide-gray-200 rounded shadow animate-pulse dark:divide-gray-700 md:p-6 dark:border-gray-700">
      <div className="flex items-center justify-between">
        <div>
          <div className="h-2.5 bg-gray-300 rounded-full dark:bg-gray-600 w-24 mb-2.5"></div>
          <LoadingElement width="32" height="2" />
        </div>
        <div className="h-2.5 bg-gray-300 rounded-full dark:bg-gray-700 w-12"></div>
      </div>

      {Array.from({ length: 7 }).map((_, index) => (
        <LoadingElement key={index} width="48" height={index === 0 ? 2.5 : 2} />
      ))}

      <span className="sr-only">Loading...</span>
    </div>
  );
};
