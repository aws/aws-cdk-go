package awsmacie

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAllowList`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAllowListProps := &cfnAllowListProps{
//   	criteria: &criteriaProperty{
//   		regex: jsii.String("regex"),
//   		s3WordsList: &s3WordsListProperty{
//   			bucketName: jsii.String("bucketName"),
//   			objectKey: jsii.String("objectKey"),
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAllowListProps struct {
	// The criteria that specify the text or text pattern to ignore.
	//
	// The criteria can be the location and name of an Amazon S3 object that lists specific text to ignore ( `S3WordsList` ), or a regular expression ( `Regex` ) that defines a text pattern to ignore.
	Criteria interface{} `field:"required" json:"criteria" yaml:"criteria"`
	// A custom name for the allow list.
	//
	// The name can contain 1-128 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A custom description of the allow list.
	//
	// The description can contain 1-512 characters.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to the allow list.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

