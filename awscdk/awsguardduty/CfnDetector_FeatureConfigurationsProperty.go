package awsguardduty


// Describes the configuration for a feature.
//
// Although the `Required` field associated with the following properties specifies `No` , if you provide information for `Name` , you will need to provide the information for `Status` too. For information about the available feature configurations, see [DetectorFeatureConfiguration](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DetectorFeatureConfiguration.html) .
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureconfigurations.html
//
type CfnDetector_FeatureConfigurationsProperty struct {
	// Additional configuration of the feature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureconfigurations.html#cfn-guardduty-detector-featureconfigurations-additionalconfiguration
	//
	AdditionalConfiguration interface{} `field:"optional" json:"additionalConfiguration" yaml:"additionalConfiguration"`
	// Name of the feature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureconfigurations.html#cfn-guardduty-detector-featureconfigurations-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Status of the feature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-featureconfigurations.html#cfn-guardduty-detector-featureconfigurations-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

