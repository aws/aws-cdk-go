package awsecs


// The deployment controller to use for the service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentControllerProperty := &DeploymentControllerProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentcontroller.html
//
type CfnService_DeploymentControllerProperty struct {
	// The deployment controller type to use. There are three deployment controller types available:.
	//
	// - **ECS** - The rolling update ( `ECS` ) deployment type involves replacing the current running version of the container with the latest version. The number of containers Amazon ECS adds or removes from the service during a rolling update is controlled by adjusting the minimum and maximum number of healthy tasks allowed during a service deployment, as specified in the [DeploymentConfiguration](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeploymentConfiguration.html) .
	// - **CODE_DEPLOY** - The blue/green ( `CODE_DEPLOY` ) deployment type uses the blue/green deployment model powered by AWS CodeDeploy , which allows you to verify a new deployment of a service before sending production traffic to it.
	// - **EXTERNAL** - The external ( `EXTERNAL` ) deployment type enables you to use any third-party deployment controller for full control over the deployment process for an Amazon ECS service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentcontroller.html#cfn-ecs-service-deploymentcontroller-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

