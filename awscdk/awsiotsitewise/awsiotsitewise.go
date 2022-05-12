package awsiotsitewise

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::IoTSiteWise::AccessPolicy`.
//
// Creates an access policy that grants the specified identity ( AWS SSO user, AWS SSO group, or IAM user) access to the specified AWS IoT SiteWise Monitor portal or project resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessPolicy := awscdk.Aws_iotsitewise.NewCfnAccessPolicy(this, jsii.String("MyCfnAccessPolicy"), &cfnAccessPolicyProps{
//   	accessPolicyIdentity: &accessPolicyIdentityProperty{
//   		iamRole: &iamRoleProperty{
//   			arn: jsii.String("arn"),
//   		},
//   		iamUser: &iamUserProperty{
//   			arn: jsii.String("arn"),
//   		},
//   		user: &userProperty{
//   			id: jsii.String("id"),
//   		},
//   	},
//   	accessPolicyPermission: jsii.String("accessPolicyPermission"),
//   	accessPolicyResource: &accessPolicyResourceProperty{
//   		portal: &portalProperty{
//   			id: jsii.String("id"),
//   		},
//   		project: &projectProperty{
//   			id: jsii.String("id"),
//   		},
//   	},
//   })
//
type CfnAccessPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The identity for this access policy.
	//
	// Choose an AWS SSO user, an AWS SSO group, or an IAM user.
	AccessPolicyIdentity() interface{}
	SetAccessPolicyIdentity(val interface{})
	// The permission level for this access policy.
	//
	// Choose either a `ADMINISTRATOR` or `VIEWER` . Note that a project `ADMINISTRATOR` is also known as a project owner.
	AccessPolicyPermission() *string
	SetAccessPolicyPermission(val *string)
	// The AWS IoT SiteWise Monitor resource for this access policy.
	//
	// Choose either a portal or a project.
	AccessPolicyResource() interface{}
	SetAccessPolicyResource(val interface{})
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the access policy, which has the following format.
	//
	// `arn:${Partition}:iotsitewise:${Region}:${Account}:access-policy/${AccessPolicyId}`.
	AttrAccessPolicyArn() *string
	// The ID of the access policy.
	AttrAccessPolicyId() *string
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnAccessPolicy
type jsiiProxy_CfnAccessPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAccessPolicy) AccessPolicyIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessPolicyIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) AccessPolicyPermission() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicyPermission",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) AccessPolicyResource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessPolicyResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) AttrAccessPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAccessPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) AttrAccessPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAccessPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTSiteWise::AccessPolicy`.
func NewCfnAccessPolicy(scope awscdk.Construct, id *string, props *CfnAccessPolicyProps) CfnAccessPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnAccessPolicy{}

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnAccessPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTSiteWise::AccessPolicy`.
func NewCfnAccessPolicy_Override(c CfnAccessPolicy, scope awscdk.Construct, id *string, props *CfnAccessPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnAccessPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAccessPolicy) SetAccessPolicyIdentity(val interface{}) {
	_jsii_.Set(
		j,
		"accessPolicyIdentity",
		val,
	)
}

func (j *jsiiProxy_CfnAccessPolicy) SetAccessPolicyPermission(val *string) {
	_jsii_.Set(
		j,
		"accessPolicyPermission",
		val,
	)
}

func (j *jsiiProxy_CfnAccessPolicy) SetAccessPolicyResource(val interface{}) {
	_jsii_.Set(
		j,
		"accessPolicyResource",
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
func CfnAccessPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnAccessPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAccessPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnAccessPolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAccessPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnAccessPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccessPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iotsitewise.CfnAccessPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccessPolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAccessPolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAccessPolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAccessPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAccessPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAccessPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAccessPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAccessPolicy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccessPolicy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccessPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAccessPolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAccessPolicy) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAccessPolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccessPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAccessPolicy) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAccessPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccessPolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccessPolicy) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAccessPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccessPolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccessPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The identity ( AWS SSO user, AWS SSO group, or IAM user) to which this access policy applies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessPolicyIdentityProperty := &accessPolicyIdentityProperty{
//   	iamRole: &iamRoleProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	iamUser: &iamUserProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	user: &userProperty{
//   		id: jsii.String("id"),
//   	},
//   }
//
type CfnAccessPolicy_AccessPolicyIdentityProperty struct {
	// An IAM role identity.
	IamRole interface{} `field:"optional" json:"iamRole" yaml:"iamRole"`
	// An IAM user identity.
	IamUser interface{} `field:"optional" json:"iamUser" yaml:"iamUser"`
	// The AWS SSO user to which this access policy maps.
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

// The AWS IoT SiteWise Monitor resource for this access policy.
//
// Choose either a portal or a project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessPolicyResourceProperty := &accessPolicyResourceProperty{
//   	portal: &portalProperty{
//   		id: jsii.String("id"),
//   	},
//   	project: &projectProperty{
//   		id: jsii.String("id"),
//   	},
//   }
//
type CfnAccessPolicy_AccessPolicyResourceProperty struct {
	// The AWS IoT SiteWise Monitor portal for this access policy.
	Portal interface{} `field:"optional" json:"portal" yaml:"portal"`
	// The AWS IoT SiteWise Monitor project for this access policy.
	Project interface{} `field:"optional" json:"project" yaml:"project"`
}

// Contains information about an AWS Identity and Access Management role.
//
// For more information, see [IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) in the *IAM User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamRoleProperty := &iamRoleProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnAccessPolicy_IamRoleProperty struct {
	// The ARN of the IAM role.
	//
	// For more information, see [IAM ARNs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html) in the *IAM User Guide* .
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

// Contains information about an AWS Identity and Access Management user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamUserProperty := &iamUserProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnAccessPolicy_IamUserProperty struct {
	// The ARN of the IAM user. For more information, see [IAM ARNs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html) in the *IAM User Guide* .
	//
	// > If you delete the IAM user, access policies that contain this identity include an empty `arn` . You can delete the access policy for the IAM user that no longer exists.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

// The `Portal` property type specifies the AWS IoT SiteWise Monitor portal for an [AWS::IoTSiteWise::AccessPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-accesspolicy.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portalProperty := &portalProperty{
//   	id: jsii.String("id"),
//   }
//
type CfnAccessPolicy_PortalProperty struct {
	// The ID of the portal.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

// The `Project` property type specifies the AWS IoT SiteWise Monitor project for an [AWS::IoTSiteWise::AccessPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-accesspolicy.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectProperty := &projectProperty{
//   	id: jsii.String("id"),
//   }
//
type CfnAccessPolicy_ProjectProperty struct {
	// The ID of the project.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

// The `User` property type specifies the AWS IoT SiteWise Monitor user for an [AWS::IoTSiteWise::AccessPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-accesspolicy.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userProperty := &userProperty{
//   	id: jsii.String("id"),
//   }
//
type CfnAccessPolicy_UserProperty struct {
	// The ID of the user.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

// Properties for defining a `CfnAccessPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessPolicyProps := &cfnAccessPolicyProps{
//   	accessPolicyIdentity: &accessPolicyIdentityProperty{
//   		iamRole: &iamRoleProperty{
//   			arn: jsii.String("arn"),
//   		},
//   		iamUser: &iamUserProperty{
//   			arn: jsii.String("arn"),
//   		},
//   		user: &userProperty{
//   			id: jsii.String("id"),
//   		},
//   	},
//   	accessPolicyPermission: jsii.String("accessPolicyPermission"),
//   	accessPolicyResource: &accessPolicyResourceProperty{
//   		portal: &portalProperty{
//   			id: jsii.String("id"),
//   		},
//   		project: &projectProperty{
//   			id: jsii.String("id"),
//   		},
//   	},
//   }
//
type CfnAccessPolicyProps struct {
	// The identity for this access policy.
	//
	// Choose an AWS SSO user, an AWS SSO group, or an IAM user.
	AccessPolicyIdentity interface{} `field:"required" json:"accessPolicyIdentity" yaml:"accessPolicyIdentity"`
	// The permission level for this access policy.
	//
	// Choose either a `ADMINISTRATOR` or `VIEWER` . Note that a project `ADMINISTRATOR` is also known as a project owner.
	AccessPolicyPermission *string `field:"required" json:"accessPolicyPermission" yaml:"accessPolicyPermission"`
	// The AWS IoT SiteWise Monitor resource for this access policy.
	//
	// Choose either a portal or a project.
	AccessPolicyResource interface{} `field:"required" json:"accessPolicyResource" yaml:"accessPolicyResource"`
}

// A CloudFormation `AWS::IoTSiteWise::Asset`.
//
// Creates an asset from an existing asset model. For more information, see [Creating assets](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-assets.html) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAsset := awscdk.Aws_iotsitewise.NewCfnAsset(this, jsii.String("MyCfnAsset"), &cfnAssetProps{
//   	assetModelId: jsii.String("assetModelId"),
//   	assetName: jsii.String("assetName"),
//
//   	// the properties below are optional
//   	assetHierarchies: []interface{}{
//   		&assetHierarchyProperty{
//   			childAssetId: jsii.String("childAssetId"),
//   			logicalId: jsii.String("logicalId"),
//   		},
//   	},
//   	assetProperties: []interface{}{
//   		&assetPropertyProperty{
//   			logicalId: jsii.String("logicalId"),
//
//   			// the properties below are optional
//   			alias: jsii.String("alias"),
//   			notificationState: jsii.String("notificationState"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnAsset interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A list of asset hierarchies that each contain a `hierarchyLogicalId` .
	//
	// A hierarchy specifies allowed parent/child asset relationships.
	AssetHierarchies() interface{}
	SetAssetHierarchies(val interface{})
	// The ID of the asset model from which to create the asset.
	AssetModelId() *string
	SetAssetModelId(val *string)
	// A unique, friendly name for the asset.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	AssetName() *string
	SetAssetName(val *string)
	// The list of asset properties for the asset.
	//
	// This object doesn't include properties that you define in composite models. You can find composite model properties in the `assetCompositeModels` object.
	AssetProperties() interface{}
	SetAssetProperties(val interface{})
	// The ARN of the asset.
	AttrAssetArn() *string
	// The ID of the asset.
	AttrAssetId() *string
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of key-value pairs that contain metadata for the asset.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
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

// The jsii proxy struct for CfnAsset
type jsiiProxy_CfnAsset struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAsset) AssetHierarchies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetHierarchies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) AssetModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) AssetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) AssetProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) AttrAssetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) AttrAssetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTSiteWise::Asset`.
func NewCfnAsset(scope awscdk.Construct, id *string, props *CfnAssetProps) CfnAsset {
	_init_.Initialize()

	j := jsiiProxy_CfnAsset{}

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnAsset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTSiteWise::Asset`.
func NewCfnAsset_Override(c CfnAsset, scope awscdk.Construct, id *string, props *CfnAssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnAsset",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAsset) SetAssetHierarchies(val interface{}) {
	_jsii_.Set(
		j,
		"assetHierarchies",
		val,
	)
}

