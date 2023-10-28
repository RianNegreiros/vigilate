"use client"

import { usePathname } from "next/navigation";
import NavBar from "../../components/Navbar";
import { useEffect, useState } from "react";
import pusher from "../../util/pusher";
import { getServerById } from "../../util/api";
import { Server } from "../../models";

export default function MonitorPage({ params }: { params: { serverId: string } }) {
  const [serverData, setServerData] = useState<Server>({
    id: "",
    name: "",
    address: "",
    is_active: false,
    last_check_time: "",
    next_check_time: "",
    user_id: "",
  });
  let pathname = usePathname();
  const [isOnline, setIsOnline] = useState(false);
  const [events, setEvents] = useState<any[]>([]);

  useEffect(() => {
    async function getServerData() {
      const serverData = await getServerById(params.serverId);
      if (serverData !== null) {
        setServerData(serverData);
        setIsOnline(isOnline);
      }
    }
    getServerData();

    const channel = pusher.subscribe(`server-${params.serverId}`);

    channel.bind('status-changed', (data: any) => {
      if (!events.some(event => event.id === data.id)) {
        setIsOnline(data.isServerUp);

        setEvents(prevEvents => [...prevEvents, data]);
      }
    });

    return () => {
      pusher.unsubscribe(params.serverId);
    }
  }, [params.serverId]); // Make sure to include params.serverId in the dependency array

  return (
    <>
      <NavBar pathname={pathname} />
      <div className="relative flex justify-center items-center overflow-x-auto sm:rounded-lg">
        <table className="w-1/2 mt-5 text-sm text-left text-gray-500 dark:text-gray-400">
          <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th scope="col" className="px-6 py-3">
                Name
              </th>
              <th scope="col" className="px-6 py-3">
                Address
              </th>
              <th scope="col" className="px-6 py-3">
                Status
              </th>
            </tr>
          </thead>
          <tbody>
            {events.map((event, index) => (
              <tr key={event.id} className={`bg-white border-b ${event.isServerUp ? 'bg-green-50' : 'bg-red-50'}`}>
                <th scope="row" className="flex items-center px-6 py-4 text-gray-900 whitespace-nowrap">
                  <div className="pl-3">
                    <div className="text-base font-semibold">{serverData.name}</div>
                  </div>
                </th>
                <td className="px-6 py-4">{serverData.address}</td>
                <td className="px-6 py-4">
                  <div className="flex items-center">
                    <div className={`h-2.5 w-2.5 rounded-full ${event.isServerUp ? 'bg-green-500' : 'bg-red-500'} mr-2`} />
                    {event.isServerUp ? 'Online' : 'Offline'}
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
}
