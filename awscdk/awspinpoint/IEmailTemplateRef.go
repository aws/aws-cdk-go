package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailTemplate.
// Experimental.
type IEmailTemplateRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EmailTemplate resource.
	// Experimental.
	EmailTemplateRef() *EmailTemplateReference
}

// The jsii proxy for IEmailTemplateRef
type jsiiProxy_IEmailTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEmailTemplateRef) EmailTemplateRef() *EmailTemplateReference {
	var returns *EmailTemplateReference
	_jsii_.Get(
		j,
		"emailTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailTemplateRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

