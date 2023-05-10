package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventInvokeConfigProperty := &EventInvokeConfigProperty{
//   	DestinationConfig: &EventInvokeDestinationConfigProperty{
//   		OnFailure: &DestinationProperty{
//   			Destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			Type: jsii.String("type"),
//   		},
//   		OnSuccess: &DestinationProperty{
//   			Destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	MaximumEventAgeInSeconds: jsii.Number(123),
//   	MaximumRetryAttempts: jsii.Number(123),
//   }
//
type CfnFunction_EventInvokeConfigProperty struct {
	// `CfnFunction.EventInvokeConfigProperty.DestinationConfig`.
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// `CfnFunction.EventInvokeConfigProperty.MaximumEventAgeInSeconds`.
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// `CfnFunction.EventInvokeConfigProperty.MaximumRetryAttempts`.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

