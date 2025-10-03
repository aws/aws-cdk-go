package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataAutomationProject.
// Experimental.
type IDataAutomationProjectRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataAutomationProjectRef
type jsiiProxy_IDataAutomationProjectRef struct {
	internal.Type__constructsIConstruct
}

