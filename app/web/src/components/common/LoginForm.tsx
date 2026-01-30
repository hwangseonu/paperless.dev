import { Input } from '@/components/ui/input.tsx'
import { Label } from '@/components/ui/label.tsx'

function LoginForm() {
  return (
    <form>
      <div className="flex flex-col gap-6">
        <div className="grid gap-2">
          <Label htmlFor="email">이메일 주소</Label>
          <Input id="email" type="email" placeholder="m@example.com" required />
        </div>
        <div className="grid gap-2">
          <div className="flex items-center">
            <Label htmlFor="password">패스워드</Label>
          </div>
          <Input id="password" type="password" required />
        </div>
      </div>
    </form>
  )
}

export default LoginForm
