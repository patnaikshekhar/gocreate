import { Button } from "@/components/ui/button"
import './App.css'

function App() {

  const fetchData = async () => {
    const response = await fetch("/api/v1/healthz")
    const data = await response.json()
    console.log(data)
  }

  return (
    <>
      <Button onClick={() => fetchData()}>Click me</Button>
    </>
  )
}

export default App
