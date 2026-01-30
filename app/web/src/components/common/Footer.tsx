function Footer() {
  return (
    <footer className="py-16 bg-white border-t border-slate-100">
      <div className="max-w-7xl mx-auto px-6 text-center">
        <div className="flex items-center justify-center gap-2 mb-6">
          <div className="w-6 h-6 bg-slate-300 rounded flex items-center justify-center text-white text-[10px] font-bold">
            R
          </div>
          <span className="font-bold text-slate-400 tracking-tight">PAPERLESS.DEV</span>
        </div>
        <p className="text-slate-400 text-sm mb-8">
          최고의 커리어를 위한 스마트한 도구. <br />© 2026 hwangseonu. All rights reserved.
        </p>
        <div className="flex justify-center gap-6 text-slate-400 text-xs font-medium">
          <a href="#" className="hover:text-indigo-600">
            이용약관
          </a>
          <a href="#" className="hover:text-indigo-600">
            개인정보처리방침
          </a>
          <a href="#" className="hover:text-indigo-600">
            고객센터
          </a>
        </div>
      </div>
    </footer>
  )
}

export default Footer