func (j *jsiiProxy_CfnAsset) SetAssetModelId(val *string) {
	_jsii_.Set(
		j,
		"assetModelId",
		val,
	)
}

func (j *jsiiProxy_CfnAsset) SetAssetName(val *string) {
	_jsii_.Set(
		j,
		"assetName",
		val,
	)
}

func (j *jsiiProxy_CfnAsset) SetAssetProperties(val interface{}) {
	_jsii_.Set(
		j,
		"assetProperties",
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
func CfnAsset_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnAsset",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAsset_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnAsset",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAsset_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iotsitewise.CfnAsset",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAsset) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAsset) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAsset) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAsset) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAsset) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAsset) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAsset) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAsset) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAsset) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAsset) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAsset) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAsset) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAsset) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAsset) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAsset) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAsset) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAsset) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAsset) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAsset) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAsset) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes an asset hierarchy that contains a `childAssetId` and `hierarchyLogicalId` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetHierarchyProperty := &assetHierarchyProperty{
//   	childAssetId: jsii.String("childAssetId"),
//   	logicalId: jsii.String("logicalId"),
//   }
//
type CfnAsset_AssetHierarchyProperty struct {
	// The Id of the child asset.
	ChildAssetId *string `field:"required" json:"childAssetId" yaml:"childAssetId"`
	// The `LogicalID` of the hierarchy. This ID is a `hierarchyLogicalId` .
	//
	// The maximum length is 256 characters, with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
}

// Contains asset property information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyProperty := &assetPropertyProperty{
//   	logicalId: jsii.String("logicalId"),
//
//   	// the properties below are optional
//   	alias: jsii.String("alias"),
//   	notificationState: jsii.String("notificationState"),
//   }
//
type CfnAsset_AssetPropertyProperty struct {
	// The `LogicalID` of the asset property.
	//
	// The maximum length is 256 characters, with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The property alias that identifies the property, such as an OPC-UA server data stream path (for example, `/company/windfarm/3/turbine/7/temperature` ).
	//
	// For more information, see [Mapping industrial data streams to asset properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html) in the *AWS IoT SiteWise User Guide* .
	//
	// The property alias must have 1-1000 characters.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The MQTT notification state (enabled or disabled) for this asset property.
	//
	// When the notification state is enabled, AWS IoT SiteWise publishes property value updates to a unique MQTT topic. For more information, see [Interacting with other services](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/interact-with-other-services.html) in the *AWS IoT SiteWise User Guide* .
	//
	// If you omit this parameter, the notification state is set to `DISABLED` .
	NotificationState *string `field:"optional" json:"notificationState" yaml:"notificationState"`
}

// A CloudFormation `AWS::IoTSiteWise::AssetModel`.
//
// Creates an asset model from specified property and hierarchy definitions. You create assets from asset models. With asset models, you can easily create assets of the same type that have standardized definitions. Each asset created from a model inherits the asset model's property and hierarchy definitions. For more information, see [Defining asset models](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/define-models.html) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssetModel := awscdk.Aws_iotsitewise.NewCfnAssetModel(this, jsii.String("MyCfnAssetModel"), &cfnAssetModelProps{
//   	assetModelName: jsii.String("assetModelName"),
//
//   	// the properties below are optional
//   	assetModelCompositeModels: []interface{}{
//   		&assetModelCompositeModelProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			compositeModelProperties: []interface{}{
//   				&assetModelPropertyProperty{
//   					dataType: jsii.String("dataType"),
//   					logicalId: jsii.String("logicalId"),
//   					name: jsii.String("name"),
//   					type: &propertyTypeProperty{
//   						typeName: jsii.String("typeName"),
//
//   						// the properties below are optional
//   						attribute: &attributeProperty{
//   							defaultValue: jsii.String("defaultValue"),
//   						},
//   						metric: &metricProperty{
//   							expression: jsii.String("expression"),
//   							variables: []interface{}{
//   								&expressionVariableProperty{
//   									name: jsii.String("name"),
//   									value: &variableValueProperty{
//   										propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   										// the properties below are optional
//   										hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   									},
//   								},
//   							},
//   							window: &metricWindowProperty{
//   								tumbling: &tumblingWindowProperty{
//   									interval: jsii.String("interval"),
//
//   									// the properties below are optional
//   									offset: jsii.String("offset"),
//   								},
//   							},
//   						},
//   						transform: &transformProperty{
//   							expression: jsii.String("expression"),
//   							variables: []interface{}{
//   								&expressionVariableProperty{
//   									name: jsii.String("name"),
//   									value: &variableValueProperty{
//   										propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   										// the properties below are optional
//   										hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					dataTypeSpec: jsii.String("dataTypeSpec"),
//   					unit: jsii.String("unit"),
//   				},
//   			},
//   			description: jsii.String("description"),
//   		},
//   	},
//   	assetModelDescription: jsii.String("assetModelDescription"),
//   	assetModelHierarchies: []interface{}{
//   		&assetModelHierarchyProperty{
//   			childAssetModelId: jsii.String("childAssetModelId"),
//   			logicalId: jsii.String("logicalId"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	assetModelProperties: []interface{}{
//   		&assetModelPropertyProperty{
//   			dataType: jsii.String("dataType"),
//   			logicalId: jsii.String("logicalId"),
//   			name: jsii.String("name"),
//   			type: &propertyTypeProperty{
//   				typeName: jsii.String("typeName"),
//
//   				// the properties below are optional
//   				attribute: &attributeProperty{
//   					defaultValue: jsii.String("defaultValue"),
//   				},
//   				metric: &metricProperty{
//   					expression: jsii.String("expression"),
//   					variables: []interface{}{
//   						&expressionVariableProperty{
//   							name: jsii.String("name"),
//   							value: &variableValueProperty{
//   								propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   								// the properties below are optional
//   								hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   							},
//   						},
//   					},
//   					window: &metricWindowProperty{
//   						tumbling: &tumblingWindowProperty{
//   							interval: jsii.String("interval"),
//
//   							// the properties below are optional
//   							offset: jsii.String("offset"),
//   						},
//   					},
//   				},
//   				transform: &transformProperty{
//   					expression: jsii.String("expression"),
//   					variables: []interface{}{
//   						&expressionVariableProperty{
//   							name: jsii.String("name"),
//   							value: &variableValueProperty{
//   								propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   								// the properties below are optional
//   								hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   							},
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			dataTypeSpec: jsii.String("dataTypeSpec"),
//   			unit: jsii.String("unit"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnAssetModel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The composite asset models that are part of this asset model.
	//
	// Composite asset models are asset models that contain specific properties. Each composite model has a type that defines the properties that the composite model supports. You can use composite asset models to define alarms on this asset model.
	AssetModelCompositeModels() interface{}
	SetAssetModelCompositeModels(val interface{})
	// A description for the asset model.
	AssetModelDescription() *string
	SetAssetModelDescription(val *string)
	// The hierarchy definitions of the asset model.
	//
	// Each hierarchy specifies an asset model whose assets can be children of any other assets created from this asset model. For more information, see [Defining relationships between assets](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html) in the *AWS IoT SiteWise User Guide* .
	//
	// You can specify up to 10 hierarchies per asset model. For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	AssetModelHierarchies() interface{}
	SetAssetModelHierarchies(val interface{})
	// A unique, friendly name for the asset model.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	AssetModelName() *string
	SetAssetModelName(val *string)
	// The property definitions of the asset model.
	//
	// For more information, see [Defining data properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html) in the *AWS IoT SiteWise User Guide* .
	//
	// You can specify up to 200 properties per asset model. For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	AssetModelProperties() interface{}
	SetAssetModelProperties(val interface{})
	AttrAssetModelArn() *string
	// The ID of the asset model.
	AttrAssetModelId() *string
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of key-value pairs that contain metadata for the asset.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
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

// The jsii proxy struct for CfnAssetModel
type jsiiProxy_CfnAssetModel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAssetModel) AssetModelCompositeModels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetModelCompositeModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AssetModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AssetModelHierarchies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetModelHierarchies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AssetModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AssetModelProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetModelProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AttrAssetModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssetModelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AttrAssetModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssetModelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTSiteWise::AssetModel`.
func NewCfnAssetModel(scope awscdk.Construct, id *string, props *CfnAssetModelProps) CfnAssetModel {
	_init_.Initialize()

	j := jsiiProxy_CfnAssetModel{}

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnAssetModel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTSiteWise::AssetModel`.
func NewCfnAssetModel_Override(c CfnAssetModel, scope awscdk.Construct, id *string, props *CfnAssetModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnAssetModel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAssetModel) SetAssetModelCompositeModels(val interface{}) {
	_jsii_.Set(
		j,
		"assetModelCompositeModels",
		val,
	)
}

func (j *jsiiProxy_CfnAssetModel) SetAssetModelDescription(val *string) {
	_jsii_.Set(
		j,
		"assetModelDescription",
		val,
	)
}

func (j *jsiiProxy_CfnAssetModel) SetAssetModelHierarchies(val interface{}) {
	_jsii_.Set(
		j,
		"assetModelHierarchies",
		val,
	)
}

func (j *jsiiProxy_CfnAssetModel) SetAssetModelName(val *string) {
	_jsii_.Set(
		j,
		"assetModelName",
		val,
	)
}

func (j *jsiiProxy_CfnAssetModel) SetAssetModelProperties(val interface{}) {
	_jsii_.Set(
		j,
		"assetModelProperties",
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
func CfnAssetModel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnAssetModel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAssetModel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnAssetModel",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAssetModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnAssetModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssetModel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iotsitewise.CfnAssetModel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssetModel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAssetModel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAssetModel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAssetModel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAssetModel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAssetModel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAssetModel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAssetModel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssetModel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssetModel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAssetModel) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAssetModel) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAssetModel) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssetModel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAssetModel) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAssetModel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssetModel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssetModel) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAssetModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssetModel) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssetModel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Contains information about a composite model in an asset model.
//
// This object contains the asset property definitions that you define in the composite model. You can use composite asset models to define alarms on this asset model.
//
// If you use the `AssetModelCompositeModel` property to create an alarm, you must use the following information to define three asset model properties:
//
// - Use an asset model property to specify the alarm type.
//
// - The name must be `AWS/ALARM_TYPE` .
// - The data type must be `STRING` .
// - For the `Type` property, the type name must be `Attribute` and the default value must be `IOT_EVENTS` .
// - Use an asset model property to specify the alarm source.
//
// - The name must be `AWS/ALARM_SOURCE` .
// - The data type must be `STRING` .
// - For the `Type` property, the type name must be `Attribute` and the default value must be the ARN of the alarm model that you created in AWS IoT Events .
//
// > For the ARN of the alarm model, you can use the `Fn::Sub` intrinsic function to substitute the `AWS::Partition` , `AWS::Region` , and `AWS::AccountId` variables in an input string with values that you specify.
// >
// > For example, `Fn::Sub: "arn:${AWS::Partition}:iotevents:${AWS::Region}:${AWS::AccountId}:alarmModel/TestAlarmModel"` .
// >
// > Replace `TestAlarmModel` with the name of your alarm model.
// >
// > For more information about using the `Fn::Sub` intrinsic function, see [Fn::Sub](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-sub.html) .
// - Use an asset model property to specify the state of the alarm.
//
// - The name must be `AWS/ALARM_STATE` .
// - The data type must be `STRUCT` .
// - The `DataTypeSpec` value must be `AWS/ALARM_STATE` .
// - For the `Type` property, the type name must be `Measurement` .
//
// At the bottom of this page, we provide a YAML example that you can modify to create an alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetModelCompositeModelProperty := &assetModelCompositeModelProperty{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	compositeModelProperties: []interface{}{
//   		&assetModelPropertyProperty{
//   			dataType: jsii.String("dataType"),
//   			logicalId: jsii.String("logicalId"),
//   			name: jsii.String("name"),
//   			type: &propertyTypeProperty{
//   				typeName: jsii.String("typeName"),
//
//   				// the properties below are optional
//   				attribute: &attributeProperty{
//   					defaultValue: jsii.String("defaultValue"),
//   				},
//   				metric: &metricProperty{
//   					expression: jsii.String("expression"),
//   					variables: []interface{}{
//   						&expressionVariableProperty{
//   							name: jsii.String("name"),
//   							value: &variableValueProperty{
//   								propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   								// the properties below are optional
//   								hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   							},
//   						},
//   					},
//   					window: &metricWindowProperty{
//   						tumbling: &tumblingWindowProperty{
//   							interval: jsii.String("interval"),
//
//   							// the properties below are optional
//   							offset: jsii.String("offset"),
//   						},
//   					},
//   				},
//   				transform: &transformProperty{
//   					expression: jsii.String("expression"),
//   					variables: []interface{}{
//   						&expressionVariableProperty{
//   							name: jsii.String("name"),
//   							value: &variableValueProperty{
//   								propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   								// the properties below are optional
//   								hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   							},
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			dataTypeSpec: jsii.String("dataTypeSpec"),
//   			unit: jsii.String("unit"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   }
//
type CfnAssetModel_AssetModelCompositeModelProperty struct {
	// The name of the composite model.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the composite model.
	//
	// For alarm composite models, this type is `AWS/ALARM` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// The asset property definitions for this composite model.
	CompositeModelProperties interface{} `field:"optional" json:"compositeModelProperties" yaml:"compositeModelProperties"`
	// The description of the composite model.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

// Describes an asset hierarchy that contains a hierarchy's name, `LogicalID` , and child asset model ID that specifies the type of asset that can be in this hierarchy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetModelHierarchyProperty := &assetModelHierarchyProperty{
//   	childAssetModelId: jsii.String("childAssetModelId"),
//   	logicalId: jsii.String("logicalId"),
//   	name: jsii.String("name"),
//   }
//
type CfnAssetModel_AssetModelHierarchyProperty struct {
	// The Id of the asset model.
	ChildAssetModelId *string `field:"required" json:"childAssetModelId" yaml:"childAssetModelId"`
	// The `LogicalID` of the asset model hierarchy. This ID is a `hierarchyLogicalId` .
	//
	// The maximum length is 256 characters, with the pattern `[^\ u0000-\ u001F\ u007F]+`.
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The name of the asset model hierarchy.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Contains information about an asset model property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetModelPropertyProperty := &assetModelPropertyProperty{
//   	dataType: jsii.String("dataType"),
//   	logicalId: jsii.String("logicalId"),
//   	name: jsii.String("name"),
//   	type: &propertyTypeProperty{
//   		typeName: jsii.String("typeName"),
//
//   		// the properties below are optional
//   		attribute: &attributeProperty{
//   			defaultValue: jsii.String("defaultValue"),
//   		},
//   		metric: &metricProperty{
//   			expression: jsii.String("expression"),
//   			variables: []interface{}{
//   				&expressionVariableProperty{
//   					name: jsii.String("name"),
//   					value: &variableValueProperty{
//   						propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   						// the properties below are optional
//   						hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   					},
//   				},
//   			},
//   			window: &metricWindowProperty{
//   				tumbling: &tumblingWindowProperty{
//   					interval: jsii.String("interval"),
//
//   					// the properties below are optional
//   					offset: jsii.String("offset"),
//   				},
//   			},
//   		},
//   		transform: &transformProperty{
//   			expression: jsii.String("expression"),
//   			variables: []interface{}{
//   				&expressionVariableProperty{
//   					name: jsii.String("name"),
//   					value: &variableValueProperty{
//   						propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   						// the properties below are optional
//   						hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	dataTypeSpec: jsii.String("dataTypeSpec"),
//   	unit: jsii.String("unit"),
//   }
//
type CfnAssetModel_AssetModelPropertyProperty struct {
	// The data type of the asset model property.
	//
	// The value can be `STRING` , `INTEGER` , `DOUBLE` , `BOOLEAN` , or `STRUCT` .
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The `LogicalID` of the asset model property.
	//
	// The maximum length is 256 characters, with the pattern `[^\\ u0000-\\ u001F\\ u007F]+` .
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The name of the asset model property.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contains a property type, which can be one of `Attribute` , `Measurement` , `Metric` , or `Transform` .
	Type interface{} `field:"required" json:"type" yaml:"type"`
	// The data type of the structure for this property.
	//
	// This parameter exists on properties that have the `STRUCT` data type.
	DataTypeSpec *string `field:"optional" json:"dataTypeSpec" yaml:"dataTypeSpec"`
	// The unit of the asset model property, such as `Newtons` or `RPM` .
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

// Contains an asset attribute property.
//
// For more information, see [Defining data properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html#attributes) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeProperty := &attributeProperty{
//   	defaultValue: jsii.String("defaultValue"),
//   }
//
type CfnAssetModel_AttributeProperty struct {
	// The default value of the asset model property attribute.
	//
	// All assets that you create from the asset model contain this attribute value. You can update an attribute's value after you create an asset. For more information, see [Updating attribute values](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/update-attribute-values.html) in the *AWS IoT SiteWise User Guide* .
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
}

// Contains expression variable information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expressionVariableProperty := &expressionVariableProperty{
//   	name: jsii.String("name"),
//   	value: &variableValueProperty{
//   		propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   		// the properties below are optional
//   		hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   	},
//   }
//
type CfnAssetModel_ExpressionVariableProperty struct {
	// The friendly name of the variable to be used in the expression.
	//
	// The maximum length is 64 characters with the pattern `^[a-z][a-z0-9_]*$` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The variable that identifies an asset property from which to use values.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

// Contains an asset metric property.
//
// With metrics, you can calculate aggregate functions, such as an average, maximum, or minimum, as specified through an expression. A metric maps several values to a single value (such as a sum).
//
// The maximum number of dependent/cascading variables used in any one metric calculation is 10. Therefore, a *root* metric can have up to 10 cascading metrics in its computational dependency tree. Additionally, a metric can only have a data type of `DOUBLE` and consume properties with data types of `INTEGER` or `DOUBLE` .
//
// For more information, see [Defining data properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html#metrics) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricProperty := &metricProperty{
//   	expression: jsii.String("expression"),
//   	variables: []interface{}{
//   		&expressionVariableProperty{
//   			name: jsii.String("name"),
//   			value: &variableValueProperty{
//   				propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   				// the properties below are optional
//   				hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   			},
//   		},
//   	},
//   	window: &metricWindowProperty{
//   		tumbling: &tumblingWindowProperty{
//   			interval: jsii.String("interval"),
//
//   			// the properties below are optional
//   			offset: jsii.String("offset"),
//   		},
//   	},
//   }
//
type CfnAssetModel_MetricProperty struct {
	// The mathematical expression that defines the metric aggregation function.
	//
	// You can specify up to 10 variables per expression. You can specify up to 10 functions per expression.
	//
	// For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The list of variables used in the expression.
	Variables interface{} `field:"required" json:"variables" yaml:"variables"`
	// The window (time interval) over which AWS IoT SiteWise computes the metric's aggregation expression.
	//
	// AWS IoT SiteWise computes one data point per `window` .
	Window interface{} `field:"required" json:"window" yaml:"window"`
}

// Contains a time interval window used for data aggregate computations (for example, average, sum, count, and so on).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricWindowProperty := &metricWindowProperty{
//   	tumbling: &tumblingWindowProperty{
//   		interval: jsii.String("interval"),
//
//   		// the properties below are optional
//   		offset: jsii.String("offset"),
//   	},
//   }
//
type CfnAssetModel_MetricWindowProperty struct {
	// The tumbling time interval window.
	Tumbling interface{} `field:"optional" json:"tumbling" yaml:"tumbling"`
}

// Contains a property type, which can be one of `Attribute` , `Measurement` , `Metric` , or `Transform` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   propertyTypeProperty := &propertyTypeProperty{
//   	typeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	attribute: &attributeProperty{
//   		defaultValue: jsii.String("defaultValue"),
//   	},
//   	metric: &metricProperty{
//   		expression: jsii.String("expression"),
//   		variables: []interface{}{
//   			&expressionVariableProperty{
//   				name: jsii.String("name"),
//   				value: &variableValueProperty{
//   					propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   					// the properties below are optional
//   					hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   				},
//   			},
//   		},
//   		window: &metricWindowProperty{
//   			tumbling: &tumblingWindowProperty{
//   				interval: jsii.String("interval"),
//
//   				// the properties below are optional
//   				offset: jsii.String("offset"),
//   			},
//   		},
//   	},
//   	transform: &transformProperty{
//   		expression: jsii.String("expression"),
//   		variables: []interface{}{
//   			&expressionVariableProperty{
//   				name: jsii.String("name"),
//   				value: &variableValueProperty{
//   					propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   					// the properties below are optional
//   					hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnAssetModel_PropertyTypeProperty struct {
	// The type of property type, which can be one of `Attribute` , `Measurement` , `Metric` , or `Transform` .
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// Specifies an asset attribute property.
	//
	// An attribute generally contains static information, such as the serial number of an [industrial IoT](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Internet_of_things#Industrial_applications) wind turbine.
	//
	// This is required if the `TypeName` is `Attribute` and has a `DefaultValue` .
	Attribute interface{} `field:"optional" json:"attribute" yaml:"attribute"`
	// Specifies an asset metric property.
	//
	// A metric contains a mathematical expression that uses aggregate functions to process all input data points over a time interval and output a single data point, such as to calculate the average hourly temperature.
	//
	// This is required if the `TypeName` is `Metric` .
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
	// Specifies an asset transform property.
	//
	// A transform contains a mathematical expression that maps a property's data points from one form to another, such as a unit conversion from Celsius to Fahrenheit.
	//
	// This is required if the `TypeName` is `Transform` .
	Transform interface{} `field:"optional" json:"transform" yaml:"transform"`
}

// Contains an asset transform property.
//
// A transform is a one-to-one mapping of a property's data points from one form to another. For example, you can use a transform to convert a Celsius data stream to Fahrenheit by applying the transformation expression to each data point of the Celsius stream. Transforms can only input properties that are `INTEGER` , `DOUBLE` , or `BOOLEAN` type. Booleans convert to `0` ( `FALSE` ) and `1` ( `TRUE` )..
//
// For more information, see [Defining data properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html#transforms) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformProperty := &transformProperty{
//   	expression: jsii.String("expression"),
//   	variables: []interface{}{
//   		&expressionVariableProperty{
//   			name: jsii.String("name"),
//   			value: &variableValueProperty{
//   				propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   				// the properties below are optional
//   				hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   			},
//   		},
//   	},
//   }
//
type CfnAssetModel_TransformProperty struct {
	// The mathematical expression that defines the transformation function.
	//
	// You can specify up to 10 variables per expression. You can specify up to 10 functions per expression.
	//
	// For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The list of variables used in the expression.
	Variables interface{} `field:"required" json:"variables" yaml:"variables"`
}

// Contains a tumbling window, which is a repeating fixed-sized, non-overlapping, and contiguous time window.
//
// You can use this window in metrics to aggregate data from properties and other assets.
//
// You can use `m` , `h` , `d` , and `w` when you specify an interval or offset. Note that `m` represents minutes, `h` represents hours, `d` represents days, and `w` represents weeks. You can also use `s` to represent seconds in `offset` .
//
// The `interval` and `offset` parameters support the [ISO 8601 format](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_8601) . For example, `PT5S` represents 5 seconds, `PT5M` represents 5 minutes, and `PT5H` represents 5 hours.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tumblingWindowProperty := &tumblingWindowProperty{
//   	interval: jsii.String("interval"),
//
//   	// the properties below are optional
//   	offset: jsii.String("offset"),
//   }
//
type CfnAssetModel_TumblingWindowProperty struct {
	// The time interval for the tumbling window. The interval time must be between 1 minute and 1 week.
	//
	// AWS IoT SiteWise computes the `1w` interval the end of Sunday at midnight each week (UTC), the `1d` interval at the end of each day at midnight (UTC), the `1h` interval at the end of each hour, and so on.
	//
	// When AWS IoT SiteWise aggregates data points for metric computations, the start of each interval is exclusive and the end of each interval is inclusive. AWS IoT SiteWise places the computed data point at the end of the interval.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// The offset for the tumbling window. The `offset` parameter accepts the following:.
	//
	// - The offset time.
	//
	// For example, if you specify `18h` for `offset` and `1d` for `interval` , AWS IoT SiteWise aggregates data in one of the following ways:
	//
	// - If you create the metric before or at 6 PM (UTC), you get the first aggregation result at 6 PM (UTC) on the day when you create the metric.
	// - If you create the metric after 6 PM (UTC), you get the first aggregation result at 6 PM (UTC) the next day.
	// - The ISO 8601 format.
	//
	// For example, if you specify `PT18H` for `offset` and `1d` for `interval` , AWS IoT SiteWise aggregates data in one of the following ways:
	//
	// - If you create the metric before or at 6 PM (UTC), you get the first aggregation result at 6 PM (UTC) on the day when you create the metric.
	// - If you create the metric after 6 PM (UTC), you get the first aggregation result at 6 PM (UTC) the next day.
	// - The 24-hour clock.
	//
	// For example, if you specify `00:03:00` for `offset` , `5m` for `interval` , and you create the metric at 2 PM (UTC), you get the first aggregation result at 2:03 PM (UTC). You get the second aggregation result at 2:08 PM (UTC).
	// - The offset time zone.
	//
	// For example, if you specify `2021-07-23T18:00-08` for `offset` and `1d` for `interval` , AWS IoT SiteWise aggregates data in one of the following ways:
	//
	// - If you create the metric before or at 6 PM (PST), you get the first aggregation result at 6 PM (PST) on the day when you create the metric.
	// - If you create the metric after 6 PM (PST), you get the first aggregation result at 6 PM (PST) the next day.
	Offset *string `field:"optional" json:"offset" yaml:"offset"`
}

// Identifies a property value used in an expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   variableValueProperty := &variableValueProperty{
//   	propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   	// the properties below are optional
//   	hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   }
//
type CfnAssetModel_VariableValueProperty struct {
	// The `LogicalID` of the property to use as the variable.
	PropertyLogicalId *string `field:"required" json:"propertyLogicalId" yaml:"propertyLogicalId"`
	// The `LogicalID` of the hierarchy to query for the `PropertyLogicalID` .
	//
	// You use a `hierarchyLogicalID` instead of a model ID because you can have several hierarchies using the same model and therefore the same property. For example, you might have separately grouped assets that come from the same asset model. For more information, see [Defining relationships between assets](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html) in the *AWS IoT SiteWise User Guide* .
	HierarchyLogicalId *string `field:"optional" json:"hierarchyLogicalId" yaml:"hierarchyLogicalId"`
}

// Properties for defining a `CfnAssetModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssetModelProps := &cfnAssetModelProps{
//   	assetModelName: jsii.String("assetModelName"),
//
//   	// the properties below are optional
//   	assetModelCompositeModels: []interface{}{
//   		&assetModelCompositeModelProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			compositeModelProperties: []interface{}{
//   				&assetModelPropertyProperty{
//   					dataType: jsii.String("dataType"),
//   					logicalId: jsii.String("logicalId"),
//   					name: jsii.String("name"),
//   					type: &propertyTypeProperty{
//   						typeName: jsii.String("typeName"),
//
//   						// the properties below are optional
//   						attribute: &attributeProperty{
//   							defaultValue: jsii.String("defaultValue"),
//   						},
//   						metric: &metricProperty{
//   							expression: jsii.String("expression"),
//   							variables: []interface{}{
//   								&expressionVariableProperty{
//   									name: jsii.String("name"),
//   									value: &variableValueProperty{
//   										propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   										// the properties below are optional
//   										hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   									},
//   								},
//   							},
//   							window: &metricWindowProperty{
//   								tumbling: &tumblingWindowProperty{
//   									interval: jsii.String("interval"),
//
//   									// the properties below are optional
//   									offset: jsii.String("offset"),
//   								},
//   							},
//   						},
//   						transform: &transformProperty{
//   							expression: jsii.String("expression"),
//   							variables: []interface{}{
//   								&expressionVariableProperty{
//   									name: jsii.String("name"),
//   									value: &variableValueProperty{
//   										propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   										// the properties below are optional
//   										hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					dataTypeSpec: jsii.String("dataTypeSpec"),
//   					unit: jsii.String("unit"),
//   				},
//   			},
//   			description: jsii.String("description"),
//   		},
//   	},
//   	assetModelDescription: jsii.String("assetModelDescription"),
//   	assetModelHierarchies: []interface{}{
//   		&assetModelHierarchyProperty{
//   			childAssetModelId: jsii.String("childAssetModelId"),
//   			logicalId: jsii.String("logicalId"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	assetModelProperties: []interface{}{
//   		&assetModelPropertyProperty{
//   			dataType: jsii.String("dataType"),
//   			logicalId: jsii.String("logicalId"),
//   			name: jsii.String("name"),
//   			type: &propertyTypeProperty{
//   				typeName: jsii.String("typeName"),
//
//   				// the properties below are optional
//   				attribute: &attributeProperty{
//   					defaultValue: jsii.String("defaultValue"),
//   				},
//   				metric: &metricProperty{
//   					expression: jsii.String("expression"),
//   					variables: []interface{}{
//   						&expressionVariableProperty{
//   							name: jsii.String("name"),
//   							value: &variableValueProperty{
//   								propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   								// the properties below are optional
//   								hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   							},
//   						},
//   					},
//   					window: &metricWindowProperty{
//   						tumbling: &tumblingWindowProperty{
//   							interval: jsii.String("interval"),
//
//   							// the properties below are optional
//   							offset: jsii.String("offset"),
//   						},
//   					},
//   				},
//   				transform: &transformProperty{
//   					expression: jsii.String("expression"),
//   					variables: []interface{}{
//   						&expressionVariableProperty{
//   							name: jsii.String("name"),
//   							value: &variableValueProperty{
//   								propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   								// the properties below are optional
//   								hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   							},
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			dataTypeSpec: jsii.String("dataTypeSpec"),
//   			unit: jsii.String("unit"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAssetModelProps struct {
	// A unique, friendly name for the asset model.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	AssetModelName *string `field:"required" json:"assetModelName" yaml:"assetModelName"`
	// The composite asset models that are part of this asset model.
	//
	// Composite asset models are asset models that contain specific properties. Each composite model has a type that defines the properties that the composite model supports. You can use composite asset models to define alarms on this asset model.
	AssetModelCompositeModels interface{} `field:"optional" json:"assetModelCompositeModels" yaml:"assetModelCompositeModels"`
	// A description for the asset model.
	AssetModelDescription *string `field:"optional" json:"assetModelDescription" yaml:"assetModelDescription"`
	// The hierarchy definitions of the asset model.
	//
	// Each hierarchy specifies an asset model whose assets can be children of any other assets created from this asset model. For more information, see [Defining relationships between assets](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html) in the *AWS IoT SiteWise User Guide* .
	//
	// You can specify up to 10 hierarchies per asset model. For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	AssetModelHierarchies interface{} `field:"optional" json:"assetModelHierarchies" yaml:"assetModelHierarchies"`
	// The property definitions of the asset model.
	//
	// For more information, see [Defining data properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html) in the *AWS IoT SiteWise User Guide* .
	//
	// You can specify up to 200 properties per asset model. For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	AssetModelProperties interface{} `field:"optional" json:"assetModelProperties" yaml:"assetModelProperties"`
	// A list of key-value pairs that contain metadata for the asset.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnAsset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssetProps := &cfnAssetProps{
//   	assetModelId: jsii.String("assetModelId"),
//   	assetName: jsii.String("assetName"),
//
//   	// the properties below are optional
//   	assetHierarchies: []interface{}{
//   		&assetHierarchyProperty{
//   			childAssetId: jsii.String("childAssetId"),
//   			logicalId: jsii.String("logicalId"),
//   		},
//   	},
//   	assetProperties: []interface{}{
//   		&assetPropertyProperty{
//   			logicalId: jsii.String("logicalId"),
//
//   			// the properties below are optional
//   			alias: jsii.String("alias"),
//   			notificationState: jsii.String("notificationState"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAssetProps struct {
	// The ID of the asset model from which to create the asset.
	AssetModelId *string `field:"required" json:"assetModelId" yaml:"assetModelId"`
	// A unique, friendly name for the asset.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	AssetName *string `field:"required" json:"assetName" yaml:"assetName"`
	// A list of asset hierarchies that each contain a `hierarchyLogicalId` .
	//
	// A hierarchy specifies allowed parent/child asset relationships.
	AssetHierarchies interface{} `field:"optional" json:"assetHierarchies" yaml:"assetHierarchies"`
	// The list of asset properties for the asset.
	//
	// This object doesn't include properties that you define in composite models. You can find composite model properties in the `assetCompositeModels` object.
	AssetProperties interface{} `field:"optional" json:"assetProperties" yaml:"assetProperties"`
	// A list of key-value pairs that contain metadata for the asset.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTSiteWise::Dashboard`.
//
// Creates a dashboard in an AWS IoT SiteWise Monitor project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDashboard := awscdk.Aws_iotsitewise.NewCfnDashboard(this, jsii.String("MyCfnDashboard"), &cfnDashboardProps{
//   	dashboardDefinition: jsii.String("dashboardDefinition"),
//   	dashboardDescription: jsii.String("dashboardDescription"),
//   	dashboardName: jsii.String("dashboardName"),
//
//   	// the properties below are optional
//   	projectId: jsii.String("projectId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDashboard interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the dashboard, which has the following format.
	//
	// `arn:${Partition}:iotsitewise:${Region}:${Account}:dashboard/${DashboardId}`.
	AttrDashboardArn() *string
	// The ID of the dashboard.
	AttrDashboardId() *string
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
	// The dashboard definition specified in a JSON literal.
	//
	// For detailed information, see [Creating dashboards (CLI)](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-dashboards-using-aws-cli.html) in the *AWS IoT SiteWise User Guide* .
	DashboardDefinition() *string
	SetDashboardDefinition(val *string)
	// A description for the dashboard.
	DashboardDescription() *string
	SetDashboardDescription(val *string)
	// A friendly name for the dashboard.
	DashboardName() *string
	SetDashboardName(val *string)
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The ID of the project in which to create the dashboard.
	ProjectId() *string
	SetProjectId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of key-value pairs that contain metadata for the dashboard.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
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

// The jsii proxy struct for CfnDashboard
type jsiiProxy_CfnDashboard struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDashboard) AttrDashboardArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDashboardArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) AttrDashboardId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDashboardId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) DashboardDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) DashboardDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) DashboardName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTSiteWise::Dashboard`.
func NewCfnDashboard(scope awscdk.Construct, id *string, props *CfnDashboardProps) CfnDashboard {
	_init_.Initialize()

	j := jsiiProxy_CfnDashboard{}

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnDashboard",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTSiteWise::Dashboard`.
func NewCfnDashboard_Override(c CfnDashboard, scope awscdk.Construct, id *string, props *CfnDashboardProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnDashboard",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDashboard) SetDashboardDefinition(val *string) {
	_jsii_.Set(
		j,
		"dashboardDefinition",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetDashboardDescription(val *string) {
	_jsii_.Set(
		j,
		"dashboardDescription",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetDashboardName(val *string) {
	_jsii_.Set(
		j,
		"dashboardName",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetProjectId(val *string) {
	_jsii_.Set(
		j,
		"projectId",
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
func CfnDashboard_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnDashboard",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDashboard_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnDashboard",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDashboard_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnDashboard",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDashboard_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iotsitewise.CfnDashboard",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDashboard) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDashboard) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDashboard) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDashboard) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDashboard) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDashboard) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDashboard) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDashboard) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDashboard) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDashboard) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDashboard) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDashboard) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDashboard) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDashboard) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDashboard) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDashboard) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDashboard) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDashboard) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDashboard) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDashboard) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDashboard) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDashboard`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDashboardProps := &cfnDashboardProps{
//   	dashboardDefinition: jsii.String("dashboardDefinition"),
//   	dashboardDescription: jsii.String("dashboardDescription"),
//   	dashboardName: jsii.String("dashboardName"),
//
//   	// the properties below are optional
//   	projectId: jsii.String("projectId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDashboardProps struct {
	// The dashboard definition specified in a JSON literal.
	//
	// For detailed information, see [Creating dashboards (CLI)](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-dashboards-using-aws-cli.html) in the *AWS IoT SiteWise User Guide* .
	DashboardDefinition *string `field:"required" json:"dashboardDefinition" yaml:"dashboardDefinition"`
	// A description for the dashboard.
	DashboardDescription *string `field:"required" json:"dashboardDescription" yaml:"dashboardDescription"`
	// A friendly name for the dashboard.
	DashboardName *string `field:"required" json:"dashboardName" yaml:"dashboardName"`
	// The ID of the project in which to create the dashboard.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// A list of key-value pairs that contain metadata for the dashboard.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTSiteWise::Gateway`.
