package previewawsecsmixins


// A deployment lifecycle hook runs custom logic at specific stages of the deployment process.
//
// Currently, you can use Lambda functions as hook targets.
//
// For more information, see [Lifecycle hooks for Amazon ECS service deployments](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-lifecycle-hooks.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var hookDetails interface{}
//
//   deploymentLifecycleHookProperty := &DeploymentLifecycleHookProperty{
//   	HookDetails: hookDetails,
//   	HookTargetArn: jsii.String("hookTargetArn"),
//   	LifecycleStages: []*string{
//   		jsii.String("lifecycleStages"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentlifecyclehook.html
//
type CfnServicePropsMixin_DeploymentLifecycleHookProperty struct {
	// Use this field to specify custom parameters that Amazon ECS passes to your hook target invocations (such as a Lambda function).
	//
	// This field must be a JSON object as a string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentlifecyclehook.html#cfn-ecs-service-deploymentlifecyclehook-hookdetails
	//
	HookDetails interface{} `field:"optional" json:"hookDetails" yaml:"hookDetails"`
	// The Amazon Resource Name (ARN) of the hook target. Currently, only Lambda function ARNs are supported.
	//
	// You must provide this parameter when configuring a deployment lifecycle hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentlifecyclehook.html#cfn-ecs-service-deploymentlifecyclehook-hooktargetarn
	//
	HookTargetArn *string `field:"optional" json:"hookTargetArn" yaml:"hookTargetArn"`
	// The lifecycle stages at which to run the hook. Choose from these valid values:.
	//
	// - RECONCILE_SERVICE
	//
	// The reconciliation stage that only happens when you start a new service deployment with more than 1 service revision in an ACTIVE state.
	//
	// You can use a lifecycle hook for this stage.
	// - PRE_SCALE_UP
	//
	// The green service revision has not started. The blue service revision is handling 100% of the production traffic. There is no test traffic.
	//
	// You can use a lifecycle hook for this stage.
	// - POST_SCALE_UP
	//
	// The green service revision has started. The blue service revision is handling 100% of the production traffic. There is no test traffic.
	//
	// You can use a lifecycle hook for this stage.
	// - TEST_TRAFFIC_SHIFT
	//
	// The blue and green service revisions are running. The blue service revision handles 100% of the production traffic. The green service revision is migrating from 0% to 100% of test traffic.
	//
	// You can use a lifecycle hook for this stage.
	// - POST_TEST_TRAFFIC_SHIFT
	//
	// The test traffic shift is complete. The green service revision handles 100% of the test traffic.
	//
	// You can use a lifecycle hook for this stage.
	// - PRODUCTION_TRAFFIC_SHIFT
	//
	// Production traffic is shifting to the green service revision. The green service revision is migrating from 0% to 100% of production traffic.
	//
	// You can use a lifecycle hook for this stage.
	// - POST_PRODUCTION_TRAFFIC_SHIFT
	//
	// The production traffic shift is complete.
	//
	// You can use a lifecycle hook for this stage.
	//
	// You must provide this parameter when configuring a deployment lifecycle hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentlifecyclehook.html#cfn-ecs-service-deploymentlifecyclehook-lifecyclestages
	//
	LifecycleStages *[]*string `field:"optional" json:"lifecycleStages" yaml:"lifecycleStages"`
	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon ECS permission to call Lambda functions on your behalf.
	//
	// For more information, see [Permissions required for Lambda functions in Amazon ECS blue/green deployments](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/blue-green-permissions.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentlifecyclehook.html#cfn-ecs-service-deploymentlifecyclehook-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

