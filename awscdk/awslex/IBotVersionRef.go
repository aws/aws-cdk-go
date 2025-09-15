package awslex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslex/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BotVersion.
// Experimental.
type IBotVersionRef interface {
	constructs.IConstruct
	// A reference to a BotVersion resource.
	// Experimental.
	BotVersionRef() *BotVersionReference
}

// The jsii proxy for IBotVersionRef
type jsiiProxy_IBotVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBotVersionRef) BotVersionRef() *BotVersionReference {
	var returns *BotVersionReference
	_jsii_.Get(
		j,
		"botVersionRef",
		&returns,
	)
	return returns
}

