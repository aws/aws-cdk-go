package awsssm


// Properties for CfnServiceSettingPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnServiceSettingMixinProps := &CfnServiceSettingMixinProps{
//   	SettingId: jsii.String("settingId"),
//   	SettingValue: jsii.String("settingValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-servicesetting.html
//
type CfnServiceSettingMixinProps struct {
	// The ID of the service setting, such as /ssm/parameter-store/high-throughput-enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-servicesetting.html#cfn-ssm-servicesetting-settingid
	//
	SettingId *string `field:"optional" json:"settingId" yaml:"settingId"`
	// The value of the service setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-servicesetting.html#cfn-ssm-servicesetting-settingvalue
	//
	SettingValue *string `field:"optional" json:"settingValue" yaml:"settingValue"`
}