//
// Creates a gateway, which is a virtual or edge device that delivers industrial data streams from local servers to AWS IoT SiteWise . For more information, see [Ingesting data using a gateway](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/gateway-connector.html) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGateway := awscdk.Aws_iotsitewise.NewCfnGateway(this, jsii.String("MyCfnGateway"), &cfnGatewayProps{
//   	gatewayName: jsii.String("gatewayName"),
//   	gatewayPlatform: &gatewayPlatformProperty{
//   		greengrass: &greengrassProperty{
//   			groupArn: jsii.String("groupArn"),
//   		},
//   		greengrassV2: &greengrassV2Property{
//   			coreDeviceThingName: jsii.String("coreDeviceThingName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	gatewayCapabilitySummaries: []interface{}{
//   		&gatewayCapabilitySummaryProperty{
//   			capabilityNamespace: jsii.String("capabilityNamespace"),
//
//   			// the properties below are optional
//   			capabilityConfiguration: jsii.String("capabilityConfiguration"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnGateway interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID for the gateway.
	AttrGatewayId() *string
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
	// A list of gateway capability summaries that each contain a namespace and status.
	//
	// Each gateway capability defines data sources for the gateway. To retrieve a capability configuration's definition, use [DescribeGatewayCapabilityConfiguration](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeGatewayCapabilityConfiguration.html) .
	GatewayCapabilitySummaries() interface{}
	SetGatewayCapabilitySummaries(val interface{})
	// A unique, friendly name for the gateway.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	GatewayName() *string
	SetGatewayName(val *string)
	// The gateway's platform.
	//
	// You can only specify one platform in a gateway.
	GatewayPlatform() interface{}
	SetGatewayPlatform(val interface{})
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of key-value pairs that contain metadata for the gateway.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
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

