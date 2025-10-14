package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmacie/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AllowList.
// Experimental.
type IAllowListRef interface {
	constructs.IConstruct
	// A reference to a AllowList resource.
	// Experimental.
	AllowListRef() *AllowListReference
}

// The jsii proxy for IAllowListRef
type jsiiProxy_IAllowListRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAllowListRef) AllowListRef() *AllowListReference {
	var returns *AllowListReference
	_jsii_.Get(
		j,
		"allowListRef",
		&returns,
	)
	return returns
}

