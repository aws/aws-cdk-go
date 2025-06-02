package awsecs


// > The deployment circuit breaker can only be used for services using the rolling update ( `ECS` ) deployment type.
//
// The *deployment circuit breaker* determines whether a service deployment will fail if the service can't reach a steady state. If it is turned on, a service deployment will transition to a failed state and stop launching new tasks. You can also configure Amazon ECS to roll back your service to the last completed deployment after a failure. For more information, see [Rolling update](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// For more information about API failure reasons, see [API failure reasons](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/api_failures_messages.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentCircuitBreakerProperty := &DeploymentCircuitBreakerProperty{
//   	Enable: jsii.Boolean(false),
//   	Rollback: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentcircuitbreaker.html
//
type CfnService_DeploymentCircuitBreakerProperty struct {
	// Determines whether to use the deployment circuit breaker logic for the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentcircuitbreaker.html#cfn-ecs-service-deploymentcircuitbreaker-enable
	//
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Determines whether to configure Amazon ECS to roll back the service if a service deployment fails.
	//
	// If rollback is on, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentcircuitbreaker.html#cfn-ecs-service-deploymentcircuitbreaker-rollback
	//
	Rollback interface{} `field:"required" json:"rollback" yaml:"rollback"`
}

