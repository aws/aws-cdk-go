package interfacesawsivschat


// A reference to a LoggingConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationReference := &LoggingConfigurationReference{
//   	LoggingConfigurationArn: jsii.String("loggingConfigurationArn"),
//   }
//
type LoggingConfigurationReference struct {
	// The Arn of the LoggingConfiguration resource.
	LoggingConfigurationArn *string `field:"required" json:"loggingConfigurationArn" yaml:"loggingConfigurationArn"`
}

