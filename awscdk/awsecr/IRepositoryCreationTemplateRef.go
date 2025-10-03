package awsecr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RepositoryCreationTemplate.
// Experimental.
type IRepositoryCreationTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRepositoryCreationTemplateRef
type jsiiProxy_IRepositoryCreationTemplateRef struct {
	internal.Type__constructsIConstruct
}

