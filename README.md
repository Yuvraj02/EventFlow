# EventFlow

## Overview

EventFlow is a backend service written in Go that ingests, controls, and processes events in a reliable and scalable manner.  
It is designed to protect the system under load, handle failures gracefully, and ensure predictable event processing.

---

## What EventFlow Does

### 1. Ingests events
- Accepts events (meaningful facts that have occurred) via an external interface (e.g., HTTP).
- Treats events as immutable records of something that already happened.

### 2. Applies backpressure at ingestion
- Accepts events only when system capacity is available.
- Rejects new events when capacity is exceeded to prevent overload.
- Ensures the system remains stable under high load.

### 3. Buffers accepted events
- Stores accepted events in a bounded queue.
- Decouples event ingestion from event processing.

### 4. Processes events using a bounded worker pool
- Uses a fixed number of workers to process events concurrently.
- Prevents unbounded concurrency and protects downstream resources.
- Workers pull events from the queue when ready.

### 5. Persists events
- Stores events for durability and later analysis.
- Initial implementation persists events to a JSON file (storage is pluggable).

### 6. Handles processing failures
- Detects failures during event processing.
- Applies defined policies such as retry, drop, or record failure.
- Prevents individual failures from crashing the entire service.

### 7. Shuts down gracefully
- Stops accepting new events on shutdown.
- Allows in-flight events to finish processing.
- Ensures no data loss during termination.

---

## What EventFlow Is Not

- Not a UI or frontend system
- Not a workflow engine
- Not a general-purpose message broker
- Not Kubernetes-specific logic

EventFlow focuses only on **controlled event ingestion and processing**.

---

## Core Design Principles

- **Events are facts**  
  Events describe what has already happened; they do not describe how the system should act.

- **Backpressure over blind acceptance**  
  The system prioritizes survivability over accepting all input.

- **Bounded concurrency**  
  All processing is explicitly limited and predictable.

- **Separation of concerns**
  - Ingestion is fast and lightweight
  - Processing is asynchronous and controlled
  - Persistence is isolated from business logic

---

## High-Level Flow

Incoming Event
↓
Ingestion (Backpressure Check)
↓
Bounded Queue
↓
Worker Pool
↓
Processing & Persistence

---

## Goal

EventFlow exists to turn uncontrolled incoming events into controlled, reliable processing while maintaining system stability, scalability, and correctness.