package awsecs


// Configuration for forcing a new deployment of the service.
//
// Example:
//   var cluster Cluster
//   var taskDefinition TaskDefinition
//
//
//   // Force a new deployment on every `cdk deploy` by using a time-based nonce
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	ForceNewDeployment: &ForceNewDeployment{
//   		Enabled: jsii.Boolean(true),
//   		Nonce: date.now().toString(),
//   	},
//   })
//
type ForceNewDeployment struct {
	// Whether to enable the force-new-deployment mechanism for the service.
	//
	// Setting this to `true` enables the mechanism, but on its own it does not
	// force a new deployment on every `cdk deploy`: CloudFormation only starts a
	// new deployment when it detects a change in the template, and the signal for
	// that is the `nonce` value changing between deployments. If `nonce` is not
	// provided or its value stays the same across deployments, no new deployment
	// is forced. When set to `false`, the `ForceNewDeployment` property is rendered
	// with `EnableForceNewDeployment: false`.
	//
	// To force a new deployment on every `cdk deploy`, provide a `nonce` with a
	// unique, time-varying value such as a timestamp, random string, or sequence
	// number (e.g. `Date.now().toString()`).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-service-forcenewdeployment.html
	//
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// A unique nonce value that signals Amazon ECS to start a new deployment.
	//
	// When you change this value, it triggers a new deployment even though no
	// other service parameters have changed. Use a stable, time-varying value
	// like a commit hash, image digest, or version string.
	//
	// If not provided and `enabled` is `true`, only `EnableForceNewDeployment`
	// is set without a nonce.
	//
	// Must be between 1 and 255 characters.
	// Default: - no nonce.
	//
	Nonce *string `field:"optional" json:"nonce" yaml:"nonce"`
}

