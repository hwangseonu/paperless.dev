import { useLoaderData } from 'react-router-dom'
import {
  Award,
  Briefcase,
  Calendar,
  Code,
  ExternalLink,
  Globe,
  GraduationCap,
  Mail,
  MapPin,
} from 'lucide-react'
import ResumeSection from '@/components/ResumeSection.tsx'
import type { Resume } from '@/utils/types.ts'

function ResumePage() {
  const data: Resume = useLoaderData()

  return (
    <main className="w-full mx-auto py-12 px-6 lg:px-8">
      <div className="grid grid-cols-1 lg:grid-cols-12 gap-12">
        {/* Left Sidebar - Bio & Info */}
        <aside className="lg:col-span-4 space-y-8 animate-fade-in-up">
          <div className="bg-white p-8 rounded-2xl shadow-sm border border-slate-100">
            <div className="w-24 h-24 bg-slate-200 rounded-2xl mb-6 overflow-hidden">
              <div className="w-full h-full flex items-center justify-center text-slate-400 text-3xl font-bold bg-slate-100 uppercase">
                {data.title.substring(0, 2)}
              </div>
            </div>
            <h1 className="text-2xl font-extrabold text-slate-900 leading-tight mb-4">
              {data.title}
            </h1>
            <p className="text-slate-600 text-sm leading-relaxed mb-6">{data.bio}</p>

            <div className="space-y-3 pt-6 border-t border-slate-100">
              <div className="flex items-center gap-3 text-sm text-slate-500">
                <Mail size={16} />
                <span>minjun.dev@example.com</span>
              </div>
              <div className="flex items-center gap-3 text-sm text-slate-500">
                <Globe size={16} />
                <span className="text-indigo-600 hover:underline cursor-pointer">portfolio.me</span>
              </div>
            </div>
          </div>

          <div className="bg-white p-8 rounded-2xl shadow-sm border border-slate-100">
            <h3 className="text-sm font-bold text-slate-400 uppercase tracking-widest mb-6">
              핵심 보유 기술
            </h3>
            <div className="flex flex-wrap gap-2">
              {data.skills.map((skill, index) => (
                <span
                  key={index}
                  className="px-3 py-1.5 bg-slate-50 text-slate-700 rounded-lg text-xs font-semibold border border-slate-100"
                >
                  {skill}
                </span>
              ))}
            </div>
          </div>

          <div className="p-8 rounded-2xl bg-indigo-600 text-white">
            <h3 className="font-bold mb-2">업데이트 정보</h3>
            <p className="text-indigo-100 text-xs">
              최종 수정: {new Date(data.updatedAt).toLocaleDateString()}
            </p>
          </div>
        </aside>

        {/* Right Content - Sections */}
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
