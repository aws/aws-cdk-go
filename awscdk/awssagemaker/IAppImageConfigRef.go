package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppImageConfig.
// Experimental.
type IAppImageConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAppImageConfigRef
type jsiiProxy_IAppImageConfigRef struct {
	internal.Type__constructsIConstruct
}

