package awsemrcontainers

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVirtualCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVirtualClusterProps := &cfnVirtualClusterProps{
//   	containerProvider: &containerProviderProperty{
//   		id: jsii.String("id"),
//   		info: &containerInfoProperty{
//   			eksInfo: &eksInfoProperty{
//   				namespace: jsii.String("namespace"),
//   			},
//   		},
//   		type: jsii.String("type"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnVirtualClusterProps struct {
	// The container provider of the virtual cluster.
	ContainerProvider interface{} `field:"required" json:"containerProvider" yaml:"containerProvider"`
	// The name of the virtual cluster.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

