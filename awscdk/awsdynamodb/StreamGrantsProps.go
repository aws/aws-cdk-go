package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdynamodb"
)

// Construction properties for StreamGrants.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key Key
//   var tableRef ITableRef
//
//   streamGrantsProps := &StreamGrantsProps{
//   	TableStreamArn: jsii.String("tableStreamArn"),
//
//   	// the properties below are optional
//   	EncryptionKey: key,
//   	Table: tableRef,
//   }
//
type StreamGrantsProps struct {
	// The ARN of the Stream.
	TableStreamArn *string `field:"required" json:"tableStreamArn" yaml:"tableStreamArn"`
	// The encryption key of the table.
	//
	// Required permissions will be added to the key as well.
	// Default: - No key.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The table this stream is for.
	// Default: - None, no longer required.
	//
	// Deprecated: This property is not used anymore.
	Table interfacesawsdynamodb.ITableRef `field:"optional" json:"table" yaml:"table"`
}

