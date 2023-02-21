package awsroute53recoverycontrol

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &CfnClusterProps{
//   	ClusterEndpoints: []interface{}{
//   		&ClusterEndpointProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			Region: jsii.String("region"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

