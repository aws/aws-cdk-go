package awsgreengrass

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Greengrass::ConnectorDefinition`.
//
// The `AWS::Greengrass::ConnectorDefinition` resource represents a connector definition for AWS IoT Greengrass . Connector definitions are used to organize your connector definition versions.
//
// Connector definitions can reference multiple connector definition versions. All connector definition versions must be associated with a connector definition. Each connector definition version can contain one or more connectors.
//
// > When you create a connector definition, you can optionally include an initial connector definition version. To associate a connector definition version later, create an [`AWS::Greengrass::ConnectorDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html) resource and specify the ID of this connector definition.
// >
// > After you create the connector definition version that contains the connectors you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var tags interface{}
//
//   cfnConnectorDefinition := awscdk.Aws_greengrass.NewCfnConnectorDefinition(this, jsii.String("MyCfnConnectorDefinition"), &cfnConnectorDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &connectorDefinitionVersionProperty{
//   		connectors: []interface{}{
//   			&connectorProperty{
//   				connectorArn: jsii.String("connectorArn"),
//   				id: jsii.String("id"),
//
//   				// the properties below are optional
//   				parameters: parameters,
//   			},
//   		},
//   	},
//   	tags: tags,
//   })
//
type CfnConnectorDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the `ConnectorDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/connectors/1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrArn() *string
	// The ID of the `ConnectorDefinition` , such as `1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrId() *string
	// The ARN of the last `ConnectorDefinitionVersion` that was added to the `ConnectorDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/connectors/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2` .
	AttrLatestVersionArn() *string
	// The name of the `ConnectorDefinition` , such as `MyConnectorDefinition` .
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The connector definition version to include when the connector definition is created.
	//
	// A connector definition version contains a list of [`connector`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinition-connector.html) property types.
	//
	// > To associate a connector definition version after the connector definition is created, create an [`AWS::Greengrass::ConnectorDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html) resource and specify the ID of this connector definition.
	InitialVersion() interface{}
	SetInitialVersion(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the connector definition.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Application-specific metadata to attach to the connector definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnConnectorDefinition
type jsiiProxy_CfnConnectorDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConnectorDefinition) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) AttrLatestVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLatestVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) InitialVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::ConnectorDefinition`.
func NewCfnConnectorDefinition(scope constructs.Construct, id *string, props *CfnConnectorDefinitionProps) CfnConnectorDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnConnectorDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnConnectorDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::ConnectorDefinition`.
func NewCfnConnectorDefinition_Override(c CfnConnectorDefinition, scope constructs.Construct, id *string, props *CfnConnectorDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnConnectorDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConnectorDefinition) SetInitialVersion(val interface{}) {
	_jsii_.Set(
		j,
		"initialVersion",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorDefinition) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConnectorDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnConnectorDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConnectorDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnConnectorDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnConnectorDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnConnectorDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectorDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnConnectorDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectorDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConnectorDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnectorDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConnectorDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConnectorDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConnectorDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConnectorDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConnectorDefinition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorDefinition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConnectorDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConnectorDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A connector definition version contains a list of connectors.
//
// > After you create a connector definition version that contains the connectors you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an AWS CloudFormation template, `ConnectorDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::ConnectorDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   connectorDefinitionVersionProperty := &connectorDefinitionVersionProperty{
//   	connectors: []interface{}{
//   		&connectorProperty{
//   			connectorArn: jsii.String("connectorArn"),
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			parameters: parameters,
//   		},
//   	},
//   }
//
type CfnConnectorDefinition_ConnectorDefinitionVersionProperty struct {
	// The connectors in this version.
	//
	// Only one instance of a given connector can be added to a connector definition version at a time.
	Connectors interface{} `field:"required" json:"connectors" yaml:"connectors"`
}

// Connectors are modules that provide built-in integration with local infrastructure, device protocols, AWS , and other cloud services.
//
// For more information, see [Integrate with Services and Protocols Using Greengrass Connectors](https://docs.aws.amazon.com/greengrass/latest/developerguide/connectors.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Connectors` property of the [`ConnectorDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinition-connectordefinitionversion.html) property type contains a list of `Connector` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   connectorProperty := &connectorProperty{
//   	connectorArn: jsii.String("connectorArn"),
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	parameters: parameters,
//   }
//
type CfnConnectorDefinition_ConnectorProperty struct {
	// The Amazon Resource Name (ARN) of the connector.
	//
	// For more information about connectors provided by AWS , see [Greengrass Connectors Provided by AWS](https://docs.aws.amazon.com/greengrass/latest/developerguide/connectors-list.html) .
	ConnectorArn *string `field:"required" json:"connectorArn" yaml:"connectorArn"`
	// A descriptive or arbitrary ID for the connector.
	//
	// This value must be unique within the connector definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The parameters or configuration used by the connector.
	//
	// For more information about connectors provided by AWS , see [Greengrass Connectors Provided by AWS](https://docs.aws.amazon.com/greengrass/latest/developerguide/connectors-list.html) .
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// Properties for defining a `CfnConnectorDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var tags interface{}
//
//   cfnConnectorDefinitionProps := &cfnConnectorDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &connectorDefinitionVersionProperty{
//   		connectors: []interface{}{
//   			&connectorProperty{
//   				connectorArn: jsii.String("connectorArn"),
//   				id: jsii.String("id"),
//
//   				// the properties below are optional
//   				parameters: parameters,
//   			},
//   		},
//   	},
//   	tags: tags,
//   }
//
type CfnConnectorDefinitionProps struct {
	// The name of the connector definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The connector definition version to include when the connector definition is created.
	//
	// A connector definition version contains a list of [`connector`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinition-connector.html) property types.
	//
	// > To associate a connector definition version after the connector definition is created, create an [`AWS::Greengrass::ConnectorDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html) resource and specify the ID of this connector definition.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the connector definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Greengrass::ConnectorDefinitionVersion`.
//
// The `AWS::Greengrass::ConnectorDefinitionVersion` resource represents a connector definition version for AWS IoT Greengrass . A connector definition version contains a list of connectors.
//
// > To create a connector definition version, you must specify the ID of the connector definition that you want to associate with the version. For information about creating a connector definition, see [`AWS::Greengrass::ConnectorDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinition.html) .
// >
// > After you create a connector definition version that contains the connectors you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnConnectorDefinitionVersion := awscdk.Aws_greengrass.NewCfnConnectorDefinitionVersion(this, jsii.String("MyCfnConnectorDefinitionVersion"), &cfnConnectorDefinitionVersionProps{
//   	connectorDefinitionId: jsii.String("connectorDefinitionId"),
//   	connectors: []interface{}{
//   		&connectorProperty{
//   			connectorArn: jsii.String("connectorArn"),
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			parameters: parameters,
//   		},
//   	},
//   })
//
type CfnConnectorDefinitionVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The ID of the connector definition associated with this version.
	//
	// This value is a GUID.
	ConnectorDefinitionId() *string
	SetConnectorDefinitionId(val *string)
	// The connectors in this version.
	//
	// Only one instance of a given connector can be added to the connector definition version at a time.
	Connectors() interface{}
	SetConnectors(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnConnectorDefinitionVersion
type jsiiProxy_CfnConnectorDefinitionVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) ConnectorDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) Connectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::ConnectorDefinitionVersion`.
func NewCfnConnectorDefinitionVersion(scope constructs.Construct, id *string, props *CfnConnectorDefinitionVersionProps) CfnConnectorDefinitionVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnConnectorDefinitionVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnConnectorDefinitionVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::ConnectorDefinitionVersion`.
func NewCfnConnectorDefinitionVersion_Override(c CfnConnectorDefinitionVersion, scope constructs.Construct, id *string, props *CfnConnectorDefinitionVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnConnectorDefinitionVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) SetConnectorDefinitionId(val *string) {
	_jsii_.Set(
		j,
		"connectorDefinitionId",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorDefinitionVersion) SetConnectors(val interface{}) {
	_jsii_.Set(
		j,
		"connectors",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConnectorDefinitionVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnConnectorDefinitionVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConnectorDefinitionVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnConnectorDefinitionVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnConnectorDefinitionVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnConnectorDefinitionVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectorDefinitionVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnConnectorDefinitionVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorDefinitionVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Connectors are modules that provide built-in integration with local infrastructure, device protocols, AWS , and other cloud services.
//
// For more information, see [Integrate with Services and Protocols Using Greengrass Connectors](https://docs.aws.amazon.com/greengrass/latest/developerguide/connectors.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Connectors` property of the [`AWS::Greengrass::ConnectorDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html) resource contains a list of `Connector` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   connectorProperty := &connectorProperty{
//   	connectorArn: jsii.String("connectorArn"),
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	parameters: parameters,
//   }
//
type CfnConnectorDefinitionVersion_ConnectorProperty struct {
	// The Amazon Resource Name (ARN) of the connector.
	//
	// For more information about connectors provided by AWS , see [Greengrass Connectors Provided by AWS](https://docs.aws.amazon.com/greengrass/latest/developerguide/connectors-list.html) .
	ConnectorArn *string `field:"required" json:"connectorArn" yaml:"connectorArn"`
	// A descriptive or arbitrary ID for the connector.
	//
	// This value must be unique within the connector definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The parameters or configuration that the connector uses.
	//
	// For more information about connectors provided by AWS , see [Greengrass Connectors Provided by AWS](https://docs.aws.amazon.com/greengrass/latest/developerguide/connectors-list.html) .
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// Properties for defining a `CfnConnectorDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnConnectorDefinitionVersionProps := &cfnConnectorDefinitionVersionProps{
//   	connectorDefinitionId: jsii.String("connectorDefinitionId"),
//   	connectors: []interface{}{
//   		&connectorProperty{
//   			connectorArn: jsii.String("connectorArn"),
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			parameters: parameters,
//   		},
//   	},
//   }
//
type CfnConnectorDefinitionVersionProps struct {
	// The ID of the connector definition associated with this version.
	//
	// This value is a GUID.
	ConnectorDefinitionId *string `field:"required" json:"connectorDefinitionId" yaml:"connectorDefinitionId"`
	// The connectors in this version.
	//
	// Only one instance of a given connector can be added to the connector definition version at a time.
	Connectors interface{} `field:"required" json:"connectors" yaml:"connectors"`
}

// A CloudFormation `AWS::Greengrass::CoreDefinition`.
//
// The `AWS::Greengrass::CoreDefinition` resource represents a core definition for AWS IoT Greengrass . Core definitions are used to organize your core definition versions.
//
// Core definitions can reference multiple core definition versions. All core definition versions must be associated with a core definition. Each core definition version can contain one Greengrass core.
//
// > When you create a core definition, you can optionally include an initial core definition version. To associate a core definition version later, create an [`AWS::Greengrass::CoreDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinitionversion.html) resource and specify the ID of this core definition.
// >
// > After you create the core definition version that contains the core you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnCoreDefinition := awscdk.Aws_greengrass.NewCfnCoreDefinition(this, jsii.String("MyCfnCoreDefinition"), &cfnCoreDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &coreDefinitionVersionProperty{
//   		cores: []interface{}{
//   			&coreProperty{
//   				certificateArn: jsii.String("certificateArn"),
//   				id: jsii.String("id"),
//   				thingArn: jsii.String("thingArn"),
//
//   				// the properties below are optional
//   				syncShadow: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	tags: tags,
//   })
//
type CfnCoreDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the `CoreDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/cores/1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrArn() *string
	// The ID of the `CoreDefinition` , such as `1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrId() *string
	// The ARN of the last `CoreDefinitionVersion` that was added to the `CoreDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/cores/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2` .
	AttrLatestVersionArn() *string
	// The name of the `CoreDefinition` , such as `MyCoreDefinition` .
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The core definition version to include when the core definition is created.
	//
	// Currently, a core definition version can contain only one [`core`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-core.html) .
	//
	// > To associate a core definition version after the core definition is created, create an [`AWS::Greengrass::CoreDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinitionversion.html) resource and specify the ID of this core definition.
	InitialVersion() interface{}
	SetInitialVersion(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the core definition.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Application-specific metadata to attach to the core definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCoreDefinition
type jsiiProxy_CfnCoreDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCoreDefinition) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) AttrLatestVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLatestVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) InitialVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::CoreDefinition`.
func NewCfnCoreDefinition(scope constructs.Construct, id *string, props *CfnCoreDefinitionProps) CfnCoreDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnCoreDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnCoreDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::CoreDefinition`.
func NewCfnCoreDefinition_Override(c CfnCoreDefinition, scope constructs.Construct, id *string, props *CfnCoreDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnCoreDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCoreDefinition) SetInitialVersion(val interface{}) {
	_jsii_.Set(
		j,
		"initialVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCoreDefinition) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCoreDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnCoreDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCoreDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnCoreDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnCoreDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnCoreDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCoreDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnCoreDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCoreDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCoreDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCoreDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCoreDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCoreDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCoreDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCoreDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCoreDefinition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCoreDefinition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCoreDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCoreDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCoreDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCoreDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCoreDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCoreDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A core definition version contains a Greengrass [core](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-core.html) .
//
// > After you create a core definition version that contains the core you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an AWS CloudFormation template, `CoreDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::CoreDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreDefinitionVersionProperty := &coreDefinitionVersionProperty{
//   	cores: []interface{}{
//   		&coreProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   			id: jsii.String("id"),
//   			thingArn: jsii.String("thingArn"),
//
//   			// the properties below are optional
//   			syncShadow: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnCoreDefinition_CoreDefinitionVersionProperty struct {
	// The Greengrass core in this version.
	//
	// Currently, the `Cores` property for a core definition version can contain only one core.
	Cores interface{} `field:"required" json:"cores" yaml:"cores"`
}

// A core is an AWS IoT device that runs the AWS IoT Greengrass core software and manages local processes for a Greengrass group.
//
// For more information, see [What Is AWS IoT Greengrass ?](https://docs.aws.amazon.com/greengrass/latest/developerguide/what-is-gg.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Cores` property of the [`CoreDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-coredefinitionversion.html) property type contains a list of `Core` property types. Currently, the list can contain only one core.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreProperty := &coreProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	id: jsii.String("id"),
//   	thingArn: jsii.String("thingArn"),
//
//   	// the properties below are optional
//   	syncShadow: jsii.Boolean(false),
//   }
//
type CfnCoreDefinition_CoreProperty struct {
	// The Amazon Resource Name (ARN) of the device certificate for the core.
	//
	// This X.509 certificate is used to authenticate the core with AWS IoT and AWS IoT Greengrass services.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// A descriptive or arbitrary ID for the core.
	//
	// This value must be unique within the core definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ARN of the core, which is an AWS IoT device (thing).
	ThingArn *string `field:"required" json:"thingArn" yaml:"thingArn"`
	// Indicates whether the core's local shadow is synced with the cloud automatically.
	//
	// The default is false.
	SyncShadow interface{} `field:"optional" json:"syncShadow" yaml:"syncShadow"`
}

// Properties for defining a `CfnCoreDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnCoreDefinitionProps := &cfnCoreDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &coreDefinitionVersionProperty{
//   		cores: []interface{}{
//   			&coreProperty{
//   				certificateArn: jsii.String("certificateArn"),
//   				id: jsii.String("id"),
//   				thingArn: jsii.String("thingArn"),
//
//   				// the properties below are optional
//   				syncShadow: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	tags: tags,
//   }
//
type CfnCoreDefinitionProps struct {
	// The name of the core definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The core definition version to include when the core definition is created.
	//
	// Currently, a core definition version can contain only one [`core`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-core.html) .
	//
	// > To associate a core definition version after the core definition is created, create an [`AWS::Greengrass::CoreDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinitionversion.html) resource and specify the ID of this core definition.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the core definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Greengrass::CoreDefinitionVersion`.
//
// The `AWS::Greengrass::CoreDefinitionVersion` resource represents a core definition version for AWS IoT Greengrass . A core definition version contains a Greengrass core.
//
// > To create a core definition version, you must specify the ID of the core definition that you want to associate with the version. For information about creating a core definition, see [`AWS::Greengrass::CoreDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinition.html) .
// >
// > After you create a core definition version that contains the core you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCoreDefinitionVersion := awscdk.Aws_greengrass.NewCfnCoreDefinitionVersion(this, jsii.String("MyCfnCoreDefinitionVersion"), &cfnCoreDefinitionVersionProps{
//   	coreDefinitionId: jsii.String("coreDefinitionId"),
//   	cores: []interface{}{
//   		&coreProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   			id: jsii.String("id"),
//   			thingArn: jsii.String("thingArn"),
//
//   			// the properties below are optional
//   			syncShadow: jsii.Boolean(false),
//   		},
//   	},
//   })
//
type CfnCoreDefinitionVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The ID of the core definition associated with this version.
	//
	// This value is a GUID.
	CoreDefinitionId() *string
	SetCoreDefinitionId(val *string)
	// The Greengrass core in this version.
	//
	// Currently, the `Cores` property for a core definition version can contain only one core.
	Cores() interface{}
	SetCores(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCoreDefinitionVersion
type jsiiProxy_CfnCoreDefinitionVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) CoreDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) Cores() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::CoreDefinitionVersion`.
func NewCfnCoreDefinitionVersion(scope constructs.Construct, id *string, props *CfnCoreDefinitionVersionProps) CfnCoreDefinitionVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnCoreDefinitionVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnCoreDefinitionVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::CoreDefinitionVersion`.
func NewCfnCoreDefinitionVersion_Override(c CfnCoreDefinitionVersion, scope constructs.Construct, id *string, props *CfnCoreDefinitionVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnCoreDefinitionVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) SetCoreDefinitionId(val *string) {
	_jsii_.Set(
		j,
		"coreDefinitionId",
		val,
	)
}

func (j *jsiiProxy_CfnCoreDefinitionVersion) SetCores(val interface{}) {
	_jsii_.Set(
		j,
		"cores",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCoreDefinitionVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnCoreDefinitionVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCoreDefinitionVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnCoreDefinitionVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnCoreDefinitionVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnCoreDefinitionVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCoreDefinitionVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnCoreDefinitionVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCoreDefinitionVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A core is an AWS IoT device that runs the AWS IoT Greengrass core software and manages local processes for a Greengrass group.
//
// For more information, see [What Is AWS IoT Greengrass ?](https://docs.aws.amazon.com/greengrass/latest/developerguide/what-is-gg.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Cores` property of the [`AWS::Greengrass::CoreDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinitionversion.html) resource contains a list of `Core` property types. Currently, the list can contain only one core.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreProperty := &coreProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	id: jsii.String("id"),
//   	thingArn: jsii.String("thingArn"),
//
//   	// the properties below are optional
//   	syncShadow: jsii.Boolean(false),
//   }
//
type CfnCoreDefinitionVersion_CoreProperty struct {
	// The ARN of the device certificate for the core.
	//
	// This X.509 certificate is used to authenticate the core with AWS IoT and AWS IoT Greengrass services.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// A descriptive or arbitrary ID for the core.
	//
	// This value must be unique within the core definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The Amazon Resource Name (ARN) of the core, which is an AWS IoT device (thing).
	ThingArn *string `field:"required" json:"thingArn" yaml:"thingArn"`
	// Indicates whether the core's local shadow is synced with the cloud automatically.
	//
	// The default is false.
	SyncShadow interface{} `field:"optional" json:"syncShadow" yaml:"syncShadow"`
}

// Properties for defining a `CfnCoreDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCoreDefinitionVersionProps := &cfnCoreDefinitionVersionProps{
//   	coreDefinitionId: jsii.String("coreDefinitionId"),
//   	cores: []interface{}{
//   		&coreProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   			id: jsii.String("id"),
//   			thingArn: jsii.String("thingArn"),
//
//   			// the properties below are optional
//   			syncShadow: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnCoreDefinitionVersionProps struct {
	// The ID of the core definition associated with this version.
	//
	// This value is a GUID.
	CoreDefinitionId *string `field:"required" json:"coreDefinitionId" yaml:"coreDefinitionId"`
	// The Greengrass core in this version.
	//
	// Currently, the `Cores` property for a core definition version can contain only one core.
	Cores interface{} `field:"required" json:"cores" yaml:"cores"`
}

// A CloudFormation `AWS::Greengrass::DeviceDefinition`.
//
// The `AWS::Greengrass::DeviceDefinition` resource represents a device definition for AWS IoT Greengrass . Device definitions are used to organize your device definition versions.
//
// Device definitions can reference multiple device definition versions. All device definition versions must be associated with a device definition. Each device definition version can contain one or more devices.
//
// > When you create a device definition, you can optionally include an initial device definition version. To associate a device definition version later, create an [`AWS::Greengrass::DeviceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html) resource and specify the ID of this device definition.
// >
// > After you create the device definition version that contains the devices you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnDeviceDefinition := awscdk.Aws_greengrass.NewCfnDeviceDefinition(this, jsii.String("MyCfnDeviceDefinition"), &cfnDeviceDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &deviceDefinitionVersionProperty{
//   		devices: []interface{}{
//   			&deviceProperty{
//   				certificateArn: jsii.String("certificateArn"),
//   				id: jsii.String("id"),
//   				thingArn: jsii.String("thingArn"),
//
//   				// the properties below are optional
//   				syncShadow: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	tags: tags,
//   })
//
type CfnDeviceDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the `DeviceDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/devices/1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrArn() *string
	// The ID of the `DeviceDefinition` , such as `1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrId() *string
	// The ARN of the last `DeviceDefinitionVersion` that was added to the `DeviceDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/devices/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2` .
	AttrLatestVersionArn() *string
	// The name of the device definition.
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The device definition version to include when the device definition is created.
	//
	// A device definition version contains a list of [`device`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-device.html) property types.
	//
	// > To associate a device definition version after the device definition is created, create an [`AWS::Greengrass::DeviceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html) resource and specify the ID of this device definition.
	InitialVersion() interface{}
	SetInitialVersion(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the device definition.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Application-specific metadata to attach to the device definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDeviceDefinition
type jsiiProxy_CfnDeviceDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDeviceDefinition) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) AttrLatestVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLatestVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) InitialVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::DeviceDefinition`.
func NewCfnDeviceDefinition(scope constructs.Construct, id *string, props *CfnDeviceDefinitionProps) CfnDeviceDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnDeviceDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnDeviceDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::DeviceDefinition`.
func NewCfnDeviceDefinition_Override(c CfnDeviceDefinition, scope constructs.Construct, id *string, props *CfnDeviceDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnDeviceDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDeviceDefinition) SetInitialVersion(val interface{}) {
	_jsii_.Set(
		j,
		"initialVersion",
		val,
	)
}

func (j *jsiiProxy_CfnDeviceDefinition) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDeviceDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnDeviceDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDeviceDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnDeviceDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnDeviceDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnDeviceDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeviceDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnDeviceDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeviceDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDeviceDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeviceDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDeviceDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDeviceDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDeviceDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDeviceDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDeviceDefinition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceDefinition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDeviceDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDeviceDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A device definition version contains a list of [devices](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-device.html) .
//
// > After you create a device definition version that contains the devices you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an AWS CloudFormation template, `DeviceDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::DeviceDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceDefinitionVersionProperty := &deviceDefinitionVersionProperty{
//   	devices: []interface{}{
//   		&deviceProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   			id: jsii.String("id"),
//   			thingArn: jsii.String("thingArn"),
//
//   			// the properties below are optional
//   			syncShadow: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnDeviceDefinition_DeviceDefinitionVersionProperty struct {
	// The devices in this version.
	Devices interface{} `field:"required" json:"devices" yaml:"devices"`
}

// A device is an AWS IoT device (thing) that's added to a Greengrass group.
//
// Greengrass devices can communicate with the Greengrass core in the same group. For more information, see [What Is AWS IoT Greengrass ?](https://docs.aws.amazon.com/greengrass/latest/developerguide/what-is-gg.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Devices` property of the [`DeviceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-devicedefinitionversion.html) property type contains a list of `Device` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceProperty := &deviceProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	id: jsii.String("id"),
//   	thingArn: jsii.String("thingArn"),
//
//   	// the properties below are optional
//   	syncShadow: jsii.Boolean(false),
//   }
//
type CfnDeviceDefinition_DeviceProperty struct {
	// The Amazon Resource Name (ARN) of the device certificate for the device.
	//
	// This X.509 certificate is used to authenticate the device with AWS IoT and AWS IoT Greengrass services.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// A descriptive or arbitrary ID for the device.
	//
	// This value must be unique within the device definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ARN of the device, which is an AWS IoT device (thing).
	ThingArn *string `field:"required" json:"thingArn" yaml:"thingArn"`
	// Indicates whether the device's local shadow is synced with the cloud automatically.
	SyncShadow interface{} `field:"optional" json:"syncShadow" yaml:"syncShadow"`
}

// Properties for defining a `CfnDeviceDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnDeviceDefinitionProps := &cfnDeviceDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &deviceDefinitionVersionProperty{
//   		devices: []interface{}{
//   			&deviceProperty{
//   				certificateArn: jsii.String("certificateArn"),
//   				id: jsii.String("id"),
//   				thingArn: jsii.String("thingArn"),
//
//   				// the properties below are optional
//   				syncShadow: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	tags: tags,
//   }
//
type CfnDeviceDefinitionProps struct {
	// The name of the device definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The device definition version to include when the device definition is created.
	//
	// A device definition version contains a list of [`device`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-device.html) property types.
	//
	// > To associate a device definition version after the device definition is created, create an [`AWS::Greengrass::DeviceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html) resource and specify the ID of this device definition.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the device definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Greengrass::DeviceDefinitionVersion`.
//
// The `AWS::Greengrass::DeviceDefinitionVersion` resource represents a device definition version for AWS IoT Greengrass . A device definition version contains a list of devices.
//
// > To create a device definition version, you must specify the ID of the device definition that you want to associate with the version. For information about creating a device definition, see [`AWS::Greengrass::DeviceDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html) .
// >
// > After you create a device definition version that contains the devices you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceDefinitionVersion := awscdk.Aws_greengrass.NewCfnDeviceDefinitionVersion(this, jsii.String("MyCfnDeviceDefinitionVersion"), &cfnDeviceDefinitionVersionProps{
//   	deviceDefinitionId: jsii.String("deviceDefinitionId"),
//   	devices: []interface{}{
//   		&deviceProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   			id: jsii.String("id"),
//   			thingArn: jsii.String("thingArn"),
//
//   			// the properties below are optional
//   			syncShadow: jsii.Boolean(false),
//   		},
//   	},
//   })
//
type CfnDeviceDefinitionVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ID of the device definition associated with this version.
	//
	// This value is a GUID.
	DeviceDefinitionId() *string
	SetDeviceDefinitionId(val *string)
	// The devices in this version.
	Devices() interface{}
	SetDevices(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDeviceDefinitionVersion
type jsiiProxy_CfnDeviceDefinitionVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) DeviceDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) Devices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::DeviceDefinitionVersion`.
func NewCfnDeviceDefinitionVersion(scope constructs.Construct, id *string, props *CfnDeviceDefinitionVersionProps) CfnDeviceDefinitionVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnDeviceDefinitionVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnDeviceDefinitionVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::DeviceDefinitionVersion`.
func NewCfnDeviceDefinitionVersion_Override(c CfnDeviceDefinitionVersion, scope constructs.Construct, id *string, props *CfnDeviceDefinitionVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnDeviceDefinitionVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) SetDeviceDefinitionId(val *string) {
	_jsii_.Set(
		j,
		"deviceDefinitionId",
		val,
	)
}

func (j *jsiiProxy_CfnDeviceDefinitionVersion) SetDevices(val interface{}) {
	_jsii_.Set(
		j,
		"devices",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDeviceDefinitionVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnDeviceDefinitionVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDeviceDefinitionVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnDeviceDefinitionVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnDeviceDefinitionVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnDeviceDefinitionVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeviceDefinitionVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnDeviceDefinitionVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceDefinitionVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A device is an AWS IoT device (thing) that's added to a Greengrass group.
//
// Greengrass devices can communicate with the Greengrass core in the same group. For more information, see [What Is AWS IoT Greengrass ?](https://docs.aws.amazon.com/greengrass/latest/developerguide/what-is-gg.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Devices` property of the [`AWS::Greengrass::DeviceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html) resource contains a list of `Device` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceProperty := &deviceProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	id: jsii.String("id"),
//   	thingArn: jsii.String("thingArn"),
//
//   	// the properties below are optional
//   	syncShadow: jsii.Boolean(false),
//   }
//
type CfnDeviceDefinitionVersion_DeviceProperty struct {
	// The ARN of the device certificate for the device.
	//
	// This X.509 certificate is used to authenticate the device with AWS IoT and AWS IoT Greengrass services.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// A descriptive or arbitrary ID for the device.
	//
	// This value must be unique within the device definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The Amazon Resource Name (ARN) of the device, which is an AWS IoT device (thing).
	ThingArn *string `field:"required" json:"thingArn" yaml:"thingArn"`
	// Indicates whether the device's local shadow is synced with the cloud automatically.
	SyncShadow interface{} `field:"optional" json:"syncShadow" yaml:"syncShadow"`
}

// Properties for defining a `CfnDeviceDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceDefinitionVersionProps := &cfnDeviceDefinitionVersionProps{
//   	deviceDefinitionId: jsii.String("deviceDefinitionId"),
//   	devices: []interface{}{
//   		&deviceProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   			id: jsii.String("id"),
//   			thingArn: jsii.String("thingArn"),
//
//   			// the properties below are optional
//   			syncShadow: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnDeviceDefinitionVersionProps struct {
	// The ID of the device definition associated with this version.
	//
	// This value is a GUID.
	DeviceDefinitionId *string `field:"required" json:"deviceDefinitionId" yaml:"deviceDefinitionId"`
	// The devices in this version.
	Devices interface{} `field:"required" json:"devices" yaml:"devices"`
}

// A CloudFormation `AWS::Greengrass::FunctionDefinition`.
//
// The `AWS::Greengrass::FunctionDefinition` resource represents a function definition for AWS IoT Greengrass . Function definitions are used to organize your function definition versions.
//
// Function definitions can reference multiple function definition versions. All function definition versions must be associated with a function definition. Each function definition version can contain one or more functions.
//
// > When you create a function definition, you can optionally include an initial function definition version. To associate a function definition version later, create an [`AWS::Greengrass::FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html) resource and specify the ID of this function definition.
// >
// > After you create the function definition version that contains the functions you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//   var variables interface{}
//
//   cfnFunctionDefinition := awscdk.Aws_greengrass.NewCfnFunctionDefinition(this, jsii.String("MyCfnFunctionDefinition"), &cfnFunctionDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &functionDefinitionVersionProperty{
//   		functions: []interface{}{
//   			&functionProperty{
//   				functionArn: jsii.String("functionArn"),
//   				functionConfiguration: &functionConfigurationProperty{
//   					encodingType: jsii.String("encodingType"),
//   					environment: &environmentProperty{
//   						accessSysfs: jsii.Boolean(false),
//   						execution: &executionProperty{
//   							isolationMode: jsii.String("isolationMode"),
//   							runAs: &runAsProperty{
//   								gid: jsii.Number(123),
//   								uid: jsii.Number(123),
//   							},
//   						},
//   						resourceAccessPolicies: []interface{}{
//   							&resourceAccessPolicyProperty{
//   								resourceId: jsii.String("resourceId"),
//
//   								// the properties below are optional
//   								permission: jsii.String("permission"),
//   							},
//   						},
//   						variables: variables,
//   					},
//   					execArgs: jsii.String("execArgs"),
//   					executable: jsii.String("executable"),
//   					memorySize: jsii.Number(123),
//   					pinned: jsii.Boolean(false),
//   					timeout: jsii.Number(123),
//   				},
//   				id: jsii.String("id"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		defaultConfig: &defaultConfigProperty{
//   			execution: &executionProperty{
//   				isolationMode: jsii.String("isolationMode"),
//   				runAs: &runAsProperty{
//   					gid: jsii.Number(123),
//   					uid: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	tags: tags,
//   })
//
type CfnFunctionDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the `FunctionDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/functions/1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrArn() *string
	// The ID of the `FunctionDefinition` , such as `1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrId() *string
	// The ARN of the last `FunctionDefinitionVersion` that was added to the `FunctionDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/functions/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2` .
	AttrLatestVersionArn() *string
	// The name of the `FunctionDefinition` , such as `MyFunctionDefinition` .
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The function definition version to include when the function definition is created.
	//
	// A function definition version contains a list of [`function`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-function.html) property types.
	//
	// > To associate a function definition version after the function definition is created, create an [`AWS::Greengrass::FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html) resource and specify the ID of this function definition.
	InitialVersion() interface{}
	SetInitialVersion(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the function definition.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Application-specific metadata to attach to the function definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnFunctionDefinition
type jsiiProxy_CfnFunctionDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFunctionDefinition) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) AttrLatestVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLatestVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) InitialVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::FunctionDefinition`.
func NewCfnFunctionDefinition(scope constructs.Construct, id *string, props *CfnFunctionDefinitionProps) CfnFunctionDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnFunctionDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnFunctionDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::FunctionDefinition`.
func NewCfnFunctionDefinition_Override(c CfnFunctionDefinition, scope constructs.Construct, id *string, props *CfnFunctionDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnFunctionDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFunctionDefinition) SetInitialVersion(val interface{}) {
	_jsii_.Set(
		j,
		"initialVersion",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionDefinition) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFunctionDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnFunctionDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFunctionDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnFunctionDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnFunctionDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnFunctionDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFunctionDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnFunctionDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFunctionDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFunctionDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFunctionDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFunctionDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFunctionDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFunctionDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFunctionDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFunctionDefinition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionDefinition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFunctionDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFunctionDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The default configuration that applies to all Lambda functions in the function definition version.
//
// Individual Lambda functions can override these settings.
//
// In an AWS CloudFormation template, `DefaultConfig` is a property of the [`FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functiondefinitionversion.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultConfigProperty := &defaultConfigProperty{
//   	execution: &executionProperty{
//   		isolationMode: jsii.String("isolationMode"),
//   		runAs: &runAsProperty{
//   			gid: jsii.Number(123),
//   			uid: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnFunctionDefinition_DefaultConfigProperty struct {
	// Configuration settings for the Lambda execution environment on the AWS IoT Greengrass core.
	Execution interface{} `field:"required" json:"execution" yaml:"execution"`
}

// The environment configuration for a Lambda function on the AWS IoT Greengrass core.
//
// In an AWS CloudFormation template, `Environment` is a property of the [`FunctionConfiguration`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functionconfiguration.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   environmentProperty := &environmentProperty{
//   	accessSysfs: jsii.Boolean(false),
//   	execution: &executionProperty{
//   		isolationMode: jsii.String("isolationMode"),
//   		runAs: &runAsProperty{
//   			gid: jsii.Number(123),
//   			uid: jsii.Number(123),
//   		},
//   	},
//   	resourceAccessPolicies: []interface{}{
//   		&resourceAccessPolicyProperty{
//   			resourceId: jsii.String("resourceId"),
//
//   			// the properties below are optional
//   			permission: jsii.String("permission"),
//   		},
//   	},
//   	variables: variables,
//   }
//
type CfnFunctionDefinition_EnvironmentProperty struct {
	// Indicates whether the function is allowed to access the `/sys` directory on the core device, which allows the read device information from `/sys` .
	//
	// > This property applies only to Lambda functions that run in a Greengrass container.
	AccessSysfs interface{} `field:"optional" json:"accessSysfs" yaml:"accessSysfs"`
	// Settings for the Lambda execution environment in AWS IoT Greengrass .
	Execution interface{} `field:"optional" json:"execution" yaml:"execution"`
	// A list of the [resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourceinstance.html) in the group that the function can access, with the corresponding read-only or read-write permissions. The maximum is 10 resources.
	//
	// > This property applies only for Lambda functions that run in a Greengrass container.
	ResourceAccessPolicies interface{} `field:"optional" json:"resourceAccessPolicies" yaml:"resourceAccessPolicies"`
	// Environment variables for the Lambda function.
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

// Configuration settings for the Lambda execution environment on the AWS IoT Greengrass core.
//
// In an AWS CloudFormation template, `Execution` is a property of the [`DefaultConfig`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-defaultconfig.html) property type for a function definition version and the [`Environment`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-environment.html) property type for a function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executionProperty := &executionProperty{
//   	isolationMode: jsii.String("isolationMode"),
//   	runAs: &runAsProperty{
//   		gid: jsii.Number(123),
//   		uid: jsii.Number(123),
//   	},
//   }
//
type CfnFunctionDefinition_ExecutionProperty struct {
	// The containerization that the Lambda function runs in.
	//
	// Valid values are `GreengrassContainer` or `NoContainer` . Typically, this is `GreengrassContainer` . For more information, see [Containerization](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html#lambda-function-containerization) in the *Developer Guide* .
	//
	// - When set on the [`DefaultConfig`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-defaultconfig.html) property of a function definition version, this setting is used as the default containerization for all Lambda functions in the function definition version.
	// - When set on the [`Environment`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html) property of a function, this setting applies to the individual function and overrides the default. Omit this value to run the function with the default containerization.
	//
	// > We recommend that you run in a Greengrass container unless your business case requires that you run without containerization.
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
	// The user and group permissions used to run the Lambda function.
	//
	// Typically, this is the ggc_user and ggc_group. For more information, see [Run as](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html#lambda-access-identity.html) in the *Developer Guide* .
	//
	// - When set on the [`DefaultConfig`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-defaultconfig.html) property of a function definition version, this setting is used as the default access identity for all Lambda functions in the function definition version.
	// - When set on the [`Environment`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html) property of a function, this setting applies to the individual function and overrides the default. You can override the user, group, or both. Omit this value to run the function with the default permissions.
	//
	// > Running as the root user increases risks to your data and device. Do not run as root (UID/GID=0) unless your business case requires it. For more information and requirements, see [Running a Lambda Function as Root](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html#lambda-running-as-root) .
	RunAs interface{} `field:"optional" json:"runAs" yaml:"runAs"`
}

// The group-specific configuration settings for a Lambda function.
//
// These settings configure the function's behavior in the Greengrass group. For more information, see [Controlling Execution of Greengrass Lambda Functions by Using Group-Specific Configuration](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `FunctionConfiguration` is a property of the [`Function`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-function.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   functionConfigurationProperty := &functionConfigurationProperty{
//   	encodingType: jsii.String("encodingType"),
//   	environment: &environmentProperty{
//   		accessSysfs: jsii.Boolean(false),
//   		execution: &executionProperty{
//   			isolationMode: jsii.String("isolationMode"),
//   			runAs: &runAsProperty{
//   				gid: jsii.Number(123),
//   				uid: jsii.Number(123),
//   			},
//   		},
//   		resourceAccessPolicies: []interface{}{
//   			&resourceAccessPolicyProperty{
//   				resourceId: jsii.String("resourceId"),
//
//   				// the properties below are optional
//   				permission: jsii.String("permission"),
//   			},
//   		},
//   		variables: variables,
//   	},
//   	execArgs: jsii.String("execArgs"),
//   	executable: jsii.String("executable"),
//   	memorySize: jsii.Number(123),
//   	pinned: jsii.Boolean(false),
//   	timeout: jsii.Number(123),
//   }
//
type CfnFunctionDefinition_FunctionConfigurationProperty struct {
	// The expected encoding type of the input payload for the function.
	//
	// Valid values are `json` (default) and `binary` .
	EncodingType *string `field:"optional" json:"encodingType" yaml:"encodingType"`
	// The environment configuration of the function.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The execution arguments.
	ExecArgs *string `field:"optional" json:"execArgs" yaml:"execArgs"`
	// The name of the function executable.
	Executable *string `field:"optional" json:"executable" yaml:"executable"`
	// The memory size (in KB) required by the function.
	//
	// > This property applies only to Lambda functions that run in a Greengrass container.
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// Indicates whether the function is pinned (or *long-lived* ).
	//
	// Pinned functions start when the core starts and process all requests in the same container. The default value is false.
	Pinned interface{} `field:"optional" json:"pinned" yaml:"pinned"`
	// The allowed execution time (in seconds) after which the function should terminate.
	//
	// For pinned functions, this timeout applies for each request.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

// A function definition version contains a list of functions.
//
// > After you create a function definition version that contains the functions you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an AWS CloudFormation template, `FunctionDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::FunctionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   functionDefinitionVersionProperty := &functionDefinitionVersionProperty{
//   	functions: []interface{}{
//   		&functionProperty{
//   			functionArn: jsii.String("functionArn"),
//   			functionConfiguration: &functionConfigurationProperty{
//   				encodingType: jsii.String("encodingType"),
//   				environment: &environmentProperty{
//   					accessSysfs: jsii.Boolean(false),
//   					execution: &executionProperty{
//   						isolationMode: jsii.String("isolationMode"),
//   						runAs: &runAsProperty{
//   							gid: jsii.Number(123),
//   							uid: jsii.Number(123),
//   						},
//   					},
//   					resourceAccessPolicies: []interface{}{
//   						&resourceAccessPolicyProperty{
//   							resourceId: jsii.String("resourceId"),
//
//   							// the properties below are optional
//   							permission: jsii.String("permission"),
//   						},
//   					},
//   					variables: variables,
//   				},
//   				execArgs: jsii.String("execArgs"),
//   				executable: jsii.String("executable"),
//   				memorySize: jsii.Number(123),
//   				pinned: jsii.Boolean(false),
//   				timeout: jsii.Number(123),
//   			},
//   			id: jsii.String("id"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	defaultConfig: &defaultConfigProperty{
//   		execution: &executionProperty{
//   			isolationMode: jsii.String("isolationMode"),
//   			runAs: &runAsProperty{
//   				gid: jsii.Number(123),
//   				uid: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnFunctionDefinition_FunctionDefinitionVersionProperty struct {
	// The functions in this version.
	Functions interface{} `field:"required" json:"functions" yaml:"functions"`
	// The default configuration that applies to all Lambda functions in the group.
	//
	// Individual Lambda functions can override these settings.
	DefaultConfig interface{} `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
}

// A function is a Lambda function that's referenced from an AWS IoT Greengrass group.
//
// The function is deployed to a Greengrass core where it runs locally. For more information, see [Run Lambda Functions on the AWS IoT Greengrass Core](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-functions.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Functions` property of the [`FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functiondefinitionversion.html) property type contains a list of `Function` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   functionProperty := &functionProperty{
//   	functionArn: jsii.String("functionArn"),
//   	functionConfiguration: &functionConfigurationProperty{
//   		encodingType: jsii.String("encodingType"),
//   		environment: &environmentProperty{
//   			accessSysfs: jsii.Boolean(false),
//   			execution: &executionProperty{
//   				isolationMode: jsii.String("isolationMode"),
//   				runAs: &runAsProperty{
//   					gid: jsii.Number(123),
//   					uid: jsii.Number(123),
//   				},
//   			},
//   			resourceAccessPolicies: []interface{}{
//   				&resourceAccessPolicyProperty{
//   					resourceId: jsii.String("resourceId"),
//
//   					// the properties below are optional
//   					permission: jsii.String("permission"),
//   				},
//   			},
//   			variables: variables,
//   		},
//   		execArgs: jsii.String("execArgs"),
//   		executable: jsii.String("executable"),
//   		memorySize: jsii.Number(123),
//   		pinned: jsii.Boolean(false),
//   		timeout: jsii.Number(123),
//   	},
//   	id: jsii.String("id"),
//   }
//
type CfnFunctionDefinition_FunctionProperty struct {
	// The Amazon Resource Name (ARN) of the alias (recommended) or version of the referenced Lambda function.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The group-specific settings of the Lambda function.
	//
	// These settings configure the function's behavior in the Greengrass group.
	FunctionConfiguration interface{} `field:"required" json:"functionConfiguration" yaml:"functionConfiguration"`
	// A descriptive or arbitrary ID for the function.
	//
	// This value must be unique within the function definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
}

// A list of the [resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourceinstance.html) in the group that the function can access, with the corresponding read-only or read-write permissions. The maximum is 10 resources.
//
// > This property applies only to Lambda functions that run in a Greengrass container.
//
// In an AWS CloudFormation template, `ResourceAccessPolicy` is a property of the [`Environment`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-environment.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceAccessPolicyProperty := &resourceAccessPolicyProperty{
//   	resourceId: jsii.String("resourceId"),
//
//   	// the properties below are optional
//   	permission: jsii.String("permission"),
//   }
//
type CfnFunctionDefinition_ResourceAccessPolicyProperty struct {
	// The ID of the resource.
	//
	// This ID is assigned to the resource when you create the resource definition.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The read-only or read-write access that the Lambda function has to the resource.
	//
	// Valid values are `ro` or `rw` .
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

// The access identity whose permissions are used to run the Lambda function.
//
// This setting overrides the default access identity that's specified for the group (by default, ggc_user and ggc_group). You can override the user, group, or both. For more information, see [Run as](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html#lambda-access-identity.html) in the *Developer Guide* .
//
// > Running as the root user increases risks to your data and device. Do not run as root (UID/GID=0) unless your business case requires it. For more information and requirements, see [Running a Lambda Function as Root](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html#lambda-running-as-root) .
//
// In an AWS CloudFormation template, `RunAs` is a property of the [`Execution`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-execution.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runAsProperty := &runAsProperty{
//   	gid: jsii.Number(123),
//   	uid: jsii.Number(123),
//   }
//
type CfnFunctionDefinition_RunAsProperty struct {
	// The group ID whose permissions are used to run the Lambda function.
	//
	// You can use the `getent group` command on your core device to look up the group ID.
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
	// The user ID whose permissions are used to run the Lambda function.
	//
	// You can use the `getent passwd` command on your core device to look up the user ID.
	Uid *float64 `field:"optional" json:"uid" yaml:"uid"`
}

// Properties for defining a `CfnFunctionDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//   var variables interface{}
//
//   cfnFunctionDefinitionProps := &cfnFunctionDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &functionDefinitionVersionProperty{
//   		functions: []interface{}{
//   			&functionProperty{
//   				functionArn: jsii.String("functionArn"),
//   				functionConfiguration: &functionConfigurationProperty{
//   					encodingType: jsii.String("encodingType"),
//   					environment: &environmentProperty{
//   						accessSysfs: jsii.Boolean(false),
//   						execution: &executionProperty{
//   							isolationMode: jsii.String("isolationMode"),
//   							runAs: &runAsProperty{
//   								gid: jsii.Number(123),
//   								uid: jsii.Number(123),
//   							},
//   						},
//   						resourceAccessPolicies: []interface{}{
//   							&resourceAccessPolicyProperty{
//   								resourceId: jsii.String("resourceId"),
//
//   								// the properties below are optional
//   								permission: jsii.String("permission"),
//   							},
//   						},
//   						variables: variables,
//   					},
//   					execArgs: jsii.String("execArgs"),
//   					executable: jsii.String("executable"),
//   					memorySize: jsii.Number(123),
//   					pinned: jsii.Boolean(false),
//   					timeout: jsii.Number(123),
//   				},
//   				id: jsii.String("id"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		defaultConfig: &defaultConfigProperty{
//   			execution: &executionProperty{
//   				isolationMode: jsii.String("isolationMode"),
//   				runAs: &runAsProperty{
//   					gid: jsii.Number(123),
//   					uid: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	tags: tags,
//   }
//
type CfnFunctionDefinitionProps struct {
	// The name of the function definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The function definition version to include when the function definition is created.
	//
	// A function definition version contains a list of [`function`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-function.html) property types.
	//
	// > To associate a function definition version after the function definition is created, create an [`AWS::Greengrass::FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html) resource and specify the ID of this function definition.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the function definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Greengrass::FunctionDefinitionVersion`.
//
// The `AWS::Greengrass::FunctionDefinitionVersion` resource represents a function definition version for AWS IoT Greengrass . A function definition version contains contain a list of functions.
//
// > To create a function definition version, you must specify the ID of the function definition that you want to associate with the version. For information about creating a function definition, see [`AWS::Greengrass::FunctionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html) .
// >
// > After you create a function definition version that contains the functions you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   cfnFunctionDefinitionVersion := awscdk.Aws_greengrass.NewCfnFunctionDefinitionVersion(this, jsii.String("MyCfnFunctionDefinitionVersion"), &cfnFunctionDefinitionVersionProps{
//   	functionDefinitionId: jsii.String("functionDefinitionId"),
//   	functions: []interface{}{
//   		&functionProperty{
//   			functionArn: jsii.String("functionArn"),
//   			functionConfiguration: &functionConfigurationProperty{
//   				encodingType: jsii.String("encodingType"),
//   				environment: &environmentProperty{
//   					accessSysfs: jsii.Boolean(false),
//   					execution: &executionProperty{
//   						isolationMode: jsii.String("isolationMode"),
//   						runAs: &runAsProperty{
//   							gid: jsii.Number(123),
//   							uid: jsii.Number(123),
//   						},
//   					},
//   					resourceAccessPolicies: []interface{}{
//   						&resourceAccessPolicyProperty{
//   							resourceId: jsii.String("resourceId"),
//
//   							// the properties below are optional
//   							permission: jsii.String("permission"),
//   						},
//   					},
//   					variables: variables,
//   				},
//   				execArgs: jsii.String("execArgs"),
//   				executable: jsii.String("executable"),
//   				memorySize: jsii.Number(123),
//   				pinned: jsii.Boolean(false),
//   				timeout: jsii.Number(123),
//   			},
//   			id: jsii.String("id"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	defaultConfig: &defaultConfigProperty{
//   		execution: &executionProperty{
//   			isolationMode: jsii.String("isolationMode"),
//   			runAs: &runAsProperty{
//   				gid: jsii.Number(123),
//   				uid: jsii.Number(123),
//   			},
//   		},
//   	},
//   })
//
type CfnFunctionDefinitionVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The default configuration that applies to all Lambda functions in the group.
	//
	// Individual Lambda functions can override these settings.
	DefaultConfig() interface{}
	SetDefaultConfig(val interface{})
	// The ID of the function definition associated with this version.
	//
	// This value is a GUID.
	FunctionDefinitionId() *string
	SetFunctionDefinitionId(val *string)
	// The functions in this version.
	Functions() interface{}
	SetFunctions(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnFunctionDefinitionVersion
type jsiiProxy_CfnFunctionDefinitionVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) DefaultConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) FunctionDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) Functions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::FunctionDefinitionVersion`.
func NewCfnFunctionDefinitionVersion(scope constructs.Construct, id *string, props *CfnFunctionDefinitionVersionProps) CfnFunctionDefinitionVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnFunctionDefinitionVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnFunctionDefinitionVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::FunctionDefinitionVersion`.
func NewCfnFunctionDefinitionVersion_Override(c CfnFunctionDefinitionVersion, scope constructs.Construct, id *string, props *CfnFunctionDefinitionVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnFunctionDefinitionVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) SetDefaultConfig(val interface{}) {
	_jsii_.Set(
		j,
		"defaultConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) SetFunctionDefinitionId(val *string) {
	_jsii_.Set(
		j,
		"functionDefinitionId",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionDefinitionVersion) SetFunctions(val interface{}) {
	_jsii_.Set(
		j,
		"functions",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFunctionDefinitionVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnFunctionDefinitionVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFunctionDefinitionVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnFunctionDefinitionVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnFunctionDefinitionVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnFunctionDefinitionVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFunctionDefinitionVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnFunctionDefinitionVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionDefinitionVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The default configuration that applies to all Lambda functions in the function definition version.
//
// Individual Lambda functions can override these settings.
//
// In an AWS CloudFormation template, `DefaultConfig` is a property of the [`AWS::Greengrass::FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultConfigProperty := &defaultConfigProperty{
//   	execution: &executionProperty{
//   		isolationMode: jsii.String("isolationMode"),
//   		runAs: &runAsProperty{
//   			gid: jsii.Number(123),
//   			uid: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnFunctionDefinitionVersion_DefaultConfigProperty struct {
	// Configuration settings for the Lambda execution environment on the AWS IoT Greengrass core.
	Execution interface{} `field:"required" json:"execution" yaml:"execution"`
}

// The environment configuration for a Lambda function on the AWS IoT Greengrass core.
//
// In an AWS CloudFormation template, `Environment` is a property of the [`FunctionConfiguration`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-functionconfiguration.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   environmentProperty := &environmentProperty{
//   	accessSysfs: jsii.Boolean(false),
//   	execution: &executionProperty{
//   		isolationMode: jsii.String("isolationMode"),
//   		runAs: &runAsProperty{
//   			gid: jsii.Number(123),
//   			uid: jsii.Number(123),
//   		},
//   	},
//   	resourceAccessPolicies: []interface{}{
//   		&resourceAccessPolicyProperty{
//   			resourceId: jsii.String("resourceId"),
//
//   			// the properties below are optional
//   			permission: jsii.String("permission"),
//   		},
//   	},
//   	variables: variables,
//   }
//
type CfnFunctionDefinitionVersion_EnvironmentProperty struct {
	// Indicates whether the function is allowed to access the `/sys` directory on the core device, which allows the read device information from `/sys` .
	//
	// > This property applies only to Lambda functions that run in a Greengrass container.
	AccessSysfs interface{} `field:"optional" json:"accessSysfs" yaml:"accessSysfs"`
	// Settings for the Lambda execution environment in AWS IoT Greengrass .
	Execution interface{} `field:"optional" json:"execution" yaml:"execution"`
	// A list of the [resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourceinstance.html) in the group that the function can access, with the corresponding read-only or read-write permissions. The maximum is 10 resources.
	//
	// > This property applies only to Lambda functions that run in a Greengrass container.
	ResourceAccessPolicies interface{} `field:"optional" json:"resourceAccessPolicies" yaml:"resourceAccessPolicies"`
	// Environment variables for the Lambda function.
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

// Configuration settings for the Lambda execution environment on the AWS IoT Greengrass core.
//
// In an AWS CloudFormation template, `Execution` is a property of the [`DefaultConfig`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-defaultconfig.html) property type for a function definition version and the [`Environment`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html) property type for a function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executionProperty := &executionProperty{
//   	isolationMode: jsii.String("isolationMode"),
//   	runAs: &runAsProperty{
//   		gid: jsii.Number(123),
//   		uid: jsii.Number(123),
//   	},
//   }
//
type CfnFunctionDefinitionVersion_ExecutionProperty struct {
	// The containerization that the Lambda function runs in.
	//
	// Valid values are `GreengrassContainer` or `NoContainer` . Typically, this is `GreengrassContainer` . For more information, see [Containerization](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html#lambda-function-containerization) in the *Developer Guide* .
	//
	// - When set on the [`DefaultConfig`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-defaultconfig.html) property of a function definition version, this setting is used as the default containerization for all Lambda functions in the function definition version.
	// - When set on the [`Environment`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html) property of a function, this setting applies to the individual function and overrides the default. Omit this value to run the function with the default containerization.
	//
	// > We recommend that you run in a Greengrass container unless your business case requires that you run without containerization.
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
	// The user and group permissions used to run the Lambda function.
	//
	// Typically, this is the ggc_user and ggc_group. For more information, see [Run as](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html#lambda-access-identity.html) in the *Developer Guide* .
	//
	// - When set on the [`DefaultConfig`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-defaultconfig.html) property of a function definition version, this setting is used as the default access identity for all Lambda functions in the function definition version.
	// - When set on the [`Environment`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html) property of a function, this setting applies to the individual function and overrides the default. You can override the user, group, or both. Omit this value to run the function with the default permissions.
	//
	// > Running as the root user increases risks to your data and device. Do not run as root (UID/GID=0) unless your business case requires it. For more information and requirements, see [Running a Lambda Function as Root](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html#lambda-running-as-root) .
	RunAs interface{} `field:"optional" json:"runAs" yaml:"runAs"`
}

// The group-specific configuration settings for a Lambda function.
//
// These settings configure the function's behavior in the Greengrass group. For more information, see [Controlling Execution of Greengrass Lambda Functions by Using Group-Specific Configuration](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `FunctionConfiguration` is a property of the [`Function`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-function.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   functionConfigurationProperty := &functionConfigurationProperty{
//   	encodingType: jsii.String("encodingType"),
//   	environment: &environmentProperty{
//   		accessSysfs: jsii.Boolean(false),
//   		execution: &executionProperty{
//   			isolationMode: jsii.String("isolationMode"),
//   			runAs: &runAsProperty{
//   				gid: jsii.Number(123),
//   				uid: jsii.Number(123),
//   			},
//   		},
//   		resourceAccessPolicies: []interface{}{
//   			&resourceAccessPolicyProperty{
//   				resourceId: jsii.String("resourceId"),
//
//   				// the properties below are optional
//   				permission: jsii.String("permission"),
//   			},
//   		},
//   		variables: variables,
//   	},
//   	execArgs: jsii.String("execArgs"),
//   	executable: jsii.String("executable"),
//   	memorySize: jsii.Number(123),
//   	pinned: jsii.Boolean(false),
//   	timeout: jsii.Number(123),
//   }
//
type CfnFunctionDefinitionVersion_FunctionConfigurationProperty struct {
	// The expected encoding type of the input payload for the function.
	//
	// Valid values are `json` (default) and `binary` .
	EncodingType *string `field:"optional" json:"encodingType" yaml:"encodingType"`
	// The environment configuration of the function.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The execution arguments.
	ExecArgs *string `field:"optional" json:"execArgs" yaml:"execArgs"`
	// The name of the function executable.
	Executable *string `field:"optional" json:"executable" yaml:"executable"`
	// The memory size (in KB) required by the function.
	//
	// > This property applies only to Lambda functions that run in a Greengrass container.
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// Indicates whether the function is pinned (or *long-lived* ).
	//
	// Pinned functions start when the core starts and process all requests in the same container. The default value is false.
	Pinned interface{} `field:"optional" json:"pinned" yaml:"pinned"`
	// The allowed execution time (in seconds) after which the function should terminate.
	//
	// For pinned functions, this timeout applies for each request.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

// A function is a Lambda function that's referenced from an AWS IoT Greengrass group.
//
// The function is deployed to a Greengrass core where it runs locally. For more information, see [Run Lambda Functions on the AWS IoT Greengrass Core](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-functions.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Functions` property of the [`AWS::Greengrass::FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html) resource contains a list of `Function` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   functionProperty := &functionProperty{
//   	functionArn: jsii.String("functionArn"),
//   	functionConfiguration: &functionConfigurationProperty{
//   		encodingType: jsii.String("encodingType"),
//   		environment: &environmentProperty{
//   			accessSysfs: jsii.Boolean(false),
//   			execution: &executionProperty{
//   				isolationMode: jsii.String("isolationMode"),
//   				runAs: &runAsProperty{
//   					gid: jsii.Number(123),
//   					uid: jsii.Number(123),
//   				},
//   			},
//   			resourceAccessPolicies: []interface{}{
//   				&resourceAccessPolicyProperty{
//   					resourceId: jsii.String("resourceId"),
//
//   					// the properties below are optional
//   					permission: jsii.String("permission"),
//   				},
//   			},
//   			variables: variables,
//   		},
//   		execArgs: jsii.String("execArgs"),
//   		executable: jsii.String("executable"),
//   		memorySize: jsii.Number(123),
//   		pinned: jsii.Boolean(false),
//   		timeout: jsii.Number(123),
//   	},
//   	id: jsii.String("id"),
//   }
//
type CfnFunctionDefinitionVersion_FunctionProperty struct {
	// The Amazon Resource Name (ARN) of the alias (recommended) or version of the referenced Lambda function.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The group-specific settings of the Lambda function.
	//
	// These settings configure the function's behavior in the Greengrass group.
	FunctionConfiguration interface{} `field:"required" json:"functionConfiguration" yaml:"functionConfiguration"`
	// A descriptive or arbitrary ID for the function.
	//
	// This value must be unique within the function definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
}

// A list of the [resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourceinstance.html) in the group that the function can access, with the corresponding read-only or read-write permissions. The maximum is 10 resources.
//
// > This property applies only to Lambda functions that run in a Greengrass container.
//
// In an AWS CloudFormation template, `ResourceAccessPolicy` is a property of the [`Environment`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceAccessPolicyProperty := &resourceAccessPolicyProperty{
//   	resourceId: jsii.String("resourceId"),
//
//   	// the properties below are optional
//   	permission: jsii.String("permission"),
//   }
//
type CfnFunctionDefinitionVersion_ResourceAccessPolicyProperty struct {
	// The ID of the resource.
	//
	// This ID is assigned to the resource when you create the resource definition.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The read-only or read-write access that the Lambda function has to the resource.
	//
	// Valid values are `ro` or `rw` .
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

// The user and group permissions used to run the Lambda function.
//
// This setting overrides the default access identity that's specified for the group (by default, ggc_user and ggc_group). You can override the user, group, or both. For more information, see [Run as](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html#lambda-access-identity.html) in the *Developer Guide* .
//
// > Running as the root user increases risks to your data and device. Do not run as root (UID/GID=0) unless your business case requires it. For more information and requirements, see [Running a Lambda Function as Root](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html#lambda-running-as-root) .
//
// In an AWS CloudFormation template, `RunAs` is a property of the [`Execution`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-execution.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runAsProperty := &runAsProperty{
//   	gid: jsii.Number(123),
//   	uid: jsii.Number(123),
//   }
//
type CfnFunctionDefinitionVersion_RunAsProperty struct {
	// The group ID whose permissions are used to run the Lambda function.
	//
	// You can use the `getent group` command on your core device to look up the group ID.
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
	// The user ID whose permissions are used to run the Lambda function.
	//
	// You can use the `getent passwd` command on your core device to look up the user ID.
	Uid *float64 `field:"optional" json:"uid" yaml:"uid"`
}

// Properties for defining a `CfnFunctionDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   cfnFunctionDefinitionVersionProps := &cfnFunctionDefinitionVersionProps{
//   	functionDefinitionId: jsii.String("functionDefinitionId"),
//   	functions: []interface{}{
//   		&functionProperty{
//   			functionArn: jsii.String("functionArn"),
//   			functionConfiguration: &functionConfigurationProperty{
//   				encodingType: jsii.String("encodingType"),
//   				environment: &environmentProperty{
//   					accessSysfs: jsii.Boolean(false),
//   					execution: &executionProperty{
//   						isolationMode: jsii.String("isolationMode"),
//   						runAs: &runAsProperty{
//   							gid: jsii.Number(123),
//   							uid: jsii.Number(123),
//   						},
//   					},
//   					resourceAccessPolicies: []interface{}{
//   						&resourceAccessPolicyProperty{
//   							resourceId: jsii.String("resourceId"),
//
//   							// the properties below are optional
//   							permission: jsii.String("permission"),
//   						},
//   					},
//   					variables: variables,
//   				},
//   				execArgs: jsii.String("execArgs"),
//   				executable: jsii.String("executable"),
//   				memorySize: jsii.Number(123),
//   				pinned: jsii.Boolean(false),
//   				timeout: jsii.Number(123),
//   			},
//   			id: jsii.String("id"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	defaultConfig: &defaultConfigProperty{
//   		execution: &executionProperty{
//   			isolationMode: jsii.String("isolationMode"),
//   			runAs: &runAsProperty{
//   				gid: jsii.Number(123),
//   				uid: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnFunctionDefinitionVersionProps struct {
	// The ID of the function definition associated with this version.
	//
	// This value is a GUID.
	FunctionDefinitionId *string `field:"required" json:"functionDefinitionId" yaml:"functionDefinitionId"`
	// The functions in this version.
	Functions interface{} `field:"required" json:"functions" yaml:"functions"`
	// The default configuration that applies to all Lambda functions in the group.
	//
	// Individual Lambda functions can override these settings.
	DefaultConfig interface{} `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
}

// A CloudFormation `AWS::Greengrass::Group`.
//
// AWS IoT Greengrass seamlessly extends AWS to edge devices so they can act locally on the data they generate, while still using the cloud for management, analytics, and durable storage. With AWS IoT Greengrass , connected devices can run AWS Lambda functions, execute predictions based on machine learning models, keep device data in sync, and communicate with other devices securely – even when not connected to the internet. For more information, see the [Developer Guide](https://docs.aws.amazon.com/greengrass/latest/developerguide/what-is-gg.html) .
//
// > For AWS Region support, see [AWS CloudFormation Support for AWS IoT Greengrass](https://docs.aws.amazon.com/greengrass/latest/developerguide/cloudformation-support.html) in the *Developer Guide* .
//
// The `AWS::Greengrass::Group` resource represents a group in AWS IoT Greengrass . In the AWS IoT Greengrass API, groups are used to organize your group versions.
//
// Groups can reference multiple group versions. All group versions must be associated with a group. A group version references a device definition version, subscription definition version, and other version types that contain the components you want to deploy to a Greengrass core device.
//
// To deploy a group version, the group version must reference a core definition version that contains one core. Other version types are optionally included, depending on your business need.
//
// > When you create a group, you can optionally include an initial group version. To associate a group version later, create a [`AWS::Greengrass::GroupVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html) resource and specify the ID of this group.
// >
// > To change group components (such as devices, subscriptions, or functions), you must create new versions. This is because versions are immutable. For example, to add a function, you create a function definition version that contains the new function (and all other functions that you want to deploy). Then you create a group version that references the new function definition version (and all other version types that you want to deploy).
//
// *Deploying a Group Version*
//
// After you create the group version in your AWS CloudFormation template, you can deploy it using the [`aws greengrass create-deployment`](https://docs.aws.amazon.com/greengrass/latest/apireference/createdeployment-post.html) command in the AWS CLI or from the *Greengrass* node in the AWS IoT console. To deploy a group version, you must have a Greengrass service role associated with your AWS account . For more information, see [AWS CloudFormation Support for AWS IoT Greengrass](https://docs.aws.amazon.com/greengrass/latest/developerguide/cloudformation-support.html) in the *Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnGroup := awscdk.Aws_greengrass.NewCfnGroup(this, jsii.String("MyCfnGroup"), &cfnGroupProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &groupVersionProperty{
//   		connectorDefinitionVersionArn: jsii.String("connectorDefinitionVersionArn"),
//   		coreDefinitionVersionArn: jsii.String("coreDefinitionVersionArn"),
//   		deviceDefinitionVersionArn: jsii.String("deviceDefinitionVersionArn"),
//   		functionDefinitionVersionArn: jsii.String("functionDefinitionVersionArn"),
//   		loggerDefinitionVersionArn: jsii.String("loggerDefinitionVersionArn"),
//   		resourceDefinitionVersionArn: jsii.String("resourceDefinitionVersionArn"),
//   		subscriptionDefinitionVersionArn: jsii.String("subscriptionDefinitionVersionArn"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//   	tags: tags,
//   })
//
type CfnGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the `Group` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/groups/1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrArn() *string
	// The ID of the `Group` , such as `1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrId() *string
	// The ARN of the last `GroupVersion` that was added to the `Group` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/groups/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2` .
	AttrLatestVersionArn() *string
	// The name of the `Group` , such as `MyGroup` .
	AttrName() *string
	// The ARN of the IAM role that's attached to the `Group` , such as `arn:aws:iam::  :role/role-name` .
	AttrRoleArn() *string
	// The time (in milliseconds since the epoch) when the group role was attached to the `Group` .
	AttrRoleAttachedAt() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The group version to include when the group is created.
	//
	// A group version references the Amazon Resource Name (ARN) of a core definition version, device definition version, subscription definition version, and other version types. The group version must reference a core definition version that contains one core. Other version types are optionally included, depending on your business need.
	//
	// > To associate a group version after the group is created, create an [`AWS::Greengrass::GroupVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html) resource and specify the ID of this group.
	InitialVersion() interface{}
	SetInitialVersion(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the group.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the IAM role attached to the group.
	//
	// This role contains the permissions that Lambda functions and connectors use to interact with other AWS services.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Application-specific metadata to attach to the group.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnGroup
type jsiiProxy_CfnGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) AttrLatestVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLatestVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) AttrRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) AttrRoleAttachedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRoleAttachedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) InitialVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::Group`.
func NewCfnGroup(scope constructs.Construct, id *string, props *CfnGroupProps) CfnGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::Group`.
func NewCfnGroup_Override(c CfnGroup, scope constructs.Construct, id *string, props *CfnGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGroup) SetInitialVersion(val interface{}) {
	_jsii_.Set(
		j,
		"initialVersion",
		val,
	)
}

func (j *jsiiProxy_CfnGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnGroup) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A group version in AWS IoT Greengrass , which references of a core definition version, device definition version, subscription definition version, and other version types that contain the components you want to deploy to a Greengrass core device.
//
// The group version must reference a core definition version that contains one core. Other version types are optionally included, depending on your business need.
//
// In an AWS CloudFormation template, `GroupVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupVersionProperty := &groupVersionProperty{
//   	connectorDefinitionVersionArn: jsii.String("connectorDefinitionVersionArn"),
//   	coreDefinitionVersionArn: jsii.String("coreDefinitionVersionArn"),
//   	deviceDefinitionVersionArn: jsii.String("deviceDefinitionVersionArn"),
//   	functionDefinitionVersionArn: jsii.String("functionDefinitionVersionArn"),
//   	loggerDefinitionVersionArn: jsii.String("loggerDefinitionVersionArn"),
//   	resourceDefinitionVersionArn: jsii.String("resourceDefinitionVersionArn"),
//   	subscriptionDefinitionVersionArn: jsii.String("subscriptionDefinitionVersionArn"),
//   }
//
type CfnGroup_GroupVersionProperty struct {
	// The Amazon Resource Name (ARN) of the connector definition version that contains the connectors you want to deploy with the group version.
	ConnectorDefinitionVersionArn *string `field:"optional" json:"connectorDefinitionVersionArn" yaml:"connectorDefinitionVersionArn"`
	// The ARN of the core definition version that contains the core you want to deploy with the group version.
	//
	// Currently, the core definition version can contain only one core.
	CoreDefinitionVersionArn *string `field:"optional" json:"coreDefinitionVersionArn" yaml:"coreDefinitionVersionArn"`
	// The ARN of the device definition version that contains the devices you want to deploy with the group version.
	DeviceDefinitionVersionArn *string `field:"optional" json:"deviceDefinitionVersionArn" yaml:"deviceDefinitionVersionArn"`
	// The ARN of the function definition version that contains the functions you want to deploy with the group version.
	FunctionDefinitionVersionArn *string `field:"optional" json:"functionDefinitionVersionArn" yaml:"functionDefinitionVersionArn"`
	// The ARN of the logger definition version that contains the loggers you want to deploy with the group version.
	LoggerDefinitionVersionArn *string `field:"optional" json:"loggerDefinitionVersionArn" yaml:"loggerDefinitionVersionArn"`
	// The ARN of the resource definition version that contains the resources you want to deploy with the group version.
	ResourceDefinitionVersionArn *string `field:"optional" json:"resourceDefinitionVersionArn" yaml:"resourceDefinitionVersionArn"`
	// The ARN of the subscription definition version that contains the subscriptions you want to deploy with the group version.
	SubscriptionDefinitionVersionArn *string `field:"optional" json:"subscriptionDefinitionVersionArn" yaml:"subscriptionDefinitionVersionArn"`
}

// Properties for defining a `CfnGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnGroupProps := &cfnGroupProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &groupVersionProperty{
//   		connectorDefinitionVersionArn: jsii.String("connectorDefinitionVersionArn"),
//   		coreDefinitionVersionArn: jsii.String("coreDefinitionVersionArn"),
//   		deviceDefinitionVersionArn: jsii.String("deviceDefinitionVersionArn"),
//   		functionDefinitionVersionArn: jsii.String("functionDefinitionVersionArn"),
//   		loggerDefinitionVersionArn: jsii.String("loggerDefinitionVersionArn"),
//   		resourceDefinitionVersionArn: jsii.String("resourceDefinitionVersionArn"),
//   		subscriptionDefinitionVersionArn: jsii.String("subscriptionDefinitionVersionArn"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//   	tags: tags,
//   }
//
type CfnGroupProps struct {
	// The name of the group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The group version to include when the group is created.
	//
	// A group version references the Amazon Resource Name (ARN) of a core definition version, device definition version, subscription definition version, and other version types. The group version must reference a core definition version that contains one core. Other version types are optionally included, depending on your business need.
	//
	// > To associate a group version after the group is created, create an [`AWS::Greengrass::GroupVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html) resource and specify the ID of this group.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// The Amazon Resource Name (ARN) of the IAM role attached to the group.
	//
	// This role contains the permissions that Lambda functions and connectors use to interact with other AWS services.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Application-specific metadata to attach to the group.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Greengrass::GroupVersion`.
//
// The `AWS::Greengrass::GroupVersion` resource represents a group version in AWS IoT Greengrass . A group version references a core definition version, device definition version, subscription definition version, and other version types that contain the components you want to deploy to a Greengrass core device. The group version must reference a core definition version that contains one core. Other version types are optionally included, depending on your business need.
//
// > To create a group version, you must specify the ID of the group that you want to associate with the version. For information about creating a group, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupVersion := awscdk.Aws_greengrass.NewCfnGroupVersion(this, jsii.String("MyCfnGroupVersion"), &cfnGroupVersionProps{
//   	groupId: jsii.String("groupId"),
//
//   	// the properties below are optional
//   	connectorDefinitionVersionArn: jsii.String("connectorDefinitionVersionArn"),
//   	coreDefinitionVersionArn: jsii.String("coreDefinitionVersionArn"),
//   	deviceDefinitionVersionArn: jsii.String("deviceDefinitionVersionArn"),
//   	functionDefinitionVersionArn: jsii.String("functionDefinitionVersionArn"),
//   	loggerDefinitionVersionArn: jsii.String("loggerDefinitionVersionArn"),
//   	resourceDefinitionVersionArn: jsii.String("resourceDefinitionVersionArn"),
//   	subscriptionDefinitionVersionArn: jsii.String("subscriptionDefinitionVersionArn"),
//   })
//
type CfnGroupVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The Amazon Resource Name (ARN) of the connector definition version that contains the connectors you want to deploy with the group version.
	ConnectorDefinitionVersionArn() *string
	SetConnectorDefinitionVersionArn(val *string)
	// The ARN of the core definition version that contains the core you want to deploy with the group version.
	//
	// Currently, the core definition version can contain only one core.
	CoreDefinitionVersionArn() *string
	SetCoreDefinitionVersionArn(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ARN of the device definition version that contains the devices you want to deploy with the group version.
	DeviceDefinitionVersionArn() *string
	SetDeviceDefinitionVersionArn(val *string)
	// The ARN of the function definition version that contains the functions you want to deploy with the group version.
	FunctionDefinitionVersionArn() *string
	SetFunctionDefinitionVersionArn(val *string)
	// The ID of the group associated with this version.
	//
	// This value is a GUID.
	GroupId() *string
	SetGroupId(val *string)
	// The ARN of the logger definition version that contains the loggers you want to deploy with the group version.
	LoggerDefinitionVersionArn() *string
	SetLoggerDefinitionVersionArn(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ARN of the resource definition version that contains the resources you want to deploy with the group version.
	ResourceDefinitionVersionArn() *string
	SetResourceDefinitionVersionArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The ARN of the subscription definition version that contains the subscriptions you want to deploy with the group version.
	SubscriptionDefinitionVersionArn() *string
	SetSubscriptionDefinitionVersionArn(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnGroupVersion
type jsiiProxy_CfnGroupVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGroupVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) ConnectorDefinitionVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorDefinitionVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) CoreDefinitionVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreDefinitionVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) DeviceDefinitionVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceDefinitionVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) FunctionDefinitionVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionDefinitionVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) LoggerDefinitionVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggerDefinitionVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) ResourceDefinitionVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceDefinitionVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) SubscriptionDefinitionVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionDefinitionVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::GroupVersion`.
func NewCfnGroupVersion(scope constructs.Construct, id *string, props *CfnGroupVersionProps) CfnGroupVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnGroupVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnGroupVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::GroupVersion`.
func NewCfnGroupVersion_Override(c CfnGroupVersion, scope constructs.Construct, id *string, props *CfnGroupVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnGroupVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGroupVersion) SetConnectorDefinitionVersionArn(val *string) {
	_jsii_.Set(
		j,
		"connectorDefinitionVersionArn",
		val,
	)
}

func (j *jsiiProxy_CfnGroupVersion) SetCoreDefinitionVersionArn(val *string) {
	_jsii_.Set(
		j,
		"coreDefinitionVersionArn",
		val,
	)
}

func (j *jsiiProxy_CfnGroupVersion) SetDeviceDefinitionVersionArn(val *string) {
	_jsii_.Set(
		j,
		"deviceDefinitionVersionArn",
		val,
	)
}

func (j *jsiiProxy_CfnGroupVersion) SetFunctionDefinitionVersionArn(val *string) {
	_jsii_.Set(
		j,
		"functionDefinitionVersionArn",
		val,
	)
}

func (j *jsiiProxy_CfnGroupVersion) SetGroupId(val *string) {
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_CfnGroupVersion) SetLoggerDefinitionVersionArn(val *string) {
	_jsii_.Set(
		j,
		"loggerDefinitionVersionArn",
		val,
	)
}

func (j *jsiiProxy_CfnGroupVersion) SetResourceDefinitionVersionArn(val *string) {
	_jsii_.Set(
		j,
		"resourceDefinitionVersionArn",
		val,
	)
}

func (j *jsiiProxy_CfnGroupVersion) SetSubscriptionDefinitionVersionArn(val *string) {
	_jsii_.Set(
		j,
		"subscriptionDefinitionVersionArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnGroupVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnGroupVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnGroupVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnGroupVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnGroupVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnGroupVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGroupVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnGroupVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGroupVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGroupVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGroupVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGroupVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGroupVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGroupVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGroupVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnGroupVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroupVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroupVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnGroupVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGroupVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroupVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroupVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroupVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnGroupVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupVersionProps := &cfnGroupVersionProps{
//   	groupId: jsii.String("groupId"),
//
//   	// the properties below are optional
//   	connectorDefinitionVersionArn: jsii.String("connectorDefinitionVersionArn"),
//   	coreDefinitionVersionArn: jsii.String("coreDefinitionVersionArn"),
//   	deviceDefinitionVersionArn: jsii.String("deviceDefinitionVersionArn"),
//   	functionDefinitionVersionArn: jsii.String("functionDefinitionVersionArn"),
//   	loggerDefinitionVersionArn: jsii.String("loggerDefinitionVersionArn"),
//   	resourceDefinitionVersionArn: jsii.String("resourceDefinitionVersionArn"),
//   	subscriptionDefinitionVersionArn: jsii.String("subscriptionDefinitionVersionArn"),
//   }
//
type CfnGroupVersionProps struct {
	// The ID of the group associated with this version.
	//
	// This value is a GUID.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The Amazon Resource Name (ARN) of the connector definition version that contains the connectors you want to deploy with the group version.
	ConnectorDefinitionVersionArn *string `field:"optional" json:"connectorDefinitionVersionArn" yaml:"connectorDefinitionVersionArn"`
	// The ARN of the core definition version that contains the core you want to deploy with the group version.
	//
	// Currently, the core definition version can contain only one core.
	CoreDefinitionVersionArn *string `field:"optional" json:"coreDefinitionVersionArn" yaml:"coreDefinitionVersionArn"`
	// The ARN of the device definition version that contains the devices you want to deploy with the group version.
	DeviceDefinitionVersionArn *string `field:"optional" json:"deviceDefinitionVersionArn" yaml:"deviceDefinitionVersionArn"`
	// The ARN of the function definition version that contains the functions you want to deploy with the group version.
	FunctionDefinitionVersionArn *string `field:"optional" json:"functionDefinitionVersionArn" yaml:"functionDefinitionVersionArn"`
	// The ARN of the logger definition version that contains the loggers you want to deploy with the group version.
	LoggerDefinitionVersionArn *string `field:"optional" json:"loggerDefinitionVersionArn" yaml:"loggerDefinitionVersionArn"`
	// The ARN of the resource definition version that contains the resources you want to deploy with the group version.
	ResourceDefinitionVersionArn *string `field:"optional" json:"resourceDefinitionVersionArn" yaml:"resourceDefinitionVersionArn"`
	// The ARN of the subscription definition version that contains the subscriptions you want to deploy with the group version.
	SubscriptionDefinitionVersionArn *string `field:"optional" json:"subscriptionDefinitionVersionArn" yaml:"subscriptionDefinitionVersionArn"`
}

// A CloudFormation `AWS::Greengrass::LoggerDefinition`.
//
// The `AWS::Greengrass::LoggerDefinition` resource represents a logger definition for AWS IoT Greengrass . Logger definitions are used to organize your logger definition versions.
//
// Logger definitions can reference multiple logger definition versions. All logger definition versions must be associated with a logger definition. Each logger definition version can contain one or more loggers.
//
// > When you create a logger definition, you can optionally include an initial logger definition version. To associate a logger definition version later, create an [`AWS::Greengrass::LoggerDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html) resource and specify the ID of this logger definition.
// >
// > After you create the logger definition version that contains the loggers you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnLoggerDefinition := awscdk.Aws_greengrass.NewCfnLoggerDefinition(this, jsii.String("MyCfnLoggerDefinition"), &cfnLoggerDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &loggerDefinitionVersionProperty{
//   		loggers: []interface{}{
//   			&loggerProperty{
//   				component: jsii.String("component"),
//   				id: jsii.String("id"),
//   				level: jsii.String("level"),
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				space: jsii.Number(123),
//   			},
//   		},
//   	},
//   	tags: tags,
//   })
//
type CfnLoggerDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the `LoggerDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/loggers/1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrArn() *string
	// The ID of the `LoggerDefinition` , such as `1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrId() *string
	// The ARN of the last `LoggerDefinitionVersion` that was added to the `LoggerDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/loggers/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2` .
	AttrLatestVersionArn() *string
	// The name of the `LoggerDefinition` , such as `MyLoggerDefinition` .
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logger definition version to include when the logger definition is created.
	//
	// A logger definition version contains a list of [`logger`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-loggerdefinition-logger.html) property types.
	//
	// > To associate a logger definition version after the logger definition is created, create an [`AWS::Greengrass::LoggerDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html) resource and specify the ID of this logger definition.
	InitialVersion() interface{}
	SetInitialVersion(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the logger definition.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Application-specific metadata to attach to the logger definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnLoggerDefinition
type jsiiProxy_CfnLoggerDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLoggerDefinition) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) AttrLatestVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLatestVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) InitialVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::LoggerDefinition`.
func NewCfnLoggerDefinition(scope constructs.Construct, id *string, props *CfnLoggerDefinitionProps) CfnLoggerDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnLoggerDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnLoggerDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::LoggerDefinition`.
func NewCfnLoggerDefinition_Override(c CfnLoggerDefinition, scope constructs.Construct, id *string, props *CfnLoggerDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnLoggerDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLoggerDefinition) SetInitialVersion(val interface{}) {
	_jsii_.Set(
		j,
		"initialVersion",
		val,
	)
}

func (j *jsiiProxy_CfnLoggerDefinition) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLoggerDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnLoggerDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLoggerDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnLoggerDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnLoggerDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnLoggerDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoggerDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnLoggerDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLoggerDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLoggerDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLoggerDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLoggerDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLoggerDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLoggerDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLoggerDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLoggerDefinition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggerDefinition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggerDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLoggerDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLoggerDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggerDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggerDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggerDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A logger definition version contains a list of [loggers](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-loggerdefinition-logger.html) .
//
// > After you create a logger definition version that contains the loggers you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an AWS CloudFormation template, `LoggerDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::LoggerDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggerDefinitionVersionProperty := &loggerDefinitionVersionProperty{
//   	loggers: []interface{}{
//   		&loggerProperty{
//   			component: jsii.String("component"),
//   			id: jsii.String("id"),
//   			level: jsii.String("level"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			space: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnLoggerDefinition_LoggerDefinitionVersionProperty struct {
	// The loggers in this version.
	Loggers interface{} `field:"required" json:"loggers" yaml:"loggers"`
}

// A logger represents logging settings for the AWS IoT Greengrass group, which can be stored in CloudWatch and the local file system of your core device.
//
// All log entries include a timestamp, log level, and information about the event. For more information, see [Monitoring with AWS IoT Greengrass Logs](https://docs.aws.amazon.com/greengrass/latest/developerguide/greengrass-logs-overview.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Loggers` property of the [`LoggerDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-loggerdefinition-loggerdefinitionversion.html) property type contains a list of `Logger` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggerProperty := &loggerProperty{
//   	component: jsii.String("component"),
//   	id: jsii.String("id"),
//   	level: jsii.String("level"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	space: jsii.Number(123),
//   }
//
type CfnLoggerDefinition_LoggerProperty struct {
	// The source of the log event.
	//
	// Valid values are `GreengrassSystem` or `Lambda` . When `GreengrassSystem` is used, events from Greengrass system components are logged. When `Lambda` is used, events from user-defined Lambda functions are logged.
	Component *string `field:"required" json:"component" yaml:"component"`
	// A descriptive or arbitrary ID for the logger.
	//
	// This value must be unique within the logger definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The log-level threshold.
	//
	// Log events below this threshold are filtered out and aren't stored. Valid values are `DEBUG` , `INFO` (recommended), `WARN` , `ERROR` , or `FATAL` .
	Level *string `field:"required" json:"level" yaml:"level"`
	// The storage mechanism for log events.
	//
	// Valid values are `FileSystem` or `AWSCloudWatch` . When `AWSCloudWatch` is used, log events are sent to CloudWatch Logs . When `FileSystem` is used, log events are stored on the local file system.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The amount of file space (in KB) to use when writing logs to the local file system.
	//
	// This property does not apply for CloudWatch Logs .
	Space *float64 `field:"optional" json:"space" yaml:"space"`
}

// Properties for defining a `CfnLoggerDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnLoggerDefinitionProps := &cfnLoggerDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &loggerDefinitionVersionProperty{
//   		loggers: []interface{}{
//   			&loggerProperty{
//   				component: jsii.String("component"),
//   				id: jsii.String("id"),
//   				level: jsii.String("level"),
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				space: jsii.Number(123),
//   			},
//   		},
//   	},
//   	tags: tags,
//   }
//
type CfnLoggerDefinitionProps struct {
	// The name of the logger definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The logger definition version to include when the logger definition is created.
	//
	// A logger definition version contains a list of [`logger`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-loggerdefinition-logger.html) property types.
	//
	// > To associate a logger definition version after the logger definition is created, create an [`AWS::Greengrass::LoggerDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html) resource and specify the ID of this logger definition.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the logger definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Greengrass::LoggerDefinitionVersion`.
//
// The `AWS::Greengrass::LoggerDefinitionVersion` resource represents a logger definition version for AWS IoT Greengrass . A logger definition version contains a list of loggers.
//
// > To create a logger definition version, you must specify the ID of the logger definition that you want to associate with the version. For information about creating a logger definition, see [`AWS::Greengrass::LoggerDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinition.html) .
// >
// > After you create a logger definition version that contains the loggers you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoggerDefinitionVersion := awscdk.Aws_greengrass.NewCfnLoggerDefinitionVersion(this, jsii.String("MyCfnLoggerDefinitionVersion"), &cfnLoggerDefinitionVersionProps{
//   	loggerDefinitionId: jsii.String("loggerDefinitionId"),
//   	loggers: []interface{}{
//   		&loggerProperty{
//   			component: jsii.String("component"),
//   			id: jsii.String("id"),
//   			level: jsii.String("level"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			space: jsii.Number(123),
//   		},
//   	},
//   })
//
type CfnLoggerDefinitionVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ID of the logger definition associated with this version.
	//
	// This value is a GUID.
	LoggerDefinitionId() *string
	SetLoggerDefinitionId(val *string)
	// The loggers in this version.
	Loggers() interface{}
	SetLoggers(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnLoggerDefinitionVersion
type jsiiProxy_CfnLoggerDefinitionVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) LoggerDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggerDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) Loggers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::LoggerDefinitionVersion`.
func NewCfnLoggerDefinitionVersion(scope constructs.Construct, id *string, props *CfnLoggerDefinitionVersionProps) CfnLoggerDefinitionVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnLoggerDefinitionVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnLoggerDefinitionVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::LoggerDefinitionVersion`.
func NewCfnLoggerDefinitionVersion_Override(c CfnLoggerDefinitionVersion, scope constructs.Construct, id *string, props *CfnLoggerDefinitionVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnLoggerDefinitionVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) SetLoggerDefinitionId(val *string) {
	_jsii_.Set(
		j,
		"loggerDefinitionId",
		val,
	)
}

func (j *jsiiProxy_CfnLoggerDefinitionVersion) SetLoggers(val interface{}) {
	_jsii_.Set(
		j,
		"loggers",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLoggerDefinitionVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnLoggerDefinitionVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLoggerDefinitionVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnLoggerDefinitionVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnLoggerDefinitionVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnLoggerDefinitionVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoggerDefinitionVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnLoggerDefinitionVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggerDefinitionVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A logger represents logging settings for the AWS IoT Greengrass group, which can be stored in CloudWatch and the local file system of your core device.
//
// All log entries include a timestamp, log level, and information about the event. For more information, see [Monitoring with AWS IoT Greengrass Logs](https://docs.aws.amazon.com/greengrass/latest/developerguide/greengrass-logs-overview.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Loggers` property of the [`AWS::Greengrass::LoggerDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html) resource contains a list of `Logger` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggerProperty := &loggerProperty{
//   	component: jsii.String("component"),
//   	id: jsii.String("id"),
//   	level: jsii.String("level"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	space: jsii.Number(123),
//   }
//
type CfnLoggerDefinitionVersion_LoggerProperty struct {
	// The source of the log event.
	//
	// Valid values are `GreengrassSystem` or `Lambda` . When `GreengrassSystem` is used, events from Greengrass system components are logged. When `Lambda` is used, events from user-defined Lambda functions are logged.
	Component *string `field:"required" json:"component" yaml:"component"`
	// A descriptive or arbitrary ID for the logger.
	//
	// This value must be unique within the logger definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The log-level threshold.
	//
	// Log events below this threshold are filtered out and aren't stored. Valid values are `DEBUG` , `INFO` (recommended), `WARN` , `ERROR` , or `FATAL` .
	Level *string `field:"required" json:"level" yaml:"level"`
	// The storage mechanism for log events.
	//
	// Valid values are `FileSystem` or `AWSCloudWatch` . When `AWSCloudWatch` is used, log events are sent to CloudWatch Logs . When `FileSystem` is used, log events are stored on the local file system.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The amount of file space (in KB) to use when writing logs to the local file system.
	//
	// This property does not apply for CloudWatch Logs .
	Space *float64 `field:"optional" json:"space" yaml:"space"`
}

// Properties for defining a `CfnLoggerDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoggerDefinitionVersionProps := &cfnLoggerDefinitionVersionProps{
//   	loggerDefinitionId: jsii.String("loggerDefinitionId"),
//   	loggers: []interface{}{
//   		&loggerProperty{
//   			component: jsii.String("component"),
//   			id: jsii.String("id"),
//   			level: jsii.String("level"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			space: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnLoggerDefinitionVersionProps struct {
	// The ID of the logger definition associated with this version.
	//
	// This value is a GUID.
	LoggerDefinitionId *string `field:"required" json:"loggerDefinitionId" yaml:"loggerDefinitionId"`
	// The loggers in this version.
	Loggers interface{} `field:"required" json:"loggers" yaml:"loggers"`
}

// A CloudFormation `AWS::Greengrass::ResourceDefinition`.
//
// The `AWS::Greengrass::ResourceDefinition` resource represents a resource definition for AWS IoT Greengrass . Resource definitions are used to organize your resource definition versions.
//
// Resource definitions can reference multiple resource definition versions. All resource definition versions must be associated with a resource definition. Each resource definition version can contain one or more resources. (In AWS CloudFormation , resources are named *resource instances* .)
//
// > When you create a resource definition, you can optionally include an initial resource definition version. To associate a resource definition version later, create an [`AWS::Greengrass::ResourceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html) resource and specify the ID of this resource definition.
// >
// > After you create the resource definition version that contains the resources you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnResourceDefinition := awscdk.Aws_greengrass.NewCfnResourceDefinition(this, jsii.String("MyCfnResourceDefinition"), &cfnResourceDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &resourceDefinitionVersionProperty{
//   		resources: []interface{}{
//   			&resourceInstanceProperty{
//   				id: jsii.String("id"),
//   				name: jsii.String("name"),
//   				resourceDataContainer: &resourceDataContainerProperty{
//   					localDeviceResourceData: &localDeviceResourceDataProperty{
//   						sourcePath: jsii.String("sourcePath"),
//
//   						// the properties below are optional
//   						groupOwnerSetting: &groupOwnerSettingProperty{
//   							autoAddGroupOwner: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							groupOwner: jsii.String("groupOwner"),
//   						},
//   					},
//   					localVolumeResourceData: &localVolumeResourceDataProperty{
//   						destinationPath: jsii.String("destinationPath"),
//   						sourcePath: jsii.String("sourcePath"),
//
//   						// the properties below are optional
//   						groupOwnerSetting: &groupOwnerSettingProperty{
//   							autoAddGroupOwner: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							groupOwner: jsii.String("groupOwner"),
//   						},
//   					},
//   					s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   						destinationPath: jsii.String("destinationPath"),
//   						s3Uri: jsii.String("s3Uri"),
//
//   						// the properties below are optional
//   						ownerSetting: &resourceDownloadOwnerSettingProperty{
//   							groupOwner: jsii.String("groupOwner"),
//   							groupPermission: jsii.String("groupPermission"),
//   						},
//   					},
//   					sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   						destinationPath: jsii.String("destinationPath"),
//   						sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   						// the properties below are optional
//   						ownerSetting: &resourceDownloadOwnerSettingProperty{
//   							groupOwner: jsii.String("groupOwner"),
//   							groupPermission: jsii.String("groupPermission"),
//   						},
//   					},
//   					secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   						arn: jsii.String("arn"),
//
//   						// the properties below are optional
//   						additionalStagingLabelsToDownload: []*string{
//   							jsii.String("additionalStagingLabelsToDownload"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	tags: tags,
//   })
//
type CfnResourceDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the `ResourceDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/resources/1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrArn() *string
	// The ID of the `ResourceDefinition` , such as `1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrId() *string
	// The ARN of the last `ResourceDefinitionVersion` that was added to the `ResourceDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/resources/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2` .
	AttrLatestVersionArn() *string
	// The name of the `ResourceDefinition` , such as `MyResourceDefinition` .
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The resource definition version to include when the resource definition is created.
	//
	// A resource definition version contains a list of [`resource instance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html) property types.
	//
	// > To associate a resource definition version after the resource definition is created, create an [`AWS::Greengrass::ResourceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html) resource and specify the ID of this resource definition.
	InitialVersion() interface{}
	SetInitialVersion(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the resource definition.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Application-specific metadata to attach to the resource definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnResourceDefinition
type jsiiProxy_CfnResourceDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResourceDefinition) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) AttrLatestVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLatestVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) InitialVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::ResourceDefinition`.
func NewCfnResourceDefinition(scope constructs.Construct, id *string, props *CfnResourceDefinitionProps) CfnResourceDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnResourceDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnResourceDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::ResourceDefinition`.
func NewCfnResourceDefinition_Override(c CfnResourceDefinition, scope constructs.Construct, id *string, props *CfnResourceDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnResourceDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResourceDefinition) SetInitialVersion(val interface{}) {
	_jsii_.Set(
		j,
		"initialVersion",
		val,
	)
}

func (j *jsiiProxy_CfnResourceDefinition) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResourceDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnResourceDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResourceDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnResourceDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnResourceDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnResourceDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnResourceDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResourceDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResourceDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResourceDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResourceDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResourceDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResourceDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResourceDefinition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefinition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResourceDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResourceDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Settings that define additional Linux OS group permissions to give to the Lambda function process.
//
// You can give the permissions of the Linux group that owns the resource or choose another Linux group. These permissions are in addition to the function's `RunAs` permissions.
//
// In an AWS CloudFormation template, `GroupOwnerSetting` is a property of the [`LocalDeviceResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-localdeviceresourcedata.html) and [`LocalVolumeResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-localvolumeresourcedata.html) property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupOwnerSettingProperty := &groupOwnerSettingProperty{
//   	autoAddGroupOwner: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	groupOwner: jsii.String("groupOwner"),
//   }
//
type CfnResourceDefinition_GroupOwnerSettingProperty struct {
	// Indicates whether to give the privileges of the Linux group that owns the resource to the Lambda process.
	//
	// This gives the Lambda process the file access permissions of the Linux group.
	AutoAddGroupOwner interface{} `field:"required" json:"autoAddGroupOwner" yaml:"autoAddGroupOwner"`
	// The name of the Linux group whose privileges you want to add to the Lambda process.
	//
	// This value is ignored if `AutoAddGroupOwner` is true.
	GroupOwner *string `field:"optional" json:"groupOwner" yaml:"groupOwner"`
}

// Settings for a local device resource, which represents a file under `/dev` .
//
// For more information, see [Access Local Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-local-resources.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `LocalDeviceResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localDeviceResourceDataProperty := &localDeviceResourceDataProperty{
//   	sourcePath: jsii.String("sourcePath"),
//
//   	// the properties below are optional
//   	groupOwnerSetting: &groupOwnerSettingProperty{
//   		autoAddGroupOwner: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		groupOwner: jsii.String("groupOwner"),
//   	},
//   }
//
type CfnResourceDefinition_LocalDeviceResourceDataProperty struct {
	// The local absolute path of the device resource.
	//
	// The source path for a device resource can refer only to a character device or block device under `/dev` .
	SourcePath *string `field:"required" json:"sourcePath" yaml:"sourcePath"`
	// Settings that define additional Linux OS group permissions to give to the Lambda function process.
	GroupOwnerSetting interface{} `field:"optional" json:"groupOwnerSetting" yaml:"groupOwnerSetting"`
}

// Settings for a local volume resource, which represents a file or directory on the root file system.
//
// For more information, see [Access Local Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-local-resources.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `LocalVolumeResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localVolumeResourceDataProperty := &localVolumeResourceDataProperty{
//   	destinationPath: jsii.String("destinationPath"),
//   	sourcePath: jsii.String("sourcePath"),
//
//   	// the properties below are optional
//   	groupOwnerSetting: &groupOwnerSettingProperty{
//   		autoAddGroupOwner: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		groupOwner: jsii.String("groupOwner"),
//   	},
//   }
//
type CfnResourceDefinition_LocalVolumeResourceDataProperty struct {
	// The absolute local path of the resource in the Lambda environment.
	DestinationPath *string `field:"required" json:"destinationPath" yaml:"destinationPath"`
	// The local absolute path of the volume resource on the host.
	//
	// The source path for a volume resource type cannot start with `/sys` .
	SourcePath *string `field:"required" json:"sourcePath" yaml:"sourcePath"`
	// Settings that define additional Linux OS group permissions to give to the Lambda function process.
	GroupOwnerSetting interface{} `field:"optional" json:"groupOwnerSetting" yaml:"groupOwnerSetting"`
}

// A container for resource data, which defines the resource type.
//
// The container takes only one of the following supported resource data types: `LocalDeviceResourceData` , `LocalVolumeResourceData` , `SageMakerMachineLearningModelResourceData` , `S3MachineLearningModelResourceData` , or `SecretsManagerSecretResourceData` .
//
// > Only one resource type can be defined for a `ResourceDataContainer` instance.
//
// In an AWS CloudFormation template, `ResourceDataContainer` is a property of the [`ResourceInstance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceDataContainerProperty := &resourceDataContainerProperty{
//   	localDeviceResourceData: &localDeviceResourceDataProperty{
//   		sourcePath: jsii.String("sourcePath"),
//
//   		// the properties below are optional
//   		groupOwnerSetting: &groupOwnerSettingProperty{
//   			autoAddGroupOwner: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			groupOwner: jsii.String("groupOwner"),
//   		},
//   	},
//   	localVolumeResourceData: &localVolumeResourceDataProperty{
//   		destinationPath: jsii.String("destinationPath"),
//   		sourcePath: jsii.String("sourcePath"),
//
//   		// the properties below are optional
//   		groupOwnerSetting: &groupOwnerSettingProperty{
//   			autoAddGroupOwner: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			groupOwner: jsii.String("groupOwner"),
//   		},
//   	},
//   	s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   		destinationPath: jsii.String("destinationPath"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		ownerSetting: &resourceDownloadOwnerSettingProperty{
//   			groupOwner: jsii.String("groupOwner"),
//   			groupPermission: jsii.String("groupPermission"),
//   		},
//   	},
//   	sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   		destinationPath: jsii.String("destinationPath"),
//   		sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   		// the properties below are optional
//   		ownerSetting: &resourceDownloadOwnerSettingProperty{
//   			groupOwner: jsii.String("groupOwner"),
//   			groupPermission: jsii.String("groupPermission"),
//   		},
//   	},
//   	secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   		arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		additionalStagingLabelsToDownload: []*string{
//   			jsii.String("additionalStagingLabelsToDownload"),
//   		},
//   	},
//   }
//
type CfnResourceDefinition_ResourceDataContainerProperty struct {
	// Settings for a local device resource.
	LocalDeviceResourceData interface{} `field:"optional" json:"localDeviceResourceData" yaml:"localDeviceResourceData"`
	// Settings for a local volume resource.
	LocalVolumeResourceData interface{} `field:"optional" json:"localVolumeResourceData" yaml:"localVolumeResourceData"`
	// Settings for a machine learning resource stored in Amazon S3 .
	S3MachineLearningModelResourceData interface{} `field:"optional" json:"s3MachineLearningModelResourceData" yaml:"s3MachineLearningModelResourceData"`
	// Settings for a machine learning resource saved as an SageMaker training job.
	SageMakerMachineLearningModelResourceData interface{} `field:"optional" json:"sageMakerMachineLearningModelResourceData" yaml:"sageMakerMachineLearningModelResourceData"`
	// Settings for a secret resource.
	SecretsManagerSecretResourceData interface{} `field:"optional" json:"secretsManagerSecretResourceData" yaml:"secretsManagerSecretResourceData"`
}

// A resource definition version contains a list of resources. (In AWS CloudFormation , resources are named *resource instances* .).
//
// > After you create a resource definition version that contains the resources you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an AWS CloudFormation template, `ResourceDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::ResourceDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceDefinitionVersionProperty := &resourceDefinitionVersionProperty{
//   	resources: []interface{}{
//   		&resourceInstanceProperty{
//   			id: jsii.String("id"),
//   			name: jsii.String("name"),
//   			resourceDataContainer: &resourceDataContainerProperty{
//   				localDeviceResourceData: &localDeviceResourceDataProperty{
//   					sourcePath: jsii.String("sourcePath"),
//
//   					// the properties below are optional
//   					groupOwnerSetting: &groupOwnerSettingProperty{
//   						autoAddGroupOwner: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						groupOwner: jsii.String("groupOwner"),
//   					},
//   				},
//   				localVolumeResourceData: &localVolumeResourceDataProperty{
//   					destinationPath: jsii.String("destinationPath"),
//   					sourcePath: jsii.String("sourcePath"),
//
//   					// the properties below are optional
//   					groupOwnerSetting: &groupOwnerSettingProperty{
//   						autoAddGroupOwner: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						groupOwner: jsii.String("groupOwner"),
//   					},
//   				},
//   				s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   					destinationPath: jsii.String("destinationPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					ownerSetting: &resourceDownloadOwnerSettingProperty{
//   						groupOwner: jsii.String("groupOwner"),
//   						groupPermission: jsii.String("groupPermission"),
//   					},
//   				},
//   				sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   					destinationPath: jsii.String("destinationPath"),
//   					sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   					// the properties below are optional
//   					ownerSetting: &resourceDownloadOwnerSettingProperty{
//   						groupOwner: jsii.String("groupOwner"),
//   						groupPermission: jsii.String("groupPermission"),
//   					},
//   				},
//   				secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   					arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					additionalStagingLabelsToDownload: []*string{
//   						jsii.String("additionalStagingLabelsToDownload"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnResourceDefinition_ResourceDefinitionVersionProperty struct {
	// The resources in this version.
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
}

