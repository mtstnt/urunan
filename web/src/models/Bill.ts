
export type User = {
    id: number
    email: string
    fullName: string
}

export type Order = {
    id: number
    item_id: number
    qty: number
}

export type Participant = {
    id: number
    user: User
    orders: Order[]
}

export type Item = {
    id: number
    name: string
    price: number
    initialQty: number
}

export type Bill = {
    id: number
    code: string
    title: string
    description: string
    participants: Participant[]
    items: Item[]
}