package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdMappingTable.
// Experimental.
type IIdMappingTableRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIdMappingTableRef
type jsiiProxy_IIdMappingTableRef struct {
	internal.Type__constructsIConstruct
}

