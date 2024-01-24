package awsdatazone


// The schedule of the data source runs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleConfigurationProperty := &ScheduleConfigurationProperty{
//   	Schedule: jsii.String("schedule"),
//   	Timezone: jsii.String("timezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-scheduleconfiguration.html
//
type CfnDataSource_ScheduleConfigurationProperty struct {
	// The schedule of the data source runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-scheduleconfiguration.html#cfn-datazone-datasource-scheduleconfiguration-schedule
	//
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-scheduleconfiguration.html#cfn-datazone-datasource-scheduleconfiguration-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

