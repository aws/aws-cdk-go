package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolRiskConfigurationAttachment.
// Experimental.
type IUserPoolRiskConfigurationAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a UserPoolRiskConfigurationAttachment resource.
	// Experimental.
	UserPoolRiskConfigurationAttachmentRef() *UserPoolRiskConfigurationAttachmentReference
}

// The jsii proxy for IUserPoolRiskConfigurationAttachmentRef
type jsiiProxy_IUserPoolRiskConfigurationAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IUserPoolRiskConfigurationAttachmentRef) UserPoolRiskConfigurationAttachmentRef() *UserPoolRiskConfigurationAttachmentReference {
	var returns *UserPoolRiskConfigurationAttachmentReference
	_jsii_.Get(
		j,
		"userPoolRiskConfigurationAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolRiskConfigurationAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolRiskConfigurationAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

