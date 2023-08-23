package cmd

import "errors"

// ErrMissingFlag — represents the error for missing required flag
var ErrMissingFlag = errors.New("missing required argument")
