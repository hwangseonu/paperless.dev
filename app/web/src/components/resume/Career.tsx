import { Briefcase, Calendar, MapPin } from 'lucide-react'
import ResumeSection from '@/components/resume/ResumeSection.tsx'
import type { Experience } from '@/utils/types.ts'
import { Badge } from '@/components/ui/badge.tsx'

type props = {
  experiences: Experience[]
}

function CareerItem(props: { experience: Experience }) {
  return (
    <div className="relative pl-6 before:content-[''] before:absolute before:left-0 before:top-1.5 before:bottom-0 before:w-0.5 before:bg-slate-100">
      <div className="absolute -left-1 top-1.5 w-2 h-2 rounded-full bg-indigo-500"></div>
      <div className="flex flex-wrap justify-between items-start mb-2 gap-2">
        <div>
          <h3 className="text-lg font-bold text-slate-900">{props.experience.title}</h3>
          <p className="text-indigo-600 font-medium text-sm">{props.experience.company}</p>
        </div>
        <Badge variant={'secondary'} className={'text-slate-400'}>
          <Calendar size={12} />
          {props.experience.startDate} - {props.experience.endDate}
        </Badge>
      </div>
      {props.experience.location && (
        <div className="flex items-center gap-1 text-slate-500 text-xs mb-3">
          <MapPin size={12} />
          <span>{props.experience.location}</span>
        </div>
      )}
      <p className="text-slate-600 text-sm leading-relaxed whitespace-pre-wrap">
        {props.experience.description}
      </p>
    </div>
  )
}

function Career({ experiences }: props) {
  return (
    <ResumeSection icon={Briefcase} title="Career">
      {experiences.map((exp) => (
        <CareerItem experience={exp} />
      ))}
    </ResumeSection>
  )
}

export default Career
