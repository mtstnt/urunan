import { Toaster } from "react-hot-toast";
import LoadingOverlayWrapper from "react-loading-overlay-ts";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Login from "./pages/auth/Login";
import JoinBill from "./pages/bills/JoinBill";
import Profile from "./pages/Profile";
import { useLoadingStore } from "./stores/loading";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Profile />,
  },
  {
    path: "/:code",
    element: <JoinBill />,
  },
  {
    path: "/auth",
    children: [
      {
        path: "signin",
        element: <Login />,
      },
      {
        path: "signup",
        element: <h1>Register</h1>,
      },
    ],
  },
]);

export default function App() {
  const { isLoading } = useLoadingStore();
  return (
    <LoadingOverlayWrapper active={isLoading}>
      <Toaster
        toastOptions={{
          style: {
            border: "1px solid #713200",
            padding: "16px",
            color: "#713200",
          },
        }}
      />
      <RouterProvider router={router} />
    </LoadingOverlayWrapper>
  );
}
