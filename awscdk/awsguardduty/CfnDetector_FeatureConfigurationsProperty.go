package awsguardduty


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   featureConfigurationsProperty := &FeatureConfigurationsProperty{
//   	AdditionalConfiguration: []interface{}{
//   		&FeatureAdditionalConfigurationProperty{
//   			Name: jsii.String("name"),
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//   }
//
type CfnDetector_FeatureConfigurationsProperty struct {
	// `CfnDetector.FeatureConfigurationsProperty.AdditionalConfiguration`.
	AdditionalConfiguration interface{} `field:"optional" json:"additionalConfiguration" yaml:"additionalConfiguration"`
	// `CfnDetector.FeatureConfigurationsProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnDetector.FeatureConfigurationsProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

