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
//   var api IApi
//   var role Role
//   var table Table
//
//   appSyncDynamoDbDataSourceProps := &AppSyncDynamoDbDataSourceProps{
//   	Api: api,
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
type AppSyncDynamoDbDataSourceProps struct {
	// The API to attach this data source to.
	Api IApi `field:"required" json:"api" yaml:"api"`
	// The description of the data source.
	// Default: - None.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source.
	//
	// The only allowed pattern is: {[_A-Za-z][_0-9A-Za-z]*}.
	// Any invalid characters will be automatically removed.
	// Default: - id of data source.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM service role to be assumed by AppSync to interact with the data source.
	// Default: -  Create a new role.
	//
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// The DynamoDB table backing this data source.
	Table awsdynamodb.ITable `field:"required" json:"table" yaml:"table"`
	// Specify whether this Data Source is read only or has read and write permissions to the DynamoDB table.
	// Default: false.
	//
	ReadOnlyAccess *bool `field:"optional" json:"readOnlyAccess" yaml:"readOnlyAccess"`
	// Use credentials of caller to access DynamoDB.
	// Default: false.
	//
	UseCallerCredentials *bool `field:"optional" json:"useCallerCredentials" yaml:"useCallerCredentials"`
}

