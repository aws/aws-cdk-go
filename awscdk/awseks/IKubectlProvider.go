package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Imported KubectlProvider that can be used in place of the default one created by CDK.
type IKubectlProvider interface {
	constructs.IConstruct
	// The IAM execution role of the handler.
	HandlerRole() awsiam.IRole
	// The IAM role to assume in order to perform kubectl operations against this cluster.
	RoleArn() *string
	// The custom resource provider's service token.
	ServiceToken() *string
}

// The jsii proxy for IKubectlProvider
type jsiiProxy_IKubectlProvider struct {
	internal.Type__constructsIConstruct
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

