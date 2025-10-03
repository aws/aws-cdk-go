package awskinesisanalytics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationReferenceDataSource.
// Experimental.
type IApplicationReferenceDataSourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApplicationReferenceDataSourceRef
type jsiiProxy_IApplicationReferenceDataSourceRef struct {
	internal.Type__constructsIConstruct
}