// The jsii proxy struct for CfnGateway
type jsiiProxy_CfnGateway struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGateway) AttrGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) GatewayCapabilitySummaries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayCapabilitySummaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) GatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) GatewayPlatform() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGateway) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTSiteWise::Gateway`.
func NewCfnGateway(scope awscdk.Construct, id *string, props *CfnGatewayProps) CfnGateway {
	_init_.Initialize()

	j := jsiiProxy_CfnGateway{}

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnGateway",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTSiteWise::Gateway`.
func NewCfnGateway_Override(c CfnGateway, scope awscdk.Construct, id *string, props *CfnGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnGateway",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGateway) SetGatewayCapabilitySummaries(val interface{}) {
	_jsii_.Set(
		j,
		"gatewayCapabilitySummaries",
		val,
	)
}

func (j *jsiiProxy_CfnGateway) SetGatewayName(val *string) {
	_jsii_.Set(
		j,
		"gatewayName",
		val,
	)
}

func (j *jsiiProxy_CfnGateway) SetGatewayPlatform(val interface{}) {
	_jsii_.Set(
		j,
		"gatewayPlatform",
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
func CfnGateway_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnGateway",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnGateway_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnGateway",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGateway_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iotsitewise.CfnGateway",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGateway) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGateway) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGateway) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGateway) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGateway) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGateway) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGateway) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnGateway) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGateway) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGateway) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnGateway) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnGateway) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnGateway) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGateway) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGateway) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnGateway) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGateway) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGateway) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGateway) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGateway) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Contains a summary of a gateway capability configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayCapabilitySummaryProperty := &gatewayCapabilitySummaryProperty{
//   	capabilityNamespace: jsii.String("capabilityNamespace"),
//
//   	// the properties below are optional
//   	capabilityConfiguration: jsii.String("capabilityConfiguration"),
//   }
//
type CfnGateway_GatewayCapabilitySummaryProperty struct {
	// The namespace of the capability configuration.
	//
	// For example, if you configure OPC-UA sources from the AWS IoT SiteWise console, your OPC-UA capability configuration has the namespace `iotsitewise:opcuacollector:version` , where `version` is a number such as `1` .
	//
	// The maximum length is 512 characters with the pattern `^[a-zA-Z]+:[a-zA-Z]+:[0-9]+$` .
	CapabilityNamespace *string `field:"required" json:"capabilityNamespace" yaml:"capabilityNamespace"`
	// The JSON document that defines the configuration for the gateway capability.
	//
	// For more information, see [Configuring data sources (CLI)](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/configure-sources.html#configure-source-cli) in the *AWS IoT SiteWise User Guide* .
	CapabilityConfiguration *string `field:"optional" json:"capabilityConfiguration" yaml:"capabilityConfiguration"`
}

// Contains a gateway's platform information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayPlatformProperty := &gatewayPlatformProperty{
//   	greengrass: &greengrassProperty{
//   		groupArn: jsii.String("groupArn"),
//   	},
//   	greengrassV2: &greengrassV2Property{
//   		coreDeviceThingName: jsii.String("coreDeviceThingName"),
//   	},
//   }
//
type CfnGateway_GatewayPlatformProperty struct {
	// A gateway that runs on AWS IoT Greengrass .
	Greengrass interface{} `field:"optional" json:"greengrass" yaml:"greengrass"`
	// A gateway that runs on AWS IoT Greengrass V2.
	GreengrassV2 interface{} `field:"optional" json:"greengrassV2" yaml:"greengrassV2"`
}

// Contains details for a gateway that runs on AWS IoT Greengrass .
//
// To create a gateway that runs on AWS IoT Greengrass , you must add the IoT SiteWise connector to a Greengrass group and deploy it. Your Greengrass group must also have permissions to upload data to AWS IoT SiteWise . For more information, see [Ingesting data using a gateway](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/gateway-connector.html) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   greengrassProperty := &greengrassProperty{
//   	groupArn: jsii.String("groupArn"),
//   }
//
type CfnGateway_GreengrassProperty struct {
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Greengrass group. For more information about how to find a group's ARN, see [ListGroups](https://docs.aws.amazon.com/greengrass/latest/apireference/listgroups-get.html) and [GetGroup](https://docs.aws.amazon.com/greengrass/latest/apireference/getgroup-get.html) in the *AWS IoT Greengrass API Reference* .
	GroupArn *string `field:"required" json:"groupArn" yaml:"groupArn"`
}

// Contains details for a gateway that runs on AWS IoT Greengrass V2.
//
// To create a gateway that runs on AWS IoT Greengrass V2, you must deploy the IoT SiteWise Edge component to your gateway device. Your [Greengrass device role](https://docs.aws.amazon.com/greengrass/v2/developerguide/device-service-role.html) must use the `AWSIoTSiteWiseEdgeAccess` policy. For more information, see [Using AWS IoT SiteWise at the edge](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/sw-gateways.html) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   greengrassV2Property := &greengrassV2Property{
//   	coreDeviceThingName: jsii.String("coreDeviceThingName"),
//   }
//
type CfnGateway_GreengrassV2Property struct {
	// The name of the AWS IoT thing for your AWS IoT Greengrass V2 core device.
	CoreDeviceThingName *string `field:"required" json:"coreDeviceThingName" yaml:"coreDeviceThingName"`
}

// Properties for defining a `CfnGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayProps := &cfnGatewayProps{
//   	gatewayName: jsii.String("gatewayName"),
//   	gatewayPlatform: &gatewayPlatformProperty{
//   		greengrass: &greengrassProperty{
//   			groupArn: jsii.String("groupArn"),
//   		},
//   		greengrassV2: &greengrassV2Property{
//   			coreDeviceThingName: jsii.String("coreDeviceThingName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	gatewayCapabilitySummaries: []interface{}{
//   		&gatewayCapabilitySummaryProperty{
//   			capabilityNamespace: jsii.String("capabilityNamespace"),
//
//   			// the properties below are optional
//   			capabilityConfiguration: jsii.String("capabilityConfiguration"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnGatewayProps struct {
	// A unique, friendly name for the gateway.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	GatewayName *string `field:"required" json:"gatewayName" yaml:"gatewayName"`
	// The gateway's platform.
	//
	// You can only specify one platform in a gateway.
	GatewayPlatform interface{} `field:"required" json:"gatewayPlatform" yaml:"gatewayPlatform"`
	// A list of gateway capability summaries that each contain a namespace and status.
	//
	// Each gateway capability defines data sources for the gateway. To retrieve a capability configuration's definition, use [DescribeGatewayCapabilityConfiguration](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeGatewayCapabilityConfiguration.html) .
	GatewayCapabilitySummaries interface{} `field:"optional" json:"gatewayCapabilitySummaries" yaml:"gatewayCapabilitySummaries"`
	// A list of key-value pairs that contain metadata for the gateway.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTSiteWise::Portal`.
//
// Creates a portal, which can contain projects and dashboards. Before you can create a portal, you must enable AWS SSO . AWS IoT SiteWise Monitor uses AWS SSO to manage user permissions. For more information, see [Enabling AWS SSO](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-get-started.html#mon-gs-sso) in the *AWS IoT SiteWise User Guide* .
//
// > Before you can sign in to a new portal, you must add at least one AWS SSO user or group to that portal. For more information, see [Adding or removing portal administrators](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/administer-portals.html#portal-change-admins) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var alarms interface{}
//
//   cfnPortal := awscdk.Aws_iotsitewise.NewCfnPortal(this, jsii.String("MyCfnPortal"), &cfnPortalProps{
//   	portalContactEmail: jsii.String("portalContactEmail"),
//   	portalName: jsii.String("portalName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	alarms: alarms,
//   	notificationSenderEmail: jsii.String("notificationSenderEmail"),
//   	portalAuthMode: jsii.String("portalAuthMode"),
//   	portalDescription: jsii.String("portalDescription"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPortal interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal.
	//
	// You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range. For more information, see [Monitoring with alarms](https://docs.aws.amazon.com/iot-sitewise/latest/appguide/monitor-alarms.html) in the *AWS IoT SiteWise Application Guide* .
	Alarms() interface{}
	SetAlarms(val interface{})
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the portal, which has the following format.
	//
	// `arn:${Partition}:iotsitewise:${Region}:${Account}:portal/${PortalId}`.
	AttrPortalArn() *string
	// The AWS SSO application generated client ID (used with AWS SSO APIs).
	AttrPortalClientId() *string
	// The ID of the created portal.
	AttrPortalId() *string
	// The public URL for the AWS IoT SiteWise Monitor portal.
	AttrPortalStartUrl() *string
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The email address that sends alarm notifications.
	//
	// > If you use the [AWS IoT Events managed Lambda function](https://docs.aws.amazon.com/iotevents/latest/developerguide/lambda-support.html) to manage your emails, you must [verify the sender email address in Amazon SES](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-email-addresses.html) .
	NotificationSenderEmail() *string
	SetNotificationSenderEmail(val *string)
	// The service to use to authenticate users to the portal. Choose from the following options:.
	//
	// - `SSO` – The portal uses AWS Single Sign-On to authenticate users and manage user permissions. Before you can create a portal that uses AWS SSO , you must enable AWS SSO . For more information, see [Enabling AWS SSO](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-get-started.html#mon-gs-sso) in the *AWS IoT SiteWise User Guide* . This option is only available in AWS Regions other than the China Regions.
	// - `IAM` – The portal uses AWS Identity and Access Management ( IAM ) to authenticate users and manage user permissions.
	//
	// You can't change this value after you create a portal.
	//
	// Default: `SSO`.
	PortalAuthMode() *string
	SetPortalAuthMode(val *string)
	// The AWS administrator's contact email address.
	PortalContactEmail() *string
	SetPortalContactEmail(val *string)
	// A description for the portal.
	PortalDescription() *string
	SetPortalDescription(val *string)
	// A friendly name for the portal.
	PortalName() *string
	SetPortalName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of a service role that allows the portal's users to access your AWS IoT SiteWise resources on your behalf. For more information, see [Using service roles for AWS IoT SiteWise Monitor](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-service-role.html) in the *AWS IoT SiteWise User Guide* .
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of key-value pairs that contain metadata for the portal.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
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

// The jsii proxy struct for CfnPortal
type jsiiProxy_CfnPortal struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPortal) Alarms() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) AttrPortalArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPortalArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) AttrPortalClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPortalClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) AttrPortalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPortalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) AttrPortalStartUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPortalStartUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) NotificationSenderEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationSenderEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) PortalAuthMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portalAuthMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) PortalContactEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portalContactEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) PortalDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portalDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) PortalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortal) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTSiteWise::Portal`.
func NewCfnPortal(scope awscdk.Construct, id *string, props *CfnPortalProps) CfnPortal {
	_init_.Initialize()

	j := jsiiProxy_CfnPortal{}

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnPortal",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTSiteWise::Portal`.
func NewCfnPortal_Override(c CfnPortal, scope awscdk.Construct, id *string, props *CfnPortalProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnPortal",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPortal) SetAlarms(val interface{}) {
	_jsii_.Set(
		j,
		"alarms",
		val,
	)
}

