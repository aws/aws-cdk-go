package awsresiliencehub

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsresiliencehub/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::ResilienceHub::App`.
//
// Creates an AWS Resilience Hub application. An AWS Resilience Hub application is a collection of AWS resources structured to prevent and recover AWS application disruptions. To describe a AWS Resilience Hub application, you provide an application name, resources from one or more–up to 20– AWS CloudFormation stacks, and an appropriate resiliency policy.
//
// After you create an AWS Resilience Hub application, you publish it so that you can run a resiliency assessment on it. You can then use recommendations from the assessment to improve resiliency by running another assessment, comparing results, and then iterating the process until you achieve your goals for recovery time objective (RTO) and recovery point objective (RPO).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApp := awscdk.Aws_resiliencehub.NewCfnApp(this, jsii.String("MyCfnApp"), &CfnAppProps{
//   	AppTemplateBody: jsii.String("appTemplateBody"),
//   	Name: jsii.String("name"),
//   	ResourceMappings: []interface{}{
//   		&ResourceMappingProperty{
//   			MappingType: jsii.String("mappingType"),
//   			PhysicalResourceId: &PhysicalResourceIdProperty{
//   				Identifier: jsii.String("identifier"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				AwsAccountId: jsii.String("awsAccountId"),
//   				AwsRegion: jsii.String("awsRegion"),
//   			},
//
//   			// the properties below are optional
//   			LogicalStackName: jsii.String("logicalStackName"),
//   			ResourceName: jsii.String("resourceName"),
//   			TerraformSourceName: jsii.String("terraformSourceName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AppAssessmentSchedule: jsii.String("appAssessmentSchedule"),
//   	Description: jsii.String("description"),
//   	ResiliencyPolicyArn: jsii.String("resiliencyPolicyArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnApp interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Assessment execution schedule with 'Daily' or 'Disabled' values.
	AppAssessmentSchedule() *string
	SetAppAssessmentSchedule(val *string)
	// A JSON string that provides information about your application structure.
	//
	// To learn more about the `appTemplateBody` template, see the sample template provided in the *Examples* section.
	//
	// The `appTemplateBody` JSON string has the following structure:
	//
	// - *`resources`*
	//
	// The list of logical resources that needs to be included in the AWS Resilience Hub application.
	//
	// Type: Array
	//
	// > Don't add the resources that you want to exclude.
	//
	// Each `resources` array item includes the following fields:
	//
	// - *`logicalResourceId`*
	//
	// The logical identifier of the resource.
	//
	// Type: Object
	//
	// Each `logicalResourceId` object includes the following fields:
	//
	// - `identifier`
	//
	// The identifier of the resource.
	//
	// Type: String
	// - `logicalStackName`
	//
	// The name of the AWS CloudFormation stack this resource belongs to.
	//
	// Type: String
	// - `resourceGroupName`
	//
	// The name of the resource group this resource belongs to.
	//
	// Type: String
	// - `terraformSourceName`
	//
	// The name of the Terraform S3 state file this resource belongs to.
	//
	// Type: String
	// - `eksSourceName`
	//
	// The name of the Amazon Elastic Kubernetes Service cluster and namespace this resource belongs to.
	//
	// > This parameter accepts values in "eks-cluster/namespace" format.
	//
	// Type: String
	// - *`type`*
	//
	// The type of resource.
	//
	// Type: string
	// - *`name`*
	//
	// The name of the resource.
	//
	// Type: String
	// - `additionalInfo`
	//
	// Additional configuration parameters for an AWS Resilience Hub application. If you want to implement `additionalInfo` through the AWS Resilience Hub console rather than using an API call, see [Configure the application configuration parameters](https://docs.aws.amazon.com//resilience-hub/latest/userguide/app-config-param.html) .
	//
	// > Currently, this parameter accepts a key-value mapping (in a string format) of only one failover region and one associated account.
	// >
	// > Key: `"failover-regions"`
	// >
	// > Value: `"[{"region":"<REGION>", "accounts":[{"id":"<ACCOUNT_ID>"}]}]"`
	// - *`appComponents`*
	//
	// The list of Application Components (AppComponent) that this resource belongs to. If an AppComponent is not part of the AWS Resilience Hub application, it will be added.
	//
	// Type: Array
	//
	// Each `appComponents` array item includes the following fields:
	//
	// - `name`
	//
	// The name of the AppComponent.
	//
	// Type: String
	// - `type`
	//
	// The type of AppComponent. For more information about the types of AppComponent, see [Grouping resources in an AppComponent](https://docs.aws.amazon.com/resilience-hub/latest/userguide/AppComponent.grouping.html) .
	//
	// Type: String
	// - `resourceNames`
	//
	// The list of included resources that are assigned to the AppComponent.
	//
	// Type: Array of strings
	// - `additionalInfo`
	//
	// Additional configuration parameters for an AWS Resilience Hub application. If you want to implement `additionalInfo` through the AWS Resilience Hub console rather than using an API call, see [Configure the application configuration parameters](https://docs.aws.amazon.com//resilience-hub/latest/userguide/app-config-param.html) .
	//
	// > Currently, this parameter accepts a key-value mapping (in a string format) of only one failover region and one associated account.
	// >
	// > Key: `"failover-regions"`
	// >
	// > Value: `"[{"region":"<REGION>", "accounts":[{"id":"<ACCOUNT_ID>"}]}]"`
	// - *`excludedResources`*
	//
	// The list of logical resource identifiers to be excluded from the application.
	//
	// Type: Array
	//
	// > Don't add the resources that you want to include.
	//
	// Each `excludedResources` array item includes the following fields:
	//
	// - *`logicalResourceIds`*
	//
	// The logical identifier of the resource.
	//
	// Type: Object
	//
	// > You can configure only one of the following fields:
	// >
	// > - `logicalStackName`
	// > - `resourceGroupName`
	// > - `terraformSourceName`
	// > - `eksSourceName`
	//
	// Each `logicalResourceIds` object includes the following fields:
	//
	// - `identifier`
	//
	// The identifier of the resource.
	//
	// Type: String
	// - `logicalStackName`
	//
	// The name of the AWS CloudFormation stack this resource belongs to.
	//
	// Type: String
	// - `resourceGroupName`
	//
	// The name of the resource group this resource belongs to.
	//
	// Type: String
	// - `terraformSourceName`
	//
	// The name of the Terraform S3 state file this resource belongs to.
	//
	// Type: String
	// - `eksSourceName`
	//
	// The name of the Amazon Elastic Kubernetes Service cluster and namespace this resource belongs to.
	//
	// > This parameter accepts values in "eks-cluster/namespace" format.
	//
	// Type: String
	// - *`version`*
	//
	// The AWS Resilience Hub application version.
	// - `additionalInfo`
	//
	// Additional configuration parameters for an AWS Resilience Hub application. If you want to implement `additionalInfo` through the AWS Resilience Hub console rather than using an API call, see [Configure the application configuration parameters](https://docs.aws.amazon.com//resilience-hub/latest/userguide/app-config-param.html) .
	//
	// > Currently, this parameter accepts a key-value mapping (in a string format) of only one failover region and one associated account.
	// >
	// > Key: `"failover-regions"`
	// >
	// > Value: `"[{"region":"<REGION>", "accounts":[{"id":"<ACCOUNT_ID>"}]}]"`.
	AppTemplateBody() *string
	SetAppTemplateBody(val *string)
	// The Amazon Resource Name (ARN) of the app.
	AttrAppArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The optional description for an app.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name for the application.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of the resiliency policy.
	ResiliencyPolicyArn() *string
	SetResiliencyPolicyArn(val *string)
	// An array of ResourceMapping objects.
	ResourceMappings() interface{}
	SetResourceMappings(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The tags assigned to the resource.
	//
	// A tag is a label that you assign to an AWS resource. Each tag consists of a key/value pair.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApp
type jsiiProxy_CfnApp struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApp) AppAssessmentSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appAssessmentSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AppTemplateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appTemplateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AttrAppArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) ResiliencyPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resiliencyPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) ResourceMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ResilienceHub::App`.
func NewCfnApp(scope awscdk.Construct, id *string, props *CfnAppProps) CfnApp {
	_init_.Initialize()

	if err := validateNewCfnAppParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApp{}

	_jsii_.Create(
		"monocdk.aws_resiliencehub.CfnApp",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ResilienceHub::App`.
func NewCfnApp_Override(c CfnApp, scope awscdk.Construct, id *string, props *CfnAppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_resiliencehub.CfnApp",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApp)SetAppAssessmentSchedule(val *string) {
	_jsii_.Set(
		j,
		"appAssessmentSchedule",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetAppTemplateBody(val *string) {
	if err := j.validateSetAppTemplateBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appTemplateBody",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetResiliencyPolicyArn(val *string) {
	_jsii_.Set(
		j,
		"resiliencyPolicyArn",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetResourceMappings(val interface{}) {
	if err := j.validateSetResourceMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceMappings",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnApp_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApp_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_resiliencehub.CfnApp",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApp_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnApp_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_resiliencehub.CfnApp",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_resiliencehub.CfnApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApp_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_resiliencehub.CfnApp",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApp) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApp) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApp) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApp) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApp) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApp) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApp) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApp) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnApp) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnApp) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApp) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApp) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnApp) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApp) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApp) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnApp) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

