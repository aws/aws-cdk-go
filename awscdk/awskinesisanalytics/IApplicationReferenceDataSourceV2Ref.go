package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationReferenceDataSource.
// Deprecated: use `aws-kinesisanalyticsv2` instead.
type IApplicationReferenceDataSourceV2Ref interface {
	constructs.IConstruct
	// A reference to a ApplicationReferenceDataSource resource.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ApplicationReferenceDataSourceRef() *ApplicationReferenceDataSourceV2Reference
}

// The jsii proxy for IApplicationReferenceDataSourceV2Ref
type jsiiProxy_IApplicationReferenceDataSourceV2Ref struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApplicationReferenceDataSourceV2Ref) ApplicationReferenceDataSourceRef() *ApplicationReferenceDataSourceV2Reference {
	var returns *ApplicationReferenceDataSourceV2Reference
	_jsii_.Get(
		j,
		"applicationReferenceDataSourceRef",
		&returns,
	)
	return returns
}

