package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudFormationProduct.
// Experimental.
type ICloudFormationProductRef interface {
	constructs.IConstruct
	// A reference to a CloudFormationProduct resource.
	// Experimental.
	CloudFormationProductRef() *CloudFormationProductReference
}

// The jsii proxy for ICloudFormationProductRef
type jsiiProxy_ICloudFormationProductRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICloudFormationProductRef) CloudFormationProductRef() *CloudFormationProductReference {
	var returns *CloudFormationProductReference
	_jsii_.Get(
		j,
		"cloudFormationProductRef",
		&returns,
	)
	return returns
}

