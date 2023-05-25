package awsguardduty


// Describes the additional configuration for a feature.
//
// If you want to specify any additional configuration for your feature, it is required to provide the `Name` and `Status` for that additional configuration. For more information, see [DetectorAdditionalConfiguration](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DetectorAdditionalConfiguration.html) .
//
// If you're providing additional configuration, ensure to provide the corresponding [FeatureConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureconfigurations.html#cfn-guardduty-detector-featureconfigurations-additionalconfiguration) .
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

