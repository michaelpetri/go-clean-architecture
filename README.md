# Todo List (Clean Architecture Example)

Example project to demonstrate a clean architecture in Go.

## Project Structure

All the code related to the subdomain **Todo List** can be found in the `todo` folder. Within this folder, the following
architectural layers are organized:

### Adapter

Contains all entry points. In this example, this is `cli`, but in the future, additional adapters such as `rest` or 
`grpc` can be added to support different interfaces.

### Application

Holds everything related to bootstrapping the application. This layer includes:

- **Use Cases:** All capabilities of our subdomain are located in the `usecase` folder.
- **Shared Logic:** If multiple use cases share the same logic, a `service` folder can be added to centralize it.

### Domain

Contains the core business logic. It is crucial that this layer does not include implementation details like databases,
API clients, or de-/serializers.

### Infrastructure

Houses implementation details such as databases, API clients, or de-/serializers.
