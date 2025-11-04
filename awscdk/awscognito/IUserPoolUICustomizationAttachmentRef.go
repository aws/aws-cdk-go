package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolUICustomizationAttachment.
// Experimental.
type IUserPoolUICustomizationAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a UserPoolUICustomizationAttachment resource.
	// Experimental.
	UserPoolUiCustomizationAttachmentRef() *UserPoolUICustomizationAttachmentReference
}

// The jsii proxy for IUserPoolUICustomizationAttachmentRef
type jsiiProxy_IUserPoolUICustomizationAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IUserPoolUICustomizationAttachmentRef) UserPoolUiCustomizationAttachmentRef() *UserPoolUICustomizationAttachmentReference {
	var returns *UserPoolUICustomizationAttachmentReference
	_jsii_.Get(
		j,
		"userPoolUiCustomizationAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolUICustomizationAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolUICustomizationAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

