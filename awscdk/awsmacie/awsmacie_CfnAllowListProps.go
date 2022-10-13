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
	// `AWS::Macie::AllowList.Criteria`.
	Criteria interface{} `field:"required" json:"criteria" yaml:"criteria"`
	// `AWS::Macie::AllowList.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::Macie::AllowList.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Macie::AllowList.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

