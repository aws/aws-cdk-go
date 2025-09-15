package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OptionGroup.
// Experimental.
type IOptionGroupRef interface {
	constructs.IConstruct
	// A reference to a OptionGroup resource.
	// Experimental.
	OptionGroupRef() *OptionGroupReference
}

// The jsii proxy for IOptionGroupRef
type jsiiProxy_IOptionGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOptionGroupRef) OptionGroupRef() *OptionGroupReference {
	var returns *OptionGroupReference
	_jsii_.Get(
		j,
		"optionGroupRef",
		&returns,
	)
	return returns
}

