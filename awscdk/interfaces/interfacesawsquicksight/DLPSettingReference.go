package interfacesawsquicksight


// A reference to a DLPSetting resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dLPSettingReference := &DLPSettingReference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	DlpSettingArn: jsii.String("dlpSettingArn"),
//   	DlpSettingId: jsii.String("dlpSettingId"),
//   }
//
type DLPSettingReference struct {
	// The AwsAccountId of the DLPSetting resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ARN of the DLPSetting resource.
	DlpSettingArn *string `field:"required" json:"dlpSettingArn" yaml:"dlpSettingArn"`
	// The DlpSettingId of the DLPSetting resource.
	DlpSettingId *string `field:"required" json:"dlpSettingId" yaml:"dlpSettingId"`
}

