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
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   deploymentCircuitBreakerProperty := &DeploymentCircuitBreakerProperty{
//   	Enable: jsii.Boolean(false),
//   	ResetOnHealthyTask: jsii.Boolean(false),
//   	Rollback: jsii.Boolean(false),
//   	ThresholdConfiguration: &ThresholdConfigurationProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentcircuitbreaker.html
//
type CfnServicePropsMixin_DeploymentCircuitBreakerProperty struct {
	// Determines whether to use the deployment circuit breaker logic for the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentcircuitbreaker.html#cfn-ecs-service-deploymentcircuitbreaker-enable
	//
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Specifies whether the deployment circuit breaker resets its failure count when a task reaches a healthy state.
	//
	// When set to ``true``, a task that reaches a healthy state resets the failure count to ``0``. When set to ``false``, Amazon ECS does not reset the failure count. The default is ``true``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentcircuitbreaker.html#cfn-ecs-service-deploymentcircuitbreaker-resetonhealthytask
	//
	ResetOnHealthyTask interface{} `field:"optional" json:"resetOnHealthyTask" yaml:"resetOnHealthyTask"`
	// Determines whether to configure Amazon ECS to roll back the service if a service deployment fails.
	//
	// If rollback is on, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentcircuitbreaker.html#cfn-ecs-service-deploymentcircuitbreaker-rollback
	//
	Rollback interface{} `field:"optional" json:"rollback" yaml:"rollback"`
	// Defines the failure threshold that the deployment circuit breaker uses to monitor a deployment.
	//
	// The ``type`` and ``value`` together determine the number of task failures that are tolerated before the circuit breaker triggers.
	//  By default, the threshold configuration uses a ``type`` of ``BOUNDED_PERCENT`` with a ``value`` of ``50``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentcircuitbreaker.html#cfn-ecs-service-deploymentcircuitbreaker-thresholdconfiguration
	//
	ThresholdConfiguration interface{} `field:"optional" json:"thresholdConfiguration" yaml:"thresholdConfiguration"`
}

