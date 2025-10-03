package awscloudtrail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudtrail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Trail.
// Experimental.
type ITrailRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITrailRef
type jsiiProxy_ITrailRef struct {
	internal.Type__constructsIConstruct
}

