import { useState } from 'react'
import type { AxiosError, AxiosRequestConfig } from 'axios'
import api from '@/lib/api.ts'

export const useAxios = <T>(config: AxiosRequestConfig) => {
  const [response, setResponse] = useState<T | null>(null)
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<AxiosError | null>(null)

  const request = async (override?: AxiosRequestConfig) => {
    setLoading(true)
    setError(null)

    try {
      const res = await api({ ...config, ...override })
      setResponse(res as T)
      return res
    } catch (error) {
      setError(error as AxiosError)
      throw error
    } finally {
      setLoading(false)
    }
  }

  return { response, error, loading, refetch: request }
}
