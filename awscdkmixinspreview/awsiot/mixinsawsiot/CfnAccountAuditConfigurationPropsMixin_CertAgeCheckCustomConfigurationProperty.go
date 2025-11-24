package mixinsawsiot


// Configuration structure containing settings for the device certificate age check.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   certAgeCheckCustomConfigurationProperty := &CertAgeCheckCustomConfigurationProperty{
//   	CertAgeThresholdInDays: jsii.String("certAgeThresholdInDays"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-certagecheckcustomconfiguration.html
//
type CfnAccountAuditConfigurationPropsMixin_CertAgeCheckCustomConfigurationProperty struct {
	// The number of days that defines when a device certificate is considered to have aged.
	//
	// The check will report a finding if a certificate has been active for a number of days greater than or equal to this threshold value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-certagecheckcustomconfiguration.html#cfn-iot-accountauditconfiguration-certagecheckcustomconfiguration-certagethresholdindays
	//
	CertAgeThresholdInDays *string `field:"optional" json:"certAgeThresholdInDays" yaml:"certAgeThresholdInDays"`
}

