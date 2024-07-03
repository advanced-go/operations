package agency

// Client side operations agency

// Assignment
// 1. Startup - read landscape and assign partitions to case officers.
//            - look for case officer conversions
// 2. Ongoing - query for case officer changes applicable to the host region, and apply those changes
//              to the appropriate case officers

// Startup sequence.
// 1. Read partitions and create case officers for each partition
// 1a. Create a status channel for each case officer that is used for communication, such as database errors.
// 1b. If database is down, communicate directly to center via HTTP
// 1c. Periodically poll for changes in assignments
// 2. Determine which assignments are active and activate case officers
// 3. Determine if a new case officer class is available and start conversions.
// 4. Update assignments whenever a case officer class is converted.
// 5. Audit conversion status, query and review new inferences and actions
