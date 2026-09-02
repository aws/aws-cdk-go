package awskendra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnQuerySuggestionsBlockListPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnQuerySuggestionsBlockListMixinProps := &CfnQuerySuggestionsBlockListMixinProps{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-querysuggestionsblocklist.html
//
type CfnQuerySuggestionsBlockListMixinProps struct {
	// A description for the block list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-querysuggestionsblocklist.html#cfn-kendra-querysuggestionsblocklist-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier of the index for the block list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-querysuggestionsblocklist.html#cfn-kendra-querysuggestionsblocklist-indexid
	//
	IndexId *string `field:"optional" json:"indexId" yaml:"indexId"`
	// The name of the block list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-querysuggestionsblocklist.html#cfn-kendra-querysuggestionsblocklist-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of an IAM role with permission to access the S3 bucket that contains the block list text file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-querysuggestionsblocklist.html#cfn-kendra-querysuggestionsblocklist-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Information required to find a specific file in an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-querysuggestionsblocklist.html#cfn-kendra-querysuggestionsblocklist-sources3path
	//
	SourceS3Path interface{} `field:"optional" json:"sourceS3Path" yaml:"sourceS3Path"`
	// A list of key-value pairs that identify or categorize the block list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-querysuggestionsblocklist.html#cfn-kendra-querysuggestionsblocklist-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

