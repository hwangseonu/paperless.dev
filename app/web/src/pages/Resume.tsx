import { useLoaderData } from 'react-router-dom'
import type { Resume } from '@/utils/types.ts'
import Information from '@/components/resume/Information.tsx'
import Skills from '@/components/resume/Skills.tsx'
import Career from '@/components/resume/Career.tsx'
import Projects from '@/components/resume/Projects.tsx'
import Educations from '@/components/resume/Educations.tsx'
import { Card } from '@/components/ui/card.tsx'

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

        <Card className="lg:col-span-8 p-8 lg:p-12">
          <Career experiences={data.experiences} />
          <Projects projects={data.projects} />
          <Educations educations={data.educations} />
        </Card>
      </div>
    </main>
  )
}

export default ResumePage
