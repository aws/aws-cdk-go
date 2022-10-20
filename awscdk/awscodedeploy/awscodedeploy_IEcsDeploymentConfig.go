package awscodedeploy


// The Deployment Configuration of an ECS Deployment Group.
//
// If you're managing the Deployment Configuration alongside the rest of your CDK resources,
// use the {@link EcsDeploymentConfig} class.
//
// If you want to reference an already existing deployment configuration,
// or one defined in a different CDK Stack,
// use the {@link EcsDeploymentConfig#fromEcsDeploymentConfigName} method.
//
// The default, pre-defined Configurations are available as constants on the {@link EcsDeploymentConfig} class
// (for example, `EcsDeploymentConfig.AllAtOnce`).
type IEcsDeploymentConfig interface {
	IBaseDeploymentConfig
}

// The jsii proxy for IEcsDeploymentConfig
type jsiiProxy_IEcsDeploymentConfig struct {
	jsiiProxy_IBaseDeploymentConfig
}

