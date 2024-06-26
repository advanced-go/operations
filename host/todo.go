package host

// Architecture should support processing in failure of host scenarios

// Startup sequence
// 1. Initialize agent to read list of resources/hosts, and start health checks.
// 2. Based on Envoy health check status, re-assign work to new case officer.
// 3. Periodically re-assign work
