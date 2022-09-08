package awskendra

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFaq`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFaqProps := &cfnFaqProps{
//   	indexId: jsii.String("indexId"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	s3Path: &s3PathProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	fileFormat: jsii.String("fileFormat"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFaqProps struct {
	// The identifier of the index that contains the FAQ.
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// The name that you assigned the FAQ when you created or updated the FAQ.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of a role with permission to access the S3 bucket that contains the FAQ.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Simple Storage Service (Amazon S3) location of the FAQ input data.
	S3Path interface{} `field:"required" json:"s3Path" yaml:"s3Path"`
	// A description of the FAQ.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The format of the input file.
	//
	// You can choose between a basic CSV format, a CSV format that includes customs attributes in a header, and a JSON format that includes custom attributes.
	//
	// The format must match the format of the file stored in the S3 bucket identified in the S3Path parameter.
	//
	// Valid values are:
	//
	// - `CSV`
	// - `CSV_WITH_HEADER`
	// - `JSON`.
	FileFormat *string `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

