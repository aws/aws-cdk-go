package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodedeploy"
)

// A DeploymentConfig that can specialize itself based on the target group it will be used for.
//
// For example, this is used for AWS-managed deployment configs: these are already
// present in every region, but we need a region-specific ARN to reference them.
// Since we might use them in conjunction with cross-region DeploymentGroups, we
// need to specialize the account and region to the DeploymentGroup before
// using.
//
// A DeploymentGroup must call `bindEnvironment()` first if it detects this type,
// before reading the DeploymentConfig ARN.
type IBindableDeploymentConfig interface {
	interfacesawscodedeploy.IDeploymentConfigRef
	// Bind the predefined deployment config to the environment of the given resource.
	BindEnvironment(deploymentGroup interfacesawscodedeploy.IDeploymentGroupRef) interfacesawscodedeploy.IDeploymentConfigRef
}

// The jsii proxy for IBindableDeploymentConfig
type jsiiProxy_IBindableDeploymentConfig struct {
	internal.Type__interfacesawscodedeployIDeploymentConfigRef
}

func (i *jsiiProxy_IBindableDeploymentConfig) BindEnvironment(deploymentGroup interfacesawscodedeploy.IDeploymentGroupRef) interfacesawscodedeploy.IDeploymentConfigRef {
	if err := i.validateBindEnvironmentParameters(deploymentGroup); err != nil {
		panic(err)
	}
	var returns interfacesawscodedeploy.IDeploymentConfigRef

	_jsii_.Invoke(
		i,
		"bindEnvironment",
		[]interface{}{deploymentGroup},
		&returns,
	)

	return returns
}

