package awsiottwinmaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SyncJob.
// Experimental.
type ISyncJobRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISyncJobRef
type jsiiProxy_ISyncJobRef struct {
	internal.Type__constructsIConstruct
}

