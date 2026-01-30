import { type LucideIcon } from 'lucide-react'
import type { PropsWithChildren } from 'react'

type props = PropsWithChildren & {
  icon: LucideIcon
  title: string
}

const ResumeSection = ({ icon: Icon, title, children }: props) => (
  <div className="mb-12">
    <div className="flex items-center gap-3 mb-6 border-b border-slate-200 pb-2">
      <div className="p-2 bg-indigo-50 rounded-lg text-indigo-600">
        <Icon size={22} />
      </div>
      <h2 className="text-xl font-bold text-slate-800">{title}</h2>
    </div>
    <div className="space-y-8">{children}</div>
  </div>
)

export default ResumeSection
