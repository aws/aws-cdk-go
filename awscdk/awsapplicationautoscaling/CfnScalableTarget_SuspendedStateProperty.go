package awsapplicationautoscaling


// `SuspendedState` is a property of the [AWS::ApplicationAutoScaling::ScalableTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html) resource that specifies whether the scaling activities for a scalable target are in a suspended state.
//
// For more information, see [Suspending and resuming scaling](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-suspend-resume-scaling.html) in the *Application Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   suspendedStateProperty := &SuspendedStateProperty{
//   	DynamicScalingInSuspended: jsii.Boolean(false),
//   	DynamicScalingOutSuspended: jsii.Boolean(false),
//   	ScheduledScalingSuspended: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-suspendedstate.html
//
type CfnScalableTarget_SuspendedStateProperty struct {
	// Whether scale in by a target tracking scaling policy or a step scaling policy is suspended.
	//
	// Set the value to `true` if you don't want Application Auto Scaling to remove capacity when a scaling policy is triggered. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-suspendedstate.html#cfn-applicationautoscaling-scalabletarget-suspendedstate-dynamicscalinginsuspended
	//
	DynamicScalingInSuspended interface{} `field:"optional" json:"dynamicScalingInSuspended" yaml:"dynamicScalingInSuspended"`
	// Whether scale out by a target tracking scaling policy or a step scaling policy is suspended.
	//
	// Set the value to `true` if you don't want Application Auto Scaling to add capacity when a scaling policy is triggered. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-suspendedstate.html#cfn-applicationautoscaling-scalabletarget-suspendedstate-dynamicscalingoutsuspended
	//
	DynamicScalingOutSuspended interface{} `field:"optional" json:"dynamicScalingOutSuspended" yaml:"dynamicScalingOutSuspended"`
	// Whether scheduled scaling is suspended.
	//
	// Set the value to `true` if you don't want Application Auto Scaling to add or remove capacity by initiating scheduled actions. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-suspendedstate.html#cfn-applicationautoscaling-scalabletarget-suspendedstate-scheduledscalingsuspended
	//
	ScheduledScalingSuspended interface{} `field:"optional" json:"scheduledScalingSuspended" yaml:"scheduledScalingSuspended"`
}

