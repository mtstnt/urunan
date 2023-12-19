import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { getTokenFromStorage } from "../utils/token";
import {
  Box,
  Button,
  Card,
  CardBody,
  Flex,
  Grid,
  GridItem,
  Heading,
  useDisclosure,
} from "@chakra-ui/react";
import toast from "react-hot-toast";
import { useLoadingStore } from "../stores/loading";
import { defaultShadow } from "../constants";
import JoinBillModal from "../components/JoinBillModal";
import { BillCard } from "../components/BillCard";

type Bill = {
  id: number;
  title: string;
  description: string;
  date: Date;
  isHost: boolean;
  isLocked: boolean;
};

type UserHomeData = {
  email: string;
  username: string;

  totalUnpaid: number;
  bills: Bill[];
};

const dummyBills = [
  {
    id: 1,
    title: "Makan Bareng di Masbro",
    description: "Transfer ke ...",
    date: new Date(2023, 12, 18),
    isHost: false,
    isLocked: false,
  },
  {
    id: 2,
    title: "Beli Minum Wedrink",
    description: "Transfer ke ...",
    date: new Date(2023, 12, 17),
    isHost: false,
    isLocked: false,
  },
];

const dummyUser: UserHomeData = {
  email: "dummy@email.com",
  username: "Dummy",
  totalUnpaid: 10_000_000,
  bills: (function () {
    let b: Bill[] = [];
    for (let i = 0; i < 100; i++) {
      b = [...b, ...dummyBills];
    }
    return b;
  })(),
};

export default function Profile() {
  const [user, setUser] = useState<UserHomeData | null>();
  const { setIsLoading } = useLoadingStore();
  const navigate = useNavigate();

  useEffect(() => {
    setIsLoading(true);
    const token = getTokenFromStorage();
    if (token == null) {
      setIsLoading(false);
      navigate("/auth/signin");
      toast.error("Must be authenticated to visit this page.");
      return;
    }
    // TODO: Dummy user.
    setUser(dummyUser);
    setIsLoading(false);
  }, [navigate, setUser, setIsLoading]);

  const handleLogout = () => {
    localStorage.clear();
    navigate("/auth/signin");
  };

  const disclosure = useDisclosure();

  return (
    <Flex
      direction={"column"}
      h={"100vh"}
      overflowY={"auto"}
      flexFlow={"column"}
    >
      <Flex
        shadow={defaultShadow}
        p={3}
        alignItems={"center"}
        h={16}
        justifyContent={"space-between"}
      >
        <Heading size={"lg"}>{user?.username ?? "-"}</Heading>
        <Button colorScheme={"red"} onClick={handleLogout}>
          Logout
        </Button>
      </Flex>

      <Box p={3}>
        <Card shadow={defaultShadow} variant={"elevated"}>
          <CardBody>
            <Grid templateColumns={"1fr 1fr"}>
              <GridItem>Total Unpaid Bills</GridItem>
              <GridItem>
                Rp. {user?.totalUnpaid.toLocaleString("id-ID") ?? "-"}
              </GridItem>
            </Grid>
          </CardBody>
        </Card>
      </Box>

      <Flex direction={"row"} px={3}>
        <Card mb={3} border={"1px dashed black"} w={"100%"} mr={1}>
          <CardBody>
            <Flex direction={"row"} justifyContent={"space-around"}>
              <Heading size={"md"}>Create</Heading>
            </Flex>
          </CardBody>
        </Card>
        <Card
          onClick={disclosure.onOpen}
          mb={3}
          border={"1px dashed black"}
          w={"100%"}
          ml={1}
        >
          <CardBody>
            <Flex direction={"row"} justifyContent={"space-around"}>
              <Heading size={"md"}>Join</Heading>
            </Flex>
          </CardBody>
        </Card>
      </Flex>

      <Box px={3} overflowY={"auto"} className="disable-scrollbar">
        {user?.bills.map((el, index) => (
          <BillCard bill={el} index={index} />
        ))}
      </Box>

      <JoinBillModal disclosure={disclosure} />
    </Flex>
  );
}
