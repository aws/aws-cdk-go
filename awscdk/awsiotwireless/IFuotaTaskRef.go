package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FuotaTask.
// Experimental.
type IFuotaTaskRef interface {
	constructs.IConstruct
	// A reference to a FuotaTask resource.
	// Experimental.
	FuotaTaskRef() *FuotaTaskReference
}

// The jsii proxy for IFuotaTaskRef
type jsiiProxy_IFuotaTaskRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFuotaTaskRef) FuotaTaskRef() *FuotaTaskReference {
	var returns *FuotaTaskReference
	_jsii_.Get(
		j,
		"fuotaTaskRef",
		&returns,
	)
	return returns
}

