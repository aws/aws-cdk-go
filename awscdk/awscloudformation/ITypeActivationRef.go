package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TypeActivation.
// Experimental.
type ITypeActivationRef interface {
	constructs.IConstruct
	// A reference to a TypeActivation resource.
	// Experimental.
	TypeActivationRef() *TypeActivationReference
}

// The jsii proxy for ITypeActivationRef
type jsiiProxy_ITypeActivationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITypeActivationRef) TypeActivationRef() *TypeActivationReference {
	var returns *TypeActivationReference
	_jsii_.Get(
		j,
		"typeActivationRef",
		&returns,
	)
	return returns
}

