package awscodedeploy


// A configuration that shifts traffic from one version of a Lambda function or Amazon ECS task set to another in two increments.
//
// The original and target Lambda function versions or ECS task sets are specified in the deployment's AppSpec file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeBasedCanaryProperty := &TimeBasedCanaryProperty{
//   	CanaryInterval: jsii.Number(123),
//   	CanaryPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-timebasedcanary.html
//
type CfnDeploymentConfig_TimeBasedCanaryProperty struct {
	// The number of minutes between the first and second traffic shifts of a `TimeBasedCanary` deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-timebasedcanary.html#cfn-codedeploy-deploymentconfig-timebasedcanary-canaryinterval
	//
	CanaryInterval *float64 `field:"required" json:"canaryInterval" yaml:"canaryInterval"`
	// The percentage of traffic to shift in the first increment of a `TimeBasedCanary` deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-timebasedcanary.html#cfn-codedeploy-deploymentconfig-timebasedcanary-canarypercentage
	//
	CanaryPercentage *float64 `field:"required" json:"canaryPercentage" yaml:"canaryPercentage"`
}

