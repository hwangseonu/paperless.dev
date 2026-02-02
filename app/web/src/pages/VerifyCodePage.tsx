import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card.tsx'
import { Input } from '@/components/ui/input.tsx'
import { Label } from '@/components/ui/label.tsx'
import { Button } from '@/components/ui/button.tsx'
import { useCallback, useState } from 'react'
import useAxios from '@/hooks/useAxios.ts'
import { ENDPOINTS } from '@/components/config/api.ts'
import * as React from 'react'
import { LoaderIcon } from 'lucide-react'
import { cn } from '@/lib/utils.ts'
import { useNavigate } from 'react-router-dom'

function VerifyCodePage() {
  const navigate = useNavigate()
  const [code, setCode] = useState<string>('')

  const userCreateApi = useAxios(ENDPOINTS.USER.WITHOUT_ID, 'POST')

  const onSubmit = useCallback(
    async (event: React.SubmitEvent) => {
      event.preventDefault()
      try {
        await userCreateApi.execute({}, { url: `${ENDPOINTS.USER.WITHOUT_ID}?code=${code}` })
        alert('인증되었습니다.')
        navigate('/login')
      } catch (error) {
        console.error(error)
      }
    },
    [code, userCreateApi, navigate],
  )

  return (
    <div className={'w-full my-24 min-h-96 flex justify-center items-center gap-4'}>
      <form className={'w-lg'} onSubmit={onSubmit}>
        <Card className={'w-full'}>
          <CardHeader className={'flex flex-col justify-center'}>
            <CardTitle className={'text-3xl'}>PAPERLESS.DEV</CardTitle>
            <CardDescription>
              이메일로 인증 코드를 보냈습니다. 인증 코드는 5분후 만료됩니다.
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className={'flex flex-col gap-6'}>
              <div className={'grid gap-2'}>
                <Label htmlFor={'code'}>인증 코드</Label>
                <Input name={'code'} required onChange={(e) => setCode(e.target.value)} />
              </div>
            </div>
          </CardContent>
          <CardFooter className={'flex-col gap-6'}>
            <CardAction className={'w-full flex flex-row justify-between items-center'}>
              <span className={'text-xs text-slate-500'}>코드가 오지 않았습니까?</span>
              <Button variant={'link'} className={'text-xs'} type={'button'}>
                재전송
              </Button>
            </CardAction>
            <Button type={'submit'} className={'w-full'}>
              {userCreateApi.loading ? (
                <LoaderIcon role={'status'} className={cn('size-4 animate-spin')} />
              ) : (
                '인증하기'
              )}
            </Button>
          </CardFooter>
        </Card>
      </form>
    </div>
  )
}

export default VerifyCodePage
