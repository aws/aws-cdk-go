package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainConfiguration.
// Experimental.
type IDomainConfigurationRef interface {
	constructs.IConstruct
	// A reference to a DomainConfiguration resource.
	// Experimental.
	DomainConfigurationRef() *DomainConfigurationReference
}

// The jsii proxy for IDomainConfigurationRef
type jsiiProxy_IDomainConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDomainConfigurationRef) DomainConfigurationRef() *DomainConfigurationReference {
	var returns *DomainConfigurationReference
	_jsii_.Get(
		j,
		"domainConfigurationRef",
		&returns,
	)
	return returns
}

