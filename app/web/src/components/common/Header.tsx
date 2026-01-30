import logo from '../../assets/logo.svg'
import { Button } from '@/components/ui/button.tsx'
import { Link } from 'react-router-dom'

function Header() {
  return (
    <header
      className={
        'sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-slate-200 px-6 py-4 flex justify-between items-center'
      }
    >
      <Link to="/">
        <div className={'flex items-center gap-2 cursor-pointer group'}>
          <img
            src={logo}
            alt={'logo'}
            className={'size-9 group-hover:rotate-6 transition-transform'}
          />
          <span className={'font-extrabold text-xl tracking-tighter'}>PAPERLESS.DEV</span>
        </div>
      </Link>

      <Link to="/login">
        <div>
          <Button>로그인</Button>
        </div>
      </Link>
    </header>
  )
}

export default Header
