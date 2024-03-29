import HeroSection from "./components/Hero"
import Navbar from "./components/Navbar"



function App() {
  return (
    <div className="flex flex-col w-screen h-screen ">
      <Navbar />
      <HeroSection />
    </div>  
  )
}

export default App
