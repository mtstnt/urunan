import { FieldErrors, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { storeTokenInStorage } from "../../utils/token";
import {
  Box,
  Button,
  Flex,
  FormControl,
  FormLabel,
  Heading,
  Input,
} from "@chakra-ui/react";
// import { baseURL } from "../../constants";
import { useLoadingStore } from "../../stores/loading";
import toast from "react-hot-toast";

type LoginForm = {
  email: string;
  password: string;
};

export default function Login() {
  const { register, handleSubmit } = useForm<LoginForm>();
  const { setIsLoading } = useLoadingStore();
  const navigate = useNavigate();

  const onSubmit = async (data: LoginForm) => {
    console.log(data);
    try {
      setIsLoading(true);
      // // TODO: Temporary bypass for development
      // const response = await fetch(baseURL("/signin"), {
      //   method: "POST",
      //   body: JSON.stringify({
      //     email: data.email,
      //     password: data.password,
      //   }),
      // });
      // if (!response.ok) {
      //   alert("Error " + response.status + ": " + response.statusText);
      //   return;
      // }
      // console.log(response);
      storeTokenInStorage("asdfg");
      navigate("/");
    } catch (e) {
      toast.error("Error: " + e);
    } finally {
      setIsLoading(false);
    }
  };

  const onInvalid = (errors: FieldErrors) => {
    console.log(errors.password);
  };

  return (
    <Flex
      as={"form"}
      direction={"column"}
      justifyContent={"center"}
      alignItems={"center"}
      height={"100vh"}
      onSubmit={handleSubmit(onSubmit, onInvalid)}
    >
      <Box w={"80%"} border={"1px solid black"} p={5} shadow={"5px 5px #400"}>
        <Heading mb={5} size={"xl"}>
          Login
        </Heading>

        {/* {hasFlashMessage("error") ? (
          <Alert mb={3} status="error">
            <AlertIcon />
            <AlertTitle>Error</AlertTitle>
            <AlertDescription>{readFlashMessage("error")}</AlertDescription>
          </Alert>
        ) : (
          <></>
        )} */}

        <FormControl mb={2}>
          <FormLabel>Email</FormLabel>
          <Input
            type="email"
            placeholder="Email"
            {...register("email", { required: true })}
          />
        </FormControl>

        <FormControl mb={5}>
          <FormLabel>Password</FormLabel>
          <Input
            type="password"
            placeholder="Password"
            {...register("password", { required: true })}
          />
        </FormControl>

        <Button w={"100%"} type="submit" colorScheme="green">
          Submit
        </Button>
      </Box>
    </Flex>
  );
}
