package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Reference to a dynamodb table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//
//   tableAttributes := &TableAttributes{
//   	EncryptionKey: key,
//   	GlobalIndexes: []*string{
//   		jsii.String("globalIndexes"),
//   	},
//   	GrantIndexPermissions: jsii.Boolean(false),
//   	LocalIndexes: []*string{
//   		jsii.String("localIndexes"),
//   	},
//   	TableArn: jsii.String("tableArn"),
//   	TableName: jsii.String("tableName"),
//   	TableStreamArn: jsii.String("tableStreamArn"),
//   }
//
type TableAttributes struct {
	// KMS encryption key, if this table uses a customer-managed encryption key.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The name of the global indexes set for this Table.
	//
	// Note that you need to set either this property,
	// or `localIndexes`,
	// if you want methods like grantReadData()
	// to grant permissions for indexes as well as the table itself.
	GlobalIndexes *[]*string `field:"optional" json:"globalIndexes" yaml:"globalIndexes"`
	// If set to true, grant methods always grant permissions for all indexes.
	//
	// If false is provided, grant methods grant the permissions
	// only when `globalIndexes` or `localIndexes` is specified.
	GrantIndexPermissions *bool `field:"optional" json:"grantIndexPermissions" yaml:"grantIndexPermissions"`
	// The name of the local indexes set for this Table.
	//
	// Note that you need to set either this property,
	// or `globalIndexes`,
	// if you want methods like grantReadData()
	// to grant permissions for indexes as well as the table itself.
	LocalIndexes *[]*string `field:"optional" json:"localIndexes" yaml:"localIndexes"`
	// The ARN of the dynamodb table.
	//
	// One of this, or `tableName`, is required.
	TableArn *string `field:"optional" json:"tableArn" yaml:"tableArn"`
	// The table name of the dynamodb table.
	//
	// One of this, or `tableArn`, is required.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// The ARN of the table's stream.
	TableStreamArn *string `field:"optional" json:"tableStreamArn" yaml:"tableStreamArn"`
}

