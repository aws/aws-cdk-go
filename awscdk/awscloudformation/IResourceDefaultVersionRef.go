package awscloudformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceDefaultVersion.
// Experimental.
type IResourceDefaultVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceDefaultVersionRef
type jsiiProxy_IResourceDefaultVersionRef struct {
	internal.Type__constructsIConstruct
}

