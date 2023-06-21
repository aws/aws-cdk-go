package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for an AppSync DynamoDB datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var graphqlApi graphqlApi
//   var role role
//   var table table
//
//   dynamoDbDataSourceProps := &DynamoDbDataSourceProps{
//   	Api: graphqlApi,
//   	Table: table,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ReadOnlyAccess: jsii.Boolean(false),
//   	ServiceRole: role,
//   	UseCallerCredentials: jsii.Boolean(false),
//   }
//
type DynamoDbDataSourceProps struct {
	// The API to attach this data source to.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// the description of the data source.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM service role to be assumed by AppSync to interact with the data source.
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// The DynamoDB table backing this data source.
	Table awsdynamodb.ITable `field:"required" json:"table" yaml:"table"`
	// Specify whether this DS is read only or has read and write permissions to the DynamoDB table.
	ReadOnlyAccess *bool `field:"optional" json:"readOnlyAccess" yaml:"readOnlyAccess"`
	// use credentials of caller to access DynamoDB.
	UseCallerCredentials *bool `field:"optional" json:"useCallerCredentials" yaml:"useCallerCredentials"`
}

