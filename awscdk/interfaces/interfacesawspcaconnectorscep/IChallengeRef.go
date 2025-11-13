package interfacesawspcaconnectorscep

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspcaconnectorscep/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Challenge.
// Experimental.
type IChallengeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Challenge resource.
	// Experimental.
	ChallengeRef() *ChallengeReference
}

// The jsii proxy for IChallengeRef
type jsiiProxy_IChallengeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IChallengeRef) ChallengeRef() *ChallengeReference {
	var returns *ChallengeReference
	_jsii_.Get(
		j,
		"challengeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChallengeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChallengeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

