package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KeyGroup.
// Experimental.
type IKeyGroupRef interface {
	constructs.IConstruct
	// A reference to a KeyGroup resource.
	// Experimental.
	KeyGroupRef() *KeyGroupReference
}

// The jsii proxy for IKeyGroupRef
type jsiiProxy_IKeyGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IKeyGroupRef) KeyGroupRef() *KeyGroupReference {
	var returns *KeyGroupReference
	_jsii_.Get(
		j,
		"keyGroupRef",
		&returns,
	)
	return returns
}

