package awsguardduty


// Information about the configuration of a feature in your account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cFNFeatureConfigurationProperty := &CFNFeatureConfigurationProperty{
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	AdditionalConfiguration: []interface{}{
//   		&CFNFeatureAdditionalConfigurationProperty{
//   			Name: jsii.String("name"),
//   			Status: jsii.String("status"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureconfiguration.html
//
type CfnDetector_CFNFeatureConfigurationProperty struct {
	// Name of the feature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureconfiguration.html#cfn-guardduty-detector-cfnfeatureconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Status of the feature configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureconfiguration.html#cfn-guardduty-detector-cfnfeatureconfiguration-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// Information about the additional configuration of a feature in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureconfiguration.html#cfn-guardduty-detector-cfnfeatureconfiguration-additionalconfiguration
	//
	AdditionalConfiguration interface{} `field:"optional" json:"additionalConfiguration" yaml:"additionalConfiguration"`
}

