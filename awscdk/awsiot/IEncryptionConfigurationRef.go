package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EncryptionConfiguration.
// Experimental.
type IEncryptionConfigurationRef interface {
	constructs.IConstruct
	// A reference to a EncryptionConfiguration resource.
	// Experimental.
	EncryptionConfigurationRef() *EncryptionConfigurationReference
}

// The jsii proxy for IEncryptionConfigurationRef
type jsiiProxy_IEncryptionConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEncryptionConfigurationRef) EncryptionConfigurationRef() *EncryptionConfigurationReference {
	var returns *EncryptionConfigurationReference
	_jsii_.Get(
		j,
		"encryptionConfigurationRef",
		&returns,
	)
	return returns
}

