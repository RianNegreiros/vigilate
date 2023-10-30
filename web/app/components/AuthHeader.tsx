interface AuthHeaderProps {
  text: string
}

export default function AuthHeader({ text }: AuthHeaderProps) {
  return (
    <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
      {text}
    </h1>
  )
}