func (j *jsiiProxy_CfnPortal) SetNotificationSenderEmail(val *string) {
	_jsii_.Set(
		j,
		"notificationSenderEmail",
		val,
	)
}

func (j *jsiiProxy_CfnPortal) SetPortalAuthMode(val *string) {
	_jsii_.Set(
		j,
		"portalAuthMode",
		val,
	)
}

func (j *jsiiProxy_CfnPortal) SetPortalContactEmail(val *string) {
	_jsii_.Set(
		j,
		"portalContactEmail",
		val,
	)
}

func (j *jsiiProxy_CfnPortal) SetPortalDescription(val *string) {
	_jsii_.Set(
		j,
		"portalDescription",
		val,
	)
}

func (j *jsiiProxy_CfnPortal) SetPortalName(val *string) {
	_jsii_.Set(
		j,
		"portalName",
		val,
	)
}

func (j *jsiiProxy_CfnPortal) SetRoleArn(val *string) {
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
// Experimental.
func CfnPortal_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnPortal",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPortal_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnPortal",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPortal_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnPortal",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPortal_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iotsitewise.CfnPortal",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPortal) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPortal) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPortal) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPortal) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPortal) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPortal) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPortal) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPortal) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortal) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortal) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPortal) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPortal) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPortal) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortal) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPortal) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPortal) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortal) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortal) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPortal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortal) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortal) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPortal`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var alarms interface{}
//
//   cfnPortalProps := &cfnPortalProps{
//   	portalContactEmail: jsii.String("portalContactEmail"),
//   	portalName: jsii.String("portalName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	alarms: alarms,
//   	notificationSenderEmail: jsii.String("notificationSenderEmail"),
//   	portalAuthMode: jsii.String("portalAuthMode"),
//   	portalDescription: jsii.String("portalDescription"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPortalProps struct {
	// The AWS administrator's contact email address.
	PortalContactEmail *string `field:"required" json:"portalContactEmail" yaml:"portalContactEmail"`
	// A friendly name for the portal.
	PortalName *string `field:"required" json:"portalName" yaml:"portalName"`
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of a service role that allows the portal's users to access your AWS IoT SiteWise resources on your behalf. For more information, see [Using service roles for AWS IoT SiteWise Monitor](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-service-role.html) in the *AWS IoT SiteWise User Guide* .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal.
	//
	// You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range. For more information, see [Monitoring with alarms](https://docs.aws.amazon.com/iot-sitewise/latest/appguide/monitor-alarms.html) in the *AWS IoT SiteWise Application Guide* .
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// The email address that sends alarm notifications.
	//
	// > If you use the [AWS IoT Events managed Lambda function](https://docs.aws.amazon.com/iotevents/latest/developerguide/lambda-support.html) to manage your emails, you must [verify the sender email address in Amazon SES](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-email-addresses.html) .
	NotificationSenderEmail *string `field:"optional" json:"notificationSenderEmail" yaml:"notificationSenderEmail"`
	// The service to use to authenticate users to the portal. Choose from the following options:.
	//
	// - `SSO` – The portal uses AWS Single Sign-On to authenticate users and manage user permissions. Before you can create a portal that uses AWS SSO , you must enable AWS SSO . For more information, see [Enabling AWS SSO](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-get-started.html#mon-gs-sso) in the *AWS IoT SiteWise User Guide* . This option is only available in AWS Regions other than the China Regions.
	// - `IAM` – The portal uses AWS Identity and Access Management ( IAM ) to authenticate users and manage user permissions.
	//
	// You can't change this value after you create a portal.
	//
	// Default: `SSO`.
	PortalAuthMode *string `field:"optional" json:"portalAuthMode" yaml:"portalAuthMode"`
	// A description for the portal.
	PortalDescription *string `field:"optional" json:"portalDescription" yaml:"portalDescription"`
	// A list of key-value pairs that contain metadata for the portal.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTSiteWise::Project`.
