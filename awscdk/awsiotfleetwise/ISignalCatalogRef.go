package awsiotfleetwise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SignalCatalog.
// Experimental.
type ISignalCatalogRef interface {
	constructs.IConstruct
	// A reference to a SignalCatalog resource.
	// Experimental.
	SignalCatalogRef() *SignalCatalogReference
}

// The jsii proxy for ISignalCatalogRef
type jsiiProxy_ISignalCatalogRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISignalCatalogRef) SignalCatalogRef() *SignalCatalogReference {
	var returns *SignalCatalogReference
	_jsii_.Get(
		j,
		"signalCatalogRef",
		&returns,
	)
	return returns
}

