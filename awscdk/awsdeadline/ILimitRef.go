package awsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Limit.
// Experimental.
type ILimitRef interface {
	constructs.IConstruct
	// A reference to a Limit resource.
	// Experimental.
	LimitRef() *LimitReference
}

// The jsii proxy for ILimitRef
type jsiiProxy_ILimitRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILimitRef) LimitRef() *LimitReference {
	var returns *LimitReference
	_jsii_.Get(
		j,
		"limitRef",
		&returns,
	)
	return returns
}

