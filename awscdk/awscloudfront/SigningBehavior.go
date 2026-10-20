package awscloudfront


// Options for which requests CloudFront signs.
//
// The recommended setting is `always`.
type SigningBehavior string

const (
	// Sign all origin requests, overwriting the Authorization header from the viewer request if one exists.
	SigningBehavior_ALWAYS SigningBehavior = "ALWAYS"
	// Sign all origin requests, and pass the Authorization header from the viewer request to the origin unchanged.
	//
	// CloudFront adds its signature in a dedicated header instead of the
	// Authorization header. This is the only signing behavior supported by
	// AWS Lambda Web Function origins, and it isn't supported by any other
	// origin type.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-lambda-web.html
	//
	SigningBehavior_ALWAYS_AMZ_AUTH SigningBehavior = "ALWAYS_AMZ_AUTH"
	// Do not sign any origin requests.
	//
	// This value turns off origin access control for all origins in all
	// distributions that use this origin access control.
	SigningBehavior_NEVER SigningBehavior = "NEVER"
	// Sign origin requests only if the viewer request doesn't contain the Authorization header.
	SigningBehavior_NO_OVERRIDE SigningBehavior = "NO_OVERRIDE"
)

