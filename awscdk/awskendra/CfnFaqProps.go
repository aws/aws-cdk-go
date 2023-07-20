package awskendra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFaq`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFaqProps := &CfnFaqProps{
//   	IndexId: jsii.String("indexId"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	S3Path: &S3PathProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	FileFormat: jsii.String("fileFormat"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html
//
type CfnFaqProps struct {
	// The identifier of the index that contains the FAQ.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-indexid
	//
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// The name that you assigned the FAQ when you created or updated the FAQ.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of a role with permission to access the S3 bucket that contains the FAQ.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Simple Storage Service (Amazon S3) location of the FAQ input data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-s3path
	//
	S3Path interface{} `field:"required" json:"s3Path" yaml:"s3Path"`
	// A description for the FAQ.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-description
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-fileformat
	//
	FileFormat *string `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

