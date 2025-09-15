package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationOutput.
// Experimental.
type IApplicationOutputV2Ref interface {
	constructs.IConstruct
	// A reference to a ApplicationOutput resource.
	// Experimental.
	ApplicationOutputRef() *ApplicationOutputV2Reference
}

// The jsii proxy for IApplicationOutputV2Ref
type jsiiProxy_IApplicationOutputV2Ref struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApplicationOutputV2Ref) ApplicationOutputRef() *ApplicationOutputV2Reference {
	var returns *ApplicationOutputV2Reference
	_jsii_.Get(
		j,
		"applicationOutputRef",
		&returns,
	)
	return returns
}

