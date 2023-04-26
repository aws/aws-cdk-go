package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsses/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::SES::EmailIdentity`.
//
// Specifies an identity for using within SES. An identity is an email address or domain that you use when you send email. Before you can use an identity to send email, you first have to verify it. By verifying an identity, you demonstrate that you're the owner of the identity, and that you've given Amazon SES API v2 permission to send email from the identity.
//
// When you verify an email address, SES sends an email to the address. Your email address is verified as soon as you follow the link in the verification email. When you verify a domain without specifying the DkimSigningAttributes properties, OR only the NextSigningKeyLength property of DkimSigningAttributes, this resource provides a set of CNAME token names and values (DkimDNSTokenName1, DkimDNSTokenValue1, DkimDNSTokenName2, DkimDNSTokenValue2, DkimDNSTokenName3, DkimDNSTokenValue3) as outputs. You can then add these to the DNS configuration for your domain. Your domain is verified when Amazon SES detects these records in the DNS configuration for your domain. This verification method is known as Easy DKIM.
//
// Alternatively, you can perform the verification process by providing your own public-private key pair. This verification method is known as Bring Your Own DKIM (BYODKIM). To use BYODKIM, your resource must include DkimSigningAttributes properties DomainSigningSelector and DomainSigningPrivateKey. When you specify this object, you provide a selector (DomainSigningSelector) (a component of the DNS record name that identifies the public key to use for DKIM authentication) and a private key (DomainSigningPrivateKey).
//
// Additionally, you can associate an existing configuration set with the email identity that you're verifying.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEmailIdentity := awscdk.Aws_ses.NewCfnEmailIdentity(this, jsii.String("MyCfnEmailIdentity"), &CfnEmailIdentityProps{
//   	EmailIdentity: jsii.String("emailIdentity"),
//
//   	// the properties below are optional
//   	ConfigurationSetAttributes: &ConfigurationSetAttributesProperty{
//   		ConfigurationSetName: jsii.String("configurationSetName"),
//   	},
//   	DkimAttributes: &DkimAttributesProperty{
//   		SigningEnabled: jsii.Boolean(false),
//   	},
//   	DkimSigningAttributes: &DkimSigningAttributesProperty{
//   		DomainSigningPrivateKey: jsii.String("domainSigningPrivateKey"),
//   		DomainSigningSelector: jsii.String("domainSigningSelector"),
//   		NextSigningKeyLength: jsii.String("nextSigningKeyLength"),
//   	},
//   	FeedbackAttributes: &FeedbackAttributesProperty{
//   		EmailForwardingEnabled: jsii.Boolean(false),
//   	},
//   	MailFromAttributes: &MailFromAttributesProperty{
//   		BehaviorOnMxFailure: jsii.String("behaviorOnMxFailure"),
//   		MailFromDomain: jsii.String("mailFromDomain"),
//   	},
//   })
//
type CfnEmailIdentity interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The host name for the first token that you have to add to the DNS configuration for your domain.
	AttrDkimDnsTokenName1() *string
	// The host name for the second token that you have to add to the DNS configuration for your domain.
	AttrDkimDnsTokenName2() *string
	// The host name for the third token that you have to add to the DNS configuration for your domain.
	AttrDkimDnsTokenName3() *string
	// The record value for the first token that you have to add to the DNS configuration for your domain.
	AttrDkimDnsTokenValue1() *string
	// The record value for the second token that you have to add to the DNS configuration for your domain.
	AttrDkimDnsTokenValue2() *string
	// The record value for the third token that you have to add to the DNS configuration for your domain.
	AttrDkimDnsTokenValue3() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Used to associate a configuration set with an email identity.
	ConfigurationSetAttributes() interface{}
	SetConfigurationSetAttributes(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// An object that contains information about the DKIM attributes for the identity.
	DkimAttributes() interface{}
	SetDkimAttributes(val interface{})
	// If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for [Easy DKIM](https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-easy.html) .
	DkimSigningAttributes() interface{}
	SetDkimSigningAttributes(val interface{})
	// The email address or domain to verify.
	EmailIdentity() *string
	SetEmailIdentity(val *string)
	// Used to enable or disable feedback forwarding for an identity.
	FeedbackAttributes() interface{}
	SetFeedbackAttributes(val interface{})
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
	// Used to enable or disable the custom Mail-From domain configuration for an email identity.
	MailFromAttributes() interface{}
	SetMailFromAttributes(val interface{})
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

// The jsii proxy struct for CfnEmailIdentity
type jsiiProxy_CfnEmailIdentity struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEmailIdentity) AttrDkimDnsTokenName1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDkimDnsTokenName1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) AttrDkimDnsTokenName2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDkimDnsTokenName2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) AttrDkimDnsTokenName3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDkimDnsTokenName3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) AttrDkimDnsTokenValue1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDkimDnsTokenValue1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) AttrDkimDnsTokenValue2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDkimDnsTokenValue2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) AttrDkimDnsTokenValue3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDkimDnsTokenValue3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) ConfigurationSetAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationSetAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) DkimAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dkimAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) DkimSigningAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dkimSigningAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) EmailIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) FeedbackAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"feedbackAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) MailFromAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mailFromAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SES::EmailIdentity`.
func NewCfnEmailIdentity(scope awscdk.Construct, id *string, props *CfnEmailIdentityProps) CfnEmailIdentity {
	_init_.Initialize()

	if err := validateNewCfnEmailIdentityParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEmailIdentity{}

	_jsii_.Create(
		"monocdk.aws_ses.CfnEmailIdentity",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SES::EmailIdentity`.
func NewCfnEmailIdentity_Override(c CfnEmailIdentity, scope awscdk.Construct, id *string, props *CfnEmailIdentityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ses.CfnEmailIdentity",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEmailIdentity)SetConfigurationSetAttributes(val interface{}) {
	if err := j.validateSetConfigurationSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationSetAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnEmailIdentity)SetDkimAttributes(val interface{}) {
	if err := j.validateSetDkimAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dkimAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnEmailIdentity)SetDkimSigningAttributes(val interface{}) {
	if err := j.validateSetDkimSigningAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dkimSigningAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnEmailIdentity)SetEmailIdentity(val *string) {
	if err := j.validateSetEmailIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailIdentity",
		val,
	)
}

func (j *jsiiProxy_CfnEmailIdentity)SetFeedbackAttributes(val interface{}) {
	if err := j.validateSetFeedbackAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"feedbackAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnEmailIdentity)SetMailFromAttributes(val interface{}) {
	if err := j.validateSetMailFromAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailFromAttributes",
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
func CfnEmailIdentity_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEmailIdentity_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ses.CfnEmailIdentity",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnEmailIdentity_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnEmailIdentity_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ses.CfnEmailIdentity",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnEmailIdentity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEmailIdentity_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ses.CfnEmailIdentity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEmailIdentity_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ses.CfnEmailIdentity",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEmailIdentity) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnEmailIdentity) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnEmailIdentity) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEmailIdentity) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailIdentity) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEmailIdentity) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnEmailIdentity) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailIdentity) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailIdentity) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailIdentity) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

