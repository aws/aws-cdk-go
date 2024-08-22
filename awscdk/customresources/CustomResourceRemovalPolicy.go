package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/customresources/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Manages removal policy for AWS-vended custom resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customResourceRemovalPolicy := awscdk.Custom_resources.NewCustomResourceRemovalPolicy(cdk.RemovalPolicy_DESTROY)
//
type CustomResourceRemovalPolicy interface {
	awscdk.IAspect
	// All aspects can visit an IConstruct.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for CustomResourceRemovalPolicy
type jsiiProxy_CustomResourceRemovalPolicy struct {
	internal.Type__awscdkIAspect
}

func NewCustomResourceRemovalPolicy(removalPolicy awscdk.RemovalPolicy) CustomResourceRemovalPolicy {
	_init_.Initialize()

	if err := validateNewCustomResourceRemovalPolicyParameters(removalPolicy); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomResourceRemovalPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.CustomResourceRemovalPolicy",
		[]interface{}{removalPolicy},
		&j,
	)

	return &j
}

func NewCustomResourceRemovalPolicy_Override(c CustomResourceRemovalPolicy, removalPolicy awscdk.RemovalPolicy) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.CustomResourceRemovalPolicy",
		[]interface{}{removalPolicy},
		c,
	)
}

func (c *jsiiProxy_CustomResourceRemovalPolicy) Visit(node constructs.IConstruct) {
	if err := c.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"visit",
		[]interface{}{node},
	)
}

