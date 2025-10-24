package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for an AppSync RDS datasource Aurora Serverless V2.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var databaseCluster DatabaseCluster
//   var graphqlApi GraphqlApi
//   var role Role
//   var secret Secret
//
//   rdsDataSourcePropsV2 := &RdsDataSourcePropsV2{
//   	Api: graphqlApi,
//   	SecretStore: secret,
//   	ServerlessCluster: databaseCluster,
//
//   	// the properties below are optional
//   	DatabaseName: jsii.String("databaseName"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ServiceRole: role,
//   }
//
type RdsDataSourcePropsV2 struct {
	// The API to attach this data source to.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// the description of the data source.
	// Default: - None.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source.
	// Default: - id of data source.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM service role to be assumed by AppSync to interact with the data source.
	// Default: -  Create a new role.
	//
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// The secret containing the credentials for the database.
	SecretStore awssecretsmanager.ISecret `field:"required" json:"secretStore" yaml:"secretStore"`
	// The serverless cluster to call to interact with this data source.
	ServerlessCluster awsrds.IDatabaseCluster `field:"required" json:"serverlessCluster" yaml:"serverlessCluster"`
	// The name of the database to use within the cluster.
	// Default: - None.
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
}

