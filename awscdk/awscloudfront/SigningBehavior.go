package awscloudfront


// Options for which requests CloudFront signs.
//
// The recommended setting is `always`.
type SigningBehavior string

const (
	// Sign all origin requests, overwriting the Authorization header from the viewer request if one exists.
	SigningBehavior_ALWAYS SigningBehavior = "ALWAYS"
	// Do not sign any origin requests.
	//
	// This value turns off origin access control for all origins in all
	// distributions that use this origin access control.
	SigningBehavior_NEVER SigningBehavior = "NEVER"
	// Sign origin requests only if the viewer request doesn't contain the Authorization header.
	SigningBehavior_NO_OVERRIDE SigningBehavior = "NO_OVERRIDE"
)

