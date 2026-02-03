import type { UserInfo } from '@/lib/types.ts'
import { type PropsWithChildren, useCallback, useEffect, useState } from 'react'
import useAxios from '@/hooks/useAxios.ts'
import { ENDPOINTS } from '@/components/config/api.ts'
import { AuthContext } from '@/hooks/useAuth'

export const AuthProvider = ({ children }: PropsWithChildren) => {
  const [user, setUser] = useState<UserInfo | null>(null)
  const [isLoading, setIsLoading] = useState<boolean>(true)
  const userInfoAPI = useAxios<{ user: UserInfo }>(ENDPOINTS.USER.WITH_ID('me'))

  const fetchUserInfo = useCallback(async () => {
    try {
      const accessToken = localStorage.getItem('access_token')
      const data = await userInfoAPI.execute(
        {},
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        },
      )

      if (data) {
        setUser(data.user)
      }
    } catch (e) {
      console.error(e)
      localStorage.removeItem('access_token')
    } finally {
      setIsLoading(false)
    }
  }, [userInfoAPI])

  useEffect(() => {
    if (localStorage.getItem('access_token')) {
      fetchUserInfo().then()
    } else {
      setIsLoading(false)
    }
  }, [])

  const login = async (tokens: { access_token: string; refresh_token: string }) => {
    localStorage.setItem('access_token', tokens.access_token)
    localStorage.setItem('refresh_token', tokens.refresh_token)
    await fetchUserInfo()
  }

  const logout = () => {
    localStorage.removeItem('access_token')
    localStorage.removeItem('refresh_token')
    setUser(null)
  }

  return (
    <AuthContext.Provider value={{ user, isAuthenticated: !!user, isLoading, login, logout }}>
      {children}
    </AuthContext.Provider>
  )
}
