package awsrefactorspaces

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentProps := &cfnEnvironmentProps{
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	networkFabricType: jsii.String("networkFabricType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEnvironmentProps struct {
	// A description of the environment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the environment.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The network fabric type of the environment.
	NetworkFabricType *string `field:"optional" json:"networkFabricType" yaml:"networkFabricType"`
	// The tags assigned to the environment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

