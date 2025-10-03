package awsemr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceGroupConfig.
// Experimental.
type IInstanceGroupConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInstanceGroupConfigRef
type jsiiProxy_IInstanceGroupConfigRef struct {
	internal.Type__constructsIConstruct
}

