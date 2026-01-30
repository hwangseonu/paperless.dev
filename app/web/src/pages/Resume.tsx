import { useLoaderData } from 'react-router-dom'
import { Award, Briefcase, Calendar, Code, ExternalLink, GraduationCap, MapPin } from 'lucide-react'
import ResumeSection from '@/components/resume/ResumeSection.tsx'
import type { Resume } from '@/utils/types.ts'
import Information from '@/components/resume/Information.tsx'
import Skills from '@/components/resume/Skills.tsx'

function ResumePage() {
  const data: Resume = useLoaderData()

  return (
    <main className="w-full mx-auto py-12 px-6 lg:px-8">
      <div className="grid grid-cols-1 lg:grid-cols-12 gap-12">
        <aside className="lg:col-span-4 space-y-8 animate-fade-in-up">
          <Information
            title={data.title}
            bio={data.bio}
            email={data.email}
            url={data.url}
            image={data.image}
          />
          <Skills skills={data.skills} />
        </aside>

        <div
          className="lg:col-span-8 bg-white p-8 lg:p-12 rounded-2xl shadow-sm border border-slate-100 animate-fade-in-up"
          style={{ animationDelay: '100ms' }}
        >
          <ResumeSection icon={Briefcase} title="Work Experience">
            {data.experiences.map((exp) => (
              <div
                key={exp.id}
                className="relative pl-6 before:content-[''] before:absolute before:left-0 before:top-1.5 before:bottom-0 before:w-0.5 before:bg-slate-100"
              >
                <div className="absolute -left-1 top-1.5 w-2 h-2 rounded-full bg-indigo-500"></div>
                <div className="flex flex-wrap justify-between items-start mb-2 gap-2">
                  <div>
                    <h3 className="text-lg font-bold text-slate-900">{exp.title}</h3>
                    <p className="text-indigo-600 font-medium text-sm">{exp.company}</p>
                  </div>
                  <div className="flex items-center gap-2 text-slate-400 text-xs font-medium bg-slate-50 px-3 py-1 rounded-full">
                    <Calendar size={12} />
                    <span>
                      {exp.startDate} - {exp.endDate}
                    </span>
                  </div>
                </div>
                <div className="flex items-center gap-1 text-slate-500 text-xs mb-3">
                  <MapPin size={12} />
                  <span>{exp.location}</span>
                </div>
                <p className="text-slate-600 text-sm leading-relaxed whitespace-pre-wrap">
                  {exp.description}
                </p>
              </div>
            ))}
          </ResumeSection>

          <ResumeSection icon={Code} title="Projects">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {data.projects.map((proj) => (
                <div
                  key={proj.id}
                  className="group p-5 rounded-xl border border-slate-200 hover:border-indigo-300 transition-all hover:shadow-md bg-white"
                >
                  <div className="flex justify-between items-start mb-3">
                    <h3 className="font-bold text-slate-900 group-hover:text-indigo-600 transition-colors">
                      {proj.title}
                    </h3>
                    <a
                      href={proj.url}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-slate-400 hover:text-indigo-600"
                    >
                      <ExternalLink size={16} />
                    </a>
                  </div>
                  <p className="text-xs text-slate-500 mb-4 line-clamp-2 leading-relaxed">
                    {proj.description}
                  </p>
                  <div className="flex flex-wrap gap-1.5">
                    {proj.skills.map((s, i) => (
                      <span
                        key={i}
                        className="px-2 py-0.5 bg-slate-100 text-slate-500 rounded text-[10px] font-bold"
                      >
                        {s}
                      </span>
                    ))}
                  </div>
                </div>
              ))}
            </div>
          </ResumeSection>

          <ResumeSection icon={GraduationCap} title="Education">
            {data.educations.map((edu) => (
              <div key={edu.id} className="flex gap-6">
                <div className="flex-1">
                  <div className="flex flex-wrap justify-between items-start mb-2 gap-2">
                    <div>
                      <h3 className="text-lg font-bold text-slate-900">{edu.school}</h3>
                      <p className="text-slate-600 text-sm">
                        {edu.major} · {edu.degree}
                      </p>
                    </div>
                    <span className="text-xs font-bold text-indigo-600 bg-indigo-50 px-3 py-1 rounded-full">
                      {edu.startDate} - {edu.endDate}
                    </span>
                  </div>
                  {edu.gpa && (
                    <div className="flex items-center gap-2 text-sm mb-3">
                      <Award size={14} className="text-amber-500" />
                      <span className="font-medium text-slate-700">GPA {edu.gpa}</span>
                    </div>
                  )}
                  {edu.activities && (
                    <div className="bg-slate-50 p-3 rounded-lg">
                      <p className="text-xs text-slate-500 leading-relaxed italic">
                        "{edu.activities}"
                      </p>
                    </div>
                  )}
                </div>
              </div>
            ))}
          </ResumeSection>
        </div>
      </div>
    </main>
  )
}

export default ResumePage
