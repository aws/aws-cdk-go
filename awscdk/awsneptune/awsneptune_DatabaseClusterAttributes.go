package awsneptune

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
//   databaseClusterAttributes := &databaseClusterAttributes{
//   	clusterEndpointAddress: jsii.String("clusterEndpointAddress"),
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   	clusterResourceIdentifier: jsii.String("clusterResourceIdentifier"),
//   	port: jsii.Number(123),
//   	readerEndpointAddress: jsii.String("readerEndpointAddress"),
//   	securityGroup: securityGroup,
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

