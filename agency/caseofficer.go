package agency

// Changes to status or assigned region need to generate case officer changes to remove the old assignment
// and create the new assignment
// When a partition is added, then the assignments should be generated.
// A corresponding case officer change should be generated

// Case officer responsibilities

// Operations - provide the following functionality related to partitions
// 1. Assign
// 2. Revoke
// 3. Pause
// 4. Resume
// 5. Start
// 6. Stop

// Functionality related to processing
// 1. Look for new services/assignments
// 2. Look for agents that appear to be tombstone or not actively processing
