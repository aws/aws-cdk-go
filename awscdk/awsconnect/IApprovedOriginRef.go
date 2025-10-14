package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApprovedOrigin.
// Experimental.
type IApprovedOriginRef interface {
	constructs.IConstruct
	// A reference to a ApprovedOrigin resource.
	// Experimental.
	ApprovedOriginRef() *ApprovedOriginReference
}

// The jsii proxy for IApprovedOriginRef
type jsiiProxy_IApprovedOriginRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApprovedOriginRef) ApprovedOriginRef() *ApprovedOriginReference {
	var returns *ApprovedOriginReference
	_jsii_.Get(
		j,
		"approvedOriginRef",
		&returns,
	)
	return returns
}

