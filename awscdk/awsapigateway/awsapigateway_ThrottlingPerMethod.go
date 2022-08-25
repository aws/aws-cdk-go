package awsapigateway


// Represents per-method throttling for a resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var method method
//
//   throttlingPerMethod := &throttlingPerMethod{
//   	method: method,
//   	throttle: &throttleSettings{
//   		burstLimit: jsii.Number(123),
//   		rateLimit: jsii.Number(123),
//   	},
//   }
//
type ThrottlingPerMethod struct {
	// [disable-awslint:ref-via-interface] The method for which you specify the throttling settings.
	Method Method `field:"required" json:"method" yaml:"method"`
	// Specifies the overall request rate (average requests per second) and burst capacity.
	Throttle *ThrottleSettings `field:"required" json:"throttle" yaml:"throttle"`
}

