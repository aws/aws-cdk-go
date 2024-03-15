package awssagemaker


// A list of features.
//
// You must include `FeatureName` and `FeatureType` . Valid feature `FeatureType` s are `Integral` , `Fractional` and `String` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   featureDefinitionProperty := &FeatureDefinitionProperty{
//   	FeatureName: jsii.String("featureName"),
//   	FeatureType: jsii.String("featureType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-featuredefinition.html
//
type CfnFeatureGroup_FeatureDefinitionProperty struct {
	// The name of a feature.
	//
	// The type must be a string. `FeatureName` cannot be any of the following: `is_deleted` , `write_time` , `api_invocation_time` .
	//
	// The name:
	//
	// - Must start and end with an alphanumeric character.
	// - Can only include alphanumeric characters, underscores, and hyphens. Spaces are not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-featuredefinition.html#cfn-sagemaker-featuregroup-featuredefinition-featurename
	//
	FeatureName *string `field:"required" json:"featureName" yaml:"featureName"`
	// The value type of a feature.
	//
	// Valid values are Integral, Fractional, or String.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-featuredefinition.html#cfn-sagemaker-featuregroup-featuredefinition-featuretype
	//
	FeatureType *string `field:"required" json:"featureType" yaml:"featureType"`
}

