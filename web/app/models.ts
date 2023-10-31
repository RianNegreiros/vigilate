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

export interface Server {
  id: string
  user_id: string
  name: string
  address: string
  is_active: boolean
  last_check_time: string
  next_check_time: string
}

export interface CreateServer {
  user_id?: number
  name: string
  address: string
}

export interface UpdateServer {
  id?: string
  name: string
  address: string
}
