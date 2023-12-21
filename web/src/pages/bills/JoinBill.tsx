import React, {
  ChangeEvent,
  MouseEvent,
  useEffect,
  useMemo,
  useState,
} from "react";
import {
  Link,
  createSearchParams,
  useNavigate,
  useParams,
  useSearchParams,
} from "react-router-dom";
import { useLoadingStore } from "../../stores/loading";
import toast from "react-hot-toast";
import {
  Alert,
  AlertIcon,
  Box,
  Button,
  Card,
  CardBody,
  Flex,
  Heading,
  Input,
  Stack,
  Text,
} from "@chakra-ui/react";
import { Bill, Item, User } from "../../models/Bill";
import { wait } from "../../utils/token";

/**
 * JoinBill handles accessing the bill through an invitation link.
 * This part will ask for the nickname and the orders the participant takes.
 * http://localhost:5173/ABCDEF
 */

async function fetchBillByCode(code: string): Promise<Bill> {
  await wait(1);
  return {
    id: 1,
    code: code,
    title: "Test",
    description: "Description",
    items: [
      {
        id: 1,
        name: "Owo",
        initialQty: 4,
        price: 20000,
      },
      {
        id: 2,
        name: "Coklat",
        initialQty: 2,
        price: 80000,
      },
    ],
    participants: [],
  };
}

async function fetchCurrentUser(): Promise<User | null> {
  await wait(1);
  return localStorage.getItem('token') != null ? {
    id: 1,
    email: "user@email.com",
    fullName: "User 1",
  } : null;
}

export default function JoinBill() {
  const params = useParams<string>();
  const { setIsLoading } = useLoadingStore();
  const [, setSearchParams] = useSearchParams();
  const navigate = useNavigate();
  const code = params.code!;
  const [user, setUser] = useState<User | null>(null);
  const [bill, setBill] = useState<Bill | null>(null);
  const nickNameState = useState("");
  const ordersState = useState<Map<number, number>>(new Map());
  const stageState = useState<number>(1);
  const [stage] = stageState;
  const handleSubmitAll = () => {};

  useEffect(() => {
    setIsLoading(true);

    Promise.all([fetchCurrentUser(), fetchBillByCode(code)]).then((results) => {
      const [user, bill] = results;
      if (user == null) {
        toast.error("Must login first.");
        setSearchParams([["next", location.pathname]]);
        setIsLoading(false);
        navigate({
          pathname: "/auth/signin",
          search: createSearchParams({
            next: location.pathname,
          }).toString(),
        });
        return;
      }
      setUser(user);
      setBill(bill);
      setIsLoading(false);
    });
  }, [setIsLoading, code, navigate, setSearchParams]);

  return (
    <>
      {stage == 1 && (
        <NicknamePromptForm
          nickNameState={nickNameState}
          stageState={stageState}
          bill={bill}
          user={user}
        />
      )}
      {stage == 2 && (
        <OrderPromptForm
          bill={bill}
          stageState={stageState}
          ordersState={ordersState}
          handleSubmitAll={handleSubmitAll}
        />
      )}
    </>
  );
}

interface NicknamePromptFormProps {
  bill: Bill | null;
  user: User | null;
  nickNameState: [string, React.Dispatch<React.SetStateAction<string>>];
  stageState: [number, React.Dispatch<React.SetStateAction<number>>];
}

function NicknamePromptForm({
  nickNameState,
  stageState,
  bill,
  user,
}: NicknamePromptFormProps) {
  const [stage, setStage] = stageState;
  const [nickname, setNickname] = nickNameState;

  const handleSubmitNickname = () => {
    setStage(stage + 1);
  };

  return (
    <Stack
      bg={""}
      justifyContent={"center"}
      height={"100vh"}
      width={"80%"}
      margin={"0 auto"}
      alignItems={"center"}
      direction={"column"}
    >
      <Heading mb={3} size={"lg"} textAlign={"center"} width={"100%"}>
        Hi, {user?.fullName} <br />
        You are joining "{bill?.title}"
      </Heading>

      <Alert status="info" rounded={"5px"} mb={3}>
        <AlertIcon />
        Is this your account and the right bill? <br />
        Click on cancel to abort joining.
      </Alert>

      <Input
        placeholder="Enter a nickname"
        mb={3}
        textAlign={"center"}
        value={nickname}
        onChange={(e) => setNickname(e.currentTarget.value)}
      />

      <Button w={"100%"} colorScheme="green" onClick={handleSubmitNickname}>
        Submit
      </Button>
      <Link to={"/"} style={{ width: "100%" }}>
        <Button colorScheme={"gray"} w={"100%"}>
          Cancel
        </Button>
      </Link>
    </Stack>
  );
}

