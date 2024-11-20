package awsautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availabilityZoneImpairmentPolicyProperty := &AvailabilityZoneImpairmentPolicyProperty{
//   	ImpairedZoneHealthCheckBehavior: jsii.String("impairedZoneHealthCheckBehavior"),
//   	ZonalShiftEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-availabilityzoneimpairmentpolicy.html
//
type CfnAutoScalingGroup_AvailabilityZoneImpairmentPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-availabilityzoneimpairmentpolicy.html#cfn-autoscaling-autoscalinggroup-availabilityzoneimpairmentpolicy-impairedzonehealthcheckbehavior
	//
	ImpairedZoneHealthCheckBehavior *string `field:"required" json:"impairedZoneHealthCheckBehavior" yaml:"impairedZoneHealthCheckBehavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-availabilityzoneimpairmentpolicy.html#cfn-autoscaling-autoscalinggroup-availabilityzoneimpairmentpolicy-zonalshiftenabled
	//
	ZonalShiftEnabled interface{} `field:"required" json:"zonalShiftEnabled" yaml:"zonalShiftEnabled"`
}

