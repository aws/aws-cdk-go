package mixinsawsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsapigateway/mixinsawsapigateway/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGateway::Method` resource creates API Gateway methods that define the parameters and body that clients must send in their requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMethodPropsMixin := awscdkmixinspreview.Mixins.NewCfnMethodPropsMixin(&CfnMethodMixinProps{
//   	ApiKeyRequired: jsii.Boolean(false),
//   	AuthorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	AuthorizationType: jsii.String("authorizationType"),
//   	AuthorizerId: jsii.String("authorizerId"),
//   	HttpMethod: jsii.String("httpMethod"),
//   	Integration: &IntegrationProperty{
//   		CacheKeyParameters: []*string{
//   			jsii.String("cacheKeyParameters"),
//   		},
//   		CacheNamespace: jsii.String("cacheNamespace"),
//   		ConnectionId: jsii.String("connectionId"),
//   		ConnectionType: jsii.String("connectionType"),
//   		ContentHandling: jsii.String("contentHandling"),
//   		Credentials: jsii.String("credentials"),
//   		IntegrationHttpMethod: jsii.String("integrationHttpMethod"),
//   		IntegrationResponses: []interface{}{
//   			&IntegrationResponseProperty{
//   				ContentHandling: jsii.String("contentHandling"),
//   				ResponseParameters: map[string]*string{
//   					"responseParametersKey": jsii.String("responseParameters"),
//   				},
//   				ResponseTemplates: map[string]*string{
//   					"responseTemplatesKey": jsii.String("responseTemplates"),
//   				},
//   				SelectionPattern: jsii.String("selectionPattern"),
//   				StatusCode: jsii.String("statusCode"),
//   			},
//   		},
//   		IntegrationTarget: jsii.String("integrationTarget"),
//   		PassthroughBehavior: jsii.String("passthroughBehavior"),
//   		RequestParameters: map[string]*string{
//   			"requestParametersKey": jsii.String("requestParameters"),
//   		},
//   		RequestTemplates: map[string]*string{
//   			"requestTemplatesKey": jsii.String("requestTemplates"),
//   		},
//   		ResponseTransferMode: jsii.String("responseTransferMode"),
//   		TimeoutInMillis: jsii.Number(123),
//   		Type: jsii.String("type"),
//   		Uri: jsii.String("uri"),
//   	},
//   	MethodResponses: []interface{}{
//   		&MethodResponseProperty{
//   			ResponseModels: map[string]*string{
//   				"responseModelsKey": jsii.String("responseModels"),
//   			},
//   			ResponseParameters: map[string]interface{}{
//   				"responseParametersKey": jsii.Boolean(false),
//   			},
//   			StatusCode: jsii.String("statusCode"),
//   		},
//   	},
//   	OperationName: jsii.String("operationName"),
//   	RequestModels: map[string]*string{
//   		"requestModelsKey": jsii.String("requestModels"),
//   	},
//   	RequestParameters: map[string]interface{}{
//   		"requestParametersKey": jsii.Boolean(false),
//   	},
//   	RequestValidatorId: jsii.String("requestValidatorId"),
//   	ResourceId: jsii.String("resourceId"),
//   	RestApiId: jsii.String("restApiId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html
//
type CfnMethodPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMethodMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMethodPropsMixin
type jsiiProxy_CfnMethodPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMethodPropsMixin) Props() *CfnMethodMixinProps {
	var returns *CfnMethodMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethodPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApiGateway::Method`.
func NewCfnMethodPropsMixin(props *CfnMethodMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMethodPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMethodPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMethodPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnMethodPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApiGateway::Method`.
func NewCfnMethodPropsMixin_Override(c CfnMethodPropsMixin, props *CfnMethodMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnMethodPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMethodPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMethodPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnMethodPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMethodPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnMethodPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMethodPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMethodPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

