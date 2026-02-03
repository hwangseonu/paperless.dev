import { type UserInfo } from '@/lib/types'
import { Avatar as AvatarWrapper, AvatarImage } from '@/components/ui/avatar'
import profile from '@/assets/profile.jpg'

type props = {
  user: UserInfo
}

function Avatar({ user }: props) {
  return (
    <div className={'flex flex-row items-center gap-2'}>
      <AvatarWrapper>
        <AvatarImage src={profile} />
      </AvatarWrapper>
      <span className={'text-base font-semibold'}>{user.nickname} 님</span>
    </div>
  )
}

export default Avatar
