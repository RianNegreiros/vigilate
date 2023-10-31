"use client"

import { useState, useRef, useEffect } from "react";

interface DropdownMenuProps {
  openModal?: () => void;
}

export default function DropdownMenu({ openModal }: DropdownMenuProps) {
  const [showDropdown, setShowDropdown] = useState(false);
  const dropdown = useRef<HTMLDivElement>(null);

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

  return (
    <>
      <button
        id="dropdownMenuIconHorizontalButton"
        data-dropdown-toggle="dropdownDotsHorizontal"
        className="absolute top-0 right-0 mt-2 mr-2 inline-flex items-center p-2 text-sm font-medium text-center text-gray-900 bg-white rounded-lg hover:bg-gray-100 focus:ring-4 focus:outline-none dark:text-white focus:ring-gray-50 dark:bg-gray-800 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
        type="button"
        onClick={toggleDropdown}
      >
        <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 16 3">
          <path d="M2 0a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3Zm6.041 0a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM14 0a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3Z" />
        </svg>
      </button>

      <div
        id="dropdownDotsHorizontal"
        className="hidden absolute top-0 right-0 mt-10 mr-2 z-10 bg-white divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700 dark:divide-gray-600"
        ref={dropdown}
      >
        <ul className="py-2 text-sm text-gray-700 dark:text-gray-200" aria-labelledby="dropdownMenuIconHorizontalButton">
          <li>
            <button onClick={openModal} className="block w-full text-left px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Update</button>
          </li>
        </ul>
        <div className="py-2">
          <button onClick={() => { }} className="block w-full text-left px-4 py-2 text-sm text-red-700 hover:bg-red-100 dark:hover:bg-red-600 dark:text-red-200 dark:hover:text-white">Delete</button>
        </div>
      </div>
    </>
  )
}