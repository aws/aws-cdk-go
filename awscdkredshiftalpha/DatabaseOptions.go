package awscdkredshiftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for accessing a Redshift database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import redshift_alpha "github.com/aws/aws-cdk-go/awscdkredshiftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var secret secret
//
//   databaseOptions := &DatabaseOptions{
//   	Cluster: cluster,
//   	DatabaseName: jsii.String("databaseName"),
//
//   	// the properties below are optional
//   	AdminUser: secret,
//   }
//
// Experimental.
type DatabaseOptions struct {
	// The cluster containing the database.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The name of the database.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The secret containing credentials to a Redshift user with administrator privileges.
	//
	// Secret JSON schema: `{ username: string; password: string }`.
	// Default: - the admin secret is taken from the cluster.
	//
	// Experimental.
	AdminUser awssecretsmanager.ISecret `field:"optional" json:"adminUser" yaml:"adminUser"`
}

