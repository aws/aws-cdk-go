package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataAutomationProject.
// Experimental.
type IDataAutomationProjectRef interface {
	constructs.IConstruct
	// A reference to a DataAutomationProject resource.
	// Experimental.
	DataAutomationProjectRef() *DataAutomationProjectReference
}

// The jsii proxy for IDataAutomationProjectRef
type jsiiProxy_IDataAutomationProjectRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDataAutomationProjectRef) DataAutomationProjectRef() *DataAutomationProjectReference {
	var returns *DataAutomationProjectReference
	_jsii_.Get(
		j,
		"dataAutomationProjectRef",
		&returns,
	)
	return returns
}

