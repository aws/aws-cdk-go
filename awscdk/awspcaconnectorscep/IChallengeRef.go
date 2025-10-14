package awspcaconnectorscep

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspcaconnectorscep/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Challenge.
// Experimental.
type IChallengeRef interface {
	constructs.IConstruct
	// A reference to a Challenge resource.
	// Experimental.
	ChallengeRef() *ChallengeReference
}

// The jsii proxy for IChallengeRef
type jsiiProxy_IChallengeRef struct {
	internal.Type__constructsIConstruct
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

