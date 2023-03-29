package awsguardduty


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
	// `CfnDetector.FeatureAdditionalConfigurationProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnDetector.FeatureAdditionalConfigurationProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

