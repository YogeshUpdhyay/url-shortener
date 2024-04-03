import { useState } from "react"
import illustration  from "../assets/illustration.svg"
import shortenLink from "../assets/shortenLink.svg"
import { motion, AnimatePresence } from "framer-motion"

const HeroSection = () => { 
    const [url, setUrl] = useState()
    const [isShortLinkOpen, setIsShortLinkOpen] = useState(false)
    const [shortUrl, setShortUrl] = useState()

    
    const onSubmit = () => {
        console.log(url)

        const options = {
            body: JSON.stringify({url: url}),
            headers: {'Content-Type': 'aplication/json'},
            method: 'POST'
        }
        
        fetch("http://localhost:8080/api/v1/shortenurl", options)
            .then(response => response.json())
            .then(data => setShortUrl(data.url))
            .catch(error => console.log(error))

        setIsShortLinkOpen(!isShortLinkOpen)
    }

    const copyToClipboard = () => {
        console.log("Copy to clipboard")
        navigator.clipboard.writeText(shortUrl).then(function() {
            console.log("Text copied to clipboard successfully");
        }).catch(function(error) {
            console.error("Unable to copy text: ", error);
        });
    }

    return (
        <div className="grow">
            {/* hero section left side */}
            <div className="flex flex-col lg:flex-row items-center justify-center w-full p-4 mx-auto lg:w-3/4 h-full">
                <div className="flex flex-col gap-4 p-4 w-full lg:w-2/4">
                    <div className="rounded-full bg-blue bg-opacity-25 text-black py-2 px-4 w-fit text-xs">
                        Easy Shortening URL
                    </div>
                    <div className="text-3xl font-bold">
                        linksy short URL generator.
                    </div>
                    <p className="text-grey">One shortlink, inifinite possibilities.</p>
                    <div className="rounded-full flex gap-2 p-2 border border-slate-200 h-14 w-full lg:w-3/4">
                        <input 
                            type="text" 
                            className="rounded-full grow focus:outline-none active:outline-none flex ps-3" 
                            placeholder="https://linksy.com" 
                            onChange={(event) => setUrl(event.target.value)}
                        />
                        <motion.button 
                            className="rounded-full bg-blue text-white text-sm px-4 min-w-24" 
                            onClick={onSubmit}
                            whileTap={{
                                scale: 0.5,
                                transition: { duration: 0.5 },
                            }}
                        >
                            { !isShortLinkOpen ? "Shorten" : "Try Another!" }
                        </motion.button>
                    </div>

                    <AnimatePresence>
                        {isShortLinkOpen && <motion.button
                                initial={{opacity: 0, scale: 0}}
                                animate={{opacity: 1, scale: 1}}
                                exit={{opacity: 0, scale: 0}}
                                // exit={{y:-200, opacity: 0}}
                                transition={{
                                    ease: "linear",
                                    duration: 0.5
                                }}
                                className="rounded-2xl relative border border-black p-4 mt-12"
                                onClick={copyToClipboard}
                                whileTap={{scale: 0.8}}
                            >
                                <img src={shortenLink} alt="" className="absolute -left-10 -top-16 h-32 w-32"/>
                                <input 
                                    type="text" 
                                    className="ps-14 h-10 rounded-xl focus:outline-none active:outline-none w-full color-inherit" 
                                    placeholder="https://linksy.com/TYubiug8"
                                    disabled={true}
                                    value={shortUrl}
                                />
                                <label htmlFor="" className="text-xs text-red">Click To Copy!</label>
                        </motion.button>}
                    </AnimatePresence>
            
                </div>

                {/* hero section right side */}
                <div className="flex w-full lg:w-2/4 lg:block hidden">
                    <img src={illustration} alt="" className="h-4/5 w-4/5 mx-auto my-auto"/>
                </div>
            </div>

        </div>
        
    )
}

export default HeroSection
