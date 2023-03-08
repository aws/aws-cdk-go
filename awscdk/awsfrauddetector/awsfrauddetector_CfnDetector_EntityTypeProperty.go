package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// The entity type details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityTypeProperty := &entityTypeProperty{
//   	arn: jsii.String("arn"),
//   	createdTime: jsii.String("createdTime"),
//   	description: jsii.String("description"),
//   	inline: jsii.Boolean(false),
//   	lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDetector_EntityTypeProperty struct {
	// The entity type ARN.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Timestamp of when the entity type was created.
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// The entity type description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether the resource is defined within this CloudFormation template and impacts the create, update, and delete behavior of the stack.
	//
	// If the value is `true` , CloudFormation will create/update/delete the resource when creating/updating/deleting the stack. If the value is `false` , CloudFormation will validate that the object exists and then use it within the resource without making changes to the object.
	//
	// For example, when creating `AWS::FraudDetector::Detector` you must define at least two variables. You can set `Inline=true` for these Variables and CloudFormation will create/update/delete the variables as part of stack operations. However, if you set `Inline=false` , CloudFormation will associate the variables to your detector but not execute any changes to the variables.
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
	// Timestamp of when the entity type was last updated.
	LastUpdatedTime *string `field:"optional" json:"lastUpdatedTime" yaml:"lastUpdatedTime"`
	// The entity type name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

