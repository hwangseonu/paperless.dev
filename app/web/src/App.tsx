import './App.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import ResumePage from '@/pages/ResumePage.tsx'

import { MOCK_RESUME_DATA as resume } from '@/assets/mock.ts'
import HomePage from '@/pages/HomePage.tsx'
import Layout from '@/components/common/Layout.tsx'
import LoginPage from '@/pages/LoginPage.tsx'

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
    ],
  },
])

function App() {
  return <RouterProvider router={router}></RouterProvider>
}

export default App
