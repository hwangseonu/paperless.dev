import './App.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import ResumePage from '@/pages/Resume.tsx'

import { MOCK_RESUME_DATA as resume } from '@/assets/mock.ts'
import HomePage from '@/pages/HomePage.tsx'
import Layout from '@/components/Layout.tsx'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      { index: true, element: <HomePage /> },
      {
        path: 'resume',
        element: <ResumePage />,
        // 여기서 데이터를 미리 가져올 수 있어요!
        loader: async () => resume.resume,
      },
    ],
  },
])

function App() {
  return <RouterProvider router={router}></RouterProvider>
}

export default App
