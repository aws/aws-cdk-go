package awsecr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PublicRepository.
// Experimental.
type IPublicRepositoryRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPublicRepositoryRef
type jsiiProxy_IPublicRepositoryRef struct {
	internal.Type__constructsIConstruct
}

