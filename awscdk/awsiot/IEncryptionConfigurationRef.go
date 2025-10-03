package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EncryptionConfiguration.
// Experimental.
type IEncryptionConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEncryptionConfigurationRef
type jsiiProxy_IEncryptionConfigurationRef struct {
	internal.Type__constructsIConstruct
}

