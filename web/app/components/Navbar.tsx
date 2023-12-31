"use client"

import { useRouter } from "next/navigation";
import { useEffect, useRef, useState } from "react";
import { logout } from "../util/api";
import Link from "next/link";
import Cookies from "js-cookie";
import Logo from "./Logo";

interface NavBarProps {
  openModal?: () => void;
  pathname: string;
}

export default function NavBar({ openModal, pathname }: NavBarProps) {
  const [showDropdown, setShowDropdown] = useState(false);
  const [user, setUser] = useState({ username: "", email: "" });
  const dropdown = useRef<HTMLDivElement>(null);
  const router = useRouter();

  useEffect(() => {
    const cookie = Cookies.get("user");
    const userData = JSON.parse(cookie || "{}");
    setUser(userData);
  }, []);

  useEffect(() => {
    function handleClickOutside(event: MouseEvent) {
      if (dropdown.current && !dropdown.current.contains(event.target as Node)) {
        setShowDropdown(false);
        dropdown.current.classList.add("hidden");
      }
    }
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [showDropdown]);

  const toggleDropdown = () => {
    setShowDropdown(!showDropdown);
    if (dropdown.current) {
      dropdown.current.classList.toggle("hidden");
    }
  }

  const handleLogout = async () => {
    await logout();
    router.push("/");
  }

  return (
    <>
      <nav className="bg-gray-800">
        <div className="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">

          <div className="relative flex h-16 items-center justify-between">
            <div className="absolute inset-y-0 left-0 flex items-center sm:hidden">
              <button type="button"
                className="relative inline-flex items-center justify-center rounded-md p-2 text-gray-400 hover:bg-gray-700 hover:text-white focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white"
                aria-controls="mobile-menu" aria-expanded="false">
                <span className="absolute -inset-0.5"></span>
                <span className="sr-only">Open main menu</span>
                <svg className="block h-6 w-6" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor"
                  aria-hidden="true">
                  <path strokeLinecap="round" strokeLinejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
                </svg>
                <svg className="hidden h-6 w-6" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor"
                  aria-hidden="true">
                  <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
            <div className="flex flex-1 items-center justify-center sm:items-stretch sm:justify-start">
              <a href="/">
                <div className="flex flex-shrink-0 items-center">
                  <Logo />
                  <span className="text-white text-2xl font-bold ml-2">Vigilate</span>
                </div>
              </a>
            </div>

            {pathname === "/dashboard" && (
              <button
                type="button"
                onClick={openModal}
                className="text-white inline-flex items-center bg-primary-700 hover:bg-primary-900 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover-bg-primary-700 dark:focus:ring-primary-800 hover:shadow-lg transition-background duration-300">
                New
                <svg className="ml-1 w-3 h-4 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 18 18">
                  <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M9 1v16M1 9h16" />
                </svg>
              </button>
            )}

            {pathname !== "/dashboard" && (
              <Link
                href="/dashboard"
                className="text-white inline-flex items-center bg-primary-700 hover:bg-primary-900 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover-bg-primary-700 dark:focus:ring-primary-800 hover:shadow-lg transition-background duration-300">
                Dashboard
              </Link>
            )}

            <div className="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">

              <div className="relative ml-3">
                <div>
                  <button type="button"
                    className="relative flex rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
                    id="user-menu-button" aria-expanded="false" aria-haspopup="true" onClick={toggleDropdown}>
                    <span className="absolute -inset-1.5"></span>
                    <span className="sr-only">Open user menu</span>

                    <div className="relative w-10 h-10 overflow-hidden bg-gray-100 rounded-full dark:bg-gray-600">
                      <svg className="absolute w-12 h-12 text-gray-400 -left-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clipRule="evenodd"></path></svg>
                    </div>

                  </button>
                </div>

                <div
                  className="hidden absolute right-0 z-10 mt-2 w-48 origin-top-right my-4 text-base list-none bg-white divide-y divide-gray-100 rounded-lg py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none dark:bg-gray-700 dark:divide-gray-600"
                  role="menu" aria-orientation="vertical" aria-labelledby="user-menu-button" tabIndex={-1} id="user-menu" ref={dropdown}>
                  {pathname === "/dashboard" && (
                    <div className="px-4 py-3">
                      <span className="block text-sm text-gray-900 dark:text-white">{user.username}</span>
                      <span className="block text-sm text-gray-500 truncate dark:text-gray-400">{user.email}</span>
                    </div>
                  )}
                  <ul className="py-2" aria-labelledby="user-menu-button">
                    {pathname === "/dashboard" && (
                      <li>
                        <a href="/account"
                          className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white">Account</a>
                      </li>
                    )}
                    <li>
                      <a onClick={handleLogout}
                        className="block px-4 py-2 text-sm text-gray-700 cursor-pointer hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white"
                      >
                        Sign out
                      </a>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
      </nav>
    </>

  );
}