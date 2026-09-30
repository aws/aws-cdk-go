package awsbedrock


// A weekly refresh on a specified day of the week.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   weeklyScheduleProperty := &WeeklyScheduleProperty{
//   	DayOfWeek: jsii.String("dayOfWeek"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-weeklyschedule.html
//
type CfnDataSourcePropsMixin_WeeklyScheduleProperty struct {
	// Day of the week.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-weeklyschedule.html#cfn-bedrock-datasource-weeklyschedule-dayofweek
	//
	DayOfWeek *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
}

