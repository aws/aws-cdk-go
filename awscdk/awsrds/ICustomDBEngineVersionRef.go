package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomDBEngineVersion.
// Experimental.
type ICustomDBEngineVersionRef interface {
	constructs.IConstruct
	// A reference to a CustomDBEngineVersion resource.
	// Experimental.
	CustomDbEngineVersionRef() *CustomDBEngineVersionReference
}

// The jsii proxy for ICustomDBEngineVersionRef
type jsiiProxy_ICustomDBEngineVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICustomDBEngineVersionRef) CustomDbEngineVersionRef() *CustomDBEngineVersionReference {
	var returns *CustomDBEngineVersionReference
	_jsii_.Get(
		j,
		"customDbEngineVersionRef",
		&returns,
	)
	return returns
}

