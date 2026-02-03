import { useState, useCallback } from 'react'
import axios, { type AxiosRequestConfig } from 'axios'

const api = axios.create({})

export function useAxios<T>(url: string, method: string = 'get') {
  const [data, setData] = useState<T | null>(null)
  const [error, setError] = useState<unknown | null>(null)
  const [loading, setLoading] = useState<boolean>(false)

  const execute = useCallback(
    async (body?: unknown, customConfig?: AxiosRequestConfig): Promise<T | null> => {
      setLoading(true)
      setError(null)

      try {
        const response = await api.request<T>({
          url,
          method,
          data: body,
          ...customConfig,
        })

        setData(response.data)
        return response.data
      } catch (err) {
        setError(err)
        throw err
      } finally {
        setLoading(false)
      }
    },
    [url, method],
  )

  return { data, error, loading, execute }
}

export default useAxios
