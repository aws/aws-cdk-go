package awssynthetics


// Different ways to clean up underlying Canary resources when the Canary is deleted.
type Cleanup string

const (
	// Clean up nothing.
	//
	// The user is responsible for cleaning up
	// all resources left behind by the Canary.
	Cleanup_NOTHING Cleanup = "NOTHING"
	// Clean up the underlying Lambda function only.
	//
	// The user is
	// responsible for cleaning up all other resources left behind
	// by the Canary.
	Cleanup_LAMBDA Cleanup = "LAMBDA"
)

