package awslicensemanager


// Details about a consumption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   consumptionConfigurationProperty := &consumptionConfigurationProperty{
//   	borrowConfiguration: &borrowConfigurationProperty{
//   		allowEarlyCheckIn: jsii.Boolean(false),
//   		maxTimeToLiveInMinutes: jsii.Number(123),
//   	},
//   	provisionalConfiguration: &provisionalConfigurationProperty{
//   		maxTimeToLiveInMinutes: jsii.Number(123),
//   	},
//   	renewType: jsii.String("renewType"),
//   }
//
type CfnLicense_ConsumptionConfigurationProperty struct {
	// Details about a borrow configuration.
	BorrowConfiguration interface{} `field:"optional" json:"borrowConfiguration" yaml:"borrowConfiguration"`
	// Details about a provisional configuration.
	ProvisionalConfiguration interface{} `field:"optional" json:"provisionalConfiguration" yaml:"provisionalConfiguration"`
	// Renewal frequency.
	RenewType *string `field:"optional" json:"renewType" yaml:"renewType"`
}