// The owner setting for a downloaded machine learning resource.
//
// For more information, see [Access Machine Learning Resources from Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-ml-resources.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `ResourceDownloadOwnerSetting` is the property type of the `OwnerSetting` property for the [`S3MachineLearningModelResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-s3machinelearningmodelresourcedata.html) and [`SageMakerMachineLearningModelResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-sagemakermachinelearningmodelresourcedata.html) property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceDownloadOwnerSettingProperty := &resourceDownloadOwnerSettingProperty{
//   	groupOwner: jsii.String("groupOwner"),
//   	groupPermission: jsii.String("groupPermission"),
//   }
//
type CfnResourceDefinition_ResourceDownloadOwnerSettingProperty struct {
	// The group owner of the machine learning resource.
	//
	// This is the group ID (GID) of an existing Linux OS group on the system. The group's permissions are added to the Lambda process.
	GroupOwner *string `field:"required" json:"groupOwner" yaml:"groupOwner"`
	// The permissions that the group owner has to the machine learning resource.
	//
	// Valid values are `rw` (read-write) or `ro` (read-only).
	GroupPermission *string `field:"required" json:"groupPermission" yaml:"groupPermission"`
}

// A local resource, machine learning resource, or secret resource.
//
// For more information, see [Access Local Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-local-resources.html) , [Perform Machine Learning Inference](https://docs.aws.amazon.com/greengrass/latest/developerguide/ml-inference.html) , and [Deploy Secrets to the AWS IoT Greengrass Core](https://docs.aws.amazon.com/greengrass/latest/developerguide/secrets.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Resources` property of the [`AWS::Greengrass::ResourceDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html) resource contains a list of `ResourceInstance` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceInstanceProperty := &resourceInstanceProperty{
//   	id: jsii.String("id"),
//   	name: jsii.String("name"),
//   	resourceDataContainer: &resourceDataContainerProperty{
//   		localDeviceResourceData: &localDeviceResourceDataProperty{
//   			sourcePath: jsii.String("sourcePath"),
//
//   			// the properties below are optional
//   			groupOwnerSetting: &groupOwnerSettingProperty{
//   				autoAddGroupOwner: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				groupOwner: jsii.String("groupOwner"),
//   			},
//   		},
//   		localVolumeResourceData: &localVolumeResourceDataProperty{
//   			destinationPath: jsii.String("destinationPath"),
//   			sourcePath: jsii.String("sourcePath"),
//
//   			// the properties below are optional
//   			groupOwnerSetting: &groupOwnerSettingProperty{
//   				autoAddGroupOwner: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				groupOwner: jsii.String("groupOwner"),
//   			},
//   		},
//   		s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   			destinationPath: jsii.String("destinationPath"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ownerSetting: &resourceDownloadOwnerSettingProperty{
//   				groupOwner: jsii.String("groupOwner"),
//   				groupPermission: jsii.String("groupPermission"),
//   			},
//   		},
//   		sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   			destinationPath: jsii.String("destinationPath"),
//   			sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   			// the properties below are optional
//   			ownerSetting: &resourceDownloadOwnerSettingProperty{
//   				groupOwner: jsii.String("groupOwner"),
//   				groupPermission: jsii.String("groupPermission"),
//   			},
//   		},
//   		secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   			arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			additionalStagingLabelsToDownload: []*string{
//   				jsii.String("additionalStagingLabelsToDownload"),
//   			},
//   		},
//   	},
//   }
//
type CfnResourceDefinition_ResourceInstanceProperty struct {
	// A descriptive or arbitrary ID for the resource.
	//
	// This value must be unique within the resource definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The descriptive resource name, which is displayed on the AWS IoT Greengrass console.
	//
	// Maximum length 128 characters with pattern [a-zA-Z0-9:_-]+. This must be unique within a Greengrass group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A container for resource data.
	//
	// The container takes only one of the following supported resource data types: `LocalDeviceResourceData` , `LocalVolumeResourceData` , `SageMakerMachineLearningModelResourceData` , `S3MachineLearningModelResourceData` , or `SecretsManagerSecretResourceData` .
	//
	// > Only one resource type can be defined for a `ResourceDataContainer` instance.
	ResourceDataContainer interface{} `field:"required" json:"resourceDataContainer" yaml:"resourceDataContainer"`
}

