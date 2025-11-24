package mixinsawsopensearchservice


// A custom start time for the off-peak window, in Coordinated Universal Time (UTC).
//
// The window length will always be 10 hours, so you can't specify an end time. For example, if you specify 11:00 P.M. UTC as a start time, the end time will automatically be set to 9:00 A.M.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   windowStartTimeProperty := &WindowStartTimeProperty{
//   	Hours: jsii.Number(123),
//   	Minutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-windowstarttime.html
//
type CfnDomainPropsMixin_WindowStartTimeProperty struct {
	// The start hour of the window in Coordinated Universal Time (UTC), using 24-hour time.
	//
	// For example, 17 refers to 5:00 P.M. UTC. The minimum value is 0 and the maximum value is 23.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-windowstarttime.html#cfn-opensearchservice-domain-windowstarttime-hours
	//
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// The start minute of the window, in UTC.
	//
	// The minimum value is 0 and the maximum value is 59.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-windowstarttime.html#cfn-opensearchservice-domain-windowstarttime-minutes
	//
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
}

