import { useState, useCallback } from 'react'
import axios, { AxiosError, type AxiosRequestConfig, type Method } from 'axios'

// API 응답 데이터의 타입을 제네릭 T로 받습니다.
interface UseLazyAxiosReturn<T> {
  execute: (body?: unknown, config?: AxiosRequestConfig) => Promise<T>
  data: T | null
  loading: boolean
  error: AxiosError | null
}

const useAxios = <T = unknown>(
  url: string,
  method: Method = 'get',
  baseConfig?: AxiosRequestConfig,
): UseLazyAxiosReturn<T> => {
  const [data, setData] = useState<T | null>(null)
  const [error, setError] = useState<AxiosError | null>(null)
  const [loading, setLoading] = useState<boolean>(false)

  const execute = useCallback(
    async (body?: unknown, config?: AxiosRequestConfig): Promise<T> => {
      setLoading(true)
      setError(null)

      try {
        const response = await axios({
          url,
          method,
          data: body,
          ...baseConfig,
          ...config,
        })

        setData(response.data)
        return response.data
      } catch (err) {
        const axiosError = err as AxiosError
        setError(axiosError)
        throw axiosError
      } finally {
        setLoading(false)
      }
    },
    [url, method, baseConfig],
  )

  return { execute, data, loading, error }
}

export default useAxios
