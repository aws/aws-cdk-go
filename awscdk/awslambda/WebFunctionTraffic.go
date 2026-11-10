package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// How a WebFunction endpoint splits its traffic between revisions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webFunctionTraffic := awscdk.Aws_lambda.WebFunctionTraffic_Single(jsii.String("revisionId"))
//
type WebFunctionTraffic interface {
}

// The jsii proxy struct for WebFunctionTraffic
type jsiiProxy_WebFunctionTraffic struct {
	_ byte // padding
}

// Send all of the endpoint's traffic to one revision.
func WebFunctionTraffic_Single(revisionId *string) WebFunctionTraffic {
	_init_.Initialize()

	if err := validateWebFunctionTraffic_SingleParameters(revisionId); err != nil {
		panic(err)
	}
	var returns WebFunctionTraffic

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.WebFunctionTraffic",
		"single",
		[]interface{}{revisionId},
		&returns,
	)

	return returns
}

