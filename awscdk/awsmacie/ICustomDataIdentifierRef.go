package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmacie/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomDataIdentifier.
// Experimental.
type ICustomDataIdentifierRef interface {
	constructs.IConstruct
	// A reference to a CustomDataIdentifier resource.
	// Experimental.
	CustomDataIdentifierRef() *CustomDataIdentifierReference
}

// The jsii proxy for ICustomDataIdentifierRef
type jsiiProxy_ICustomDataIdentifierRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICustomDataIdentifierRef) CustomDataIdentifierRef() *CustomDataIdentifierReference {
	var returns *CustomDataIdentifierReference
	_jsii_.Get(
		j,
		"customDataIdentifierRef",
		&returns,
	)
	return returns
}

