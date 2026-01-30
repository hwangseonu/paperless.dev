import { Code, ExternalLink } from 'lucide-react'
import ResumeSection from '@/components/resume/ResumeSection.tsx'
import type { Project } from '@/utils/types.ts'
import { Card, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card.tsx'
import { Badge } from '@/components/ui/badge.tsx'

type props = {
  projects: Project[]
}

function Projects({ projects }: props) {
  return (
    <ResumeSection icon={Code} title="Projects">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {projects.map((project) => (
          <Card
            key={project.id}
            className="group p-4 hover:border-indigo-300 transition-all hover:shadow-md "
          >
            <CardHeader className="flex justify-between items-start p-0 mt-2">
              <CardTitle className="text-slate-900 group-hover:text-indigo-600 transition-colors">
                {project.title}
              </CardTitle>
              {project.url && (
                <a
                  href={project.url}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-slate-400 hover:text-indigo-600"
                >
                  <ExternalLink size={16} />
                </a>
              )}
            </CardHeader>
            <CardDescription className="text-xs text-slate-500">
              {project.description}
            </CardDescription>
            <CardFooter className="flex flex-wrap gap-1.5 p-0">
              {project.skills.map((s, i) => (
                <Badge
                  key={i}
                  className="bg-slate-100 text-slate-500 rounded text-[10px] font-bold"
                >
                  {s}
                </Badge>
              ))}
            </CardFooter>
          </Card>
        ))}
      </div>
    </ResumeSection>
  )
}

export default Projects
