package awsquicksight


// The liner axis scale setup.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   axisLinearScaleProperty := &AxisLinearScaleProperty{
//   	StepCount: jsii.Number(123),
//   	StepSize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axislinearscale.html
//
type CfnDashboard_AxisLinearScaleProperty struct {
	// The step count setup of a linear axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axislinearscale.html#cfn-quicksight-dashboard-axislinearscale-stepcount
	//
	StepCount *float64 `field:"optional" json:"stepCount" yaml:"stepCount"`
	// The step size setup of a linear axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axislinearscale.html#cfn-quicksight-dashboard-axislinearscale-stepsize
	//
	StepSize *float64 `field:"optional" json:"stepSize" yaml:"stepSize"`
}

