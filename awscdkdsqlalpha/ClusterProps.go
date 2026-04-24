package awscdkdsqlalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a new Aurora DSQL cluster.
//
// Example:
//   cluster := dsql.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	ClusterName: jsii.String("my-dsql-cluster"),
//   	DeletionProtection: jsii.Boolean(true),
//   })
//
// Experimental.
type ClusterProps struct {
	// The name of the DSQL cluster.
	//
	// This is applied via the `Name` tag.
	// Default: - No name specified.
	//
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Specifies whether this cluster can be deleted.
	//
	// If deletionProtection is
	// enabled, the cluster cannot be deleted unless it is modified and
	// deletionProtection is disabled. deletionProtection protects clusters from
	// being accidentally deleted.
	// Default: - true if `removalPolicy` is RETAIN, undefined otherwise.
	//
	// Experimental.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The removal policy to apply when the cluster is removed or replaced during a stack update, or when the stack is deleted.
	// Default: - Retain cluster.
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

