# How to run?

There are multiple ways to run this. You can run it using kubernetes, docker compose or go. 

## docker-compose

Run the following command in the root.
```bash
docker compose up
```

Service is available at:
grpc://localhost:50051

Traces are available at
http://localhost:16686/

## Go

Run the following command in the root.
```bash
go run main.go
```

Service is available at:
grpc://localhost:50051

## Kubernetes

Run the following commands in the root.
```bash
minikube start

eval $(minikube docker-env)

docker build -t zendesk-homework .

kubectl apply -f k8s/

kubectl port-forward service/app 50051:50051
```

Service is available at:
grpc://localhost:50051

If you wish to access Jaeger.
```bash
kubectl port-forward service/jaeger 16686:16686
```

Traces are available at
http://localhost:16686/

# About tests and testing

I have added some unit tests. But I've used mainly integration tests to cover the application.

```bash
go test ./...
```

If you want to test this by hand please generate a JWT. For this you can use:
https://jwt.io/

Default JWT secret is `very-secret-key`

I have also added server reflection. So you can use that to set up POSTMAN. Remember to add the Authorization.

# What I did and why?

I implemented a gRPC service that calculates ticket scores based on weighted categories. Here's a detailed breakdown of the implementation:

## Score Calculation Algorithm
I used a weighted average algorithm to calculate the percentages. This approach was chosen because:
- It properly accounts for the importance of different rating categories through their weights
- It provides a fair representation of overall performance by considering both the rating values and their relative importance
- The formula `(rating * weight) / (sum of weights) / 5.0 * 100.0` ensures scores are normalized to a 0-100% scale

## Service Architecture
The service is built with:
- gRPC for efficient communication and type safety
- SQLite for data storage (with the option to switch to other databases in production)
- OpenTelemetry for observability and tracing
- JWT authentication for secure access

## Key Features
1. **Aggregated Category Scores**
   - Calculates daily scores for periods under a month
   - Automatically switches to weekly aggregation for longer periods
   - Provides both individual category scores and overall performance metrics

2. **Ticket-Level Analysis**
   - Breaks down scores by individual tickets
   - Shows category-specific performance for each ticket
   - Helps identify patterns and areas for improvement

3. **Overall Quality Score**
   - Calculates a single percentage score representing overall performance
   - Considers all ratings and their weights
   - Provides a quick snapshot of quality across all categories

4. **Period-over-Period Comparison**
   - Calculates percentage changes between time periods
   - Useful for tracking performance trends
   - Handles edge cases like zero scores in previous periods

## Technical Decisions
- Used Go as the primary language due to Klaus using GO for new services
- Implemented server reflection for easier testing and debugging
- Added comprehensive integration tests to ensure reliability
- Designed the service to be container-friendly for easy deployment

## Future Improvements
While the current implementation is functional, there are several areas that could be enhanced:
- Database migration system for schema changes
- Response caching for better performance
- Query optimization with proper indexing
- Chunking for handling large datasets efficiently