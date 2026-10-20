package awscloudfront


// The signing protocol of the Origin Access Control.
type SigningProtocol string

const (
	// The AWS Signature Version 4 signing protocol.
	SigningProtocol_SIGV4 SigningProtocol = "SIGV4"
	// The AWS Signature Version 4A signing protocol.
	//
	// Signs requests for multiple AWS Regions, as needed by origins such as
	// Amazon S3 Multi-Region Access Points.
	// See: https://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html
	//
	SigningProtocol_SIGV4A SigningProtocol = "SIGV4A"
)

