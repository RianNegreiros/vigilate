"use client"

import Image from 'next/image';
import Link from 'next/link';
import Cookies from 'js-cookie';
import { useEffect, useState } from 'react';

export default function Hero() {
  const [userLoggedIn, setUserLoggedIn] = useState(false);

  useEffect(() => {
    const token = Cookies.get('user');
    if (token) {
      setUserLoggedIn(true);
    }
  }, []);

  return (
    <>
      <nav className="bg-white border-gray-200 dark:bg-gray-900">
        <div className="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl p-4">
          <Link href="https://github.com/RianNegreiros/vigilate" className="flex items-center">
            <svg className="h-6 mr-1 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg"
              fill="currentColor" viewBox="0 0 20 14">
              <path
                d="M10 0C4.612 0 0 5.336 0 7c0 1.742 3.546 7 10 7 6.454 0 10-5.258 10-7 0-1.664-4.612-7-10-7Zm0 10a3 3 0 1 1 0-6 3 3 0 0 1 0 6Z" />
            </svg>
            <span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Vigilate</span>
          </Link>
          <div className="flex items-center">
            {userLoggedIn ? (
              <Link href="/account" className="text-sm text-blue-600 dark:text-blue-500 hover:underline">Account</Link>
            ) : (
              <Link href="/login" className="text-sm text-blue-600 dark:text-blue-500 hover:underline">Login</Link>
            )}
          </div>
        </div>
      </nav>

      <div className='flex flex-col min-h-screen'>
        <section className="bg-white dark:bg-gray-900 flex-grow">
          <div className="gap-8 items-center py-8 px-4 mx-auto max-w-screen-xl xl:gap-16 md:grid md:grid-cols-2 sm:py-16 lg:px-6">
            <div className="w-full dark:hidden">
              <Image src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/cta/cta-dashboard-mockup.svg" alt="dashboard image"
                width={400} height={300} priority />
            </div>
            <div className="w-full hidden dark:block">
              <Image src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/cta/cta-dashboard-mockup-dark.svg" alt="dashboard image"
                width={400} height={300} />
            </div>
            <div className="mt-4 md:mt-0">
              <h2 className="mb-4 text-4xl tracking-tight font-extrabold text-gray-900 dark:text-white">Monitor Remote Servers in Real Time and Stay Informed</h2>
              <p className="mb-6 font-light text-gray-500 md:text-lg dark:text-gray-400">Our platform allows you to monitor remote servers in real time, providing you with critical data and immediate notifications. Keep your systems running smoothly and avoid downtime with our monitoring tool.</p>
              {userLoggedIn ? (
                <Link href="/dashboard" className="inline-flex items-center text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:focus:ring-primary-900">
                  Dashboard
                  <svg className="ml-2 -mr-1 w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clipRule="evenodd"></path></svg>
                </Link>
              ) : (
                <Link href="/register" className="inline-flex items-center text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:focus:ring-primary-900">
                  Get Started
                  <svg className="ml-2 -mr-1 w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clipRule="evenodd"></path></svg>
                </Link>
              )}
            </div>
          </div>
        </section>
      </div>

      <footer className="bg-white rounded-lg shadow dark:bg-gray-900">
        <div className="w-full max-w-screen-xl mx-auto p-4 md:py-8">
          <div className="sm:flex sm:items-center sm:justify-between">
            <a href="https://github.com/RianNegreiros/vigilate" className="flex items-center mb-4 sm:mb-0">
              <svg className="h-6 mr-3 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg"
                fill="currentColor" viewBox="0 0 20 14">
                <path
                  d="M10 0C4.612 0 0 5.336 0 7c0 1.742 3.546 7 10 7 6.454 0 10-5.258 10-7 0-1.664-4.612-7-10-7Zm0 10a3 3 0 1 1 0-6 3 3 0 0 1 0 6Z" />
              </svg>
              <span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Vigilate</span>
            </a>
            <ul className="flex flex-wrap items-center mb-6 text-sm font-medium text-gray-500 sm:mb-0 dark:text-gray-400">
              <li>
                <a href="https://github.com/RianNegreiros/vigilate" className="mr-4 hover:underline md:mr-6 ">About</a>
              </li>
              <li>
                <a href="/terms" className="mr-4 hover:underline md:mr-6">Privacy Policy</a>
              </li>
              <li>
                <a href="https://github.com/RianNegreiros/vigilate/blob/main/LICENSE" className="mr-4 hover:underline md:mr-6 ">Licensing</a>
              </li>
              <li>
                <a href="mailto:riannegreiros@gmail.com" className="hover:underline">Contact</a>
              </li>
            </ul>
          </div>
          <hr className="my-6 border-gray-200 sm:mx-auto dark:border-gray-700 lg:my-8" />
          <span className="block text-sm text-gray-500 sm:text-center dark:text-gray-400">© 2023 <a href="https://github.com/RianNegreiros/vigilate" className="hover:underline">Vigilate™</a>. All Rights Reserved.</span>
        </div>
      </footer>
    </>
  );
}