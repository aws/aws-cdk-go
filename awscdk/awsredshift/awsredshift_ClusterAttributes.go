package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties that describe an existing cluster instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   clusterAttributes := &ClusterAttributes{
//   	ClusterEndpointAddress: jsii.String("clusterEndpointAddress"),
//   	ClusterEndpointPort: jsii.Number(123),
//   	ClusterName: jsii.String("clusterName"),
//
//   	// the properties below are optional
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
// Experimental.
type ClusterAttributes struct {
	// Cluster endpoint address.
	// Experimental.
	ClusterEndpointAddress *string `field:"required" json:"clusterEndpointAddress" yaml:"clusterEndpointAddress"`
	// Cluster endpoint port.
	// Experimental.
	ClusterEndpointPort *float64 `field:"required" json:"clusterEndpointPort" yaml:"clusterEndpointPort"`
	// Identifier for the cluster.
	// Experimental.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The security groups of the redshift cluster.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

