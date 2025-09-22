package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupProfile.
// Experimental.
type IGroupProfileRef interface {
	constructs.IConstruct
	// A reference to a GroupProfile resource.
	// Experimental.
	GroupProfileRef() *GroupProfileReference
}

// The jsii proxy for IGroupProfileRef
type jsiiProxy_IGroupProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGroupProfileRef) GroupProfileRef() *GroupProfileReference {
	var returns *GroupProfileReference
	_jsii_.Get(
		j,
		"groupProfileRef",
		&returns,
	)
	return returns
}

