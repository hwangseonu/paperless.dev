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
import { useCallback } from 'react'
import * as React from 'react'
import useAxios from '@/hooks/useAxios.ts'
import { ENDPOINTS } from '@/components/config/api.ts'
import { LoaderIcon } from 'lucide-react'
import { cn } from '@/lib/utils.ts'
import { Label } from '@/components/ui/label.tsx'
import { Input } from '@/components/ui/input.tsx'

export function RegisterPage() {
  const navigate = useNavigate()
  const registerAPI = useAxios(ENDPOINTS.AUTH.REGISTER, 'post')

  const onSubmit = useCallback(
    async (event: React.SubmitEvent) => {
      event.preventDefault()
      const target = event.currentTarget as HTMLFormElement
      const formData = new FormData(target)
      const data = Object.fromEntries(formData.entries())

      try {
        await registerAPI.execute(data)
        navigate('/verify')
      } catch (error) {
        console.error(error)
      }
    },
    [navigate, registerAPI],
  )

  return (
    <div className={'w-full my-24 min-h-96 flex justify-center items-center gap-4'}>
      <form className={'w-lg'} onSubmit={onSubmit}>
        <Card className={'w-full'}>
          <CardHeader className={'flex flex-col justify-center'}>
            <CardTitle className={'text-3xl'}>PAPERLESS.DEV</CardTitle>
            <CardDescription>
              회원가입을 완료하고 종이없이 최고의 이력서를 만들어보세요.
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className={'flex flex-col gap-6'}>
              <div className={'grid gap-2'}>
                <Label htmlFor={'email'}>닉네임</Label>
                <Input name={'nickname'} required />
              </div>
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
          <CardFooter className={'flex-col gap-2'}>
            <CardAction className={'w-full flex flex-row justify-between items-center'}>
              <span className={'text-xs text-slate-500'}>이미 PAPERLESS 회원이신가요?</span>
              <Button variant={'link'} className={'text-xs'}>
                <Link to={'/login'}>로그인</Link>
              </Button>
            </CardAction>
            <CardAction className={'w-full flex flex-row justify-between items-center'}>
              <Button type={'submit'} className={'w-full'} disabled={registerAPI.loading}>
                {registerAPI.loading ? (
                  <LoaderIcon role={'status'} className={cn('size-4 animate-spin')} />
                ) : (
                  '회원가입'
                )}
              </Button>
            </CardAction>
          </CardFooter>
        </Card>
      </form>
    </div>
  )
}

export default RegisterPage
