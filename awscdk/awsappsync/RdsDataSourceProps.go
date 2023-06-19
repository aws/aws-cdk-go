package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Properties for an AppSync RDS datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var graphqlApi graphqlApi
//   var role role
//   var secret secret
//   var serverlessCluster serverlessCluster
//
//   rdsDataSourceProps := &RdsDataSourceProps{
//   	Api: graphqlApi,
//   	SecretStore: secret,
//   	ServerlessCluster: serverlessCluster,
//
//   	// the properties below are optional
//   	DatabaseName: jsii.String("databaseName"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ServiceRole: role,
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

