
import { useNavigate } from "react-router-dom"
import { motion } from "framer-motion"

import notFound from "../assets/notFound.svg"

// not required remove in next build
const NotFound404 = () => {
    const navigate = useNavigate()
    
    return (
        <div className="flex flex-col justify-center items-center w-screen h-screen">
            <img src={notFound} alt="" className="w-96"/>
            <motion.button
                transition={{
                    ease: "linear",
                    duration: 0.5
                }}
                className="rounded-2xl relative border border-black p-4 mt-12"
                onClick={() => navigate("/app")}
                whileTap={{scale: 0.8}}
            >
                Go Back Home
            </motion.button>
        </div>
    )
}

export default NotFound404
