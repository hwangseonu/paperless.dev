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
import LoginForm from '@/components/common/LoginForm.tsx'

export function LoginPage() {
  return (
    <div className="w-full my-24 min-h-96 flex justify-center items-center gap-4">
      <Card className="w-full max-w-sm">
        <CardHeader className={'flex flex-col justify-center'}>
          <CardTitle className={'text-3xl'}>PAPERLESS.DEV</CardTitle>
          <CardDescription>서비스 이용을 위해 로그인을 진행해주세요.</CardDescription>
        </CardHeader>
        <CardContent>
          <LoginForm />
          <CardAction className={'w-full flex flex-row justify-between'}>
            <Button variant={'link'}>회원가입</Button>
            <Button variant={'link'} className={'text-slate-500'}>
              <a href="#">비밀번호를 잊으셨습니까?</a>
            </Button>
          </CardAction>
        </CardContent>
        <CardFooter className="flex-col gap-2">
          <Button type="submit" className="w-full">
            로그인
          </Button>
        </CardFooter>
      </Card>
    </div>
  )
}

export default LoginPage
