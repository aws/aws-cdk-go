package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerRelay.
// Experimental.
type IMailManagerRelayRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MailManagerRelay resource.
	// Experimental.
	MailManagerRelayRef() *MailManagerRelayReference
}

// The jsii proxy for IMailManagerRelayRef
type jsiiProxy_IMailManagerRelayRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMailManagerRelayRef) MailManagerRelayRef() *MailManagerRelayReference {
	var returns *MailManagerRelayReference
	_jsii_.Get(
		j,
		"mailManagerRelayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMailManagerRelayRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMailManagerRelayRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

