package awsamplify

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplify/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies the AWS::Amplify::Domain resource that enables you to connect a custom domain to your app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomain := awscdk.Aws_amplify.NewCfnDomain(this, jsii.String("MyCfnDomain"), &CfnDomainProps{
//   	AppId: jsii.String("appId"),
//   	DomainName: jsii.String("domainName"),
//   	SubDomainSettings: []interface{}{
//   		&SubDomainSettingProperty{
//   			BranchName: jsii.String("branchName"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AutoSubDomainCreationPatterns: []*string{
//   		jsii.String("autoSubDomainCreationPatterns"),
//   	},
//   	AutoSubDomainIamRole: jsii.String("autoSubDomainIamRole"),
//   	CertificateSettings: &CertificateSettingsProperty{
//   		CertificateType: jsii.String("certificateType"),
//   		CustomCertificateArn: jsii.String("customCertificateArn"),
//   	},
//   	EnableAutoSubDomain: jsii.Boolean(false),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html
//
type CfnDomain interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique ID for an Amplify app.
	AppId() *string
	SetAppId(val *string)
	// ARN for the Domain Association.
	AttrArn() *string
	// Branch patterns for the automatically created subdomain.
	AttrAutoSubDomainCreationPatterns() *[]*string
	// The IAM service role for the subdomain.
	AttrAutoSubDomainIamRole() *string
	AttrCertificate() awscdk.IResolvable
	// DNS Record for certificate verification.
	AttrCertificateRecord() *string
	// Name of the domain.
	AttrDomainName() *string
	// Status for the Domain Association.
	AttrDomainStatus() *string
	// Specifies whether the automated creation of subdomains for branches is enabled.
	AttrEnableAutoSubDomain() awscdk.IResolvable
	// Reason for the current status of the domain.
	AttrStatusReason() *string
	// The status of the domain update operation that is currently in progress.
	//
	// The following list describes the valid update states.
	//
	// - **REQUESTING_CERTIFICATE** - The certificate is in the process of being updated.
	// - **PENDING_VERIFICATION** - Indicates that an Amplify managed certificate is in the process of being verified. This occurs during the creation of a custom domain or when a custom domain is updated to use a managed certificate.
	// - **IMPORTING_CUSTOM_CERTIFICATE** - Indicates that an Amplify custom certificate is in the process of being imported. This occurs during the creation of a custom domain or when a custom domain is updated to use a custom certificate.
	// - **PENDING_DEPLOYMENT** - Indicates that the subdomain or certificate changes are being propagated.
	// - **AWAITING_APP_CNAME** - Amplify is waiting for CNAME records corresponding to subdomains to be propagated. If your custom domain is on Route 53, Amplify handles this for you automatically. For more information about custom domains, see [Setting up custom domains](https://docs.aws.amazon.com/amplify/latest/userguide/custom-domains.html) in the *Amplify Hosting User Guide* .
	// - **UPDATE_COMPLETE** - The certificate has been associated with a domain.
	// - **UPDATE_FAILED** - The certificate has failed to be provisioned or associated, and there is no existing active certificate to roll back to.
	AttrUpdateStatus() *string
	// Sets the branch patterns for automatic subdomain creation.
	AutoSubDomainCreationPatterns() *[]*string
	SetAutoSubDomainCreationPatterns(val *[]*string)
	// The required AWS Identity and Access Management (IAMlong) service role for the Amazon Resource Name (ARN) for automatically creating subdomains.
	AutoSubDomainIamRole() *string
	SetAutoSubDomainIamRole(val *string)
	// The type of SSL/TLS certificate to use for your custom domain.
	CertificateSettings() interface{}
	SetCertificateSettings(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The domain name for the domain association.
	DomainName() *string
	SetDomainName(val *string)
	// Enables the automated creation of subdomains for branches.
	EnableAutoSubDomain() interface{}
	SetEnableAutoSubDomain(val interface{})
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
	// The setting for the subdomain.
	SubDomainSettings() interface{}
	SetSubDomainSettings(val interface{})
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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

// The jsii proxy struct for CfnDomain
type jsiiProxy_CfnDomain struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDomain) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrAutoSubDomainCreationPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrAutoSubDomainCreationPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrAutoSubDomainIamRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAutoSubDomainIamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrCertificate() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrCertificateRecord() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCertificateRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrDomainStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrEnableAutoSubDomain() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrEnableAutoSubDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrStatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrUpdateStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AutoSubDomainCreationPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoSubDomainCreationPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AutoSubDomainIamRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoSubDomainIamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CertificateSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) EnableAutoSubDomain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoSubDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) SubDomainSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subDomainSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnDomain(scope constructs.Construct, id *string, props *CfnDomainProps) CfnDomain {
	_init_.Initialize()

	if err := validateNewCfnDomainParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDomain{}

	_jsii_.Create(
		"aws-cdk-lib.aws_amplify.CfnDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnDomain_Override(c CfnDomain, scope constructs.Construct, id *string, props *CfnDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_amplify.CfnDomain",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomain)SetAppId(val *string) {
	if err := j.validateSetAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetAutoSubDomainCreationPatterns(val *[]*string) {
	_jsii_.Set(
		j,
		"autoSubDomainCreationPatterns",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetAutoSubDomainIamRole(val *string) {
	_jsii_.Set(
		j,
		"autoSubDomainIamRole",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetCertificateSettings(val interface{}) {
	if err := j.validateSetCertificateSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateSettings",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetEnableAutoSubDomain(val interface{}) {
	if err := j.validateSetEnableAutoSubDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutoSubDomain",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetSubDomainSettings(val interface{}) {
	if err := j.validateSetSubDomainSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subDomainSettings",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDomain_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomain_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplify.CfnDomain",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDomain_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomain_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplify.CfnDomain",
		"isCfnResource",
		[]interface{}{x},
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
func CfnDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplify.CfnDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomain_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_amplify.CfnDomain",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomain) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDomain) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomain) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomain) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDomain) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDomain) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDomain) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDomain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDomain) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnDomain) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDomain) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDomain) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomain) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnDomain) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnDomain) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

