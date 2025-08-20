package awscdksagemakeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
	"github.com/aws/aws-cdk-go/awscdksagemakeralpha/v2/internal"
)

// The Interface for a SageMaker Endpoint resource.
// Experimental.
type IEndpoint interface {
	awssagemaker.IEndpoint
}

// The jsii proxy for IEndpoint
type jsiiProxy_IEndpoint struct {
	internal.Type__awssagemakerIEndpoint
}

