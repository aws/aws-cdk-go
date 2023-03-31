package awsecs


// The deployment circuit breaker to use for the service.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	circuitBreaker: &deploymentCircuitBreaker{
//   		rollback: jsii.Boolean(true),
//   	},
//   })
//
// Experimental.
type DeploymentCircuitBreaker struct {
	// Whether to enable rollback on deployment failure.
	// Experimental.
	Rollback *bool `field:"optional" json:"rollback" yaml:"rollback"`
}

