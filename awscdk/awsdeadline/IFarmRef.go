package awsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Farm.
// Experimental.
type IFarmRef interface {
	constructs.IConstruct
	// A reference to a Farm resource.
	// Experimental.
	FarmRef() *FarmReference
}

// The jsii proxy for IFarmRef
type jsiiProxy_IFarmRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFarmRef) FarmRef() *FarmReference {
	var returns *FarmReference
	_jsii_.Get(
		j,
		"farmRef",
		&returns,
	)
	return returns
}

