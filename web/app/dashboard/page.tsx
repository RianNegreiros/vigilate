"use client"

import { useEffect, useState } from "react";
import NavBar from "../components/Navbar";
import CreateServerModal from "../components/CreateServerModal";
import { usePathname, useRouter } from "next/navigation";
import { createServer, deleteServer, getServers, updateServer } from "../util/api";
import { CreateServer, Server, UpdateServer } from "../models";
import { AxiosError } from "axios";
import ServerCard from "../components/ServerCard";
import ServerLoadingSkeleton from "../components/ServerLoadingSkeleton";

export default function DashboardPage() {
  const [isModalOpen, setIsModalOpen] = useState(true);
  const [servers, setServers] = useState<Server[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  let pathname = usePathname();

  const router = useRouter();

  const openModal = () => {
    setIsModalOpen(false);
  };

  const closeModal = () => {
    setIsModalOpen(true);
  };

  useEffect(() => {
    async function fetchData() {
      try {
        setIsLoading(true);
        const response = await getServers();
        if (response !== null) {
          setServers(response);
        }
      } catch (error: AxiosError | any) {
        if (error.response.status === 401) {
          router.push("/login");
        }
      } finally {
        setIsLoading(false);
      }
    }
    fetchData();
  }, [router]);

  const handleCreateServer = async (formData: CreateServer) => {
    try {
      await createServer(formData);
      const updatedServers = await getServers();
      setServers(updatedServers);
    } catch (error) {
      console.log(error);
    }
  }

  const handleUpdateServer = async (formData: UpdateServer, id:string) => {
    try {
      await updateServer(formData, id);
      const updatedServers = await getServers();
      setServers(updatedServers);
    } catch (error) {
      console.log(error);
    }
  }

  const handleDeleteServer = async (id: string) => {
    try {
      await deleteServer(id);
      const updatedServers = await getServers();
      if (updatedServers === null) {
        setServers([]);
      }
      setServers(updatedServers);
    } catch (error) {
      console.log(error);
    }
  }

  return (
    <>
      <NavBar openModal={openModal} pathname={pathname} />

      <CreateServerModal hideModal={isModalOpen} closeModal={closeModal} createServer={handleCreateServer} />

      <div className="min-h-screen max-w-6xl mx-auto grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 p-4">
        {isLoading ? (
          <ServerLoadingSkeleton />
        ) : (
          servers.map((server) => (
            <ServerCard key={server.id} server={server} deleteServer={handleDeleteServer} updateServer={handleUpdateServer} />
          ))
        )}
      </div>
    </>
  );
}
