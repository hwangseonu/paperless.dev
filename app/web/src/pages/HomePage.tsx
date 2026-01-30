import { ArrowRight, Zap } from 'lucide-react'
import { Link } from 'react-router-dom'
import { Button } from '@/components/ui/button.tsx'
import { Badge } from '@/components/ui/badge.tsx'

function HomePage() {
  return (
    <main>
      <section className="relative overflow-hidden pt-16 pb-24 lg:pt-32 lg:pb-40 bg-linear-to-b from-slate-50 to-white">
        <div className="max-w-7xl mx-auto px-6 lg:px-8 text-center relative z-10">
          <Badge
            variant={'outline'}
            className="gap-2 px-4 py-2 bg-indigo-50 text-indigo-700 text-sm font-semibold mb-8 border"
          >
            <Zap size={16} />
            <span>3분 만에 완성하는 전문 이력서</span>
          </Badge>
          <h1 className="text-5xl lg:text-7xl font-extrabold text-slate-900 tracking-tight mb-8">
            당신의 커리어를 <br />
            <span className="text-transparent bg-clip-text bg-linear-to-r from-indigo-600 to-purple-600">
              가장 빛나는 방식
            </span>
            으로.
          </h1>
          <p className="max-w-3xl mx-auto text-lg lg:text-xl text-slate-600 leading-relaxed">
            복잡한 서식 고민 없이, 데이터만 입력하세요.
          </p>
          <p className="max-w-3xl mx-auto text-lg lg:text-xl text-slate-600 mb-12 leading-relaxed">
            기업이 선호하는 현대적이고 전문적인 이력서 템플릿을 실시간으로 제공합니다.
          </p>
          <div className="flex flex-col sm:flex-row items-center justify-center gap-4">
            <Link to={'/resume'}>
              <Button className="font-bold text-lg p-8 has-[>svg]:px-8">
                내 이력서 보기{' '}
                <ArrowRight className="group-hover:translate-x-1 transition-transform" />
              </Button>
            </Link>
            <Button variant={'outline'} className="font-bold text-lg p-8 has-[>svg]:px-8">
              샘플 템플릿 탐색
            </Button>
          </div>
        </div>

        <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-200 h-200 bg-indigo-100/30 rounded-full blur-3xl z-0"></div>
      </section>
    </main>
  )
}

export default HomePage
