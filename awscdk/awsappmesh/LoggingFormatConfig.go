package awsappmesh


// All Properties for Envoy Access Logging Format for mesh endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingFormatConfig := &LoggingFormatConfig{
//   	FormatConfig: &LoggingFormatProperty{
//   		Json: []interface{}{
//   			&JsonFormatRefProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Text: jsii.String("text"),
//   	},
//   }
//
type LoggingFormatConfig struct {
	// CFN configuration for Access Logging Format.
	// Default: - no access logging format.
	//
	FormatConfig *CfnVirtualNode_LoggingFormatProperty `field:"optional" json:"formatConfig" yaml:"formatConfig"`
}

