package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudFormationProduct.
// Experimental.
type ICloudFormationProductRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICloudFormationProductRef
type jsiiProxy_ICloudFormationProductRef struct {
	internal.Type__constructsIConstruct
}

