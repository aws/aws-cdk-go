package awsdirectoryservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdirectoryservice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SimpleAD.
// Experimental.
type ISimpleADRef interface {
	constructs.IConstruct
	// A reference to a SimpleAD resource.
	// Experimental.
	SimpleAdRef() *SimpleADReference
}

// The jsii proxy for ISimpleADRef
type jsiiProxy_ISimpleADRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISimpleADRef) SimpleAdRef() *SimpleADReference {
	var returns *SimpleADReference
	_jsii_.Get(
		j,
		"simpleAdRef",
		&returns,
	)
	return returns
}