// Settings for an Amazon S3 machine learning resource.
//
// For more information, see [Perform Machine Learning Inference](https://docs.aws.amazon.com/greengrass/latest/developerguide/ml-inference.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `S3MachineLearningModelResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3MachineLearningModelResourceDataProperty := &s3MachineLearningModelResourceDataProperty{
//   	destinationPath: jsii.String("destinationPath"),
//   	s3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	ownerSetting: &resourceDownloadOwnerSettingProperty{
//   		groupOwner: jsii.String("groupOwner"),
//   		groupPermission: jsii.String("groupPermission"),
//   	},
//   }
//
type CfnResourceDefinition_S3MachineLearningModelResourceDataProperty struct {
	// The absolute local path of the resource inside the Lambda environment.
	DestinationPath *string `field:"required" json:"destinationPath" yaml:"destinationPath"`
	// The URI of the source model in an Amazon S3 bucket.
	//
	// The model package must be in `tar.gz` or `.zip` format.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The owner setting for the downloaded machine learning resource.
	//
	// For more information, see [Access Machine Learning Resources from Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-ml-resources.html) in the *Developer Guide* .
	OwnerSetting interface{} `field:"optional" json:"ownerSetting" yaml:"ownerSetting"`
}

// Settings for an Secrets Manager machine learning resource.
//
// For more information, see [Perform Machine Learning Inference](https://docs.aws.amazon.com/greengrass/latest/developerguide/ml-inference.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `SageMakerMachineLearningModelResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sageMakerMachineLearningModelResourceDataProperty := &sageMakerMachineLearningModelResourceDataProperty{
//   	destinationPath: jsii.String("destinationPath"),
//   	sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   	// the properties below are optional
//   	ownerSetting: &resourceDownloadOwnerSettingProperty{
//   		groupOwner: jsii.String("groupOwner"),
//   		groupPermission: jsii.String("groupPermission"),
//   	},
//   }
//
type CfnResourceDefinition_SageMakerMachineLearningModelResourceDataProperty struct {
	// The absolute local path of the resource inside the Lambda environment.
	DestinationPath *string `field:"required" json:"destinationPath" yaml:"destinationPath"`
	// The Amazon Resource Name (ARN) of the Amazon SageMaker training job that represents the source model.
	SageMakerJobArn *string `field:"required" json:"sageMakerJobArn" yaml:"sageMakerJobArn"`
	// The owner setting for the downloaded machine learning resource.
	//
	// For more information, see [Access Machine Learning Resources from Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-ml-resources.html) in the *Developer Guide* .
	OwnerSetting interface{} `field:"optional" json:"ownerSetting" yaml:"ownerSetting"`
}

