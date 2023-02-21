package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventInvokeDestinationConfigProperty := &EventInvokeDestinationConfigProperty{
//   	OnFailure: &DestinationProperty{
//   		Destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		Type: jsii.String("type"),
//   	},
//   	OnSuccess: &DestinationProperty{
//   		Destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		Type: jsii.String("type"),
//   	},
//   }
//
type CfnFunction_EventInvokeDestinationConfigProperty struct {
	// `CfnFunction.EventInvokeDestinationConfigProperty.OnFailure`.
	OnFailure interface{} `field:"required" json:"onFailure" yaml:"onFailure"`
	// `CfnFunction.EventInvokeDestinationConfigProperty.OnSuccess`.
	OnSuccess interface{} `field:"required" json:"onSuccess" yaml:"onSuccess"`
}

