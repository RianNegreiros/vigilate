"use client"

import { usePathname, useRouter } from "next/navigation";
import NavBar from "../components/Navbar";
import Cookies from "js-cookie";
import { getUserById, updateEmailNotifications } from "../util/api";
import { useEffect, useState } from "react";
import { User } from "../models";
import { AxiosError } from "axios";

export default function AccountPage() {
  const [userData, setUserData] = useState<User>({
    id: "",
    username: "",
    email: "",
    notification_preferences: {
      email_enabled: false,
    },
  });
  const pathname = usePathname();
  const router = useRouter();

  const cookie = Cookies.get("user");
  const user = cookie ? JSON.parse(cookie) : null;
  const id = user ? user.id : "";

  useEffect(() => {
    async function getUser() {
      try {
        const userData = await getUserById(id);
        if (userData !== null) {
          setUserData(userData);
        }
      } catch (error: AxiosError | any) {
        if (error.response.status === 401 || error.response.status === 404) {
          router.push("/login");
        }
      }
    }
    getUser();
  });

  const handleEmailChange = async () => {
    await updateEmailNotifications();
    setUserData({
      ...userData,
      notification_preferences: {
        ...userData.notification_preferences,
        email_enabled: !userData.notification_preferences.email_enabled,
      },
    });
  };

  return (
    <>
      <NavBar pathname={pathname} />

      <div className="p-4 min-h-screen bg-gray-100 dark:bg-gray-900">
        <div className="max-w-6xl mx-auto">
          <div className="bg-white rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 hover:shadow-lg hover:border-gray-300 p-6">
            <div className="mb-4">
              <h3 className="text-2xl font-semibold text-gray-900 md:text-3xl dark:text-white">
                Account
              </h3>
              <p className="text-base font-bold text-gray-700 md:text-lg dark:text-gray-400">
                Account details and settings
              </p>
            </div>
            <div className="py-4">
              <p className="text-sm font-semibold text-gray-900 dark:text-white">
                Username
              </p>
              <p className="text-sm text-gray-500 dark:text-gray-400">
                {userData.username}
              </p>
            </div>
            <div className="py-4">
              <p className="text-sm font-semibold text-gray-900 dark:text-white">
                Email address
              </p>
              <p className="text-sm text-gray-500 dark:text-gray-400">
                {userData.email}
              </p>
            </div>
            <div className="py-4">
              <p className="text-sm font-semibold text-gray-900 dark:text-white">
                Notifications Settings
              </p>
              <ul
                role="list"
                className="divide-y mt-4 divide-gray-100 rounded-md border border-gray-200"
              >
                <li className="flex items-center justify-between py-4 pl-4 pr-5 text-sm leading-6">
                  <div className="flex w-0 flex-1 items-center">
                    <div className="ml-4 flex min-w-0 flex-1 gap-2">
                      <span className="truncate font-medium">Email</span>
                    </div>
                  </div>
                  <div className="ml-4 flex-shrink-0">
                    <label className="relative inline-flex items-center cursor-pointer">
                      <input
                        type="checkbox"
                        value=""
                        className="sr-only peer"
                        checked={userData.notification_preferences.email_enabled}
                        onChange={handleEmailChange}
                      />
                      <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"></div>
                    </label>
                  </div>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
