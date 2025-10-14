package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workteam.
// Experimental.
type IWorkteamRef interface {
	constructs.IConstruct
	// A reference to a Workteam resource.
	// Experimental.
	WorkteamRef() *WorkteamReference
}

// The jsii proxy for IWorkteamRef
type jsiiProxy_IWorkteamRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IWorkteamRef) WorkteamRef() *WorkteamReference {
	var returns *WorkteamReference
	_jsii_.Get(
		j,
		"workteamRef",
		&returns,
	)
	return returns
}

