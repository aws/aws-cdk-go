package previewawsguarddutymixins


// Information about the configuration of a feature in your account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cFNFeatureConfigurationProperty := &CFNFeatureConfigurationProperty{
//   	AdditionalConfiguration: []interface{}{
//   		&CFNFeatureAdditionalConfigurationProperty{
//   			Name: jsii.String("name"),
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureconfiguration.html
//
type CfnDetectorPropsMixin_CFNFeatureConfigurationProperty struct {
	// Information about the additional configuration of a feature in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureconfiguration.html#cfn-guardduty-detector-cfnfeatureconfiguration-additionalconfiguration
	//
	AdditionalConfiguration interface{} `field:"optional" json:"additionalConfiguration" yaml:"additionalConfiguration"`
	// Name of the feature.
	//
	// For a list of allowed values, see [DetectorFeatureConfiguration](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DetectorFeatureConfiguration.html#guardduty-Type-DetectorFeatureConfiguration-name) in the *GuardDuty API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureconfiguration.html#cfn-guardduty-detector-cfnfeatureconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Status of the feature configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureconfiguration.html#cfn-guardduty-detector-cfnfeatureconfiguration-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

