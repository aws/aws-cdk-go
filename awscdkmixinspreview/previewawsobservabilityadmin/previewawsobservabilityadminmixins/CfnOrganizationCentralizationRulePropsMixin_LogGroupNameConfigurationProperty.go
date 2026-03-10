package previewawsobservabilityadminmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logGroupNameConfigurationProperty := &LogGroupNameConfigurationProperty{
//   	LogGroupNamePattern: jsii.String("logGroupNamePattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-loggroupnameconfiguration.html
//
type CfnOrganizationCentralizationRulePropsMixin_LogGroupNameConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-loggroupnameconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-loggroupnameconfiguration-loggroupnamepattern
	//
	LogGroupNamePattern *string `field:"optional" json:"logGroupNamePattern" yaml:"logGroupNamePattern"`
}

