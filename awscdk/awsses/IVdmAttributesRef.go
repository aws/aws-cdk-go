package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VdmAttributes.
// Experimental.
type IVdmAttributesRef interface {
	constructs.IConstruct
	// A reference to a VdmAttributes resource.
	// Experimental.
	VdmAttributesRef() *VdmAttributesReference
}

// The jsii proxy for IVdmAttributesRef
type jsiiProxy_IVdmAttributesRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVdmAttributesRef) VdmAttributesRef() *VdmAttributesReference {
	var returns *VdmAttributesReference
	_jsii_.Get(
		j,
		"vdmAttributesRef",
		&returns,
	)
	return returns
}

