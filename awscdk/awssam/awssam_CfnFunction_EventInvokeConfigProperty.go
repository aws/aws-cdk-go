package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventInvokeConfigProperty := &eventInvokeConfigProperty{
//   	destinationConfig: &eventInvokeDestinationConfigProperty{
//   		onFailure: &destinationProperty{
//   			destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			type: jsii.String("type"),
//   		},
//   		onSuccess: &destinationProperty{
//   			destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			type: jsii.String("type"),
//   		},
//   	},
//   	maximumEventAgeInSeconds: jsii.Number(123),
//   	maximumRetryAttempts: jsii.Number(123),
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

