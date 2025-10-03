package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudFormationProvisionedProduct.
// Experimental.
type ICloudFormationProvisionedProductRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICloudFormationProvisionedProductRef
type jsiiProxy_ICloudFormationProvisionedProductRef struct {
	internal.Type__constructsIConstruct
}

