import {
  Alert,
  AlertIcon,
  Button,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
} from "@chakra-ui/react";
import { useState } from "react";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";

type DisclosureState = {
  isOpen: boolean;
  onOpen: () => void;
  onClose: () => void;
};

type Props = {
  disclosure: DisclosureState;
};

export default function JoinBillModal({ disclosure }: Props) {
  const [code, setCode] = useState("");
  const navigate = useNavigate();
  const handleJoinClicked = () => {
    toast.success("Joining " + code);
    navigate("/" + code);
  };
  return (
    <Modal isOpen={disclosure.isOpen} onClose={disclosure.onClose}>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>Join an Existing Bill</ModalHeader>
        <ModalCloseButton />
        <ModalBody>
          <Alert borderRadius={"5px"} mb={3} status="info">
            <AlertIcon />
            Code is a 6 letter long word in each unique link.
          </Alert>
          <Input
            size={"lg"}
            placeholder='e.g: "A1B3C45"'
            required
            maxLength={6}
            value={code}
            onInput={(e) => setCode(e.currentTarget.value)}
          />
        </ModalBody>
        <ModalFooter>
          <Button colorScheme="red" mr={3} onClick={disclosure.onClose}>
            Close
          </Button>
          <Button
            colorScheme="blue"
            variant="solid"
            onClick={handleJoinClicked}
          >
            Submit
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
}
