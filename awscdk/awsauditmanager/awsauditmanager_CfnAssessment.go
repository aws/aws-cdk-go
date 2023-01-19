package awsauditmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsauditmanager/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::AuditManager::Assessment`.
//
// The `AWS::AuditManager::Assessment` resource is an Audit Manager resource type that defines the scope of audit evidence collected by Audit Manager . An Audit Manager assessment is an implementation of an Audit Manager framework.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssessment := awscdk.Aws_auditmanager.NewCfnAssessment(this, jsii.String("MyCfnAssessment"), &cfnAssessmentProps{
//   	assessmentReportsDestination: &assessmentReportsDestinationProperty{
//   		destination: jsii.String("destination"),
//   		destinationType: jsii.String("destinationType"),
//   	},
//   	awsAccount: &aWSAccountProperty{
//   		emailAddress: jsii.String("emailAddress"),
//   		id: jsii.String("id"),
//   		name: jsii.String("name"),
//   	},
//   	delegations: []interface{}{
//   		&delegationProperty{
//   			assessmentId: jsii.String("assessmentId"),
//   			assessmentName: jsii.String("assessmentName"),
//   			comment: jsii.String("comment"),
//   			controlSetId: jsii.String("controlSetId"),
//   			createdBy: jsii.String("createdBy"),
//   			creationTime: jsii.Number(123),
//   			id: jsii.String("id"),
//   			lastUpdated: jsii.Number(123),
//   			roleArn: jsii.String("roleArn"),
//   			roleType: jsii.String("roleType"),
//   			status: jsii.String("status"),
//   		},
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
	// The AWS account that's associated with the assessment.
	AwsAccount() interface{}
	SetAwsAccount(val interface{})
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
	// The delegations that are associated with the assessment.
	Delegations() interface{}
	SetDelegations(val interface{})
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
	// Experimental.
	LogicalId() *string
	// The name of the assessment.
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
	// The roles that are associated with the assessment.
	Roles() interface{}
	SetRoles(val interface{})
	// The wrapper of AWS accounts and services that are in scope for the assessment.
	Scope() interface{}
	SetScope(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The overall status of the assessment.
	//
	// When you create a new assessment, the initial `Status` value is always `ACTIVE` . When you create an assessment, even if you specify the value as `INACTIVE` , the value overrides to `ACTIVE` .
	//
	// After you create an assessment, you can change the value of the `Status` property at any time. For example, when you want to stop collecting evidence for your assessment, you can change the assessment status to `INACTIVE` .
	Status() *string
	SetStatus(val *string)
	// The tags that are associated with the assessment.
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

func (j *jsiiProxy_CfnAssessment) Delegations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delegations",
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

func (j *jsiiProxy_CfnAssessment) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnAssessment(scope awscdk.Construct, id *string, props *CfnAssessmentProps) CfnAssessment {
	_init_.Initialize()

	if err := validateNewCfnAssessmentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAssessment{}

	_jsii_.Create(
		"monocdk.aws_auditmanager.CfnAssessment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AuditManager::Assessment`.
func NewCfnAssessment_Override(c CfnAssessment, scope awscdk.Construct, id *string, props *CfnAssessmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_auditmanager.CfnAssessment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAssessment)SetAssessmentReportsDestination(val interface{}) {
	if err := j.validateSetAssessmentReportsDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assessmentReportsDestination",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment)SetAwsAccount(val interface{}) {
	if err := j.validateSetAwsAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccount",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment)SetDelegations(val interface{}) {
	if err := j.validateSetDelegationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delegations",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment)SetFrameworkId(val *string) {
	_jsii_.Set(
		j,
		"frameworkId",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment)SetRoles(val interface{}) {
	if err := j.validateSetRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roles",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment)SetScope(val interface{}) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_CfnAssessment)SetStatus(val *string) {
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
// Experimental.
func CfnAssessment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssessment_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_auditmanager.CfnAssessment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAssessment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnAssessment_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_auditmanager.CfnAssessment",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAssessment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssessment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_auditmanager.CfnAssessment",
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
		"monocdk.aws_auditmanager.CfnAssessment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssessment) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAssessment) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAssessment) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAssessment) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAssessment) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAssessment) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAssessment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAssessment) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnAssessment) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnAssessment) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAssessment) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAssessment) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAssessment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssessment) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAssessment) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAssessment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnAssessment) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnAssessment) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssessment) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

