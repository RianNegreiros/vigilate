"use client"

import { useState } from "react";
import { UpdateServer } from "../models";
import Input from "./Input";
import Label from "./Label";

interface ModalProps {
  hideModal: boolean;
  closeModal: () => void;
  updateServer: (server: UpdateServer) => void;
  server: UpdateServer;
}

export default function UpdateServerModal({ hideModal, closeModal, updateServer, server }: ModalProps) {
  const [serverData, setServerData] = useState<UpdateServer>({
    id: server.id,
    name: server.name,
    address: server.address
  })

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const formData = new FormData(e.currentTarget);

    try {
      updateServer({
        name: formData.get("name") as string,
        address: formData.get("address") as string,
      });
    } catch (error) {
      console.log(error);
    } finally {
      closeModal();
    }
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setServerData({
      ...serverData,
      [e.target.id]: e.target.value
    })
  }

  return (
    <div id="default-modal" tabIndex={-1} aria-hidden="true"
      className={`fixed inset-0 z-50 overflow-y-auto ${hideModal ? "hidden" : ""}`}>
      <div className="relative w-full max-w-2xl max-h-full top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
        <div className="relative bg-white rounded-lg shadow dark:bg-gray-700">
          <div className="flex items-start justify-between p-4 border-b rounded-t dark:border-gray-600">
            <h3 className="text-xl font-semibold text-gray-900 dark:text-white">
              Update server
            </h3>
            <button type="button"
              onClick={closeModal}
              className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ml-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white"
              data-modal-hide="default-modal">
              <svg className="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
                <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"
                  d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6" />
              </svg>
              <span className="sr-only">Close modal</span>
            </button>
          </div>

          <form className="p-6 space-y-6" onSubmit={handleSubmit}>
            <div className="mb-6">
              <Label htmlFor="name" text="Server Name" />
              <Input
                id="name"
                type="text"
                name="name"
                value={serverData.name}
                onChange={handleChange}
                placeholder="My Server"
                required />
            </div>
            <div className="mb-6">
              <Label htmlFor="address" text="Server Address" />
              <Input
                id="address"
                type="text"
                name="address"
                value={serverData.address}
                onChange={handleChange}
                placeholder="https://example.com"
                required />
            </div>
            <div className="flex items-center p-6 space-x-2 border-t border-gray-200 rounded-b dark:border-gray-600">
              <button data-modal-hide="default-modal" type="submit"
                className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                Update
              </button>
              <button data-modal-hide="default-modal" type="button"
                onClick={closeModal}
                className="text-gray-500 bg-white hover:bg-gray-100 focus:ring-4 focus:outline-none focus:ring-blue-300 rounded-lg border border-gray-200 text-sm font-medium px-5 py-2.5 hover:text-gray-900 focus:z-10 dark:bg-gray-700 dark:text-gray-300 dark:border-gray-500 dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-gray-600"
              >
                Cancel
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}