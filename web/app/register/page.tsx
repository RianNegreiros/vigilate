"use client"

import { FormEvent, useState } from "react"
import { RegisterData } from "../models"
import { useRouter } from "next/navigation"
import { login, register } from "../util/api"
import FormButton from "../components/FormButton"
import AuthInput from "../components/Input"
import Label from "../components/Label"
import AuthHeader from "../components/AuthHeader"
import Logo from "../components/Logo"

export default function RegisterPage() {
  const [formData, setFormData] = useState<RegisterData>({
    username: "",
    email: "",
    password: "",
    confirmPassword: ""
  })

  const router = useRouter()

  const handleInputChange = (event: FormEvent<HTMLInputElement>) => {
    const { name, value } = event.currentTarget
    setFormData({ ...formData, [name]: value })
  }

  const handleRegister = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const formData = new FormData(event.currentTarget);

    try {
      await register({
        username: formData.get("username") as string,
        email: formData.get("email") as string,
        password: formData.get("password") as string,
        confirmPassword: formData.get("confirmPassword") as string,
      });

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
        <div className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <AuthHeader text="Sign up for an account" />
            <form className="space-y-4 md:space-y-6" onSubmit={handleRegister}>
              <div>
                <Label htmlFor="email" text="Your email" />
                <AuthInput value={formData.email}
                  onChange={handleInputChange}
                  type="email" name="email" id="email" placeholder="johndoe@mail.com"
                  required
                />
              </div>
              <div>
                <Label htmlFor="username" text="Username" />
                <AuthInput value={formData.username}
                  onChange={handleInputChange}
                  type="text" name="username" id="username" placeholder="John Doe"
                  required
                />
              </div>
              <div>
                <Label htmlFor="password" text="Password" />
                <AuthInput value={formData.password}
                  onChange={handleInputChange}
                  type="password" name="password" id="password" placeholder="••••••••"
                  required
                />
              </div>
              <div>
                <Label htmlFor="confirmPassword" text="Confirm Password" />
                <AuthInput value={formData.confirmPassword}
                  onChange={handleInputChange}
                  type="password" name="confirmPassword" id="confirmPassword" placeholder="••••••••"
                  required
                />
              </div>
              <div className="flex items-start">
                <div className="flex items-center h-5">
                  <input id="terms" aria-describedby="terms" type="checkbox" className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-primary-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-primary-600 dark:ring-offset-gray-800" required />
                </div>
                <div className="ml-3 text-sm">
                  <label htmlFor="terms" className="font-light text-gray-500 dark:text-gray-300">I accept the <a className="font-medium text-primary-600 hover:underline dark:text-primary-500" href="/terms">Terms and Conditions</a></label>
                </div>
              </div>
              <FormButton text="Create account" />
              <p className="text-sm font-light text-gray-500 dark:text-gray-400">
                Already have an account? <a href="/login" className="font-medium text-primary-600 hover:underline dark:text-primary-500">Login here</a>
              </p>
            </form>
          </div>
        </div>
      </div>
    </section>
  )
}