import { Button } from '@/components/ui/button'
import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Link, useNavigate } from 'react-router-dom'
import { Label } from '@/components/ui/label.tsx'
import { Input } from '@/components/ui/input.tsx'
import { useCallback } from 'react'
import * as React from 'react'
import useAxios from '@/hooks/useAxios.ts'
import { ENDPOINTS } from '@/components/config/api.ts'
import { useAuth } from '@/hooks/useAuth.ts'

export function LoginPage() {
  const navigate = useNavigate()
  const loginAPI = useAxios<{ access_token: string; refresh_token: string }>(
    ENDPOINTS.AUTH.LOGIN,
    'post',
  )
  const { login } = useAuth()

  const onSubmit = useCallback(
    async (event: React.SubmitEvent) => {
      event.preventDefault()
      const target = event.currentTarget as HTMLFormElement
      const formData = new FormData(target)
      const data = Object.fromEntries(formData.entries())

      try {
        const tokens = await loginAPI.execute(data)
        if (tokens) {
          await login(tokens)
          navigate('/')
        }
      } catch (e) {
        console.error(e)
      }
    },
    [login, loginAPI, navigate],
  )

  return (
    <div className={'w-full my-24 min-h-96 flex justify-center items-center gap-4'}>
      <form className={'w-lg'} onSubmit={onSubmit}>
        <Card className={'w-full'}>
          <CardHeader className={'flex flex-col justify-center'}>
            <CardTitle className={'text-3xl'}>PAPERLESS.DEV</CardTitle>
            <CardDescription>서비스 이용을 위해 로그인을 진행해주세요.</CardDescription>
          </CardHeader>
          <CardContent>
            <div className={'flex flex-col gap-6'}>
              <div className={'grid gap-2'}>
                <Label htmlFor={'email'}>이메일 주소</Label>
                <Input name={'email'} type={'email'} placeholder={'me@example.com'} required />
              </div>
              <div className={'grid gap-2'}>
                <div className={'flex items-center justify-between'}>
                  <Label htmlFor={'password'}>패스워드</Label>
                </div>
                <Input name={'password'} type={'password'} required />
              </div>
            </div>
          </CardContent>
          <CardFooter className={'flex flex-col gap-2'}>
            <div className={'w-full flex flex-col'}>
              <CardAction className={'w-full flex flex-row justify-between items-center'}>
                <span className={'text-xs text-slate-500'}>비밀번호를 잊으셨습니까?</span>
                <Button variant={'link'} className={'text-xs'} type={'button'}>
                  비밀번호 변경
                </Button>
              </CardAction>
              <CardAction className={'w-full flex flex-row justify-between items-center'}>
                <span className={'text-xs text-slate-500'}>아직 PAPERLESS 회원이 아니신가요?</span>
                <Button variant={'link'} className={'text-xs'} type={'button'}>
                  <Link to={'/register'}>회원가입</Link>
                </Button>
              </CardAction>
            </div>
            <CardAction className={'w-full flex flex-row justify-between items-center'}>
              <Button type={'submit'} className={'w-full'}>
                로그인
              </Button>
            </CardAction>
          </CardFooter>
        </Card>
      </form>
    </div>
  )
}

export default LoginPage