// Settings for a secret resource, which references a secret from AWS Secrets Manager .
//
// AWS IoT Greengrass stores a local, encrypted copy of the secret on the Greengrass core, where it can be securely accessed by connectors and Lambda functions. For more information, see [Deploy Secrets to the AWS IoT Greengrass Core](https://docs.aws.amazon.com/greengrass/latest/developerguide/secrets.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `SecretsManagerSecretResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretsManagerSecretResourceDataProperty := &secretsManagerSecretResourceDataProperty{
//   	arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	additionalStagingLabelsToDownload: []*string{
//   		jsii.String("additionalStagingLabelsToDownload"),
//   	},
//   }
//
type CfnResourceDefinition_SecretsManagerSecretResourceDataProperty struct {
	// The Amazon Resource Name (ARN) of the Secrets Manager secret to make available on the core.
	//
	// The value of the secret's latest version (represented by the `AWSCURRENT` staging label) is included by default.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The staging labels whose values you want to make available on the core, in addition to `AWSCURRENT` .
	AdditionalStagingLabelsToDownload *[]*string `field:"optional" json:"additionalStagingLabelsToDownload" yaml:"additionalStagingLabelsToDownload"`
}

// Properties for defining a `CfnResourceDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnResourceDefinitionProps := &cfnResourceDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &resourceDefinitionVersionProperty{
//   		resources: []interface{}{
//   			&resourceInstanceProperty{
//   				id: jsii.String("id"),
//   				name: jsii.String("name"),
//   				resourceDataContainer: &resourceDataContainerProperty{
//   					localDeviceResourceData: &localDeviceResourceDataProperty{
//   						sourcePath: jsii.String("sourcePath"),
//
//   						// the properties below are optional
//   						groupOwnerSetting: &groupOwnerSettingProperty{
//   							autoAddGroupOwner: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							groupOwner: jsii.String("groupOwner"),
//   						},
//   					},
//   					localVolumeResourceData: &localVolumeResourceDataProperty{
//   						destinationPath: jsii.String("destinationPath"),
//   						sourcePath: jsii.String("sourcePath"),
//
//   						// the properties below are optional
//   						groupOwnerSetting: &groupOwnerSettingProperty{
//   							autoAddGroupOwner: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							groupOwner: jsii.String("groupOwner"),
//   						},
//   					},
//   					s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   						destinationPath: jsii.String("destinationPath"),
//   						s3Uri: jsii.String("s3Uri"),
//
//   						// the properties below are optional
//   						ownerSetting: &resourceDownloadOwnerSettingProperty{
//   							groupOwner: jsii.String("groupOwner"),
//   							groupPermission: jsii.String("groupPermission"),
//   						},
//   					},
//   					sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   						destinationPath: jsii.String("destinationPath"),
//   						sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   						// the properties below are optional
//   						ownerSetting: &resourceDownloadOwnerSettingProperty{
//   							groupOwner: jsii.String("groupOwner"),
//   							groupPermission: jsii.String("groupPermission"),
//   						},
//   					},
//   					secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   						arn: jsii.String("arn"),
//
//   						// the properties below are optional
//   						additionalStagingLabelsToDownload: []*string{
//   							jsii.String("additionalStagingLabelsToDownload"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	tags: tags,
//   }
//
type CfnResourceDefinitionProps struct {
	// The name of the resource definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The resource definition version to include when the resource definition is created.
	//
	// A resource definition version contains a list of [`resource instance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html) property types.
	//
	// > To associate a resource definition version after the resource definition is created, create an [`AWS::Greengrass::ResourceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html) resource and specify the ID of this resource definition.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the resource definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Greengrass::ResourceDefinitionVersion`.
//
// The `AWS::Greengrass::ResourceDefinitionVersion` resource represents a resource definition version for AWS IoT Greengrass . A resource definition version contains a list of resources. (In AWS CloudFormation , resources are named *resource instances* .)
//
// > To create a resource definition version, you must specify the ID of the resource definition that you want to associate with the version. For information about creating a resource definition, see [`AWS::Greengrass::ResourceDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html) .
// >
// > After you create a resource definition version that contains the resources you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceDefinitionVersion := awscdk.Aws_greengrass.NewCfnResourceDefinitionVersion(this, jsii.String("MyCfnResourceDefinitionVersion"), &cfnResourceDefinitionVersionProps{
//   	resourceDefinitionId: jsii.String("resourceDefinitionId"),
//   	resources: []interface{}{
//   		&resourceInstanceProperty{
//   			id: jsii.String("id"),
//   			name: jsii.String("name"),
//   			resourceDataContainer: &resourceDataContainerProperty{
//   				localDeviceResourceData: &localDeviceResourceDataProperty{
//   					sourcePath: jsii.String("sourcePath"),
//
//   					// the properties below are optional
//   					groupOwnerSetting: &groupOwnerSettingProperty{
//   						autoAddGroupOwner: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						groupOwner: jsii.String("groupOwner"),
//   					},
//   				},
//   				localVolumeResourceData: &localVolumeResourceDataProperty{
//   					destinationPath: jsii.String("destinationPath"),
//   					sourcePath: jsii.String("sourcePath"),
//
//   					// the properties below are optional
//   					groupOwnerSetting: &groupOwnerSettingProperty{
//   						autoAddGroupOwner: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						groupOwner: jsii.String("groupOwner"),
//   					},
//   				},
//   				s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   					destinationPath: jsii.String("destinationPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					ownerSetting: &resourceDownloadOwnerSettingProperty{
//   						groupOwner: jsii.String("groupOwner"),
//   						groupPermission: jsii.String("groupPermission"),
//   					},
//   				},
//   				sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   					destinationPath: jsii.String("destinationPath"),
//   					sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   					// the properties below are optional
//   					ownerSetting: &resourceDownloadOwnerSettingProperty{
//   						groupOwner: jsii.String("groupOwner"),
//   						groupPermission: jsii.String("groupPermission"),
//   					},
//   				},
//   				secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   					arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					additionalStagingLabelsToDownload: []*string{
//   						jsii.String("additionalStagingLabelsToDownload"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   })
//
type CfnResourceDefinitionVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the resource definition associated with this version.
	//
	// This value is a GUID.
	ResourceDefinitionId() *string
	SetResourceDefinitionId(val *string)
	// The resources in this version.
	Resources() interface{}
	SetResources(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnResourceDefinitionVersion
type jsiiProxy_CfnResourceDefinitionVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) ResourceDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) Resources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::ResourceDefinitionVersion`.
func NewCfnResourceDefinitionVersion(scope constructs.Construct, id *string, props *CfnResourceDefinitionVersionProps) CfnResourceDefinitionVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnResourceDefinitionVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnResourceDefinitionVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::ResourceDefinitionVersion`.
func NewCfnResourceDefinitionVersion_Override(c CfnResourceDefinitionVersion, scope constructs.Construct, id *string, props *CfnResourceDefinitionVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnResourceDefinitionVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) SetResourceDefinitionId(val *string) {
	_jsii_.Set(
		j,
		"resourceDefinitionId",
		val,
	)
}

