"use client"

import { FormEvent, useState } from "react"
import { LoginData } from "../models"
import { useRouter } from "next/navigation"
import { login } from "../util/api"
import FormButton from "../components/FormButton"
import AuthInput from "../components/Input"
import Label from "../components/Label"
import AuthHeader from "../components/AuthHeader"
import Logo from "../components/Logo"

export default function LoginPage() {
  const [loginData, setLoginData] = useState<LoginData>({
    email: '',
    password: ''
  })

  const router = useRouter()

  const handleInputChange = (event: FormEvent<HTMLInputElement>) => {
    const { name, value } = event.currentTarget
    setLoginData({ ...loginData, [name]: value })
  }

  const handleLogin = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
  
    const formData = new FormData(event.currentTarget);
  
    try {
      await login({
        email: formData.get("email") as string,
        password: formData.get("password") as string,
      });
      router.push("/dashboard");
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <section className="bg-gray-50 dark:bg-gray-900">
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
        <a href="/" className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
          <Logo />
          Vigilate
        </a>
        <div
          className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <AuthHeader text="Sign in to your account" />
            <form className="space-y-4 md:space-y-6" onSubmit={handleLogin}>
              <div>
                <Label htmlFor="email" text="Your email" />
                <AuthInput value={loginData.email}
                  onChange={handleInputChange}
                  type="email" name="email" id="email" placeholder="johndoe@mail.com"
                  required
                />
              </div>
              <div>
                <Label htmlFor="password" text="Your password" />
                <AuthInput value={loginData.password}
                  onChange={handleInputChange}
                  type="password" name="password" id="password" placeholder="************"
                  required
                />
              </div>
              {/* <div className="flex items-center justify-between">
                        <div className="flex items-start">
                            <div className="flex items-center h-5">
                              <input id="remember" aria-describedby="remember" type="checkbox" className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-primary-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-primary-600 dark:ring-offset-gray-800" required>
                            </div>
                            <div className="ml-3 text-sm">
                              <label htmlFor="remember" className="text-gray-500 dark:text-gray-300">Remember me</label>
                            </div>
                        </div>
                        <a href="#" className="text-sm font-medium text-primary-600 hover:underline dark:text-primary-500">Forgot password?</a>
                    </div> */}
              <FormButton text="Sign in" />
              <p className="text-sm font-light text-gray-500 dark:text-gray-400">
                Don’t have an account yet? <a href="/register"
                  className="font-medium text-primary-600 hover:underline dark:text-primary-500">Sign up</a>
              </p>
            </form>
          </div>
        </div>
      </div>
    </section>
  )
}