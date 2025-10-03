package awsiotanalytics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Datastore.
// Experimental.
type IDatastoreRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDatastoreRef
type jsiiProxy_IDatastoreRef struct {
	internal.Type__constructsIConstruct
}

