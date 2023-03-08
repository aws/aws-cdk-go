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
//   featureDefinitionProperty := &featureDefinitionProperty{
//   	featureName: jsii.String("featureName"),
//   	featureType: jsii.String("featureType"),
//   }
//
type CfnFeatureGroup_FeatureDefinitionProperty struct {
	// The name of a feature.
	//
	// The type must be a string. `FeatureName` cannot be any of the following: `is_deleted` , `write_time` , `api_invocation_time` .
	FeatureName *string `field:"required" json:"featureName" yaml:"featureName"`
	// The value type of a feature.
	//
	// Valid values are Integral, Fractional, or String.
	FeatureType *string `field:"required" json:"featureType" yaml:"featureType"`
}

