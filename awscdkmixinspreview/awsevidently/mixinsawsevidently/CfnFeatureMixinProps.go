package mixinsawsevidently

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFeaturePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFeatureMixinProps := &CfnFeatureMixinProps{
//   	DefaultVariation: jsii.String("defaultVariation"),
//   	Description: jsii.String("description"),
//   	EntityOverrides: []interface{}{
//   		&EntityOverrideProperty{
//   			EntityId: jsii.String("entityId"),
//   			Variation: jsii.String("variation"),
//   		},
//   	},
//   	EvaluationStrategy: jsii.String("evaluationStrategy"),
//   	Name: jsii.String("name"),
//   	Project: jsii.String("project"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Variations: []interface{}{
//   		&VariationObjectProperty{
//   			BooleanValue: jsii.Boolean(false),
//   			DoubleValue: jsii.Number(123),
//   			LongValue: jsii.Number(123),
//   			StringValue: jsii.String("stringValue"),
//   			VariationName: jsii.String("variationName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-feature.html
//
type CfnFeatureMixinProps struct {
	// The name of the variation to use as the default variation.
	//
	// The default variation is served to users who are not allocated to any ongoing launches or experiments of this feature.
	//
	// This variation must also be listed in the `Variations` structure.
	//
	// If you omit `DefaultVariation` , the first variation listed in the `Variations` structure is used as the default variation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-feature.html#cfn-evidently-feature-defaultvariation
	//
	DefaultVariation *string `field:"optional" json:"defaultVariation" yaml:"defaultVariation"`
	// An optional description of the feature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-feature.html#cfn-evidently-feature-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specify users that should always be served a specific variation of a feature.
	//
	// Each user is specified by a key-value pair . For each key, specify a user by entering their user ID, account ID, or some other identifier. For the value, specify the name of the variation that they are to be served.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-feature.html#cfn-evidently-feature-entityoverrides
	//
	EntityOverrides interface{} `field:"optional" json:"entityOverrides" yaml:"entityOverrides"`
	// Specify `ALL_RULES` to activate the traffic allocation specified by any ongoing launches or experiments.
	//
	// Specify `DEFAULT_VARIATION` to serve the default variation to all users instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-feature.html#cfn-evidently-feature-evaluationstrategy
	//
	EvaluationStrategy *string `field:"optional" json:"evaluationStrategy" yaml:"evaluationStrategy"`
	// The name for the feature.
	//
	// It can include up to 127 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-feature.html#cfn-evidently-feature-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name or ARN of the project that is to contain the new feature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-feature.html#cfn-evidently-feature-project
	//
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Assigns one or more tags (key-value pairs) to the feature.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with a feature.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-feature.html#cfn-evidently-feature-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An array of structures that contain the configuration of the feature's different variations.
	//
	// Each `VariationObject` in the `Variations` array for a feature must have the same type of value ( `BooleanValue` , `DoubleValue` , `LongValue` or `StringValue` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-feature.html#cfn-evidently-feature-variations
	//
	Variations interface{} `field:"optional" json:"variations" yaml:"variations"`
}