func (j *jsiiProxy_CfnResourceDefinitionVersion) SetResources(val interface{}) {
	_jsii_.Set(
		j,
		"resources",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResourceDefinitionVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnResourceDefinitionVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResourceDefinitionVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnResourceDefinitionVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnResourceDefinitionVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnResourceDefinitionVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceDefinitionVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnResourceDefinitionVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceDefinitionVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Settings that define additional Linux OS group permissions to give to the Lambda function process.
//
// You can give the permissions of the Linux group that owns the resource or choose another Linux group. These permissions are in addition to the function's `RunAs` permissions.
//
// In an AWS CloudFormation template, `GroupOwnerSetting` is a property of the [`LocalDeviceResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localdeviceresourcedata.html) and [`LocalVolumeResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localvolumeresourcedata.html) property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupOwnerSettingProperty := &groupOwnerSettingProperty{
//   	autoAddGroupOwner: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	groupOwner: jsii.String("groupOwner"),
//   }
//
type CfnResourceDefinitionVersion_GroupOwnerSettingProperty struct {
	// Indicates whether to give the privileges of the Linux group that owns the resource to the Lambda process.
	//
	// This gives the Lambda process the file access permissions of the Linux group.
	AutoAddGroupOwner interface{} `field:"required" json:"autoAddGroupOwner" yaml:"autoAddGroupOwner"`
	// The name of the Linux group whose privileges you want to add to the Lambda process.
	//
	// This value is ignored if `AutoAddGroupOwner` is true.
	GroupOwner *string `field:"optional" json:"groupOwner" yaml:"groupOwner"`
}

// Settings for a local device resource, which represents a file under `/dev` .
//
// For more information, see [Access Local Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-local-resources.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `LocalDeviceResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localDeviceResourceDataProperty := &localDeviceResourceDataProperty{
//   	sourcePath: jsii.String("sourcePath"),
//
//   	// the properties below are optional
//   	groupOwnerSetting: &groupOwnerSettingProperty{
//   		autoAddGroupOwner: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		groupOwner: jsii.String("groupOwner"),
//   	},
//   }
//
type CfnResourceDefinitionVersion_LocalDeviceResourceDataProperty struct {
	// The local absolute path of the device resource.
	//
	// The source path for a device resource can refer only to a character device or block device under `/dev` .
	SourcePath *string `field:"required" json:"sourcePath" yaml:"sourcePath"`
	// Settings that define additional Linux OS group permissions to give to the Lambda function process.
	GroupOwnerSetting interface{} `field:"optional" json:"groupOwnerSetting" yaml:"groupOwnerSetting"`
}

// Settings for a local volume resource, which represents a file or directory on the root file system.
//
// For more information, see [Access Local Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-local-resources.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `LocalVolumeResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localVolumeResourceDataProperty := &localVolumeResourceDataProperty{
//   	destinationPath: jsii.String("destinationPath"),
//   	sourcePath: jsii.String("sourcePath"),
//
//   	// the properties below are optional
//   	groupOwnerSetting: &groupOwnerSettingProperty{
//   		autoAddGroupOwner: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		groupOwner: jsii.String("groupOwner"),
//   	},
//   }
//
type CfnResourceDefinitionVersion_LocalVolumeResourceDataProperty struct {
	// The absolute local path of the resource in the Lambda environment.
	DestinationPath *string `field:"required" json:"destinationPath" yaml:"destinationPath"`
	// The local absolute path of the volume resource on the host.
	//
	// The source path for a volume resource type cannot start with `/sys` .
	SourcePath *string `field:"required" json:"sourcePath" yaml:"sourcePath"`
	// Settings that define additional Linux OS group permissions to give to the Lambda function process.
	GroupOwnerSetting interface{} `field:"optional" json:"groupOwnerSetting" yaml:"groupOwnerSetting"`
}

// A container for resource data, which defines the resource type.
//
// The container takes only one of the following supported resource data types: `LocalDeviceResourceData` , `LocalVolumeResourceData` , `SageMakerMachineLearningModelResourceData` , `S3MachineLearningModelResourceData` , or `SecretsManagerSecretResourceData` .
//
// > Only one resource type can be defined for a `ResourceDataContainer` instance.
//
// In an AWS CloudFormation template, `ResourceDataContainer` is a property of the [`ResourceInstance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourceinstance.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceDataContainerProperty := &resourceDataContainerProperty{
//   	localDeviceResourceData: &localDeviceResourceDataProperty{
//   		sourcePath: jsii.String("sourcePath"),
//
//   		// the properties below are optional
//   		groupOwnerSetting: &groupOwnerSettingProperty{
//   			autoAddGroupOwner: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			groupOwner: jsii.String("groupOwner"),
//   		},
//   	},
//   	localVolumeResourceData: &localVolumeResourceDataProperty{
//   		destinationPath: jsii.String("destinationPath"),
//   		sourcePath: jsii.String("sourcePath"),
//
//   		// the properties below are optional
//   		groupOwnerSetting: &groupOwnerSettingProperty{
//   			autoAddGroupOwner: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			groupOwner: jsii.String("groupOwner"),
//   		},
//   	},
//   	s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   		destinationPath: jsii.String("destinationPath"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		ownerSetting: &resourceDownloadOwnerSettingProperty{
//   			groupOwner: jsii.String("groupOwner"),
//   			groupPermission: jsii.String("groupPermission"),
//   		},
//   	},
//   	sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   		destinationPath: jsii.String("destinationPath"),
//   		sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   		// the properties below are optional
//   		ownerSetting: &resourceDownloadOwnerSettingProperty{
//   			groupOwner: jsii.String("groupOwner"),
//   			groupPermission: jsii.String("groupPermission"),
//   		},
//   	},
//   	secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   		arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		additionalStagingLabelsToDownload: []*string{
//   			jsii.String("additionalStagingLabelsToDownload"),
//   		},
//   	},
//   }
//
type CfnResourceDefinitionVersion_ResourceDataContainerProperty struct {
	// Settings for a local device resource.
	LocalDeviceResourceData interface{} `field:"optional" json:"localDeviceResourceData" yaml:"localDeviceResourceData"`
	// Settings for a local volume resource.
	LocalVolumeResourceData interface{} `field:"optional" json:"localVolumeResourceData" yaml:"localVolumeResourceData"`
	// Settings for a machine learning resource stored in Amazon S3 .
	S3MachineLearningModelResourceData interface{} `field:"optional" json:"s3MachineLearningModelResourceData" yaml:"s3MachineLearningModelResourceData"`
	// Settings for a machine learning resource saved as an SageMaker training job.
	SageMakerMachineLearningModelResourceData interface{} `field:"optional" json:"sageMakerMachineLearningModelResourceData" yaml:"sageMakerMachineLearningModelResourceData"`
	// Settings for a secret resource.
	SecretsManagerSecretResourceData interface{} `field:"optional" json:"secretsManagerSecretResourceData" yaml:"secretsManagerSecretResourceData"`
}

// The owner setting for a downloaded machine learning resource.
//
// For more information, see [Access Machine Learning Resources from Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-ml-resources.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `ResourceDownloadOwnerSetting` is the property type of the `OwnerSetting` property for the [`S3MachineLearningModelResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-s3machinelearningmodelresourcedata.html) and [`SageMakerMachineLearningModelResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-sagemakermachinelearningmodelresourcedata.html) property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceDownloadOwnerSettingProperty := &resourceDownloadOwnerSettingProperty{
//   	groupOwner: jsii.String("groupOwner"),
//   	groupPermission: jsii.String("groupPermission"),
//   }
//
type CfnResourceDefinitionVersion_ResourceDownloadOwnerSettingProperty struct {
	// The group owner of the machine learning resource.
	//
	// This is the group ID (GID) of an existing Linux OS group on the system. The group's permissions are added to the Lambda process.
	GroupOwner *string `field:"required" json:"groupOwner" yaml:"groupOwner"`
	// The permissions that the group owner has to the machine learning resource.
	//
	// Valid values are `rw` (read-write) or `ro` (read-only).
	GroupPermission *string `field:"required" json:"groupPermission" yaml:"groupPermission"`
}

// A local resource, machine learning resource, or secret resource.
//
// For more information, see [Access Local Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-local-resources.html) , [Perform Machine Learning Inference](https://docs.aws.amazon.com/greengrass/latest/developerguide/ml-inference.html) , and [Deploy Secrets to the AWS IoT Greengrass Core](https://docs.aws.amazon.com/greengrass/latest/developerguide/secrets.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Resources` property of the [`AWS::Greengrass::ResourceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html) resource contains a list of `ResourceInstance` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceInstanceProperty := &resourceInstanceProperty{
//   	id: jsii.String("id"),
//   	name: jsii.String("name"),
//   	resourceDataContainer: &resourceDataContainerProperty{
//   		localDeviceResourceData: &localDeviceResourceDataProperty{
//   			sourcePath: jsii.String("sourcePath"),
//
//   			// the properties below are optional
//   			groupOwnerSetting: &groupOwnerSettingProperty{
//   				autoAddGroupOwner: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				groupOwner: jsii.String("groupOwner"),
//   			},
//   		},
//   		localVolumeResourceData: &localVolumeResourceDataProperty{
//   			destinationPath: jsii.String("destinationPath"),
//   			sourcePath: jsii.String("sourcePath"),
//
//   			// the properties below are optional
//   			groupOwnerSetting: &groupOwnerSettingProperty{
//   				autoAddGroupOwner: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				groupOwner: jsii.String("groupOwner"),
//   			},
//   		},
//   		s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   			destinationPath: jsii.String("destinationPath"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ownerSetting: &resourceDownloadOwnerSettingProperty{
//   				groupOwner: jsii.String("groupOwner"),
//   				groupPermission: jsii.String("groupPermission"),
//   			},
//   		},
//   		sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   			destinationPath: jsii.String("destinationPath"),
//   			sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   			// the properties below are optional
//   			ownerSetting: &resourceDownloadOwnerSettingProperty{
//   				groupOwner: jsii.String("groupOwner"),
//   				groupPermission: jsii.String("groupPermission"),
//   			},
//   		},
//   		secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   			arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			additionalStagingLabelsToDownload: []*string{
//   				jsii.String("additionalStagingLabelsToDownload"),
//   			},
//   		},
//   	},
//   }
//
type CfnResourceDefinitionVersion_ResourceInstanceProperty struct {
	// A descriptive or arbitrary ID for the resource.
	//
	// This value must be unique within the resource definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The descriptive resource name, which is displayed on the AWS IoT Greengrass console.
	//
	// Maximum length 128 characters with pattern [a-zA-Z0-9:_-]+. This must be unique within a Greengrass group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A container for resource data.
	//
	// The container takes only one of the following supported resource data types: `LocalDeviceResourceData` , `LocalVolumeResourceData` , `SageMakerMachineLearningModelResourceData` , `S3MachineLearningModelResourceData` , or `SecretsManagerSecretResourceData` .
	//
	// > Only one resource type can be defined for a `ResourceDataContainer` instance.
	ResourceDataContainer interface{} `field:"required" json:"resourceDataContainer" yaml:"resourceDataContainer"`
}

// Settings for an Amazon S3 machine learning resource.
//
// For more information, see [Perform Machine Learning Inference](https://docs.aws.amazon.com/greengrass/latest/developerguide/ml-inference.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `S3MachineLearningModelResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3MachineLearningModelResourceDataProperty := &s3MachineLearningModelResourceDataProperty{
//   	destinationPath: jsii.String("destinationPath"),
//   	s3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	ownerSetting: &resourceDownloadOwnerSettingProperty{
//   		groupOwner: jsii.String("groupOwner"),
//   		groupPermission: jsii.String("groupPermission"),
//   	},
//   }
//
type CfnResourceDefinitionVersion_S3MachineLearningModelResourceDataProperty struct {
	// The absolute local path of the resource inside the Lambda environment.
	DestinationPath *string `field:"required" json:"destinationPath" yaml:"destinationPath"`
	// The URI of the source model in an Amazon S3 bucket.
	//
	// The model package must be in `tar.gz` or `.zip` format.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The owner setting for the downloaded machine learning resource.
	//
	// For more information, see [Access Machine Learning Resources from Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-ml-resources.html) in the *Developer Guide* .
	OwnerSetting interface{} `field:"optional" json:"ownerSetting" yaml:"ownerSetting"`
}

// Settings for an Secrets Manager machine learning resource.
//
// For more information, see [Perform Machine Learning Inference](https://docs.aws.amazon.com/greengrass/latest/developerguide/ml-inference.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `SageMakerMachineLearningModelResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sageMakerMachineLearningModelResourceDataProperty := &sageMakerMachineLearningModelResourceDataProperty{
//   	destinationPath: jsii.String("destinationPath"),
//   	sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   	// the properties below are optional
//   	ownerSetting: &resourceDownloadOwnerSettingProperty{
//   		groupOwner: jsii.String("groupOwner"),
//   		groupPermission: jsii.String("groupPermission"),
//   	},
//   }
//
type CfnResourceDefinitionVersion_SageMakerMachineLearningModelResourceDataProperty struct {
	// The absolute local path of the resource inside the Lambda environment.
	DestinationPath *string `field:"required" json:"destinationPath" yaml:"destinationPath"`
	// The Amazon Resource Name (ARN) of the Amazon SageMaker training job that represents the source model.
	SageMakerJobArn *string `field:"required" json:"sageMakerJobArn" yaml:"sageMakerJobArn"`
	// The owner setting for the downloaded machine learning resource.
	//
	// For more information, see [Access Machine Learning Resources from Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-ml-resources.html) in the *Developer Guide* .
	OwnerSetting interface{} `field:"optional" json:"ownerSetting" yaml:"ownerSetting"`
}

// Settings for a secret resource, which references a secret from AWS Secrets Manager .
//
// AWS IoT Greengrass stores a local, encrypted copy of the secret on the Greengrass core, where it can be securely accessed by connectors and Lambda functions. For more information, see [Deploy Secrets to the AWS IoT Greengrass Core](https://docs.aws.amazon.com/greengrass/latest/developerguide/secrets.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `SecretsManagerSecretResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretsManagerSecretResourceDataProperty := &secretsManagerSecretResourceDataProperty{
//   	arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	additionalStagingLabelsToDownload: []*string{
//   		jsii.String("additionalStagingLabelsToDownload"),
//   	},
//   }
//
type CfnResourceDefinitionVersion_SecretsManagerSecretResourceDataProperty struct {
	// The Amazon Resource Name (ARN) of the Secrets Manager secret to make available on the core.
	//
	// The value of the secret's latest version (represented by the `AWSCURRENT` staging label) is included by default.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The staging labels whose values you want to make available on the core, in addition to `AWSCURRENT` .
	AdditionalStagingLabelsToDownload *[]*string `field:"optional" json:"additionalStagingLabelsToDownload" yaml:"additionalStagingLabelsToDownload"`
}

// Properties for defining a `CfnResourceDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceDefinitionVersionProps := &cfnResourceDefinitionVersionProps{
//   	resourceDefinitionId: jsii.String("resourceDefinitionId"),
//   	resources: []interface{}{
//   		&resourceInstanceProperty{
//   			id: jsii.String("id"),
//   			name: jsii.String("name"),
//   			resourceDataContainer: &resourceDataContainerProperty{
//   				localDeviceResourceData: &localDeviceResourceDataProperty{
//   					sourcePath: jsii.String("sourcePath"),
//
//   					// the properties below are optional
//   					groupOwnerSetting: &groupOwnerSettingProperty{
//   						autoAddGroupOwner: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						groupOwner: jsii.String("groupOwner"),
//   					},
//   				},
//   				localVolumeResourceData: &localVolumeResourceDataProperty{
//   					destinationPath: jsii.String("destinationPath"),
//   					sourcePath: jsii.String("sourcePath"),
//
//   					// the properties below are optional
//   					groupOwnerSetting: &groupOwnerSettingProperty{
//   						autoAddGroupOwner: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						groupOwner: jsii.String("groupOwner"),
//   					},
//   				},
//   				s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   					destinationPath: jsii.String("destinationPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					ownerSetting: &resourceDownloadOwnerSettingProperty{
//   						groupOwner: jsii.String("groupOwner"),
//   						groupPermission: jsii.String("groupPermission"),
//   					},
//   				},
//   				sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   					destinationPath: jsii.String("destinationPath"),
//   					sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   					// the properties below are optional
//   					ownerSetting: &resourceDownloadOwnerSettingProperty{
//   						groupOwner: jsii.String("groupOwner"),
//   						groupPermission: jsii.String("groupPermission"),
//   					},
//   				},
//   				secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   					arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					additionalStagingLabelsToDownload: []*string{
//   						jsii.String("additionalStagingLabelsToDownload"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnResourceDefinitionVersionProps struct {
	// The ID of the resource definition associated with this version.
	//
	// This value is a GUID.
	ResourceDefinitionId *string `field:"required" json:"resourceDefinitionId" yaml:"resourceDefinitionId"`
	// The resources in this version.
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
}

// A CloudFormation `AWS::Greengrass::SubscriptionDefinition`.
//
// The `AWS::Greengrass::SubscriptionDefinition` resource represents a subscription definition for AWS IoT Greengrass . Subscription definitions are used to organize your subscription definition versions.
//
// Subscription definitions can reference multiple subscription definition versions. All subscription definition versions must be associated with a subscription definition. Each subscription definition version can contain one or more subscriptions.
//
// > When you create a subscription definition, you can optionally include an initial subscription definition version. To associate a subscription definition version later, create an [`AWS::Greengrass::SubscriptionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html) resource and specify the ID of this subscription definition.
// >
// > After you create the subscription definition version that contains the subscriptions you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnSubscriptionDefinition := awscdk.Aws_greengrass.NewCfnSubscriptionDefinition(this, jsii.String("MyCfnSubscriptionDefinition"), &cfnSubscriptionDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &subscriptionDefinitionVersionProperty{
//   		subscriptions: []interface{}{
//   			&subscriptionProperty{
//   				id: jsii.String("id"),
//   				source: jsii.String("source"),
//   				subject: jsii.String("subject"),
//   				target: jsii.String("target"),
//   			},
//   		},
//   	},
//   	tags: tags,
//   })
//
type CfnSubscriptionDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the `SubscriptionDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/subscriptions/1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrArn() *string
	// The ID of the `SubscriptionDefinition` , such as `1234a5b6-78cd-901e-2fgh-3i45j6k178l9` .
	AttrId() *string
	// The ARN of the last `SubscriptionDefinitionVersion` that was added to the `SubscriptionDefinition` , such as `arn:aws:greengrass:us-east-1:  :/greengrass/definition/subscriptions/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2` .
	AttrLatestVersionArn() *string
	// The name of the `SubscriptionDefinition` , such as `MySubscriptionDefinition` .
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The subscription definition version to include when the subscription definition is created.
	//
	// A subscription definition version contains a list of [`subscription`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscription.html) property types.
	//
	// > To associate a subscription definition version after the subscription definition is created, create an [`AWS::Greengrass::SubscriptionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html) resource and specify the ID of this subscription definition.
	InitialVersion() interface{}
	SetInitialVersion(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the subscription definition.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Application-specific metadata to attach to the subscription definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnSubscriptionDefinition
type jsiiProxy_CfnSubscriptionDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSubscriptionDefinition) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) AttrLatestVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLatestVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) InitialVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::SubscriptionDefinition`.
func NewCfnSubscriptionDefinition(scope constructs.Construct, id *string, props *CfnSubscriptionDefinitionProps) CfnSubscriptionDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnSubscriptionDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnSubscriptionDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::SubscriptionDefinition`.
func NewCfnSubscriptionDefinition_Override(c CfnSubscriptionDefinition, scope constructs.Construct, id *string, props *CfnSubscriptionDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnSubscriptionDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSubscriptionDefinition) SetInitialVersion(val interface{}) {
	_jsii_.Set(
		j,
		"initialVersion",
		val,
	)
}

