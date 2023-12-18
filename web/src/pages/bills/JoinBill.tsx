import { useEffect } from "react";
import { useParams } from "react-router-dom";

type Params = {
    code: string,
};

export default function JoinBill() {
    const params = useParams<Params>();
    console.log(params.code);

    useEffect(() => {
        // Check if bill for that code exists.

        // If it does not exist, redirect to 404.

        // If it does exist, show the Nickname form page.
    }, []);

    return <h1>Join Bill Form</h1>
}