package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
)

// Represents a reference to a CodeDeploy Application deploying to EC2/on-premise instances.
//
// If you're managing the Application alongside the rest of your CDK resources,
// use the `ServerApplication` class.
//
// If you want to reference an already existing Application,
// or one defined in a different CDK Stack,
// use the `#fromServerApplicationName` method.
type IServerApplication interface {
	awscdk.IResource
	ApplicationArn() *string
	ApplicationName() *string
}

// The jsii proxy for IServerApplication
type jsiiProxy_IServerApplication struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IServerApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

