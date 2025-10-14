package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationReferenceDataSource.
// Experimental.
type IApplicationReferenceDataSourceRef interface {
	constructs.IConstruct
	// A reference to a ApplicationReferenceDataSource resource.
	// Experimental.
	ApplicationReferenceDataSourceRef() *ApplicationReferenceDataSourceReference
}

// The jsii proxy for IApplicationReferenceDataSourceRef
type jsiiProxy_IApplicationReferenceDataSourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApplicationReferenceDataSourceRef) ApplicationReferenceDataSourceRef() *ApplicationReferenceDataSourceReference {
	var returns *ApplicationReferenceDataSourceReference
	_jsii_.Get(
		j,
		"applicationReferenceDataSourceRef",
		&returns,
	)
	return returns
}

