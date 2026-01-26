# Todo List API (Go)

This project is a REST API written in Go, intended to be the backend for a simple Todo List application.

Although the domain is simple, the goal of this project is to explore production-level backend concepts while learning Go deeply and idiomatically.

## Goals

- Learn Go by building a real-world backend
- Apply clean and maintainable architecture principles
- Use production-oriented tooling and patterns
- Evolve the system incrementally, without rushing

This project is intentionally designed to grow over time.

## Architecture

The API follows a **Clean Layered Architecture**, separating concerns into well-defined layers:

- HTTP / transport layer (handlers, middleware)
- Application layer (use cases)
- Domain layer (business rules)
- Infrastructure layer (database, external services)

The architecture aims to keep the core business logic independent from frameworks and external tools.

### Database Access

- PostgreSQL is used for persistence
- SQL access is handled with **sqlc**
- Domain models are not ORM structs
- SQL queries are the source of truth for data models

This approach favors explicitness, type safety, and performance.

## Features

Current / planned features include:

- RESTful API
- Authentication
- PostgreSQL persistence
- Structured logging
- Graceful shutdown
- Context propagation
- Background processing and concurrency patterns

## Logging

The application uses **structured logging** with Go's `slog`.

Logs are designed to be machine-readable and ready for centralized log aggregation.

## Observability (Planned)

The project is designed with observability in mind and will later integrate **OpenTelemetry**.

Planned stack:

- OpenTelemetry for instrumentation
- Grafana Tempo for distributed tracing
- Prometheus for metrics
- Loki for log aggregation
- Grafana as a unified visualization layer

At the moment, only structured logging is implemented. Tracing and metrics will be added incrementally.

## Why This Project Exists

This is not meant to be the simplest possible Todo API.

It is a learning project focused on:

- Writing idiomatic Go
- Understanding concurrency and context propagation
- Designing systems that resemble real production backends
- Making architectural trade-offs consciously

The project prioritizes correctness, clarity, and learning over speed.

## Status

Work in progress.

The system will evolve over time as new concepts are introduced and refined.
