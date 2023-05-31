package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
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
//   	LocalIndexes: []*string{
//   		jsii.String("localIndexes"),
//   	},
//   	TableArn: jsii.String("tableArn"),
//   	TableName: jsii.String("tableName"),
//   	TableStreamArn: jsii.String("tableStreamArn"),
//   }
//
// Experimental.
type TableAttributes struct {
	// KMS encryption key, if this table uses a customer-managed encryption key.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The name of the global indexes set for this Table.
	//
	// Note that you need to set either this property,
	// or {@link localIndexes},
	// if you want methods like grantReadData()
	// to grant permissions for indexes as well as the table itself.
	// Experimental.
	GlobalIndexes *[]*string `field:"optional" json:"globalIndexes" yaml:"globalIndexes"`
	// The name of the local indexes set for this Table.
	//
	// Note that you need to set either this property,
	// or {@link globalIndexes},
	// if you want methods like grantReadData()
	// to grant permissions for indexes as well as the table itself.
	// Experimental.
	LocalIndexes *[]*string `field:"optional" json:"localIndexes" yaml:"localIndexes"`
	// The ARN of the dynamodb table.
	//
	// One of this, or {@link tableName}, is required.
	// Experimental.
	TableArn *string `field:"optional" json:"tableArn" yaml:"tableArn"`
	// The table name of the dynamodb table.
	//
	// One of this, or {@link tableArn}, is required.
	// Experimental.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// The ARN of the table's stream.
	// Experimental.
	TableStreamArn *string `field:"optional" json:"tableStreamArn" yaml:"tableStreamArn"`
}

