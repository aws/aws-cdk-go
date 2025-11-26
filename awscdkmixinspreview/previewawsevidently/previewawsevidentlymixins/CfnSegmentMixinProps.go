package previewawsevidentlymixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSegmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSegmentMixinProps := &CfnSegmentMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Pattern: jsii.String("pattern"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-segment.html
//
type CfnSegmentMixinProps struct {
	// An optional description for this segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-segment.html#cfn-evidently-segment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-segment.html#cfn-evidently-segment-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The pattern to use for the segment.
	//
	// For more information about pattern syntax, see [Segment rule pattern syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html#CloudWatch-Evidently-segments-syntax) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-segment.html#cfn-evidently-segment-pattern
	//
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Assigns one or more tags (key-value pairs) to the feature.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with a feature.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-segment.html#cfn-evidently-segment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

