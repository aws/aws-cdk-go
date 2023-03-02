package awsapigatewayv2


// Container for defining throttling parameters to API stages.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   throttleSettings := &throttleSettings{
//   	burstLimit: jsii.Number(123),
//   	rateLimit: jsii.Number(123),
//   }
//
// Experimental.
type ThrottleSettings struct {
	// The maximum API request rate limit over a time ranging from one to a few seconds.
	// Experimental.
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// The API request steady-state rate limit (average requests per second over an extended period of time).
	// Experimental.
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

