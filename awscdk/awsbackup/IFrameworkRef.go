package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Framework.
// Experimental.
type IFrameworkRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFrameworkRef
type jsiiProxy_IFrameworkRef struct {
	internal.Type__constructsIConstruct
}

