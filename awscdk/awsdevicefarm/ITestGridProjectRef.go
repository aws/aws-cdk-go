package awsdevicefarm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TestGridProject.
// Experimental.
type ITestGridProjectRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITestGridProjectRef
type jsiiProxy_ITestGridProjectRef struct {
	internal.Type__constructsIConstruct
}

