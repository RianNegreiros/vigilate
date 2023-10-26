export default function WebServersPage() {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 p-4">
    <div className="grid-card">
      <div
        className="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 hover:shadow-lg hover:border-gray-300">
        <div className="flex justify-between mb-4 rounded-t sm:mb-5">
          <div className="text-lg text-gray-900 md:text-xl dark:text-white">
            <h3 className="font-semibold">Name</h3>
            <p className="font-bold">Address</p>
          </div>
          <div className="flex items-center space-x-2">
            <span
              className="inline-flex items-center bg-green-100 text-green-800 text-xs font-medium mr-2 px-2.5 py-0.5 rounded-full dark:bg-green-900 dark:text-green-300">
              <span className="w-2 h-2 mr-1 bg-green-500 rounded-full"></span>
              Available
            </span>
          </div>
        </div>
        <ul role="list" className="max-w-sm divide-y divide-gray-200 dark:divide-gray-700">
          <li className="py-3 sm:py-4">
            <div className="flex items-center space-x-3">
              <div className="flex-1 min-w-0">
                <p className="text-sm font-semibold text-gray-900 truncate dark:text-white">
                  Last Time Checked
                </p>
                <p className="text-sm text-gray-500 truncate dark:text-gray-400">
                  2021-04-20 12:00:00
                </p>
              </div>
            </div>
          </li>
          <li className="py-3 sm:py-4">
            <div className="flex items-center space-x-3">
              <div className="flex-1 min-w-0">
                <p className="text-sm font-semibold text-gray-900 truncate dark:text-white">
                  Next Time Check
                </p>
                <p className="text-sm text-gray-500 truncate dark:text-gray-400">
                  2021-04-20 12:00:00
                </p>
              </div>
            </div>
          </li>
        </ul>
        <div className="flex justify-between items-center">
          <div className="flex items-center space-x-3 sm:space-x-4">
            <button type="button"
              className="text-white inline-flex items-center bg-primary-700 hover:bg-primary-900 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover-bg-primary-700 dark:focus:ring-primary-800 hover:shadow-lg transition-background duration-300">
              <svg className="mr-1 -ml-1 w-5 h-5" fill="currentColor" xmlns="http://www.w3.org/2000/svg" height="1em"
                viewBox="0 0 512 512">
                <path
                  d="M32 32c17.7 0 32 14.3 32 32V400c0 8.8 7.2 16 16 16H480c17.7 0 32 14.3 32 32s-14.3 32-32 32H80c-44.2 0-80-35.8-80-80V64C0 46.3 14.3 32 32 32zm96 96c0-17.7 14.3-32 32-32l192 0c17.7 0 32 14.3 32 32s-14.3 32-32 32l-192 0c-17.7 0-32-14.3-32-32zm32 64H288c17.7 0 32 14.3 32 32s-14.3 32-32 32H160c-17.7 0-32-14.3-32-32s14.3-32 32-32zm0 96H416c17.7 0 32 14.3 32 32s-14.3 32-32 32H160c-17.7 0-32-14.3-32-32s14.3-32 32-32z" />
              </svg>
              Stat Monitoring
            </button>
          </div>
        </div>
      </div>
    </div>

    <button type="button"
      className="flex items-center justify-center px-6 py-2 text-sm font-medium text-center max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 hover:border-gray-300 text-gray-400 hover:bg-gray-700 hover:text-white focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white hover:shadow-lg transition-background duration-300"
      >
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor"
        className="h-8">
        <path strokeLinecap="round" strokeLinejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
      </svg>
    </button>
  </div>
  )
}