package awslookoutequipment


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputNameConfigurationProperty := &InputNameConfigurationProperty{
//   	ComponentTimestampDelimiter: jsii.String("componentTimestampDelimiter"),
//   	TimestampFormat: jsii.String("timestampFormat"),
//   }
//
type CfnInferenceScheduler_InputNameConfigurationProperty struct {
	// `CfnInferenceScheduler.InputNameConfigurationProperty.ComponentTimestampDelimiter`.
	ComponentTimestampDelimiter *string `field:"optional" json:"componentTimestampDelimiter" yaml:"componentTimestampDelimiter"`
	// `CfnInferenceScheduler.InputNameConfigurationProperty.TimestampFormat`.
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
}

