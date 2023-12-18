import { useNavigate } from "react-router-dom";
import { useEffect } from "react";
import { getTokenFromStorage } from "../utils/token";
import { useLoadingStore } from "../stores/loading";

export default function Profile() {
    const { setIsLoading } = useLoadingStore();
    const navigate = useNavigate();

    useEffect(() => {
        const token = getTokenFromStorage();
        if (token == null) {
            navigate("/auth/signin");
            return;
        }
        setIsLoading(false);
    }, [navigate, setIsLoading]);

    return (
        <div>
            <h1>Hello world</h1>
            <p>{localStorage.getItem('token')}</p>
        </div>
    )
}