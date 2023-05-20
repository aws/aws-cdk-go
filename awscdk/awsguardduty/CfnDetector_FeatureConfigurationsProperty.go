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
type CfnDetector_FeatureConfigurationsProperty struct {
	// Additional configuration of the feature.
	AdditionalConfiguration interface{} `field:"optional" json:"additionalConfiguration" yaml:"additionalConfiguration"`
	// Name of the feature.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Status of the feature.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

