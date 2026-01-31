import ResumeSection from '@/components/resume/ResumeSection.tsx'
import { Award, GraduationCap } from 'lucide-react'
import type { Education } from '@/lib/types.ts'
import { Badge } from '@/components/ui/badge.tsx'

type props = {
  educations: Education[]
}

function Educations({ educations }: props) {
  return (
    <ResumeSection icon={GraduationCap} title="Education">
      {educations.map((edu) => (
        <div key={edu.id} className="flex gap-6">
          <div className="flex-1">
            <div className="flex flex-wrap justify-between items-start mb-2 gap-2">
              <div>
                <h3 className="text-lg font-bold text-slate-900">{edu.school}</h3>
                <p className="text-slate-600 text-sm">
                  {edu.major} {edu.major && edu.degree ? '·' : ''} {edu.degree}
                </p>
              </div>
              <Badge className="text-indigo-600 bg-indigo-50 px-3 py-1 ">
                {edu.startDate} - {edu.endDate}
              </Badge>
            </div>
            {edu.gpa && (
              <div className="flex items-center gap-2 text-sm mb-3">
                <Award size={14} className="text-amber-500" />
                <span className="font-medium text-slate-700">GPA {edu.gpa}</span>
              </div>
            )}
            {edu.activities && (
              <div className="bg-slate-50 p-3 rounded-lg">
                <p className="text-xs text-slate-500 leading-relaxed italic">"{edu.activities}"</p>
              </div>
            )}
          </div>
        </div>
      ))}
    </ResumeSection>
  )
}

export default Educations
