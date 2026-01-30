import { Globe, Mail } from 'lucide-react'
import { Card, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'

type props = {
  title: string
  description: string
  email: string
  url: string
  image: string
}

export function Information({ title, description, email, url, image }: props) {
  return (
    <Card className="relative mx-auto w-full pt-0 overflow-hidden">
      <div className="absolute inset-0 z-30 aspect-video bg-black/35" />
      <img
        src={image}
        alt=""
        className="relative z-20 aspect-video w-full object-cover brightness-60 grayscale dark:brightness-40"
      />
      <CardHeader>
        <CardTitle className={'text-2xl'}>{title}</CardTitle>
        <CardDescription>{description}</CardDescription>
      </CardHeader>
      <CardFooter className={'flex flex-col items-start gap-3 border-t mx-6'}>
        <div className="flex items-center gap-3 text-sm text-slate-500">
          <Mail size={16} />
          <span>{email}</span>
        </div>
        <div className="flex items-center gap-3 text-sm text-slate-500">
          <Globe size={16} />
          <a
            href={url}
            target={'_blank'}
            className="text-indigo-600 hover:underline cursor-pointer"
          >
            {url}
          </a>
        </div>
      </CardFooter>
    </Card>
  )
}

export default Information
