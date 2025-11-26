package previewawskendramixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFaqPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFaqMixinProps := &CfnFaqMixinProps{
//   	Description: jsii.String("description"),
//   	FileFormat: jsii.String("fileFormat"),
//   	IndexId: jsii.String("indexId"),
//   	LanguageCode: jsii.String("languageCode"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	S3Path: &S3PathProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html
//
type CfnFaqMixinProps struct {
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
	// The identifier of the index that contains the FAQ.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-indexid
	//
	IndexId *string `field:"optional" json:"indexId" yaml:"indexId"`
	// The code for a language.
	//
	// This shows a supported language for the FAQ document as part of the summary information for FAQs. English is supported by default. For more information on supported languages, including their codes, see [Adding documents in languages other than English](https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The name that you assigned the FAQ when you created or updated the FAQ.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of a role with permission to access the S3 bucket that contains the FAQ.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The Amazon Simple Storage Service (Amazon S3) location of the FAQ input data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-s3path
	//
	S3Path interface{} `field:"optional" json:"s3Path" yaml:"s3Path"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html#cfn-kendra-faq-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

