package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataRepositoryAssociation.
// Experimental.
type IDataRepositoryAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DataRepositoryAssociation resource.
	// Experimental.
	DataRepositoryAssociationRef() *DataRepositoryAssociationReference
}

// The jsii proxy for IDataRepositoryAssociationRef
type jsiiProxy_IDataRepositoryAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDataRepositoryAssociationRef) DataRepositoryAssociationRef() *DataRepositoryAssociationReference {
	var returns *DataRepositoryAssociationReference
	_jsii_.Get(
		j,
		"dataRepositoryAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataRepositoryAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataRepositoryAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

