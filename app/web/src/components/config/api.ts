export const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000/api'

export const ENDPOINTS = {
  AUTH: {
    LOGIN: '/api/v1/auth/login',
    REFRESH_TOKEN: '/api/v1/auth/refresh',
  },
  USER: {
    WITHOUT_ID: '/api/v1/users',
    WITH_ID: (id: string) => `/api/v1/users/${id}`,
  },
  RESUME: {
    WITHOUT_ID: '/api/v1/resumes',
    WITH_ID: (id: string) => `/api/v1/resumes/${id}`,
  },
} as const
