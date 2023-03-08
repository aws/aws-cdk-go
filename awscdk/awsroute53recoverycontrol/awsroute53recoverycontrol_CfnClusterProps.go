package awsroute53recoverycontrol

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &cfnClusterProps{
//   	clusterEndpoints: []interface{}{
//   		&clusterEndpointProperty{
//   			endpoint: jsii.String("endpoint"),
//   			region: jsii.String("region"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnClusterProps struct {
	// `AWS::Route53RecoveryControl::Cluster.ClusterEndpoints`.
	ClusterEndpoints interface{} `field:"optional" json:"clusterEndpoints" yaml:"clusterEndpoints"`
	// Name of the cluster.
	//
	// You can use any non-white space character in the name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value for a tag.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

