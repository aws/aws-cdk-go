package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties that describe an existing cluster instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var clusterEngine iClusterEngine
//   var securityGroup securityGroup
//
//   databaseClusterAttributes := &DatabaseClusterAttributes{
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//
//   	// the properties below are optional
//   	ClusterEndpointAddress: jsii.String("clusterEndpointAddress"),
//   	ClusterResourceIdentifier: jsii.String("clusterResourceIdentifier"),
//   	Engine: clusterEngine,
//   	InstanceEndpointAddresses: []*string{
//   		jsii.String("instanceEndpointAddresses"),
//   	},
//   	InstanceIdentifiers: []*string{
//   		jsii.String("instanceIdentifiers"),
//   	},
//   	Port: jsii.Number(123),
//   	ReaderEndpointAddress: jsii.String("readerEndpointAddress"),
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
type DatabaseClusterAttributes struct {
	// Identifier for the cluster.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Cluster endpoint address.
	// Default: - no endpoint address.
	//
	ClusterEndpointAddress *string `field:"optional" json:"clusterEndpointAddress" yaml:"clusterEndpointAddress"`
	// The immutable identifier for the cluster; for example: cluster-ABCD1234EFGH5678IJKL90MNOP.
	//
	// This AWS Region-unique identifier is used to grant access to the cluster.
	// Default: none.
	//
	ClusterResourceIdentifier *string `field:"optional" json:"clusterResourceIdentifier" yaml:"clusterResourceIdentifier"`
	// The engine of the existing Cluster.
	// Default: - the imported Cluster's engine is unknown.
	//
	Engine IClusterEngine `field:"optional" json:"engine" yaml:"engine"`
	// Endpoint addresses of individual instances.
	// Default: - no instance endpoints.
	//
	InstanceEndpointAddresses *[]*string `field:"optional" json:"instanceEndpointAddresses" yaml:"instanceEndpointAddresses"`
	// Identifier for the instances.
	// Default: - no instance identifiers.
	//
	InstanceIdentifiers *[]*string `field:"optional" json:"instanceIdentifiers" yaml:"instanceIdentifiers"`
	// The database port.
	// Default: - none.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Reader endpoint address.
	// Default: - no reader address.
	//
	ReaderEndpointAddress *string `field:"optional" json:"readerEndpointAddress" yaml:"readerEndpointAddress"`
	// The security groups of the database cluster.
	// Default: - no security groups.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

