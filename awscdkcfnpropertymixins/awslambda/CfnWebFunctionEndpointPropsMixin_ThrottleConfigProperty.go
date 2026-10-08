package awslambda


// The throttling configuration for the endpoint.
//
// Optionally constrains the request rate that the endpoint accepts, in addition to your account's rate limit quota.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   throttleConfigProperty := &ThrottleConfigProperty{
//   	RateLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-throttleconfig.html
//
type CfnWebFunctionEndpointPropsMixin_ThrottleConfigProperty struct {
	// The maximum request rate per second for the endpoint, up to a maximum of 10000.
	//
	// This optional limit further constrains the endpoint's request rate. When omitted, the endpoint's request rate is limited only by your account's rate limit quota. Specify 0 to reject all new requests. Other supported values are 100 through 1000 in increments of 100, and 2000 through 10000 in increments of 1000. Supported values can vary by Region; if you specify an unsupported value, the error lists the values available in that Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-throttleconfig.html#cfn-lambda-webfunctionendpoint-throttleconfig-ratelimit
	//
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

