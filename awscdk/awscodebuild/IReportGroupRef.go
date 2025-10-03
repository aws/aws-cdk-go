package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReportGroup.
// Experimental.
type IReportGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReportGroupRef
type jsiiProxy_IReportGroupRef struct {
	internal.Type__constructsIConstruct
}

