package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses"
)

// Properties of a reference to an existing configuration set event destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationSetRef IConfigurationSetRef
//
//   configurationSetEventDestinationAttributes := &ConfigurationSetEventDestinationAttributes{
//   	ConfigurationSet: configurationSetRef,
//   	ConfigurationSetEventDestinationId: jsii.String("configurationSetEventDestinationId"),
//   }
//
type ConfigurationSetEventDestinationAttributes struct {
	// The configuration set that the event destination belongs to.
	ConfigurationSet interfacesawsses.IConfigurationSetRef `field:"required" json:"configurationSet" yaml:"configurationSet"`
	// The ID of the configuration set event destination.
	ConfigurationSetEventDestinationId *string `field:"required" json:"configurationSetEventDestinationId" yaml:"configurationSetEventDestinationId"`
}

