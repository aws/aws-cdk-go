package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodedeploy/internal"
)

// Represents a reference to a CodeDeploy Application deploying to AWS Lambda.
//
// If you're managing the Application alongside the rest of your CDK resources,
// use the {@link LambdaApplication} class.
//
// If you want to reference an already existing Application,
// or one defined in a different CDK Stack,
// use the {@link LambdaApplication#fromLambdaApplicationName} method.
// Experimental.
type ILambdaApplication interface {
	awscdk.IResource
	// Experimental.
	ApplicationArn() *string
	// Experimental.
	ApplicationName() *string
}

// The jsii proxy for ILambdaApplication
type jsiiProxy_ILambdaApplication struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ILambdaApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILambdaApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

