package awsmediatailor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VodSource.
// Experimental.
type IVodSourceRef interface {
	constructs.IConstruct
	// A reference to a VodSource resource.
	// Experimental.
	VodSourceRef() *VodSourceReference
}

// The jsii proxy for IVodSourceRef
type jsiiProxy_IVodSourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVodSourceRef) VodSourceRef() *VodSourceReference {
	var returns *VodSourceReference
	_jsii_.Get(
		j,
		"vodSourceRef",
		&returns,
	)
	return returns
}

