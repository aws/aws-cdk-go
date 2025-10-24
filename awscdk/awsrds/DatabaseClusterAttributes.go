package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties that describe an existing cluster instance.
//
// Example:
//   var vpc Vpc
//   var fn Function
//   var secret Secret
//
//
//   // Create a serverless V1 cluster
//   serverlessV1Cluster := rds.NewServerlessCluster(this, jsii.String("AnotherCluster"), &ServerlessClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA_MYSQL(),
//   	Vpc: Vpc,
//   	 // this parameter is optional for serverless Clusters
//   	EnableDataApi: jsii.Boolean(true),
//   })
//   serverlessV1Cluster.grantDataApiAccess(fn)
//
//   // Create an Aurora cluster
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Cluster"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA_MYSQL(),
//   	Vpc: Vpc,
//   	EnableDataApi: jsii.Boolean(true),
//   })
//   cluster.GrantDataApiAccess(fn)
//
//   // Import an Aurora cluster
//   importedCluster := rds.DatabaseCluster_FromDatabaseClusterAttributes(this, jsii.String("ImportedCluster"), &DatabaseClusterAttributes{
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	Secret: Secret,
//   	DataApiEnabled: jsii.Boolean(true),
//   })
//   importedCluster.GrantDataApiAccess(fn)
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
	// Whether the Data API for the cluster is enabled.
	// Default: false.
	//
	DataApiEnabled *bool `field:"optional" json:"dataApiEnabled" yaml:"dataApiEnabled"`
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
	// The secret attached to the database cluster.
	// Default: - the imported Cluster's secret is unknown.
	//
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
	// The security groups of the database cluster.
	// Default: - no security groups.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

