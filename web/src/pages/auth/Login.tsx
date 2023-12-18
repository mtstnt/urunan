import { FieldErrors, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { storeTokenInStorage } from "../../utils/token";

type LoginForm = {
  email: string;
  password: string;
};

export default function Login() {
  const {
    register,
    handleSubmit,
    watch,
  } = useForm<LoginForm>();

  const navigate = useNavigate();

  console.log(watch("email"));
  const onSubmit = (data: LoginForm) => {
    console.log(data);
    storeTokenInStorage("asdfg");
    navigate("/");
  };
  const onInvalid = (errors: FieldErrors) => console.log(errors);

  return (
    <form onSubmit={handleSubmit(onSubmit, onInvalid)}>
      <h1>Login</h1>

      <input
        type="email"
        placeholder="Email"
        {...register("email", { required: true })}
      />

      <br />

      <input
        type="password"
        placeholder="Password"
        {...register("password", { required: true })}
      />

      <br />

      <button type="submit">Submit</button>
    </form>
  );
}