interface OrderPromptFormProps {
  stageState: [number, React.Dispatch<React.SetStateAction<number>>];
  ordersState: [
    Map<number, number>,
    React.Dispatch<React.SetStateAction<Map<number, number>>>
  ];
  bill: Bill | null;
  handleSubmitAll: () => void;
}

function OrderPromptForm({
  stageState,
  ordersState,
  handleSubmitAll,
  bill,
}: OrderPromptFormProps) {
  const [stage, setStage] = stageState;
  const [orders, setOrders] = ordersState;

  const itemsMap = useMemo(() => {
    const m: { [itemId: number]: Item } = {};
    for (const item of bill?.items ?? []) {
      m[item.id] = item;
    }
    return m;
  }, [bill]);

  const updateQty = <T extends object>(
    itemId: number,
    remainingQty: number,
    updateFn: (event: T, previous: number) => number
  ) => {
    return (event: T) => {
      let newValue = updateFn(event, orders.get(itemId) ?? 0);
      newValue = Math.max(newValue, 0);
      newValue = Math.min(newValue, remainingQty);
      setOrders(new Map(orders.set(itemId, newValue)));
    };
  };

  const total = useMemo(
    () =>
      Array.from(orders.entries()).reduce(
        (acc, current) => acc + current[1] * itemsMap[current[0]].price,
        0
      ),
    [orders, itemsMap]
  );

  return (
    <>
      <Stack direction={"column"} height={"100vh"} p={5}>
        <Heading size={"lg"} height={"5%"}>
          Select your Orders
        </Heading>
        <Flex
          className="disable-scrollbar"
          direction={"column"}
          height={"100%"}
          overflowY={"auto"}
        >
          {bill?.items.map((item, index) => (
            <Card key={index} mb={2}>
              <CardBody>
                <Flex
                  direction={"row"}
                  justifyContent={"space-between"}
                  alignItems={"center"}
                >
                  <Box>
                    <Heading size={"sm"}>{item.name}</Heading>
                    Remaining: {item.initialQty} <br />
                    Price: Rp{item.price.toLocaleString("id-ID")}
                  </Box>
                  <Flex justifyItems={"end"}>
                    <Button
                      mx={2}
                      colorScheme="red"
                      // TODO: Replace initialQty with remainingQty from backend.
                      onClick={updateQty<MouseEvent>(
                        item.id,
                        item.initialQty,
                        (_, a) => a - 1
                      )}
                    >
                      -
                    </Button>
                    <Input
                      width={"60px"}
                      value={orders.get(item.id) ?? 0}
                      type="number"
                      min={0}
                      onChange={updateQty<ChangeEvent<HTMLInputElement>>(
                        item.id,
                        item.initialQty,
                        (e) => parseInt(e.currentTarget.value)
                      )}
                      max={item.initialQty}
                    />
                    <Button
                      mx={2}
                      colorScheme="green"
                      onClick={updateQty<MouseEvent>(
                        item.id,
                        item.initialQty,
                        (_, a) => a + 1
                      )}
                    >
                      +
                    </Button>
                  </Flex>
                </Flex>
              </CardBody>
            </Card>
          ))}
        </Flex>
        <Flex
          direction={"row"}
          width={"100%"}
          justifyContent={"space-between"}
          alignItems={"center"}
        >
          <Button colorScheme={"red"} onClick={() => setStage(stage - 1)}>
            Go Back
          </Button>
          <Text>Total: Rp{total.toLocaleString("id-ID")}</Text>
          <Button colorScheme={"blue"} onClick={handleSubmitAll}>
            Submit
          </Button>
        </Flex>
      </Stack>
    </>
  );
}
