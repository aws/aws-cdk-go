package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Attributes of a DynamoDB table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key Key
//
//   tableAttributesV2 := &TableAttributesV2{
//   	EncryptionKey: key,
//   	GlobalIndexes: []*string{
//   		jsii.String("globalIndexes"),
//   	},
//   	GrantIndexPermissions: jsii.Boolean(false),
//   	LocalIndexes: []*string{
//   		jsii.String("localIndexes"),
//   	},
//   	TableArn: jsii.String("tableArn"),
//   	TableId: jsii.String("tableId"),
//   	TableName: jsii.String("tableName"),
//   	TableStreamArn: jsii.String("tableStreamArn"),
//   }
//
type TableAttributesV2 struct {
	// KMS encryption key for the table.
	// Default: - no KMS encryption key.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The name of the global indexes set for the table.
	//
	// Note: You must set either this property or `localIndexes` if you want permissions
	// to be granted for indexes as well as the table itself.
	// Default: - no global indexes.
	//
	GlobalIndexes *[]*string `field:"optional" json:"globalIndexes" yaml:"globalIndexes"`
	// Whether or not to grant permissions for all indexes of the table.
	//
	// Note: If false, permissions will only be granted to indexes when `globalIndexes`
	// or `localIndexes` is specified.
	// Default: false.
	//
	GrantIndexPermissions *bool `field:"optional" json:"grantIndexPermissions" yaml:"grantIndexPermissions"`
	// The name of the local indexes set for the table.
	//
	// Note: You must set either this property or `globalIndexes` if you want permissions
	// to be granted for indexes as well as the table itself.
	// Default: - no local indexes.
	//
	LocalIndexes *[]*string `field:"optional" json:"localIndexes" yaml:"localIndexes"`
	// The ARN of the table.
	//
	// Note: You must specify this or the `tableName`.
	// Default: - table arn generated using `tableName` and region of stack.
	//
	TableArn *string `field:"optional" json:"tableArn" yaml:"tableArn"`
	// The ID of the table.
	// Default: - no table id.
	//
	TableId *string `field:"optional" json:"tableId" yaml:"tableId"`
	// The name of the table.
	//
	// Note: You must specify this or the `tableArn`.
	// Default: - table name retrieved from provided `tableArn`.
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// The stream ARN of the table.
	// Default: - no table stream ARN.
	//
	TableStreamArn *string `field:"optional" json:"tableStreamArn" yaml:"tableStreamArn"`
}

