package awscdkeks-v2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkeks-v2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Imported KubectlProvider that can be used in place of the default one created by CDK.
// Experimental.
type IKubectlProvider interface {
	constructs.IConstruct
	// The role of the provider lambda function.
	//
	// If undefined,
	// you cannot use this provider to deploy helm charts.
	// Experimental.
	Role() awsiam.IRole
	// The custom resource provider's service token.
	// Experimental.
	ServiceToken() *string
}

// The jsii proxy for IKubectlProvider
type jsiiProxy_IKubectlProvider struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IKubectlProvider) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
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

