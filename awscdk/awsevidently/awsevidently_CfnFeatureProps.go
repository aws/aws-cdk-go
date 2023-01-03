package awsevidently

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFeature`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFeatureProps := &cfnFeatureProps{
//   	name: jsii.String("name"),
//   	project: jsii.String("project"),
//   	variations: []interface{}{
//   		&variationObjectProperty{
//   			variationName: jsii.String("variationName"),
//
//   			// the properties below are optional
//   			booleanValue: jsii.Boolean(false),
//   			doubleValue: jsii.Number(123),
//   			longValue: jsii.Number(123),
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	defaultVariation: jsii.String("defaultVariation"),
//   	description: jsii.String("description"),
//   	entityOverrides: []interface{}{
//   		&entityOverrideProperty{
//   			entityId: jsii.String("entityId"),
//   			variation: jsii.String("variation"),
//   		},
//   	},
//   	evaluationStrategy: jsii.String("evaluationStrategy"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFeatureProps struct {
	// The name for the feature.
	//
	// It can include up to 127 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name or ARN of the project that is to contain the new feature.
	Project *string `field:"required" json:"project" yaml:"project"`
	// An array of structures that contain the configuration of the feature's different variations.
	//
	// Each `VariationObject` in the `Variations` array for a feature must have the same type of value ( `BooleanValue` , `DoubleValue` , `LongValue` or `StringValue` ).
	Variations interface{} `field:"required" json:"variations" yaml:"variations"`
	// The name of the variation to use as the default variation.
	//
	// The default variation is served to users who are not allocated to any ongoing launches or experiments of this feature.
	//
	// This variation must also be listed in the `Variations` structure.
	//
	// If you omit `DefaultVariation` , the first variation listed in the `Variations` structure is used as the default variation.
	DefaultVariation *string `field:"optional" json:"defaultVariation" yaml:"defaultVariation"`
	// An optional description of the feature.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specify users that should always be served a specific variation of a feature.
	//
	// Each user is specified by a key-value pair . For each key, specify a user by entering their user ID, account ID, or some other identifier. For the value, specify the name of the variation that they are to be served.
	EntityOverrides interface{} `field:"optional" json:"entityOverrides" yaml:"entityOverrides"`
	// Specify `ALL_RULES` to activate the traffic allocation specified by any ongoing launches or experiments.
	//
	// Specify `DEFAULT_VARIATION` to serve the default variation to all users instead.
	EvaluationStrategy *string `field:"optional" json:"evaluationStrategy" yaml:"evaluationStrategy"`
	// Assigns one or more tags (key-value pairs) to the feature.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with a feature.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

