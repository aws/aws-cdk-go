package previewawsdatazonemixins


// The details of the schedule of the data source runs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scheduleConfigurationProperty := &ScheduleConfigurationProperty{
//   	Schedule: jsii.String("schedule"),
//   	Timezone: jsii.String("timezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-scheduleconfiguration.html
//
type CfnDataSourcePropsMixin_ScheduleConfigurationProperty struct {
	// The schedule of the data source runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-scheduleconfiguration.html#cfn-datazone-datasource-scheduleconfiguration-schedule
	//
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// The timezone of the data source run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-scheduleconfiguration.html#cfn-datazone-datasource-scheduleconfiguration-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

