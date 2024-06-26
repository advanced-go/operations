package envoy

// Startup sequence.
// 1. Read assignments and create case officers for each assignment
// 1b. Periodically poll for changes in assignments
// 2. Determine which assignments are active and activate case officers
// 3. Determine if a new case officer class is available and start conversions.
// 4. Update assignments whenever a case officer class is converted.
// 5. Audit conversion status, query and review new inferences and actions
