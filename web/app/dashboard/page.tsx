"use client"

import { useEffect, useState } from "react";
import NavBar from "../components/Navbar";
import Modal from "../components/Modal";
import { usePathname, useRouter } from "next/navigation";
import { createServer, getServers, startMonitoring } from "../util/api";
import { CreateServer, Server } from "../models";
import { format } from "date-fns";
import Link from "next/link";

export default function DashboardPage() {
  const [isModalOpen, setIsModalOpen] = useState(true);
  const [servers, setServers] = useState<Server[]>([]);
  let pathname = usePathname();

  const router = useRouter();

  const openModal = () => {
    setIsModalOpen(false);
  };

  const closeModal = () => {
    setIsModalOpen(true);
  };

  const getDomain = (address: string) => {
    const url = new URL(address);
    return url.hostname;
  };

  useEffect(() => {
    async function getServersData() {
      const servers = await getServers();
      if (servers !== null) {
        setServers(servers);
      }
    }
    getServersData();
  }, []);

  const handleCreateServer = async (formData: CreateServer) => {
    try {
      await createServer(formData);
      const updatedServers = await getServers();
      setServers(updatedServers);
    } catch (error) {
      console.log(error);
    }
  }

  const handleStartMonitoring = async (serverId: string) => {
    const monitorURL = `/monitor/${serverId}`;
    await startMonitoring(serverId);
    router.push(monitorURL);
  }

  return (
    <>
      <NavBar openModal={openModal} pathname={pathname} />

      <Modal hideModal={isModalOpen} closeModal={closeModal} createServer={handleCreateServer} />

      <div className="min-h-screen max-w-6xl mx-auto grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 p-4">
        {servers.map((server) => (
          <div key={server.id} className="grid-card w-64 h-80 overflow-hidden">
            <div className="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 hover:shadow-lg hover:border-gray-300">
              <div className="flex justify-between mb-4 rounded-t sm:mb-5">
                <div className="text-lg text-gray-900 md:text-xl dark:text-white">
                  <h3 className="font-semibold text-xl leading-6 text-gray-900 dark:text-white overflow-ellipsis">
                    {server.name}
                  </h3>
                  <div className="flex items-center space-x-2 py-1">
                    <span
                      className={`inline-flex items-center bg-green-100 ${server.is_active ? "text-green-800 dark:bg-green-900 dark:text-green-300" : "text-red-800 dark:bg-red-900 dark:text-red-300"} text-xs font-medium mr-2 px-2.5 py-0.5 rounded-full`}>
                      <span className={`w-2 h-2 mr-1 rounded-full ${server.is_active ? "bg-green-500" : "bg-red-500"}`}></span>
                      {server.is_active ? "Online" : "Offline"}
                    </span>
                  </div>
                  <p className="flex items-center text-base mt-1 leading-4 text-gray-600 dark:text-gray-400" title={server.address}>
                    <span className="overflow-ellipsis" style={{ maxWidth: "90%" }}>{getDomain(server.address)}</span>
                    <Link target="_blank" rel="noopener noreferrer" href={server.address}>
                      <svg className="ml-1 h-3 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                        <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M7.529 7.988a2.502 2.502 0 0 1 5 .191A2.441 2.441 0 0 1 10 10.582V12m-.01 3.008H10M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
                      </svg>
                    </Link>
                  </p>
                </div>
              </div>
              <ul role="list" className="max-w-sm divide-y divide-gray-200 dark:divide-gray-700">
                <li className="py-3 sm:py-3">
                  <div className="flex items-center space-x-3">
                    <div className="flex-1 min-w-0">
                      <p className="text-sm font-semibold text-gray-900 truncate dark:text-white">
                        Last Time Checked
                      </p>
                      <p className="text-sm text-gray-500 truncate dark:text-gray-400">
                        {format(new Date(server.last_check_time), 'dd/MM/yyyy HH:mm:ss')}
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
                        {format(new Date(server.next_check_time), 'dd/MM/yyyy HH:mm:ss')}
                      </p>
                    </div>
                  </div>
                </li>
              </ul>
              <div className="flex justify-between items-center">
                <div className="flex items-center space-x-3 sm:space-x-4">
                  <Link href={`/monitor/${server.id}`} passHref
                    onClick={() => handleStartMonitoring(server.id)}
                    className="text-white inline-flex items-center bg-primary-700 hover:bg-primary-900 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover-bg-primary-700 dark:focus:ring-primary-800 hover:shadow-lg transition-background duration-300">
                    <svg className="mr-1 -ml-1 w-5 h-5" fill="currentColor" xmlns="http://www.w3.org/2000/svg" height="1em"
                      viewBox="0 0 512 512">
                      <path
                        d="M32 32c17.7 0 32 14.3 32 32V400c0 8.8 7.2 16 16 16H480c17.7 0 32 14.3 32 32s-14.3 32-32 32H80c-44.2 0-80-35.8-80-80V64C0 46.3 14.3 32 32 32zm96 96c0-17.7 14.3-32 32-32l192 0c17.7 0 32 14.3 32 32s-14.3 32-32 32l-192 0c-17.7 0-32-14.3-32-32zm32 64H288c17.7 0 32 14.3 32 32s-14.3 32-32 32H160c-17.7 0-32-14.3-32-32s14.3-32 32-32zm0 96H416c17.7 0 32 14.3 32 32s-14.3 32-32 32H160c-17.7 0-32-14.3-32-32s14.3-32 32-32z" />
                    </svg>
                    Stat Monitoring
                  </Link>
                </div>
              </div>
            </div>
          </div>
        ))}
      </div>
    </>
  );
}
