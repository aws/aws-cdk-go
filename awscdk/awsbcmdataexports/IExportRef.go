package awsbcmdataexports

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbcmdataexports/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Export.
// Experimental.
type IExportRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Export resource.
	// Experimental.
	ExportRef() *ExportReference
}

// The jsii proxy for IExportRef
type jsiiProxy_IExportRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IExportRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExportRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

