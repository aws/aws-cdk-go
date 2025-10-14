package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BrowserCustom.
// Experimental.
type IBrowserCustomRef interface {
	constructs.IConstruct
	// A reference to a BrowserCustom resource.
	// Experimental.
	BrowserCustomRef() *BrowserCustomReference
}

// The jsii proxy for IBrowserCustomRef
type jsiiProxy_IBrowserCustomRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBrowserCustomRef) BrowserCustomRef() *BrowserCustomReference {
	var returns *BrowserCustomReference
	_jsii_.Get(
		j,
		"browserCustomRef",
		&returns,
	)
	return returns
}

