package interfacesawsrolesanywhere

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrolesanywhere/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrustAnchor.
// Experimental.
type ITrustAnchorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TrustAnchor resource.
	// Experimental.
	TrustAnchorRef() *TrustAnchorReference
}

// The jsii proxy for ITrustAnchorRef
type jsiiProxy_ITrustAnchorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITrustAnchorRef) TrustAnchorRef() *TrustAnchorReference {
	var returns *TrustAnchorReference
	_jsii_.Get(
		j,
		"trustAnchorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrustAnchorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrustAnchorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

