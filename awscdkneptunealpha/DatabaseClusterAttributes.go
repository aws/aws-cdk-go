package awscdkneptunealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties that describe an existing cluster instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import neptune_alpha "github.com/aws/aws-cdk-go/awscdkneptunealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup SecurityGroup
//
//   databaseClusterAttributes := &DatabaseClusterAttributes{
//   	ClusterEndpointAddress: jsii.String("clusterEndpointAddress"),
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	ClusterResourceIdentifier: jsii.String("clusterResourceIdentifier"),
//   	Port: jsii.Number(123),
//   	ReaderEndpointAddress: jsii.String("readerEndpointAddress"),
//   	SecurityGroup: securityGroup,
//   }
//
// Experimental.
type DatabaseClusterAttributes struct {
	// Cluster endpoint address.
	// Experimental.
	ClusterEndpointAddress *string `field:"required" json:"clusterEndpointAddress" yaml:"clusterEndpointAddress"`
	// Identifier for the cluster.
	// Experimental.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Resource Identifier for the cluster.
	// Experimental.
	ClusterResourceIdentifier *string `field:"required" json:"clusterResourceIdentifier" yaml:"clusterResourceIdentifier"`
	// The database port.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Reader endpoint address.
	// Experimental.
	ReaderEndpointAddress *string `field:"required" json:"readerEndpointAddress" yaml:"readerEndpointAddress"`
	// The security group of the database cluster.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"required" json:"securityGroup" yaml:"securityGroup"`
}

