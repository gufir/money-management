import { reactive, readonly } from 'vue'
import type { AuthState } from './types/auth_state'
import type { User } from './types/user'
import Cookies from 'js-cookie'

const state = reactive<AuthState>({
  user: JSON.parse(Cookies.get('user') || 'null'),
  accessToken: Cookies.get('accessToken') || null,
  refreshToken: Cookies.get('refreshToken') || null,
})

function setUser(user: User, accessToken: string, refreshToken: string) {
  state.user = user
  state.accessToken = accessToken
  state.refreshToken = refreshToken

  Cookies.set('user', JSON.stringify(user), { secure: true, sameSite: 'Strict' })
  Cookies.set('accessToken', accessToken, { secure: true, sameSite: 'Strict' })
  Cookies.set('refreshToken', refreshToken, { secure: true, sameSite: 'Strict' })
}

function clearUser() {
  Cookies.remove('user')
  Cookies.remove('accessToken')
  Cookies.remove('refreshToken')

  state.user = null
  state.accessToken = null
  state.refreshToken = null
}

export default {
  state: readonly(state),
  setUser,
  clearUser,
}
