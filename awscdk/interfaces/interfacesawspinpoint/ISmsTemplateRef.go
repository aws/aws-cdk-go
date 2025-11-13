package interfacesawspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SmsTemplate.
// Experimental.
type ISmsTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SmsTemplate resource.
	// Experimental.
	SmsTemplateRef() *SmsTemplateReference
}

// The jsii proxy for ISmsTemplateRef
type jsiiProxy_ISmsTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ISmsTemplateRef) SmsTemplateRef() *SmsTemplateReference {
	var returns *SmsTemplateReference
	_jsii_.Get(
		j,
		"smsTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISmsTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISmsTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