func (j *jsiiProxy_CfnSubscriptionDefinition) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSubscriptionDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnSubscriptionDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSubscriptionDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnSubscriptionDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnSubscriptionDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnSubscriptionDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSubscriptionDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnSubscriptionDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A subscription definition version contains a list of [subscriptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscription.html) .
//
// > After you create a subscription definition version that contains the subscriptions you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an AWS CloudFormation template, `SubscriptionDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::SubscriptionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriptionDefinitionVersionProperty := &subscriptionDefinitionVersionProperty{
//   	subscriptions: []interface{}{
//   		&subscriptionProperty{
//   			id: jsii.String("id"),
//   			source: jsii.String("source"),
//   			subject: jsii.String("subject"),
//   			target: jsii.String("target"),
//   		},
//   	},
//   }
//
type CfnSubscriptionDefinition_SubscriptionDefinitionVersionProperty struct {
	// The subscriptions in this version.
	Subscriptions interface{} `field:"required" json:"subscriptions" yaml:"subscriptions"`
}

// Subscriptions define how MQTT messages can be exchanged between devices, functions, and connectors in the group, and with AWS IoT or the local shadow service.
//
// A subscription defines a message source, message target, and a topic (or subject) that's used to route messages from the source to the target. A subscription defines the message flow in one direction, from the source to the target. For two-way communication, you must set up two subscriptions, one for each direction.
//
// In an AWS CloudFormation template, the `Subscriptions` property of the [`SubscriptionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscriptiondefinitionversion.html) property type contains a list of `Subscription` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriptionProperty := &subscriptionProperty{
//   	id: jsii.String("id"),
//   	source: jsii.String("source"),
//   	subject: jsii.String("subject"),
//   	target: jsii.String("target"),
//   }
//
type CfnSubscriptionDefinition_SubscriptionProperty struct {
	// A descriptive or arbitrary ID for the subscription.
	//
	// This value must be unique within the subscription definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The originator of the message.
	//
	// The value can be a thing ARN, the ARN of a Lambda function alias (recommended) or version, a connector ARN, `cloud` (which represents the AWS IoT cloud), or `GGShadowService` .
	Source *string `field:"required" json:"source" yaml:"source"`
	// The MQTT topic used to route the message.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// The destination of the message.
	//
	// The value can be a thing ARN, the ARN of a Lambda function alias (recommended) or version, a connector ARN, `cloud` (which represents the AWS IoT cloud), or `GGShadowService` .
	Target *string `field:"required" json:"target" yaml:"target"`
}

