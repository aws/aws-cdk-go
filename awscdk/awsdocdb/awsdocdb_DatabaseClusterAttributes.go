package awsdocdb

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
//   databaseClusterAttributes := &DatabaseClusterAttributes{
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//
//   	// the properties below are optional
//   	ClusterEndpointAddress: jsii.String("clusterEndpointAddress"),
//   	InstanceEndpointAddresses: []*string{
//   		jsii.String("instanceEndpointAddresses"),
//   	},
//   	InstanceIdentifiers: []*string{
//   		jsii.String("instanceIdentifiers"),
//   	},
//   	Port: jsii.Number(123),
//   	ReaderEndpointAddress: jsii.String("readerEndpointAddress"),
//   	SecurityGroup: securityGroup,
//   }
//
// Experimental.
type DatabaseClusterAttributes struct {
	// Identifier for the cluster.
	// Experimental.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Cluster endpoint address.
	// Experimental.
	ClusterEndpointAddress *string `field:"optional" json:"clusterEndpointAddress" yaml:"clusterEndpointAddress"`
	// Endpoint addresses of individual instances.
	// Experimental.
	InstanceEndpointAddresses *[]*string `field:"optional" json:"instanceEndpointAddresses" yaml:"instanceEndpointAddresses"`
	// Identifier for the instances.
	// Experimental.
	InstanceIdentifiers *[]*string `field:"optional" json:"instanceIdentifiers" yaml:"instanceIdentifiers"`
	// The database port.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Reader endpoint address.
	// Experimental.
	ReaderEndpointAddress *string `field:"optional" json:"readerEndpointAddress" yaml:"readerEndpointAddress"`
	// The security group of the database cluster.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

