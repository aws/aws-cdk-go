package awsiot


// Configuration for the device certificate age audit check.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceCertAgeAuditCheckConfigurationProperty := &DeviceCertAgeAuditCheckConfigurationProperty{
//   	Configuration: &CertAgeCheckCustomConfigurationProperty{
//   		CertAgeThresholdInDays: jsii.String("certAgeThresholdInDays"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-devicecertageauditcheckconfiguration.html
//
type CfnAccountAuditConfiguration_DeviceCertAgeAuditCheckConfigurationProperty struct {
	// Configuration settings for the device certificate age check, including the threshold in days for certificate age.
	//
	// This configuration is of type `CertAgeCheckCustomConfiguration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-devicecertageauditcheckconfiguration.html#cfn-iot-accountauditconfiguration-devicecertageauditcheckconfiguration-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// True if this audit check is enabled for this account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-devicecertageauditcheckconfiguration.html#cfn-iot-accountauditconfiguration-devicecertageauditcheckconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

