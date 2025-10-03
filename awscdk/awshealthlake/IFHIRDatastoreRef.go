package awshealthlake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awshealthlake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FHIRDatastore.
// Experimental.
type IFHIRDatastoreRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFHIRDatastoreRef
type jsiiProxy_IFHIRDatastoreRef struct {
	internal.Type__constructsIConstruct
}

