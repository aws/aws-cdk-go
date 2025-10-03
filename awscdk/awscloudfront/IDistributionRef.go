package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Distribution.
// Experimental.
type IDistributionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDistributionRef
type jsiiProxy_IDistributionRef struct {
	internal.Type__constructsIConstruct
}

