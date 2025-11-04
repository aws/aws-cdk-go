package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CidrCollection.
// Experimental.
type ICidrCollectionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CidrCollection resource.
	// Experimental.
	CidrCollectionRef() *CidrCollectionReference
}

// The jsii proxy for ICidrCollectionRef
type jsiiProxy_ICidrCollectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICidrCollectionRef) CidrCollectionRef() *CidrCollectionReference {
	var returns *CidrCollectionReference
	_jsii_.Get(
		j,
		"cidrCollectionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICidrCollectionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICidrCollectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

