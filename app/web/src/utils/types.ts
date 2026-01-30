/** * 이력서 데이터의 최상위 루트 인터페이스
 */
export interface ResumeData {
  resume: Resume
}

/** * 이력서 상세 정보
 */
export interface Resume {
  id: string
  title: string
  bio: string
  public: boolean
  template: string
  skills: string[]
  createdAt: string // ISO 8601 Date string
  updatedAt: string
  educations: Education[]
  experiences: Experience[]
  projects: Project[]
}

/** * 학력 사항
 */
export interface Education {
  id: string
  school: string
  major: string
  degree: string
  startDate: string // YYYY-MM
  endDate: string
  gpa?: string // 선택적 필드
  activities?: string
}

/** * 경력 사항
 */
export interface Experience {
  id: string
  company: string
  title: string
  location: string
  startDate: string
  endDate: string // '현재' 또는 YYYY-MM
  description: string
}

/** * 프로젝트 경험
 */
export interface Project {
  id: string
  title: string
  startDate: string
  endDate: string
  description: string
  skills: string[]
  url?: string
}
