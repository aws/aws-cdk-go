package awswafv2


// A reference to a LoggingConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationReference := &LoggingConfigurationReference{
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
type LoggingConfigurationReference struct {
	// The ResourceArn of the LoggingConfiguration resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

