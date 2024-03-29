import { useState } from "react"
import illustration  from "../assets/illustration.svg"
import { motion, AnimatePresence } from "framer-motion"

const HeroSection = () => { 
    const [url, setUrl] = useState()
    const [isShortLinkOpen, setIsShortLinkOpen] = useState(false)
    const onSubmit = () => { 
        console.log(url)
        setIsShortLinkOpen(!isShortLinkOpen)
    }

    return (
        <div className="grow flex container mx-auto items-center w-3/4">
            {/* hero section left side */}
            <div className="flex flex-col gap-4 p-4 w-2/4">
                <div className="rounded-full bg-blue bg-opacity-25 text-black py-2 px-4 w-fit text-xs">
                    Easy Shortening URL
                </div>
                <div className="text-3xl font-bold">
                    linksy short URL generator.
                </div>
                <p className="text-grey">One shortlink, inifinite possibilities.</p>
                <div className="rounded-full flex gap-2 p-2 border border-slate-200 h-14 w-3/4">
                    <input type="text" className="rounded-full grow focus:outline-none active:outline-none flex ps-3" placeholder="https://linksy.com" onChange={(event) => setUrl(event.target.value)}/>
                    <motion.button 
                        className="rounded-full bg-blue text-white text-sm px-4 min-w-24" 
                        onClick={onSubmit}
                        whileTap={{ scale: 0.9 }}
                    >
                        Try Another!
                    </motion.button>
                </div>
                <AnimatePresence>
                    {isShortLinkOpen && <motion.div
                            initial={{y: -200, opacity: 0}}
                            animate={{y: 0, opacity: 1}}
                            transition={{
                                ease: "linear",
                                duration: 2,
                                y: { duration: 1 }
                            }}
                            className="rounded-2xl bg-blue bg-opacity-25 h-24 p-4"
                        >
                            This is shortlink section
                        </motion.div>}
                </AnimatePresence>
            </div>

            {/* hero section right side */}
            <div className="flex w-2/4">
                <img src={illustration} alt="" className="h-4/5 w-4/5 mx-auto my-auto"/>
            </div>
        </div>
    )
}

export default HeroSection
