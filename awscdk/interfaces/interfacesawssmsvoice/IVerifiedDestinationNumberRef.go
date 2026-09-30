package interfacesawssmsvoice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VerifiedDestinationNumber.
// Experimental.
type IVerifiedDestinationNumberRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VerifiedDestinationNumber resource.
	// Experimental.
	VerifiedDestinationNumberRef() *VerifiedDestinationNumberReference
}

// The jsii proxy for IVerifiedDestinationNumberRef
type jsiiProxy_IVerifiedDestinationNumberRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IVerifiedDestinationNumberRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IVerifiedDestinationNumberRef) VerifiedDestinationNumberRef() *VerifiedDestinationNumberReference {
	var returns *VerifiedDestinationNumberReference
	_jsii_.Get(
		j,
		"verifiedDestinationNumberRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVerifiedDestinationNumberRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVerifiedDestinationNumberRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

