package awsauditmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsauditmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AuditManager::Assessment`.
//
// The `AWS::AuditManager::Assessment` resource is an AWS Audit Manager resource type that defines the scope of audit evidence collected by Audit Manager . An Audit Manager assessment is an implementation of an Audit Manager framework.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import auditmanager "github.com/aws/aws-cdk-go/awscdk/aws_auditmanager"
//   cfnAssessment := auditmanager.NewCfnAssessment(this, jsii.String("MyCfnAssessment"), &cfnAssessmentProps{
//   	assessmentReportsDestination: &assessmentReportsDestinationProperty{
//   		destination: jsii.String("destination"),
//   		destinationType: jsii.String("destinationType"),
//   	},
//   	awsAccount: &aWSAccountProperty{
//   		emailAddress: jsii.String("emailAddress"),
//   		id: jsii.String("id"),
//   		name: jsii.String("name"),
//   	},
//   	description: jsii.String("description"),
//   	frameworkId: jsii.String("frameworkId"),
//   	name: jsii.String("name"),
//   	roles: []interface{}{
//   		&roleProperty{
//   			roleArn: jsii.String("roleArn"),
//   			roleType: jsii.String("roleType"),
//   		},
//   	},
//   	scope: &scopeProperty{
//   		awsAccounts: []interface{}{
//   			&aWSAccountProperty{
//   				emailAddress: jsii.String("emailAddress"),
//   				id: jsii.String("id"),
//   				name: jsii.String("name"),
//   			},
//   		},
//   		awsServices: []interface{}{
//   			&aWSServiceProperty{
//   				serviceName: jsii.String("serviceName"),
//   			},
//   		},
//   	},
//   	status: jsii.String("status"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnAssessment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The destination that evidence reports are stored in for the assessment.
	AssessmentReportsDestination() interface{}
	SetAssessmentReportsDestination(val interface{})
	// The Amazon Resource Name (ARN) of the assessment.
	//
	// For example, `arn:aws:auditmanager:us-east-1:123456789012:assessment/111A1A1A-22B2-33C3-DDD4-55E5E5E555E5` .
	AttrArn() *string
	// The unique identifier for the assessment.
	//
	// For example, `111A1A1A-22B2-33C3-DDD4-55E5E5E555E5` .
	AttrAssessmentId() *string
	// The time when the assessment was created.
	//
	// For example, `1607582033.373` .
	AttrCreationTime() awscdk.IResolvable
	// The delegations associated with the assessment.
	AttrDelegations() awscdk.IResolvable
	// The AWS account that's associated with the assessment.
	AwsAccount() interface{}
	SetAwsAccount(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the assessment.
	Description() *string
	SetDescription(val *string)
	// The unique identifier for the framework.
	FrameworkId() *string
	SetFrameworkId(val *string)
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
	// The name of the assessment.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The roles that are associated with the assessment.
	Roles() interface{}
	SetRoles(val interface{})
	// The wrapper of AWS accounts and services that are in scope for the assessment.
	Scope() interface{}
	SetScope(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The overall status of the assessment.
	Status() *string
	SetStatus(val *string)
	// The tags that are associated with the assessment.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnAssessment
type jsiiProxy_CfnAssessment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAssessment) AssessmentReportsDestination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assessmentReportsDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) AttrAssessmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssessmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) AttrCreationTime() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) AttrDelegations() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrDelegations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) AwsAccount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) FrameworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) Roles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) Scope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AuditManager::Assessment`.
func NewCfnAssessment(scope constructs.Construct, id *string, props *CfnAssessmentProps) CfnAssessment {
	_init_.Initialize()

	j := jsiiProxy_CfnAssessment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_auditmanager.CfnAssessment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AuditManager::Assessment`.
func NewCfnAssessment_Override(c CfnAssessment, scope constructs.Construct, id *string, props *CfnAssessmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_auditmanager.CfnAssessment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAssessment) SetAssessmentReportsDestination(val interface{}) {
	_jsii_.Set(
		j,
		"assessmentReportsDestination",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment) SetAwsAccount(val interface{}) {
	_jsii_.Set(
		j,
		"awsAccount",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment) SetFrameworkId(val *string) {
	_jsii_.Set(
		j,
		"frameworkId",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment) SetRoles(val interface{}) {
	_jsii_.Set(
		j,
		"roles",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment) SetScope(val interface{}) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAssessment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_auditmanager.CfnAssessment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAssessment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_auditmanager.CfnAssessment",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnAssessment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_auditmanager.CfnAssessment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssessment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_auditmanager.CfnAssessment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssessment) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAssessment) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAssessment) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAssessment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAssessment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAssessment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAssessment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAssessment) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssessment) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssessment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAssessment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAssessment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssessment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssessment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssessment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `AWSAccount` property type specifies the wrapper of the AWS account details, such as account ID, email address, and so on.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import auditmanager "github.com/aws/aws-cdk-go/awscdk/aws_auditmanager"
//   aWSAccountProperty := &aWSAccountProperty{
//   	emailAddress: jsii.String("emailAddress"),
//   	id: jsii.String("id"),
//   	name: jsii.String("name"),
//   }
//
type CfnAssessment_AWSAccountProperty struct {
	// The email address that's associated with the AWS account .
	EmailAddress *string `json:"emailAddress" yaml:"emailAddress"`
	// The identifier for the AWS account .
	Id *string `json:"id" yaml:"id"`
	// The name of the AWS account .
	Name *string `json:"name" yaml:"name"`
}

// The `AWSService` property type specifies an AWS service such as Amazon S3 , AWS CloudTrail , and so on.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import auditmanager "github.com/aws/aws-cdk-go/awscdk/aws_auditmanager"
//   aWSServiceProperty := &aWSServiceProperty{
//   	serviceName: jsii.String("serviceName"),
//   }
//
type CfnAssessment_AWSServiceProperty struct {
	// The name of the AWS service .
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
}

// The `AssessmentReportsDestination` property type specifies the location in which AWS Audit Manager saves assessment reports for the given assessment.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import auditmanager "github.com/aws/aws-cdk-go/awscdk/aws_auditmanager"
//   assessmentReportsDestinationProperty := &assessmentReportsDestinationProperty{
//   	destination: jsii.String("destination"),
//   	destinationType: jsii.String("destinationType"),
//   }
//
type CfnAssessment_AssessmentReportsDestinationProperty struct {
	// The destination of the assessment report.
	Destination *string `json:"destination" yaml:"destination"`
	// The destination type, such as Amazon S3.
	DestinationType *string `json:"destinationType" yaml:"destinationType"`
}

// The `Delegation` property type specifies the assignment of a control set to a delegate for review.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import auditmanager "github.com/aws/aws-cdk-go/awscdk/aws_auditmanager"
//   delegationProperty := &delegationProperty{
//   	assessmentId: jsii.String("assessmentId"),
//   	assessmentName: jsii.String("assessmentName"),
//   	comment: jsii.String("comment"),
//   	controlSetId: jsii.String("controlSetId"),
//   	createdBy: jsii.String("createdBy"),
//   	creationTime: jsii.Number(123),
//   	id: jsii.String("id"),
//   	lastUpdated: jsii.Number(123),
//   	roleArn: jsii.String("roleArn"),
//   	roleType: jsii.String("roleType"),
//   	status: jsii.String("status"),
//   }
//
type CfnAssessment_DelegationProperty struct {
	// The identifier for the assessment that's associated with the delegation.
	AssessmentId *string `json:"assessmentId" yaml:"assessmentId"`
	// The name of the assessment that's associated with the delegation.
	AssessmentName *string `json:"assessmentName" yaml:"assessmentName"`
	// The comment that's related to the delegation.
	Comment *string `json:"comment" yaml:"comment"`
	// The identifier for the control set that's associated with the delegation.
	ControlSetId *string `json:"controlSetId" yaml:"controlSetId"`
	// The IAM user or role that created the delegation.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `100`
	//
	// *Pattern* : `^[a-zA-Z0-9-_()\\[\\]\\s]+$`.
	CreatedBy *string `json:"createdBy" yaml:"createdBy"`
	// Specifies when the delegation was created.
	CreationTime *float64 `json:"creationTime" yaml:"creationTime"`
	// The unique identifier for the delegation.
	Id *string `json:"id" yaml:"id"`
	// Specifies when the delegation was last updated.
	LastUpdated *float64 `json:"lastUpdated" yaml:"lastUpdated"`
	// The Amazon Resource Name (ARN) of the IAM role.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The type of customer persona.
	//
	// > In `CreateAssessment` , `roleType` can only be `PROCESS_OWNER` .
	// >
	// > In `UpdateSettings` , `roleType` can only be `PROCESS_OWNER` .
	// >
	// > In `BatchCreateDelegationByAssessment` , `roleType` can only be `RESOURCE_OWNER` .
	RoleType *string `json:"roleType" yaml:"roleType"`
	// The status of the delegation.
	Status *string `json:"status" yaml:"status"`
}

// The `Role` property type specifies the wrapper that contains AWS Audit Manager role information, such as the role type and IAM Amazon Resource Name (ARN).
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import auditmanager "github.com/aws/aws-cdk-go/awscdk/aws_auditmanager"
//   roleProperty := &roleProperty{
//   	roleArn: jsii.String("roleArn"),
//   	roleType: jsii.String("roleType"),
//   }
//
type CfnAssessment_RoleProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The type of customer persona.
	//
	// > In `CreateAssessment` , `roleType` can only be `PROCESS_OWNER` .
	// >
	// > In `UpdateSettings` , `roleType` can only be `PROCESS_OWNER` .
	// >
	// > In `BatchCreateDelegationByAssessment` , `roleType` can only be `RESOURCE_OWNER` .
	RoleType *string `json:"roleType" yaml:"roleType"`
}

// The `Scope` property type specifies the wrapper that contains the AWS accounts and services in scope for the assessment.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import auditmanager "github.com/aws/aws-cdk-go/awscdk/aws_auditmanager"
//   scopeProperty := &scopeProperty{
//   	awsAccounts: []interface{}{
//   		&aWSAccountProperty{
//   			emailAddress: jsii.String("emailAddress"),
//   			id: jsii.String("id"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	awsServices: []interface{}{
//   		&aWSServiceProperty{
//   			serviceName: jsii.String("serviceName"),
//   		},
//   	},
//   }
//
type CfnAssessment_ScopeProperty struct {
	// The AWS accounts that are included in the scope of the assessment.
	AwsAccounts interface{} `json:"awsAccounts" yaml:"awsAccounts"`
	// The AWS services that are included in the scope of the assessment.
	AwsServices interface{} `json:"awsServices" yaml:"awsServices"`
}

// Properties for defining a `CfnAssessment`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import auditmanager "github.com/aws/aws-cdk-go/awscdk/aws_auditmanager"
//   cfnAssessmentProps := &cfnAssessmentProps{
//   	assessmentReportsDestination: &assessmentReportsDestinationProperty{
//   		destination: jsii.String("destination"),
//   		destinationType: jsii.String("destinationType"),
//   	},
//   	awsAccount: &aWSAccountProperty{
//   		emailAddress: jsii.String("emailAddress"),
//   		id: jsii.String("id"),
//   		name: jsii.String("name"),
//   	},
//   	description: jsii.String("description"),
//   	frameworkId: jsii.String("frameworkId"),
//   	name: jsii.String("name"),
//   	roles: []interface{}{
//   		&roleProperty{
//   			roleArn: jsii.String("roleArn"),
//   			roleType: jsii.String("roleType"),
//   		},
//   	},
//   	scope: &scopeProperty{
//   		awsAccounts: []interface{}{
//   			&aWSAccountProperty{
//   				emailAddress: jsii.String("emailAddress"),
//   				id: jsii.String("id"),
//   				name: jsii.String("name"),
//   			},
//   		},
//   		awsServices: []interface{}{
//   			&aWSServiceProperty{
//   				serviceName: jsii.String("serviceName"),
//   			},
//   		},
//   	},
//   	status: jsii.String("status"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAssessmentProps struct {
	// The destination that evidence reports are stored in for the assessment.
	AssessmentReportsDestination interface{} `json:"assessmentReportsDestination" yaml:"assessmentReportsDestination"`
	// The AWS account that's associated with the assessment.
	AwsAccount interface{} `json:"awsAccount" yaml:"awsAccount"`
	// The description of the assessment.
	Description *string `json:"description" yaml:"description"`
	// The unique identifier for the framework.
	FrameworkId *string `json:"frameworkId" yaml:"frameworkId"`
	// The name of the assessment.
	Name *string `json:"name" yaml:"name"`
	// The roles that are associated with the assessment.
	Roles interface{} `json:"roles" yaml:"roles"`
	// The wrapper of AWS accounts and services that are in scope for the assessment.
	Scope interface{} `json:"scope" yaml:"scope"`
	// The overall status of the assessment.
	Status *string `json:"status" yaml:"status"`
	// The tags that are associated with the assessment.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

