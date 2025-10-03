package awspcaconnectorad

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspcaconnectorad/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectoryRegistration.
// Experimental.
type IDirectoryRegistrationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDirectoryRegistrationRef
type jsiiProxy_IDirectoryRegistrationRef struct {
	internal.Type__constructsIConstruct
}

