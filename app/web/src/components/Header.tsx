import logo from '../assets/logo.svg'
import { Button } from '@/components/ui/button.tsx'

function Header() {
  return (
    <header
      className={
        'sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-slate-200 px-6 py-4 flex justify-between items-center'
      }
    >
      <div className={'flex items-center gap-2 cursor-pointer group'}>
        <img
          src={logo}
          alt={'logo'}
          className={
            'w-9 h-9 rounded-xl flex items-center justify-center text-white font-bold group-hover:rotate-6 transition-transform'
          }
        />
        <span className={'font-extrabold text-xl tracking-tighter'}>PAPERLESS.DEV</span>
      </div>

      <div>
        <Button>로그인</Button>
      </div>
    </header>
  )
}

export default Header
