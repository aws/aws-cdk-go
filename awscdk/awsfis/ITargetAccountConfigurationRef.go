package awsfis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsfis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TargetAccountConfiguration.
// Experimental.
type ITargetAccountConfigurationRef interface {
	constructs.IConstruct
	// A reference to a TargetAccountConfiguration resource.
	// Experimental.
	TargetAccountConfigurationRef() *TargetAccountConfigurationReference
}

// The jsii proxy for ITargetAccountConfigurationRef
type jsiiProxy_ITargetAccountConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITargetAccountConfigurationRef) TargetAccountConfigurationRef() *TargetAccountConfigurationReference {
	var returns *TargetAccountConfigurationReference
	_jsii_.Get(
		j,
		"targetAccountConfigurationRef",
		&returns,
	)
	return returns
}

