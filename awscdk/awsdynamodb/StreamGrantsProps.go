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
//   	Table: tableRef,
//   	TableStreamArn: jsii.String("tableStreamArn"),
//
//   	// the properties below are optional
//   	EncryptionKey: key,
//   }
//
type StreamGrantsProps struct {
	// The table this stream is for.
	Table interfacesawsdynamodb.ITableRef `field:"required" json:"table" yaml:"table"`
	// The ARN of the Stream.
	TableStreamArn *string `field:"required" json:"tableStreamArn" yaml:"tableStreamArn"`
	// The encryption key of the table.
	//
	// Required permissions will be added to the key as well.
	// Default: - No key.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

