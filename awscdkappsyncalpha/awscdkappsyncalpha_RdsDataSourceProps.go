// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for an AppSync RDS datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var graphqlApi graphqlApi
//   var role role
//   var secret secret
//   var serverlessCluster serverlessCluster
//
//   rdsDataSourceProps := &rdsDataSourceProps{
//   	api: graphqlApi,
//   	secretStore: secret,
//   	serverlessCluster: serverlessCluster,
//
//   	// the properties below are optional
//   	databaseName: jsii.String("databaseName"),
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	serviceRole: role,
//   }
//
// Experimental.
type RdsDataSourceProps struct {
	// The API to attach this data source to.
	// Experimental.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// the description of the data source.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM service role to be assumed by AppSync to interact with the data source.
	// Experimental.
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// The secret containing the credentials for the database.
	// Experimental.
	SecretStore awssecretsmanager.ISecret `field:"required" json:"secretStore" yaml:"secretStore"`
	// The serverless cluster to call to interact with this data source.
	// Experimental.
	ServerlessCluster awsrds.IServerlessCluster `field:"required" json:"serverlessCluster" yaml:"serverlessCluster"`
	// The name of the database to use within the cluster.
	// Experimental.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
}

