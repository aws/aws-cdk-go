package awsguardduty


// Describes the additional configuration for a feature.
//
// If you provide any additional configuration for your feature, it is required to also provide `Name` and `Status` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   featureAdditionalConfigurationProperty := &FeatureAdditionalConfigurationProperty{
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//   }
//
type CfnDetector_FeatureAdditionalConfigurationProperty struct {
	// Name of the additional configuration of a feature.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Status of the additional configuration of a feature.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

