package mixinsawsdatazone


// The lineage sync schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lineageSyncScheduleProperty := &LineageSyncScheduleProperty{
//   	Schedule: jsii.String("schedule"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-lineagesyncschedule.html
//
type CfnConnectionPropsMixin_LineageSyncScheduleProperty struct {
	// The lineage sync schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-lineagesyncschedule.html#cfn-datazone-connection-lineagesyncschedule-schedule
	//
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
}

