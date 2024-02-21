package awsappstream


// The desired capacity for a fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeCapacityProperty := &ComputeCapacityProperty{
//   	DesiredInstances: jsii.Number(123),
//   	DesiredSessions: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-computecapacity.html
//
type CfnFleet_ComputeCapacityProperty struct {
	// The desired number of streaming instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-computecapacity.html#cfn-appstream-fleet-computecapacity-desiredinstances
	//
	DesiredInstances *float64 `field:"optional" json:"desiredInstances" yaml:"desiredInstances"`
	// The desired capacity in terms of number of user sessions, for the multi-session fleet.
	//
	// This is not allowed for single-session fleets.
	//
	// When you create a fleet, you must set define either the DesiredSessions or DesiredInstances attribute, based on the type of fleet you create. You can’t define both attributes or leave both attributes blank.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-computecapacity.html#cfn-appstream-fleet-computecapacity-desiredsessions
	//
	DesiredSessions *float64 `field:"optional" json:"desiredSessions" yaml:"desiredSessions"`
}

