package interfacesawspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailChannel.
// Experimental.
type IEmailChannelRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EmailChannel resource.
	// Experimental.
	EmailChannelRef() *EmailChannelReference
}

// The jsii proxy for IEmailChannelRef
type jsiiProxy_IEmailChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEmailChannelRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IEmailChannelRef) EmailChannelRef() *EmailChannelReference {
	var returns *EmailChannelReference
	_jsii_.Get(
		j,
		"emailChannelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailChannelRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

