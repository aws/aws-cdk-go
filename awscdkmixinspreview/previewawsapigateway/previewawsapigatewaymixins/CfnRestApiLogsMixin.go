package previewawsapigatewaymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsapigateway/previewawsapigatewaymixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGateway::RestApi` resource creates a REST API.
//
// For more information, see [restapi:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateRestApi.html) in the *Amazon API Gateway REST API Reference* .
//
// > On January 1, 2016, the Swagger Specification was donated to the [OpenAPI initiative](https://docs.aws.amazon.com/https://www.openapis.org/) , becoming the foundation of the OpenAPI Specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnRestApiLogsMixin := awscdkmixinspreview.Mixins.NewCfnRestApiLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
//
type CfnRestApiLogsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRestApiLogsMixin
type jsiiProxy_CfnRestApiLogsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRestApiLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApiLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::ApiGateway::RestApi`.
func NewCfnRestApiLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnRestApiLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnRestApiLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRestApiLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::ApiGateway::RestApi`.
func NewCfnRestApiLogsMixin_Override(c CfnRestApiLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRestApiLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRestApiLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRestApiLogsMixin_ACCESS_LOGS() CfnRestApiAccessLogs {
	_init_.Initialize()
	var returns CfnRestApiAccessLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiLogsMixin",
		"ACCESS_LOGS",
		&returns,
	)
	return returns
}

func CfnRestApiLogsMixin_EXECUTION_LOGS() CfnRestApiExecutionLogs {
	_init_.Initialize()
	var returns CfnRestApiExecutionLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiLogsMixin",
		"EXECUTION_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRestApiLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRestApiLogsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

