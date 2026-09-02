package awskendra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnThesaurusPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnThesaurusMixinProps := &CfnThesaurusMixinProps{
//   	Description: jsii.String("description"),
//   	IndexId: jsii.String("indexId"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	SourceS3Path: &S3PathProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-thesaurus.html
//
type CfnThesaurusMixinProps struct {
	// A description for the thesaurus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-thesaurus.html#cfn-kendra-thesaurus-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier of the index for the thesaurus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-thesaurus.html#cfn-kendra-thesaurus-indexid
	//
	IndexId *string `field:"optional" json:"indexId" yaml:"indexId"`
	// A name for the thesaurus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-thesaurus.html#cfn-kendra-thesaurus-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An IAM role that gives Amazon Kendra permissions to access the thesaurus file specified in SourceS3Path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-thesaurus.html#cfn-kendra-thesaurus-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Information required to find a specific file in an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-thesaurus.html#cfn-kendra-thesaurus-sources3path
	//
	SourceS3Path interface{} `field:"optional" json:"sourceS3Path" yaml:"sourceS3Path"`
	// A list of key-value pairs that identify or categorize the thesaurus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-thesaurus.html#cfn-kendra-thesaurus-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

