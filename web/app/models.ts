export interface RegisterData {
  username: string
  email: string
  password: string
  confirmPassword: string
}

export interface LoginData {
  email: string
  password: string
}

export interface User {
  id: string
  username: string
  email: string
  notification_preferences: {
    email_enabled: boolean
  }
}
