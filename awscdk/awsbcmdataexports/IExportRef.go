package awsbcmdataexports

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbcmdataexports/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Export.
// Experimental.
type IExportRef interface {
	constructs.IConstruct
	// A reference to a Export resource.
	// Experimental.
	ExportRef() *ExportReference
}

// The jsii proxy for IExportRef
type jsiiProxy_IExportRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IExportRef) ExportRef() *ExportReference {
	var returns *ExportReference
	_jsii_.Get(
		j,
		"exportRef",
		&returns,
	)
	return returns
}

