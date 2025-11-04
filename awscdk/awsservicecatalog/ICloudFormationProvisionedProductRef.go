package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudFormationProvisionedProduct.
// Experimental.
type ICloudFormationProvisionedProductRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CloudFormationProvisionedProduct resource.
	// Experimental.
	CloudFormationProvisionedProductRef() *CloudFormationProvisionedProductReference
}

// The jsii proxy for ICloudFormationProvisionedProductRef
type jsiiProxy_ICloudFormationProvisionedProductRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICloudFormationProvisionedProductRef) CloudFormationProvisionedProductRef() *CloudFormationProvisionedProductReference {
	var returns *CloudFormationProvisionedProductReference
	_jsii_.Get(
		j,
		"cloudFormationProvisionedProductRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudFormationProvisionedProductRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudFormationProvisionedProductRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

