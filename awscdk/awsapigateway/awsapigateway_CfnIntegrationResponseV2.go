package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::ApiGatewayV2::IntegrationResponse`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var responseParameters interface{}
//   var responseTemplates interface{}
//
//   cfnIntegrationResponseV2 := awscdk.Aws_apigateway.NewCfnIntegrationResponseV2(this, jsii.String("MyCfnIntegrationResponseV2"), &CfnIntegrationResponseV2Props{
//   	ApiId: jsii.String("apiId"),
//   	IntegrationId: jsii.String("integrationId"),
//   	IntegrationResponseKey: jsii.String("integrationResponseKey"),
//
//   	// the properties below are optional
//   	ContentHandlingStrategy: jsii.String("contentHandlingStrategy"),
//   	ResponseParameters: responseParameters,
//   	ResponseTemplates: responseTemplates,
//   	TemplateSelectionExpression: jsii.String("templateSelectionExpression"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnIntegrationResponseV2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::ApiGatewayV2::IntegrationResponse.ApiId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-apiid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiId() *string
	// Deprecated: moved to package aws-apigatewayv2.
	SetApiId(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Deprecated: moved to package aws-apigatewayv2.
	CfnOptions() awscdk.ICfnResourceOptions
	// Deprecated: moved to package aws-apigatewayv2.
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Deprecated: moved to package aws-apigatewayv2.
	CfnResourceType() *string
	// `AWS::ApiGatewayV2::IntegrationResponse.ContentHandlingStrategy`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-contenthandlingstrategy
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ContentHandlingStrategy() *string
	// Deprecated: moved to package aws-apigatewayv2.
	SetContentHandlingStrategy(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Deprecated: moved to package aws-apigatewayv2.
	CreationStack() *[]*string
	// `AWS::ApiGatewayV2::IntegrationResponse.IntegrationId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-integrationid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	IntegrationId() *string
	// Deprecated: moved to package aws-apigatewayv2.
	SetIntegrationId(val *string)
	// `AWS::ApiGatewayV2::IntegrationResponse.IntegrationResponseKey`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-integrationresponsekey
	//
	// Deprecated: moved to package aws-apigatewayv2.
	IntegrationResponseKey() *string
	// Deprecated: moved to package aws-apigatewayv2.
	SetIntegrationResponseKey(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Deprecated: moved to package aws-apigatewayv2.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Deprecated: moved to package aws-apigatewayv2.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Deprecated: moved to package aws-apigatewayv2.
	Ref() *string
	// `AWS::ApiGatewayV2::IntegrationResponse.ResponseParameters`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-responseparameters
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ResponseParameters() interface{}
	// Deprecated: moved to package aws-apigatewayv2.
	SetResponseParameters(val interface{})
	// `AWS::ApiGatewayV2::IntegrationResponse.ResponseTemplates`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-responsetemplates
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ResponseTemplates() interface{}
	// Deprecated: moved to package aws-apigatewayv2.
	SetResponseTemplates(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Deprecated: moved to package aws-apigatewayv2.
	Stack() awscdk.Stack
	// `AWS::ApiGatewayV2::IntegrationResponse.TemplateSelectionExpression`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-templateselectionexpression
	//
	// Deprecated: moved to package aws-apigatewayv2.
	TemplateSelectionExpression() *string
	// Deprecated: moved to package aws-apigatewayv2.
	SetTemplateSelectionExpression(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Deprecated: moved to package aws-apigatewayv2.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Deprecated: moved to package aws-apigatewayv2.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Deprecated: moved to package aws-apigatewayv2.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Deprecated: moved to package aws-apigatewayv2.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Deprecated: moved to package aws-apigatewayv2.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Deprecated: moved to package aws-apigatewayv2.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Deprecated: moved to package aws-apigatewayv2.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Deprecated: moved to package aws-apigatewayv2.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Deprecated: moved to package aws-apigatewayv2.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	// Deprecated: moved to package aws-apigatewayv2.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: moved to package aws-apigatewayv2.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: moved to package aws-apigatewayv2.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: moved to package aws-apigatewayv2.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Deprecated: moved to package aws-apigatewayv2.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: moved to package aws-apigatewayv2.
	Prepare()
	// Deprecated: moved to package aws-apigatewayv2.
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Deprecated: moved to package aws-apigatewayv2.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: moved to package aws-apigatewayv2.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Deprecated: moved to package aws-apigatewayv2.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: moved to package aws-apigatewayv2.
	Validate() *[]*string
	// Deprecated: moved to package aws-apigatewayv2.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnIntegrationResponseV2
type jsiiProxy_CfnIntegrationResponseV2 struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIntegrationResponseV2) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) ContentHandlingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) IntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) IntegrationResponseKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationResponseKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) ResponseParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) ResponseTemplates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) TemplateSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponseV2) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::IntegrationResponse`.
// Deprecated: moved to package aws-apigatewayv2.
func NewCfnIntegrationResponseV2(scope awscdk.Construct, id *string, props *CfnIntegrationResponseV2Props) CfnIntegrationResponseV2 {
	_init_.Initialize()

	if err := validateNewCfnIntegrationResponseV2Parameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIntegrationResponseV2{}

	_jsii_.Create(
		"monocdk.aws_apigateway.CfnIntegrationResponseV2",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::IntegrationResponse`.
// Deprecated: moved to package aws-apigatewayv2.
func NewCfnIntegrationResponseV2_Override(c CfnIntegrationResponseV2, scope awscdk.Construct, id *string, props *CfnIntegrationResponseV2Props) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigateway.CfnIntegrationResponseV2",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIntegrationResponseV2)SetApiId(val *string) {
	if err := j.validateSetApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnIntegrationResponseV2)SetContentHandlingStrategy(val *string) {
	_jsii_.Set(
		j,
		"contentHandlingStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnIntegrationResponseV2)SetIntegrationId(val *string) {
	if err := j.validateSetIntegrationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationId",
		val,
	)
}

func (j *jsiiProxy_CfnIntegrationResponseV2)SetIntegrationResponseKey(val *string) {
	if err := j.validateSetIntegrationResponseKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationResponseKey",
		val,
	)
}

func (j *jsiiProxy_CfnIntegrationResponseV2)SetResponseParameters(val interface{}) {
	if err := j.validateSetResponseParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseParameters",
		val,
	)
}

func (j *jsiiProxy_CfnIntegrationResponseV2)SetResponseTemplates(val interface{}) {
	if err := j.validateSetResponseTemplatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseTemplates",
		val,
	)
}

func (j *jsiiProxy_CfnIntegrationResponseV2)SetTemplateSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"templateSelectionExpression",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Deprecated: moved to package aws-apigatewayv2.
func CfnIntegrationResponseV2_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIntegrationResponseV2_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.CfnIntegrationResponseV2",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Deprecated: moved to package aws-apigatewayv2.
func CfnIntegrationResponseV2_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnIntegrationResponseV2_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.CfnIntegrationResponseV2",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Deprecated: moved to package aws-apigatewayv2.
func CfnIntegrationResponseV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIntegrationResponseV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.CfnIntegrationResponseV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIntegrationResponseV2_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_apigateway.CfnIntegrationResponseV2",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIntegrationResponseV2) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegrationResponseV2) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegrationResponseV2) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegrationResponseV2) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegrationResponseV2) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegrationResponseV2) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnIntegrationResponseV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegrationResponseV2) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegrationResponseV2) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

