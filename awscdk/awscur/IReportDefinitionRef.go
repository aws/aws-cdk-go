package awscur

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscur/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReportDefinition.
// Experimental.
type IReportDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReportDefinitionRef
type jsiiProxy_IReportDefinitionRef struct {
	internal.Type__constructsIConstruct
}

