package interfacesawscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FieldLevelEncryptionConfig.
// Experimental.
type IFieldLevelEncryptionConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FieldLevelEncryptionConfig resource.
	// Experimental.
	FieldLevelEncryptionConfigRef() *FieldLevelEncryptionConfigReference
}

// The jsii proxy for IFieldLevelEncryptionConfigRef
type jsiiProxy_IFieldLevelEncryptionConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IFieldLevelEncryptionConfigRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IFieldLevelEncryptionConfigRef) FieldLevelEncryptionConfigRef() *FieldLevelEncryptionConfigReference {
	var returns *FieldLevelEncryptionConfigReference
	_jsii_.Get(
		j,
		"fieldLevelEncryptionConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFieldLevelEncryptionConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFieldLevelEncryptionConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

