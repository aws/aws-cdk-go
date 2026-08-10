package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataLakeAssociation.
// Experimental.
type IDataLakeAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataLakeAssociation resource.
	// Experimental.
	DataLakeAssociationRef() *DataLakeAssociationReference
}

// The jsii proxy for IDataLakeAssociationRef
type jsiiProxy_IDataLakeAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDataLakeAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDataLakeAssociationRef) DataLakeAssociationRef() *DataLakeAssociationReference {
	var returns *DataLakeAssociationReference
	_jsii_.Get(
		j,
		"dataLakeAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataLakeAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataLakeAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

