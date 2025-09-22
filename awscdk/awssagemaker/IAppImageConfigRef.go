package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppImageConfig.
// Experimental.
type IAppImageConfigRef interface {
	constructs.IConstruct
	// A reference to a AppImageConfig resource.
	// Experimental.
	AppImageConfigRef() *AppImageConfigReference
}

// The jsii proxy for IAppImageConfigRef
type jsiiProxy_IAppImageConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAppImageConfigRef) AppImageConfigRef() *AppImageConfigReference {
	var returns *AppImageConfigReference
	_jsii_.Get(
		j,
		"appImageConfigRef",
		&returns,
	)
	return returns
}

