package awselasticbeanstalk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticbeanstalk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationVersion.
// Experimental.
type IApplicationVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApplicationVersionRef
type jsiiProxy_IApplicationVersionRef struct {
	internal.Type__constructsIConstruct
}

