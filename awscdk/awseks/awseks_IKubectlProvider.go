package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awseks/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Imported KubectlProvider that can be used in place of the default one created by CDK.
// Experimental.
type IKubectlProvider interface {
	awscdk.IConstruct
	// The IAM execution role of the handler.
	// Experimental.
	HandlerRole() awsiam.IRole
	// The IAM role to assume in order to perform kubectl operations against this cluster.
	// Experimental.
	RoleArn() *string
	// The custom resource provider's service token.
	// Experimental.
	ServiceToken() *string
}

// The jsii proxy for IKubectlProvider
type jsiiProxy_IKubectlProvider struct {
	internal.Type__awscdkIConstruct
}

func (j *jsiiProxy_IKubectlProvider) HandlerRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"handlerRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKubectlProvider) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKubectlProvider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

