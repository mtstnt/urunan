import { Card, CardBody, Heading, Text } from "@chakra-ui/react";
import { defaultShadow } from "../constants";

type Props = {
    index: number,
    bill: {
        title: string,
        date: Date,
        description: string,
    }
}

export function BillCard({ index, bill }: Props) {
    return <Card shadow={defaultShadow} mb={3} key={index}>
    <CardBody>
      <Heading mb={1} size={"md"}>
        {bill.title}
      </Heading>
      <Text>{bill.date.toDateString()}</Text>
      <Text>{bill.description}</Text>
    </CardBody>
  </Card>
}