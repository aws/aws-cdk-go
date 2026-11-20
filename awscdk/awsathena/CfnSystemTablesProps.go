package awsathena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSystemTables`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSystemTablesProps := &CfnSystemTablesProps{
//   	TableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-systemtables.html
//
type CfnSystemTablesProps struct {
	// The reserved keyword identifying the account-level System Tables configuration.
	//
	// The only accepted value is ALL, which enables all system tables for the account. This is the resource's primary identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-systemtables.html#cfn-athena-systemtables-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The ARN of the customer managed AWS KMS key used to encrypt the system table data.
	//
	// Specify a key ARN (not an alias); the key must be in the same account and Region. The encryption choice is permanent for the account and Region: it cannot be changed or removed after creation, and it persists after this resource is deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-systemtables.html#cfn-athena-systemtables-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-systemtables.html#cfn-athena-systemtables-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

