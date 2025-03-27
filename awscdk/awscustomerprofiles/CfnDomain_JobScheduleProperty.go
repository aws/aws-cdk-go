package awscustomerprofiles


// The day and time when do you want to start the Identity Resolution Job every week.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobScheduleProperty := &JobScheduleProperty{
//   	DayOfTheWeek: jsii.String("dayOfTheWeek"),
//   	Time: jsii.String("time"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-jobschedule.html
//
type CfnDomain_JobScheduleProperty struct {
	// The day when the Identity Resolution Job should run every week.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-jobschedule.html#cfn-customerprofiles-domain-jobschedule-dayoftheweek
	//
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// The time when the Identity Resolution Job should run every week.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-jobschedule.html#cfn-customerprofiles-domain-jobschedule-time
	//
	Time *string `field:"required" json:"time" yaml:"time"`
}

