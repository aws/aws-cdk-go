package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MultiRegionAccessPointPolicy.
// Experimental.
type IMultiRegionAccessPointPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MultiRegionAccessPointPolicy resource.
	// Experimental.
	MultiRegionAccessPointPolicyRef() *MultiRegionAccessPointPolicyReference
}

// The jsii proxy for IMultiRegionAccessPointPolicyRef
type jsiiProxy_IMultiRegionAccessPointPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMultiRegionAccessPointPolicyRef) MultiRegionAccessPointPolicyRef() *MultiRegionAccessPointPolicyReference {
	var returns *MultiRegionAccessPointPolicyReference
	_jsii_.Get(
		j,
		"multiRegionAccessPointPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiRegionAccessPointPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiRegionAccessPointPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

