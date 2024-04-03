import { useEffect, useState } from "react"
import { useNavigate, useParams } from "react-router-dom"
import { motion } from "framer-motion"

import notFound from "../assets/notFound.svg"

// not required remove in next build
const LinkHandler = () => {
    const navigate = useNavigate()
    let { linkId } = useParams()
    const [redirectUrl, setRedirectUrl] = useState(null);
    const [isErrorOpen, setIsErrorOpen] = useState(false)

    useEffect(() => {
        fetch(`http://localhost:8080/api/v1/${linkId}`, {redirect: 'follow'})
            .then(response => {
                if (response.status === 301) {
                    // If the response is a 301 redirect, get the redirected URL
                    const redirectUrl = response.headers.get('Location');
                    setRedirectUrl(redirectUrl);
                }

                if (response.status == 404) {
                    setIsErrorOpen(true)
                }
            })
            .catch(error => {
                console.log(error)
                console.log('Error')
                setIsErrorOpen(true)
            })
    }, [])

    // Redirect to the redirect URL if available
    useEffect(() => {
        if (redirectUrl) {
        window.location.href = redirectUrl; // Perform the redirect
        }
    }, [redirectUrl]);
    
    return (
        <>
        {isErrorOpen && <div className="flex flex-col justify-center items-center w-screen h-screen">
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
        }
        </>
    )
}

export default LinkHandler
