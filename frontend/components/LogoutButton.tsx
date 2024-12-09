import { useContext } from 'react'
import { useRouter } from 'next/router'
import { AuthContext } from '../modules/auth_provider'
import { API_URL } from '../constants'

const LogoutButton = () => {
  const router = useRouter()
  const { setAuthenticated } = useContext(AuthContext)

  const handleLogout = async () => {
    try {
      const res = await fetch(`${API_URL}/logout`, {
        method: 'GET',
        credentials: 'include'
      })

      if (res.ok) {
        localStorage.removeItem('user_info')
        setAuthenticated(false)
        router.push('/login')
      }
    } catch (err) {
      console.log(err)
    }
  }

  return (
    <button
      onClick={handleLogout}
      className="px-4 py-2 text-white bg-red rounded-md hover:bg-red"
    >
      Logout
    </button>
  )
}

export default LogoutButton
