package awsmacie

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAllowList`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAllowListProps := &CfnAllowListProps{
//   	Criteria: &CriteriaProperty{
//   		Regex: jsii.String("regex"),
//   		S3WordsList: &S3WordsListProperty{
//   			BucketName: jsii.String("bucketName"),
//   			ObjectKey: jsii.String("objectKey"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-allowlist.html
//
type CfnAllowListProps struct {
	// The criteria that specify the text or text pattern to ignore.
	//
	// The criteria can be the location and name of an Amazon S3 object that lists specific text to ignore ( `S3WordsList` ), or a regular expression ( `Regex` ) that defines a text pattern to ignore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-allowlist.html#cfn-macie-allowlist-criteria
	//
	Criteria interface{} `field:"required" json:"criteria" yaml:"criteria"`
	// A custom name for the allow list.
	//
	// The name can contain 1-128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-allowlist.html#cfn-macie-allowlist-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A custom description of the allow list.
	//
	// The description can contain 1-512 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-allowlist.html#cfn-macie-allowlist-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to the allow list.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-allowlist.html#cfn-macie-allowlist-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

