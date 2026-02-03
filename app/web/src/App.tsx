import './App.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import ResumePage from '@/pages/ResumePage.tsx'

import { MOCK_RESUME_DATA as resume } from '@/assets/mock.ts'
import HomePage from '@/pages/HomePage.tsx'
import Layout from '@/components/common/Layout.tsx'
import LoginPage from '@/pages/LoginPage.tsx'
import RegisterPage from '@/pages/RegisterPage.tsx'
import VerifyCodePage from '@/pages/VerifyCodePage.tsx'
import { AuthProvider } from '@/context/AuthContext.tsx'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      { index: true, element: <HomePage /> },
      {
        path: 'resume',
        element: <ResumePage />,
        loader: async () => resume.resume,
      },
      {
        path: 'login',
        element: <LoginPage />,
      },
      {
        path: 'register',
        element: <RegisterPage />,
      },
      {
        path: 'verify',
        element: <VerifyCodePage />,
      },
    ],
  },
])

function App() {
  return (
    <AuthProvider>
      <RouterProvider router={router}></RouterProvider>
    </AuthProvider>
  )
}

export default App
