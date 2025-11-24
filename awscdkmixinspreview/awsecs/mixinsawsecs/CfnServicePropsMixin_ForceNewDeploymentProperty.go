package mixinsawsecs


// Determines whether to force a new deployment of the service.
//
// By default, deployments aren't forced. You can use this option to start a new deployment with no service definition changes. For example, you can update a service's tasks to use a newer Docker image with the same image/tag combination ( `my_image:latest` ) or to roll Fargate tasks onto a newer platform version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   forceNewDeploymentProperty := &ForceNewDeploymentProperty{
//   	EnableForceNewDeployment: jsii.Boolean(false),
//   	ForceNewDeploymentNonce: jsii.String("forceNewDeploymentNonce"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-forcenewdeployment.html
//
type CfnServicePropsMixin_ForceNewDeploymentProperty struct {
	// Determines whether to force a new deployment of the service.
	//
	// By default, deployments aren't forced. You can use this option to start a new deployment with no service definition changes. For example, you can update a service's tasks to use a newer Docker image with the same image/tag combination ( `my_image:latest` ) or to roll Fargate tasks onto a newer platform version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-forcenewdeployment.html#cfn-ecs-service-forcenewdeployment-enableforcenewdeployment
	//
	EnableForceNewDeployment interface{} `field:"optional" json:"enableForceNewDeployment" yaml:"enableForceNewDeployment"`
	// When you change the `ForceNewDeploymentNonce` value in your template, it signals Amazon ECS to start a new deployment even though no other service parameters have changed.
	//
	// The value must be a unique, time- varying value like a timestamp, random string, or sequence number. Use this property when you want to ensure your tasks pick up the latest version of a Docker image that uses the same tag but has been updated in the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-forcenewdeployment.html#cfn-ecs-service-forcenewdeployment-forcenewdeploymentnonce
	//
	ForceNewDeploymentNonce *string `field:"optional" json:"forceNewDeploymentNonce" yaml:"forceNewDeploymentNonce"`
}

