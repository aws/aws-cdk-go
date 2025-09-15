package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Application.
// Experimental.
type IApplicationV2Ref interface {
	constructs.IConstruct
	// A reference to a Application resource.
	// Experimental.
	ApplicationRef() *ApplicationV2Reference
}

// The jsii proxy for IApplicationV2Ref
type jsiiProxy_IApplicationV2Ref struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApplicationV2Ref) ApplicationRef() *ApplicationV2Reference {
	var returns *ApplicationV2Reference
	_jsii_.Get(
		j,
		"applicationRef",
		&returns,
	)
	return returns
}