// Properties for defining a `CfnSubscriptionDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnSubscriptionDefinitionProps := &cfnSubscriptionDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &subscriptionDefinitionVersionProperty{
//   		subscriptions: []interface{}{
//   			&subscriptionProperty{
//   				id: jsii.String("id"),
//   				source: jsii.String("source"),
//   				subject: jsii.String("subject"),
//   				target: jsii.String("target"),
//   			},
//   		},
//   	},
//   	tags: tags,
//   }
//
type CfnSubscriptionDefinitionProps struct {
	// The name of the subscription definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The subscription definition version to include when the subscription definition is created.
	//
	// A subscription definition version contains a list of [`subscription`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscription.html) property types.
	//
	// > To associate a subscription definition version after the subscription definition is created, create an [`AWS::Greengrass::SubscriptionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html) resource and specify the ID of this subscription definition.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the subscription definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Greengrass::SubscriptionDefinitionVersion`.
//
// The `AWS::Greengrass::SubscriptionDefinitionVersion` resource represents a subscription definition version for AWS IoT Greengrass . A subscription definition version contains a list of subscriptions.
//
// > To create a subscription definition version, you must specify the ID of the subscription definition that you want to associate with the version. For information about creating a subscription definition, see [`AWS::Greengrass::SubscriptionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinition.html) .
// >
// > After you create a subscription definition version that contains the subscriptions you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubscriptionDefinitionVersion := awscdk.Aws_greengrass.NewCfnSubscriptionDefinitionVersion(this, jsii.String("MyCfnSubscriptionDefinitionVersion"), &cfnSubscriptionDefinitionVersionProps{
//   	subscriptionDefinitionId: jsii.String("subscriptionDefinitionId"),
//   	subscriptions: []interface{}{
//   		&subscriptionProperty{
//   			id: jsii.String("id"),
//   			source: jsii.String("source"),
//   			subject: jsii.String("subject"),
//   			target: jsii.String("target"),
//   		},
//   	},
//   })
//
type CfnSubscriptionDefinitionVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The ID of the subscription definition associated with this version.
	//
	// This value is a GUID.
	SubscriptionDefinitionId() *string
	SetSubscriptionDefinitionId(val *string)
	// The subscriptions in this version.
	Subscriptions() interface{}
	SetSubscriptions(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnSubscriptionDefinitionVersion
type jsiiProxy_CfnSubscriptionDefinitionVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) SubscriptionDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) Subscriptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Greengrass::SubscriptionDefinitionVersion`.
func NewCfnSubscriptionDefinitionVersion(scope constructs.Construct, id *string, props *CfnSubscriptionDefinitionVersionProps) CfnSubscriptionDefinitionVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnSubscriptionDefinitionVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnSubscriptionDefinitionVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Greengrass::SubscriptionDefinitionVersion`.
func NewCfnSubscriptionDefinitionVersion_Override(c CfnSubscriptionDefinitionVersion, scope constructs.Construct, id *string, props *CfnSubscriptionDefinitionVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_greengrass.CfnSubscriptionDefinitionVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) SetSubscriptionDefinitionId(val *string) {
	_jsii_.Set(
		j,
		"subscriptionDefinitionId",
		val,
	)
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersion) SetSubscriptions(val interface{}) {
	_jsii_.Set(
		j,
		"subscriptions",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSubscriptionDefinitionVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnSubscriptionDefinitionVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSubscriptionDefinitionVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnSubscriptionDefinitionVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnSubscriptionDefinitionVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_greengrass.CfnSubscriptionDefinitionVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSubscriptionDefinitionVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_greengrass.CfnSubscriptionDefinitionVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Subscriptions define how MQTT messages can be exchanged between devices, functions, and connectors in the group, and with AWS IoT or the local shadow service.
//
// A subscription defines a message source, message target, and a topic (or subject) that's used to route messages from the source to the target. A subscription defines the message flow in one direction, from the source to the target. For two-way communication, you must set up two subscriptions, one for each direction.
//
// In an AWS CloudFormation template, the `Subscriptions` property of the [`AWS::Greengrass::SubscriptionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html) resource contains a list of `Subscription` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriptionProperty := &subscriptionProperty{
//   	id: jsii.String("id"),
//   	source: jsii.String("source"),
//   	subject: jsii.String("subject"),
//   	target: jsii.String("target"),
//   }
//
type CfnSubscriptionDefinitionVersion_SubscriptionProperty struct {
	// A descriptive or arbitrary ID for the subscription.
	//
	// This value must be unique within the subscription definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The originator of the message.
	//
	// The value can be a thing ARN, the ARN of a Lambda function alias (recommended) or version, a connector ARN, `cloud` (which represents the AWS IoT cloud), or `GGShadowService` .
	Source *string `field:"required" json:"source" yaml:"source"`
	// The MQTT topic used to route the message.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// The destination of the message.
	//
	// The value can be a thing ARN, the ARN of a Lambda function alias (recommended) or version, a connector ARN, `cloud` (which represents the AWS IoT cloud), or `GGShadowService` .
	Target *string `field:"required" json:"target" yaml:"target"`
}

// Properties for defining a `CfnSubscriptionDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubscriptionDefinitionVersionProps := &cfnSubscriptionDefinitionVersionProps{
//   	subscriptionDefinitionId: jsii.String("subscriptionDefinitionId"),
//   	subscriptions: []interface{}{
//   		&subscriptionProperty{
//   			id: jsii.String("id"),
//   			source: jsii.String("source"),
//   			subject: jsii.String("subject"),
//   			target: jsii.String("target"),
//   		},
//   	},
//   }
//
type CfnSubscriptionDefinitionVersionProps struct {
	// The ID of the subscription definition associated with this version.
	//
	// This value is a GUID.
	SubscriptionDefinitionId *string `field:"required" json:"subscriptionDefinitionId" yaml:"subscriptionDefinitionId"`
	// The subscriptions in this version.
	Subscriptions interface{} `field:"required" json:"subscriptions" yaml:"subscriptions"`
}

