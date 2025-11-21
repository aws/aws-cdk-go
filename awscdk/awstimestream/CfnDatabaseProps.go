package awstimestream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDatabase`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatabaseProps := &CfnDatabaseProps{
//   	DatabaseName: jsii.String("databaseName"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-database.html
//
type CfnDatabaseProps struct {
	// The name of the Timestream database.
	//
	// *Length Constraints* : Minimum length of 3 bytes. Maximum length of 256 bytes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-database.html#cfn-timestream-database-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The identifier of the AWS  key used to encrypt the data stored in the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-database.html#cfn-timestream-database-kmskeyid
	//
	KmsKeyId interface{} `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The tags to add to the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-database.html#cfn-timestream-database-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

