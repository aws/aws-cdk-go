package awsses


// A reference to a ConfigurationSetEventDestination resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationSetEventDestinationReference := &ConfigurationSetEventDestinationReference{
//   	ConfigurationSetEventDestinationId: jsii.String("configurationSetEventDestinationId"),
//   }
//
type ConfigurationSetEventDestinationReference struct {
	// The Id of the ConfigurationSetEventDestination resource.
	ConfigurationSetEventDestinationId *string `field:"required" json:"configurationSetEventDestinationId" yaml:"configurationSetEventDestinationId"`
}

