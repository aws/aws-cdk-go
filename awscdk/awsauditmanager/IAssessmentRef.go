package awsauditmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsauditmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Assessment.
// Experimental.
type IAssessmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAssessmentRef
type jsiiProxy_IAssessmentRef struct {
	internal.Type__constructsIConstruct
}

