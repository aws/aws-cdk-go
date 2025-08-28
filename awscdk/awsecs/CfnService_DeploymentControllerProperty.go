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
	// The deployment controller type to use.
	//
	// The deployment controller is the mechanism that determines how tasks are deployed for your service. The valid options are:
	//
	// - ECS
	//
	// When you create a service which uses the `ECS` deployment controller, you can choose between the following deployment strategies:
	//
	// - `ROLLING` : When you create a service which uses the *rolling update* ( `ROLLING` ) deployment strategy, the Amazon ECS service scheduler replaces the currently running tasks with new tasks. The number of tasks that Amazon ECS adds or removes from the service during a rolling update is controlled by the service deployment configuration.
	//
	// Rolling update deployments are best suited for the following scenarios:
	//
	// - Gradual service updates: You need to update your service incrementally without taking the entire service offline at once.
	// - Limited resource requirements: You want to avoid the additional resource costs of running two complete environments simultaneously (as required by blue/green deployments).
	// - Acceptable deployment time: Your application can tolerate a longer deployment process, as rolling updates replace tasks one by one.
	// - No need for instant roll back: Your service can tolerate a rollback process that takes minutes rather than seconds.
	// - Simple deployment process: You prefer a straightforward deployment approach without the complexity of managing multiple environments, target groups, and listeners.
	// - No load balancer requirement: Your service doesn't use or require a load balancer, Application Load Balancer , Network Load Balancer , or Service Connect (which are required for blue/green deployments).
	// - Stateful applications: Your application maintains state that makes it difficult to run two parallel environments.
	// - Cost sensitivity: You want to minimize deployment costs by not running duplicate environments during deployment.
	//
	// Rolling updates are the default deployment strategy for services and provide a balance between deployment safety and resource efficiency for many common application scenarios.
	// - `BLUE_GREEN` : A *blue/green* deployment strategy ( `BLUE_GREEN` ) is a release methodology that reduces downtime and risk by running two identical production environments called blue and green. With Amazon ECS blue/green deployments, you can validate new service revisions before directing production traffic to them. This approach provides a safer way to deploy changes with the ability to quickly roll back if needed.
	//
	// Amazon ECS blue/green deployments are best suited for the following scenarios:
	//
	// - Service validation: When you need to validate new service revisions before directing production traffic to them
	// - Zero downtime: When your service requires zero-downtime deployments
	// - Instant roll back: When you need the ability to quickly roll back if issues are detected
	// - Load balancer requirement: When your service uses Application Load Balancer , Network Load Balancer , or Service Connect
	// - External
	//
	// Use a third-party deployment controller.
	// - Blue/green deployment (powered by CodeDeploy )
	//
	// CodeDeploy installs an updated version of the application as a new replacement task set and reroutes production traffic from the original application task set to the replacement task set. The original task set is terminated after a successful deployment. Use this deployment controller to verify a new deployment of a service before sending production traffic to it.
	//
	// When updating the deployment controller for a service, consider the following depending on the type of migration you're performing.
	//
	// - If you have a template that contains the `EXTERNAL` deployment controller information as well as `TaskSet` and `PrimaryTaskSet` resources, and you remove the task set resources from the template when updating from `EXTERNAL` to `ECS` , the `DescribeTaskSet` and `DeleteTaskSet` API calls will return a 400 error after the deployment controller is updated to `ECS` . This results in a delete failure on the task set resources, even though the stack transitions to `UPDATE_COMPLETE` status. For more information, see [Resource removed from stack but not deleted](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/troubleshooting.html#troubleshooting-errors-resource-removed-not-deleted) in the AWS CloudFormation User Guide. To fix this issue, delete the task sets directly using the Amazon ECS `DeleteTaskSet` API. For more information about how to delete a task set, see [DeleteTaskSet](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeleteTaskSet.html) in the Amazon Elastic Container Service API Reference.
	// - If you're migrating from `CODE_DEPLOY` to `ECS` with a new task definition and AWS CloudFormation performs a rollback operation, the Amazon ECS `UpdateService` request fails with the following error:
	//
	// Resource handler returned message: "Invalid request provided: Unable to update task definition on services with a CODE_DEPLOY deployment controller.
	// - After a successful migration from `ECS` to `EXTERNAL` deployment controller, you need to manually remove the `ACTIVE` task set, because Amazon ECS no longer manages the deployment. For information about how to delete a task set, see [DeleteTaskSet](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeleteTaskSet.html) in the Amazon Elastic Container Service API Reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentcontroller.html#cfn-ecs-service-deploymentcontroller-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

