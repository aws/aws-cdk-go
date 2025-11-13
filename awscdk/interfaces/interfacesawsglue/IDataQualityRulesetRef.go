package interfacesawsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataQualityRuleset.
// Experimental.
type IDataQualityRulesetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataQualityRuleset resource.
	// Experimental.
	DataQualityRulesetRef() *DataQualityRulesetReference
}

// The jsii proxy for IDataQualityRulesetRef
type jsiiProxy_IDataQualityRulesetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDataQualityRulesetRef) DataQualityRulesetRef() *DataQualityRulesetReference {
	var returns *DataQualityRulesetReference
	_jsii_.Get(
		j,
		"dataQualityRulesetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataQualityRulesetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataQualityRulesetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

