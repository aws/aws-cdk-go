package awsbcmdataexports

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbcmdataexports/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Export.
// Experimental.
type IExportRef interface {
	constructs.IConstruct
}

// The jsii proxy for IExportRef
type jsiiProxy_IExportRef struct {
	internal.Type__constructsIConstruct
}

