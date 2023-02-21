package awslambda


// The function's [AWS X-Ray](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html) tracing configuration. To sample and record incoming requests, set `Mode` to `Active` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tracingConfigProperty := &tracingConfigProperty{
//   	mode: jsii.String("mode"),
//   }
//
type CfnFunction_TracingConfigProperty struct {
	// The tracing mode.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

