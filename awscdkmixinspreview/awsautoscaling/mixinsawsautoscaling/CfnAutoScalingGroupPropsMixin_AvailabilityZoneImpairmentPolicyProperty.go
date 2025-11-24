package mixinsawsautoscaling


// Describes an Availability Zone impairment policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   availabilityZoneImpairmentPolicyProperty := &AvailabilityZoneImpairmentPolicyProperty{
//   	ImpairedZoneHealthCheckBehavior: jsii.String("impairedZoneHealthCheckBehavior"),
//   	ZonalShiftEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-availabilityzoneimpairmentpolicy.html
//
type CfnAutoScalingGroupPropsMixin_AvailabilityZoneImpairmentPolicyProperty struct {
	// Specifies the health check behavior for the impaired Availability Zone in an active zonal shift.
	//
	// If you select `Replace unhealthy` , instances that appear unhealthy will be replaced in all Availability Zones. If you select `Ignore unhealthy` , instances will not be replaced in the Availability Zone with the active zonal shift. For more information, see [Auto Scaling group zonal shift](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-zonal-shift.html) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-availabilityzoneimpairmentpolicy.html#cfn-autoscaling-autoscalinggroup-availabilityzoneimpairmentpolicy-impairedzonehealthcheckbehavior
	//
	ImpairedZoneHealthCheckBehavior *string `field:"optional" json:"impairedZoneHealthCheckBehavior" yaml:"impairedZoneHealthCheckBehavior"`
	// If `true` , enable zonal shift for your Auto Scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-availabilityzoneimpairmentpolicy.html#cfn-autoscaling-autoscalinggroup-availabilityzoneimpairmentpolicy-zonalshiftenabled
	//
	ZonalShiftEnabled interface{} `field:"optional" json:"zonalShiftEnabled" yaml:"zonalShiftEnabled"`
}

