package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BatchScramSecret.
// Experimental.
type IBatchScramSecretRef interface {
	constructs.IConstruct
	// A reference to a BatchScramSecret resource.
	// Experimental.
	BatchScramSecretRef() *BatchScramSecretReference
}

// The jsii proxy for IBatchScramSecretRef
type jsiiProxy_IBatchScramSecretRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBatchScramSecretRef) BatchScramSecretRef() *BatchScramSecretReference {
	var returns *BatchScramSecretReference
	_jsii_.Get(
		j,
		"batchScramSecretRef",
		&returns,
	)
	return returns
}

