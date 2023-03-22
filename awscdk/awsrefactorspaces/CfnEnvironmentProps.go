package awsrefactorspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentProps := &CfnEnvironmentProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	NetworkFabricType: jsii.String("networkFabricType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

