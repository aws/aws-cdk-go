package awspcaconnectorad

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspcaconnectorad/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TemplateGroupAccessControlEntry.
// Experimental.
type ITemplateGroupAccessControlEntryRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TemplateGroupAccessControlEntry resource.
	// Experimental.
	TemplateGroupAccessControlEntryRef() *TemplateGroupAccessControlEntryReference
}

// The jsii proxy for ITemplateGroupAccessControlEntryRef
type jsiiProxy_ITemplateGroupAccessControlEntryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITemplateGroupAccessControlEntryRef) TemplateGroupAccessControlEntryRef() *TemplateGroupAccessControlEntryReference {
	var returns *TemplateGroupAccessControlEntryReference
	_jsii_.Get(
		j,
		"templateGroupAccessControlEntryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateGroupAccessControlEntryRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateGroupAccessControlEntryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

