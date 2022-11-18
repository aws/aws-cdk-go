package awscodedeploy


// The Deployment Configuration of a Lambda Deployment Group.
//
// If you're managing the Deployment Configuration alongside the rest of your CDK resources,
// use the {@link LambdaDeploymentConfig} class.
//
// If you want to reference an already existing deployment configuration,
// or one defined in a different CDK Stack,
// use the {@link LambdaDeploymentConfig#fromLambdaDeploymentConfigName} method.
//
// The default, pre-defined Configurations are available as constants on the {@link LambdaDeploymentConfig} class
// (`LambdaDeploymentConfig.AllAtOnce`, `LambdaDeploymentConfig.Canary10Percent30Minutes`, etc.).
type ILambdaDeploymentConfig interface {
	IBaseDeploymentConfig
}

// The jsii proxy for ILambdaDeploymentConfig
type jsiiProxy_ILambdaDeploymentConfig struct {
	jsiiProxy_IBaseDeploymentConfig
}

