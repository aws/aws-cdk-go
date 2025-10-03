package awspcaconnectorscep

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspcaconnectorscep/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Challenge.
// Experimental.
type IChallengeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IChallengeRef
type jsiiProxy_IChallengeRef struct {
	internal.Type__constructsIConstruct
}

