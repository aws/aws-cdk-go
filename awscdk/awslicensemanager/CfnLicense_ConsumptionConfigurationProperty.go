package awslicensemanager


// Details about a consumption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   consumptionConfigurationProperty := &ConsumptionConfigurationProperty{
//   	BorrowConfiguration: &BorrowConfigurationProperty{
//   		AllowEarlyCheckIn: jsii.Boolean(false),
//   		MaxTimeToLiveInMinutes: jsii.Number(123),
//   	},
//   	ProvisionalConfiguration: &ProvisionalConfigurationProperty{
//   		MaxTimeToLiveInMinutes: jsii.Number(123),
//   	},
//   	RenewType: jsii.String("renewType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html
//
type CfnLicense_ConsumptionConfigurationProperty struct {
	// Details about a borrow configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html#cfn-licensemanager-license-consumptionconfiguration-borrowconfiguration
	//
	BorrowConfiguration interface{} `field:"optional" json:"borrowConfiguration" yaml:"borrowConfiguration"`
	// Details about a provisional configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html#cfn-licensemanager-license-consumptionconfiguration-provisionalconfiguration
	//
	ProvisionalConfiguration interface{} `field:"optional" json:"provisionalConfiguration" yaml:"provisionalConfiguration"`
	// Renewal frequency.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html#cfn-licensemanager-license-consumptionconfiguration-renewtype
	//
	RenewType *string `field:"optional" json:"renewType" yaml:"renewType"`
}

