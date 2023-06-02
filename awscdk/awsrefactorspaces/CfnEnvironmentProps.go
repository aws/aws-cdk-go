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
//   	Name: jsii.String("name"),
//   	NetworkFabricType: jsii.String("networkFabricType"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEnvironmentProps struct {
	// The name of the environment.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network fabric type of the environment.
	NetworkFabricType *string `field:"required" json:"networkFabricType" yaml:"networkFabricType"`
	// A description of the environment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags assigned to the environment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

