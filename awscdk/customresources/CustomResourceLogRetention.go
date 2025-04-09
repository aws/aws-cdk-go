package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/customresources/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Manages log retention for AWS-vended custom resources.
//
// This feature is currently experimental.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customResourceLogRetention := awscdk.Custom_resources.NewCustomResourceLogRetention(awscdk.Aws_logs.RetentionDays_ONE_DAY)
//
type CustomResourceLogRetention interface {
	awscdk.IAspect
	// All aspects can visit an IConstruct.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for CustomResourceLogRetention
type jsiiProxy_CustomResourceLogRetention struct {
	internal.Type__awscdkIAspect
}

func NewCustomResourceLogRetention(setLogRetention awslogs.RetentionDays) CustomResourceLogRetention {
	_init_.Initialize()

	if err := validateNewCustomResourceLogRetentionParameters(setLogRetention); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomResourceLogRetention{}

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.CustomResourceLogRetention",
		[]interface{}{setLogRetention},
		&j,
	)

	return &j
}

func NewCustomResourceLogRetention_Override(c CustomResourceLogRetention, setLogRetention awslogs.RetentionDays) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.CustomResourceLogRetention",
		[]interface{}{setLogRetention},
		c,
	)
}

func (c *jsiiProxy_CustomResourceLogRetention) Visit(node constructs.IConstruct) {
	if err := c.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"visit",
		[]interface{}{node},
	)
}

