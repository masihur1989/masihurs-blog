package common

import "errors"

// ErrorIntConverstion Used in D parsing
var ErrorIntConverstion = errors.New("Invalid id Passed")

// ErrorJWTStringGeneration jwt error generated
var ErrorJWTStringGeneration = errors.New("JWT String genration Error")

// ErrorHashing Used in D parsing
var ErrorHashing = errors.New("Hashing Error")

// ErrorPasswordMatching godoc
var ErrorPasswordMatching = errors.New("Password Mismatch")

// ErrorQuery Query Execution Error
var ErrorQuery = errors.New("Query Execution Error")

// ErrorTransaction Transaction Execution Error
var ErrorTransaction = errors.New("Transaction Error")

// ErrorScanning Scanning DB result Error
var ErrorScanning = errors.New("Failed to Scan DB result")

// ErrorCreatingStmnt Statement Creation Error
var ErrorCreatingStmnt = errors.New("Error Creating Statement")
