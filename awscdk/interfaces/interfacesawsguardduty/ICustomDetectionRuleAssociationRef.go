package interfacesawsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomDetectionRuleAssociation.
// Experimental.
type ICustomDetectionRuleAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CustomDetectionRuleAssociation resource.
	// Experimental.
	CustomDetectionRuleAssociationRef() *CustomDetectionRuleAssociationReference
}

// The jsii proxy for ICustomDetectionRuleAssociationRef
type jsiiProxy_ICustomDetectionRuleAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICustomDetectionRuleAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ICustomDetectionRuleAssociationRef) CustomDetectionRuleAssociationRef() *CustomDetectionRuleAssociationReference {
	var returns *CustomDetectionRuleAssociationReference
	_jsii_.Get(
		j,
		"customDetectionRuleAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomDetectionRuleAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomDetectionRuleAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

