import Header from '@/components/common/Header.tsx'
import { Outlet } from 'react-router-dom'
import Footer from '@/components/common/Footer.tsx'

function Layout() {
  return (
    <div className="layout-wrapper">
      <Header />

      <main>
        <Outlet />
      </main>

      <Footer />
    </div>
  )
}

export default Layout
