package interfacesawsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudFormationProduct.
// Experimental.
type ICloudFormationProductRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CloudFormationProduct resource.
	// Experimental.
	CloudFormationProductRef() *CloudFormationProductReference
}

// The jsii proxy for ICloudFormationProductRef
type jsiiProxy_ICloudFormationProductRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ICloudFormationProductRef) CloudFormationProductRef() *CloudFormationProductReference {
	var returns *CloudFormationProductReference
	_jsii_.Get(
		j,
		"cloudFormationProductRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudFormationProductRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudFormationProductRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

