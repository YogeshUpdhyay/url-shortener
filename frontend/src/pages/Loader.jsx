import { useEffect, useState } from "react";
import { useNavigate } from 'react-router-dom';
import logoLarge from "../assets/logoLarge.png"
import { Bars } from "react-loader-spinner"

const Loader = () => {
    const navigate = useNavigate();
    const [redirect, setRedirect] = useState(false);

    useEffect(() => {
        const timer = setTimeout(() => setRedirect(true), 1000);
        return () => clearTimeout(timer)
    }, [])

    // Redirect to the destination screen after 1 second
    useEffect(() => {
        if (redirect) {
            navigate("/app");
        }
    }, [redirect, navigate]);
    
    return <div className="flex flex-col gap-8 justify-center items-center h-screen w-screen">
        <img src={logoLarge} alt="" className="w-36 h-36"/>
        <Bars
            height="35"
            width="75"
            color="#2161E0"
            ariaLabel="bars-loading"
            wrapperStyle={{}}
            wrapperClass=""
            visible={true}
        />
    </div>
}

export default Loader
