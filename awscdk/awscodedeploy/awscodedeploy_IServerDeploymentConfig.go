package awscodedeploy


// The Deployment Configuration of an EC2/on-premise Deployment Group.
//
// The default, pre-defined Configurations are available as constants on the {@link ServerDeploymentConfig} class
// (`ServerDeploymentConfig.HALF_AT_A_TIME`, `ServerDeploymentConfig.ALL_AT_ONCE`, etc.).
// To create a custom Deployment Configuration,
// instantiate the {@link ServerDeploymentConfig} Construct.
type IServerDeploymentConfig interface {
	IBaseDeploymentConfig
}

// The jsii proxy for IServerDeploymentConfig
type jsiiProxy_IServerDeploymentConfig struct {
	jsiiProxy_IBaseDeploymentConfig
}

