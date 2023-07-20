package awscodedeploy


// A configuration that shifts traffic from one version of a Lambda function or ECS task set to another in equal increments, with an equal number of minutes between each increment.
//
// The original and target Lambda function versions or ECS task sets are specified in the deployment's AppSpec file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeBasedLinearProperty := &TimeBasedLinearProperty{
//   	LinearInterval: jsii.Number(123),
//   	LinearPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-timebasedlinear.html
//
type CfnDeploymentConfig_TimeBasedLinearProperty struct {
	// The number of minutes between each incremental traffic shift of a `TimeBasedLinear` deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-timebasedlinear.html#cfn-codedeploy-deploymentconfig-timebasedlinear-linearinterval
	//
	LinearInterval *float64 `field:"required" json:"linearInterval" yaml:"linearInterval"`
	// The percentage of traffic that is shifted at the start of each increment of a `TimeBasedLinear` deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-timebasedlinear.html#cfn-codedeploy-deploymentconfig-timebasedlinear-linearpercentage
	//
	LinearPercentage *float64 `field:"required" json:"linearPercentage" yaml:"linearPercentage"`
}

