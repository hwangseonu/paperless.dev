export interface Resume {
  id: string
  title: string
  description: string
  email: string
  url: string
  image: string
  public: boolean
  template: string
  skills: string[]
  createdAt: string // ISO 8601 Date string
  updatedAt: string
  educations: Education[]
  experiences: Experience[]
  projects: Project[]
}

export interface Education {
  id: string
  school: string
  major: string
  degree: string
  startDate: string
  endDate: string
  gpa?: string
  activities?: string
}

export interface Experience {
  id: string
  company: string
  title: string
  location: string
  startDate: string
  endDate: string
  description: string
}

export interface Project {
  id: string
  title: string
  startDate: string
  endDate: string
  description: string
  skills: string[]
  url?: string
}