//
// Creates a project in the specified portal.
//
// > Make sure that the project name and description don't contain confidential information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProject := awscdk.Aws_iotsitewise.NewCfnProject(this, jsii.String("MyCfnProject"), &cfnProjectProps{
//   	portalId: jsii.String("portalId"),
//   	projectName: jsii.String("projectName"),
//
//   	// the properties below are optional
//   	assetIds: []*string{
//   		jsii.String("assetIds"),
//   	},
//   	projectDescription: jsii.String("projectDescription"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnProject interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A list that contains the IDs of each asset associated with the project.
	AssetIds() *[]*string
	SetAssetIds(val *[]*string)
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the project, which has the following format.
	//
	// `arn:${Partition}:iotsitewise:${Region}:${Account}:project/${ProjectId}`.
	AttrProjectArn() *string
	// The ID of the project.
	AttrProjectId() *string
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The ID of the portal in which to create the project.
	PortalId() *string
	SetPortalId(val *string)
	// A description for the project.
	ProjectDescription() *string
	SetProjectDescription(val *string)
	// A friendly name for the project.
	ProjectName() *string
	SetProjectName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of key-value pairs that contain metadata for the project.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
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

// The jsii proxy struct for CfnProject
type jsiiProxy_CfnProject struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnProject) AssetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) AttrProjectArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProjectArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) AttrProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) PortalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) ProjectDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTSiteWise::Project`.
func NewCfnProject(scope awscdk.Construct, id *string, props *CfnProjectProps) CfnProject {
	_init_.Initialize()

	j := jsiiProxy_CfnProject{}

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnProject",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTSiteWise::Project`.
func NewCfnProject_Override(c CfnProject, scope awscdk.Construct, id *string, props *CfnProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotsitewise.CfnProject",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnProject) SetAssetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"assetIds",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetPortalId(val *string) {
	_jsii_.Set(
		j,
		"portalId",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetProjectDescription(val *string) {
	_jsii_.Set(
		j,
		"projectDescription",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetProjectName(val *string) {
	_jsii_.Set(
		j,
		"projectName",
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
func CfnProject_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnProject",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnProject_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnProject",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotsitewise.CfnProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProject_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iotsitewise.CfnProject",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProject) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnProject) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnProject) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnProject) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnProject) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnProject) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnProject) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnProject) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnProject) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnProject) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnProject) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnProject) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnProject) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProps := &cfnProjectProps{
//   	portalId: jsii.String("portalId"),
//   	projectName: jsii.String("projectName"),
//
//   	// the properties below are optional
//   	assetIds: []*string{
//   		jsii.String("assetIds"),
//   	},
//   	projectDescription: jsii.String("projectDescription"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProjectProps struct {
	// The ID of the portal in which to create the project.
	PortalId *string `field:"required" json:"portalId" yaml:"portalId"`
	// A friendly name for the project.
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// A list that contains the IDs of each asset associated with the project.
	AssetIds *[]*string `field:"optional" json:"assetIds" yaml:"assetIds"`
	// A description for the project.
	ProjectDescription *string `field:"optional" json:"projectDescription" yaml:"projectDescription"`
	// A list of key-value pairs that contain metadata for the project.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

