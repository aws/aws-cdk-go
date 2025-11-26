package previewawscodedeploymixins


// The configuration that specifies how traffic is shifted from one version of a Lambda function to another version during an AWS Lambda deployment, or from one Amazon ECS task set to another during an Amazon ECS deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   trafficRoutingConfigProperty := &TrafficRoutingConfigProperty{
//   	TimeBasedCanary: &TimeBasedCanaryProperty{
//   		CanaryInterval: jsii.Number(123),
//   		CanaryPercentage: jsii.Number(123),
//   	},
//   	TimeBasedLinear: &TimeBasedLinearProperty{
//   		LinearInterval: jsii.Number(123),
//   		LinearPercentage: jsii.Number(123),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-trafficroutingconfig.html
//
type CfnDeploymentConfigPropsMixin_TrafficRoutingConfigProperty struct {
	// A configuration that shifts traffic from one version of a Lambda function or ECS task set to another in two increments.
	//
	// The original and target Lambda function versions or ECS task sets are specified in the deployment's AppSpec file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-trafficroutingconfig.html#cfn-codedeploy-deploymentconfig-trafficroutingconfig-timebasedcanary
	//
	TimeBasedCanary interface{} `field:"optional" json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// A configuration that shifts traffic from one version of a Lambda function or Amazon ECS task set to another in equal increments, with an equal number of minutes between each increment.
	//
	// The original and target Lambda function versions or Amazon ECS task sets are specified in the deployment's AppSpec file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-trafficroutingconfig.html#cfn-codedeploy-deploymentconfig-trafficroutingconfig-timebasedlinear
	//
	TimeBasedLinear interface{} `field:"optional" json:"timeBasedLinear" yaml:"timeBasedLinear"`
	// The type of traffic shifting ( `TimeBasedCanary` or `TimeBasedLinear` ) used by a deployment configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-trafficroutingconfig.html#cfn-codedeploy-deploymentconfig-trafficroutingconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

