package interfacesawsssm


// A reference to a DefaultHostManagementConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultHostManagementConfigurationReference := &DefaultHostManagementConfigurationReference{
//   	SettingArn: jsii.String("settingArn"),
//   	SettingId: jsii.String("settingId"),
//   }
//
type DefaultHostManagementConfigurationReference struct {
	// The ARN of the DefaultHostManagementConfiguration resource.
	SettingArn *string `field:"required" json:"settingArn" yaml:"settingArn"`
	// The SettingId of the DefaultHostManagementConfiguration resource.
	SettingId *string `field:"required" json:"settingId" yaml:"settingId"`
}

