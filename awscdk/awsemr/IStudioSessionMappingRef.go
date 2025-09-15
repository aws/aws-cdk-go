package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StudioSessionMapping.
// Experimental.
type IStudioSessionMappingRef interface {
	constructs.IConstruct
	// A reference to a StudioSessionMapping resource.
	// Experimental.
	StudioSessionMappingRef() *StudioSessionMappingReference
}

// The jsii proxy for IStudioSessionMappingRef
type jsiiProxy_IStudioSessionMappingRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStudioSessionMappingRef) StudioSessionMappingRef() *StudioSessionMappingReference {
	var returns *StudioSessionMappingReference
	_jsii_.Get(
		j,
		"studioSessionMappingRef",
		&returns,
	)
	return returns
}

