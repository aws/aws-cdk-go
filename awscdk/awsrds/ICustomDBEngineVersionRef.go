package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomDBEngineVersion.
// Experimental.
type ICustomDBEngineVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICustomDBEngineVersionRef
type jsiiProxy_ICustomDBEngineVersionRef struct {
	internal.Type__constructsIConstruct
}

