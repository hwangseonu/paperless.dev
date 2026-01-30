import { Badge } from '@/components/ui/badge.tsx'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card.tsx'

type props = {
  skills: string[]
}

function Skills({ skills }: props) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>보유 기술</CardTitle>
      </CardHeader>
      <CardContent className={'flex flex-wrap gap-2'}>
        {skills.map((skill, index) => (
          <Badge key={index}>{skill}</Badge>
        ))}
      </CardContent>
    </Card>
  )
}

export default Skills
