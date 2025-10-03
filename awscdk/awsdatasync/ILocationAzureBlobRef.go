package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationAzureBlob.
// Experimental.
type ILocationAzureBlobRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILocationAzureBlobRef
type jsiiProxy_ILocationAzureBlobRef struct {
	internal.Type__constructsIConstruct
}

