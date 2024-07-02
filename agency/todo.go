package agency

// Architecture should support processing in failure of agency scenarios

// Startup sequence
//
// 2. Based on Envoy health check status, re-assign work to new case officer.
// 3. Periodically re-assign work
