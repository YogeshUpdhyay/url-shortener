import logo from "../assets/logo.png"

const Navbar = () => {
    return (
        <nav className="flex justify-between items-center mt-2 mx-auto w-full lg:w-3/4">
            <img src={logo} alt="" className="flex h-20"/>
            {/* <button className="flex rounded-full bg-red text-white px-4 py-1">Get Started</button> */}
        </nav>
    )
}

export default Navbar
