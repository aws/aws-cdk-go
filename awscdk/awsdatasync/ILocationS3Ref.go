package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationS3.
// Experimental.
type ILocationS3Ref interface {
	constructs.IConstruct
}

// The jsii proxy for ILocationS3Ref
type jsiiProxy_ILocationS3Ref struct {
	internal.Type__constructsIConstruct
}

