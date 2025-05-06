# E-Commerce Platform

## Services
- **user-service**: Manages user registration and authentication.
- **inventory-service**: Handles product catalog and stock management.
- **order-service**: Manages order creation and status.
- **api-gateway**: Routes requests to appropriate services.
- **producer-service**: Publishes `order.created` events to NATS.
- **consumer-service**: Subscribes to `order.created` events and updates inventory.

## How to Run
1. Start all services using Docker Compose:
   ```
   docker-compose up --build
   ```

2. Test the event flow:
   - Create an order via `order-service`.
   - Verify that `producer-service` publishes an event to NATS.
   - Confirm that `consumer-service` processes the event and updates inventory.

## Architecture
- **gRPC**: Used for internal communication between services.
- **NATS**: Used for asynchronous messaging between `producer-service` and `consumer-service`.

