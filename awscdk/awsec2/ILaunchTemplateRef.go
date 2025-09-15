package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchTemplate.
// Experimental.
type ILaunchTemplateRef interface {
	constructs.IConstruct
	// A reference to a LaunchTemplate resource.
	// Experimental.
	LaunchTemplateRef() *LaunchTemplateReference
}

// The jsii proxy for ILaunchTemplateRef
type jsiiProxy_ILaunchTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILaunchTemplateRef) LaunchTemplateRef() *LaunchTemplateReference {
	var returns *LaunchTemplateReference
	_jsii_.Get(
		j,
		"launchTemplateRef",
		&returns,
	)
	return returns
}

