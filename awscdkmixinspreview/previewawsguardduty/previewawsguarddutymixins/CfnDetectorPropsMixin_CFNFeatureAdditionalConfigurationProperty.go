package previewawsguarddutymixins


// Information about the additional configuration of a feature in your account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cFNFeatureAdditionalConfigurationProperty := &CFNFeatureAdditionalConfigurationProperty{
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureadditionalconfiguration.html
//
type CfnDetectorPropsMixin_CFNFeatureAdditionalConfigurationProperty struct {
	// Name of the additional configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureadditionalconfiguration.html#cfn-guardduty-detector-cfnfeatureadditionalconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Status of the additional configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureadditionalconfiguration.html#cfn-guardduty-detector-cfnfeatureadditionalconfiguration-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

