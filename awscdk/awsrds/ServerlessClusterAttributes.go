package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties that describe an existing cluster instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//   var securityGroup securityGroup
//
//   serverlessClusterAttributes := &ServerlessClusterAttributes{
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//
//   	// the properties below are optional
//   	ClusterEndpointAddress: jsii.String("clusterEndpointAddress"),
//   	Port: jsii.Number(123),
//   	ReaderEndpointAddress: jsii.String("readerEndpointAddress"),
//   	Secret: secret,
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
type ServerlessClusterAttributes struct {
	// Identifier for the cluster.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Cluster endpoint address.
	// Default: - no endpoint address.
	//
	ClusterEndpointAddress *string `field:"optional" json:"clusterEndpointAddress" yaml:"clusterEndpointAddress"`
	// The database port.
	// Default: - none.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Reader endpoint address.
	// Default: - no reader address.
	//
	ReaderEndpointAddress *string `field:"optional" json:"readerEndpointAddress" yaml:"readerEndpointAddress"`
	// The secret attached to the database cluster.
	// Default: - no secret.
	//
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
	// The security groups of the database cluster.
	// Default: - no security groups.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

