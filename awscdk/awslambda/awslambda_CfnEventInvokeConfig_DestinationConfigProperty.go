package awslambda


// A configuration object that specifies the destination of an event after Lambda processes it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConfigProperty := &destinationConfigProperty{
//   	onFailure: &onFailureProperty{
//   		destination: jsii.String("destination"),
//   	},
//   	onSuccess: &onSuccessProperty{
//   		destination: jsii.String("destination"),
//   	},
//   }
//
type CfnEventInvokeConfig_DestinationConfigProperty struct {
	// The destination configuration for failed invocations.
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination configuration for successful invocations.
	OnSuccess interface{} `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

