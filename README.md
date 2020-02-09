# grpc-up-and-running

## Services And Remote Methods
### Product Info

| Method | Pattern | Description | 
|---|---|---|
| AddProduct | Unary RPC | Add a product. |
| GetProduct | Unary RPC | Get a product by product ID. |

### Order Management

| Method | Pattern | Description | 
|---|---|---|
| AddOrder | Unary RPC | Add a new order. |
| GetOrder | Unary RPC | Get a order by order ID. |
| SearchOrders | Server-side streaming | Get all the orders which has a certain item. |
| UpdateOrders | Client-side streaming | Update multiple orders. |
| ProcessOrders | Bidirectional streaming | Process multiple orders. <li>All the order IDs will be sent from client as a stream.<li>A combined shipment will contains all the orders which will be delivered to the same destination.<li>When the max batch size is reached, all the currently created combined shipments will be sent back to the client. |
