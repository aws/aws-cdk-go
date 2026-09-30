package awseventsv2


// Storage (retention) configuration for the event bus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageConfigurationProperty := &StorageConfigurationProperty{
//   	RetentionPeriodInDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventbus-storageconfiguration.html
//
type CfnEventBus_StorageConfigurationProperty struct {
	// The number of days events are retained on the event bus for replay, 1-365.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventbus-storageconfiguration.html#cfn-eventsv2-eventbus-storageconfiguration-retentionperiodindays
	//
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
}

