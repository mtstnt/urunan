import { Toaster } from "react-hot-toast";
import LoadingOverlayWrapper from "react-loading-overlay-ts";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Login from "./pages/auth/Login";
import JoinBill from "./pages/bills/JoinBill";
import Profile from "./pages/Profile";
import { useLoadingStore } from "./stores/loading";
import { Box, ChakraProvider, Spinner } from "@chakra-ui/react";

import "./styles/index.css";
import { CustomErrorBoundary } from "./components/ErrorBoundary";

const router = createBrowserRouter([
  { path: "/", element: <Profile /> },
  { path: "/create", element: <h1>Create Bill Form</h1> },
  {
    path: "/:code",
    element: <JoinBill />,
    children: [
      { path: "nickname", element: <h1>Nickname Select</h1> },
      { path: "orders", element: <h1>Choose Orders!</h1> },
    ],
  },
  {
    path: "/auth",
    children: [
      { path: "signin", element: <Login /> },
      { path: "signup", element: <h1>Register</h1> },
    ],
  },
]);

export default function App() {
  const { isLoading } = useLoadingStore();
  return (
    <ChakraProvider>
      <Box bg={"#eee"}>
        <Box
          bg={"white"}
          width={{
            base: "100%",
            sm: "100%",
            md: "80%",
            lg: "40%",
            xl: "30%",
          }}
          minH={"100vh"}
          margin={"0 auto"}
        >
          <CustomErrorBoundary>
            <LoadingOverlayWrapper
              active={isLoading}
              spinner={<Spinner size={"xl"} />}
            >
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
          </CustomErrorBoundary>
        </Box>
      </Box>
    </ChakraProvider>
  );
}
