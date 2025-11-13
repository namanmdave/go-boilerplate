Real-Time Chat (Full-stack, Real-time)
What to Build:
A live chat application that allows users to send/receive messages in real-time.


Backend with WebSocket or Server-Sent Events (SSE) for pushing messages.


Frontend chat interface with real-time updates.


Expected Output:
Backend server with a WebSocket or SSE endpoint that accepts and broadcasts chat messages.


Frontend that renders incoming messages, handles sending messages, and gracefully manages socket state.


Goals:
Demonstrate real-time streaming, typed message contracts, and state handling.


Handle message delivery, reconnections, and display latency-free updates.


Tests Required:
Unit tests for broadcast logic and message shape validation.


Backend integration test with multiple simulated clients.


Frontend tests for rendering chat stream and managing input state.
vqv-appn-nzi



