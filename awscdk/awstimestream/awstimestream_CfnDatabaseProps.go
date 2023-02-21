package awstimestream

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDatabase`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatabaseProps := &cfnDatabaseProps{
//   	databaseName: jsii.String("databaseName"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDatabaseProps struct {
	// The name of the Timestream database.
	//
	// *Length Constraints* : Minimum length of 3 bytes. Maximum length of 256 bytes.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The identifier of the AWS KMS key used to encrypt the data stored in the database.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The tags to add to the database.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

