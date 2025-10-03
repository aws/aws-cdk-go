package awsopsworks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Volume.
// Experimental.
type IVolumeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVolumeRef
type jsiiProxy_IVolumeRef struct {
	internal.Type__constructsIConstruct
}

