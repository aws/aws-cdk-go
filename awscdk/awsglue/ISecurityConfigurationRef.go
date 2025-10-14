package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityConfiguration.
// Experimental.
type ISecurityConfigurationRef interface {
	constructs.IConstruct
	// A reference to a SecurityConfiguration resource.
	// Experimental.
	SecurityConfigurationRef() *SecurityConfigurationReference
}

// The jsii proxy for ISecurityConfigurationRef
type jsiiProxy_ISecurityConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISecurityConfigurationRef) SecurityConfigurationRef() *SecurityConfigurationReference {
	var returns *SecurityConfigurationReference
	_jsii_.Get(
		j,
		"securityConfigurationRef",
		&returns,
	)
	return returns
}

