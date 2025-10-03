package awsemr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StudioSessionMapping.
// Experimental.
type IStudioSessionMappingRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStudioSessionMappingRef
type jsiiProxy_IStudioSessionMappingRef struct {
	internal.Type__constructsIConstruct
}

