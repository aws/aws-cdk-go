package awscertificatemanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
)

// A certificate managed by AWS Certificate Manager.
//
// TODO: EXAMPLE
//
type Certificate interface {
	awscdk.Resource
	ICertificate
	CertificateArn() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Region() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	MetricDaysToExpiry(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for Certificate
type jsiiProxy_Certificate struct {
	internal.Type__awscdkResource
	jsiiProxy_ICertificate
}

func (j *jsiiProxy_Certificate) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCertificate(scope constructs.Construct, id *string, props *CertificateProps) Certificate {
	_init_.Initialize()

	j := jsiiProxy_Certificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_certificatemanager.Certificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCertificate_Override(c Certificate, scope constructs.Construct, id *string, props *CertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_certificatemanager.Certificate",
		[]interface{}{scope, id, props},
		c,
	)
}

// Import a certificate.
func Certificate_FromCertificateArn(scope constructs.Construct, id *string, certificateArn *string) ICertificate {
	_init_.Initialize()

	var returns ICertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.Certificate",
		"fromCertificateArn",
		[]interface{}{scope, id, certificateArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Certificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.Certificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Certificate_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.Certificate",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_Certificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_Certificate) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (c *jsiiProxy_Certificate) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (c *jsiiProxy_Certificate) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return the DaysToExpiry metric for this AWS Certificate Manager Certificate. By default, this is the minimum value over 1 day.
//
// This metric is no longer emitted once the certificate has effectively
// expired, so alarms configured on this metric should probably treat missing
// data as "breaching".
func (c *jsiiProxy_Certificate) MetricDaysToExpiry(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricDaysToExpiry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_Certificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for your certificate.
//
// TODO: EXAMPLE
//
type CertificateProps struct {
	// Fully-qualified domain name to request a certificate for.
	//
	// May contain wildcards, such as ``*.domain.com``.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Alternative domain names on your certificate.
	//
	// Use this to register alternative domain names that represent the same site.
	SubjectAlternativeNames *[]*string `json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// How to validate this certificate.
	Validation CertificateValidation `json:"validation" yaml:"validation"`
}

// How to validate a certificate.
//
// TODO: EXAMPLE
//
type CertificateValidation interface {
	Method() ValidationMethod
	Props() *CertificationValidationProps
}

// The jsii proxy struct for CertificateValidation
type jsiiProxy_CertificateValidation struct {
	_ byte // padding
}

func (j *jsiiProxy_CertificateValidation) Method() ValidationMethod {
	var returns ValidationMethod
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateValidation) Props() *CertificationValidationProps {
	var returns *CertificationValidationProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Validate the certificate with DNS.
//
// IMPORTANT: If `hostedZone` is not specified, DNS records must be added
// manually and the stack will not complete creating until the records are
// added.
func CertificateValidation_FromDns(hostedZone awsroute53.IHostedZone) CertificateValidation {
	_init_.Initialize()

	var returns CertificateValidation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.CertificateValidation",
		"fromDns",
		[]interface{}{hostedZone},
		&returns,
	)

	return returns
}

// Validate the certificate with automatically created DNS records in multiple Amazon Route 53 hosted zones.
func CertificateValidation_FromDnsMultiZone(hostedZones *map[string]awsroute53.IHostedZone) CertificateValidation {
	_init_.Initialize()

	var returns CertificateValidation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.CertificateValidation",
		"fromDnsMultiZone",
		[]interface{}{hostedZones},
		&returns,
	)

	return returns
}

// Validate the certificate with Email.
//
// IMPORTANT: if you are creating a certificate as part of your stack, the stack
// will not complete creating until you read and follow the instructions in the
// email that you will receive.
//
// ACM will send validation emails to the following addresses:
//
//   admin@domain.com
//   administrator@domain.com
//   hostmaster@domain.com
//   postmaster@domain.com
//   webmaster@domain.com
//
// For every domain that you register.
func CertificateValidation_FromEmail(validationDomains *map[string]*string) CertificateValidation {
	_init_.Initialize()

	var returns CertificateValidation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.CertificateValidation",
		"fromEmail",
		[]interface{}{validationDomains},
		&returns,
	)

	return returns
}

// Properties for certificate validation.
//
// TODO: EXAMPLE
//
type CertificationValidationProps struct {
	// Hosted zone to use for DNS validation.
	HostedZone awsroute53.IHostedZone `json:"hostedZone" yaml:"hostedZone"`
	// A map of hosted zones to use for DNS validation.
	HostedZones *map[string]awsroute53.IHostedZone `json:"hostedZones" yaml:"hostedZones"`
	// Validation method.
	Method ValidationMethod `json:"method" yaml:"method"`
	// Validation domains to use for email validation.
	ValidationDomains *map[string]*string `json:"validationDomains" yaml:"validationDomains"`
}

// A CloudFormation `AWS::CertificateManager::Account`.
//
// The `AWS::CertificateManager::Account` resource defines the expiry event configuration that determines the number of days prior to expiry when ACM starts generating EventBridge events.
//
// TODO: EXAMPLE
//
type CfnAccount interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAccountId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	ExpiryEventsConfiguration() interface{}
	SetExpiryEventsConfiguration(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnAccount
type jsiiProxy_CfnAccount struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAccount) AttrAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) ExpiryEventsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expiryEventsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CertificateManager::Account`.
func NewCfnAccount(scope constructs.Construct, id *string, props *CfnAccountProps) CfnAccount {
	_init_.Initialize()

	j := jsiiProxy_CfnAccount{}

	_jsii_.Create(
		"aws-cdk-lib.aws_certificatemanager.CfnAccount",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CertificateManager::Account`.
func NewCfnAccount_Override(c CfnAccount, scope constructs.Construct, id *string, props *CfnAccountProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_certificatemanager.CfnAccount",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAccount) SetExpiryEventsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"expiryEventsConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAccount_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.CfnAccount",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAccount_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.CfnAccount",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.CfnAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccount_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_certificatemanager.CfnAccount",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnAccount) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnAccount) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnAccount) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnAccount) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnAccount) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnAccount) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnAccount) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnAccount) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnAccount) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnAccount) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnAccount) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAccount) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnAccount) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccount) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Object containing expiration events options associated with an AWS account .
//
// For more information, see [ExpiryEventsConfiguration](https://docs.aws.amazon.com/acm/latest/APIReference/API_ExpiryEventsConfiguration.html) in the API reference.
//
// TODO: EXAMPLE
//
type CfnAccount_ExpiryEventsConfigurationProperty struct {
	// This option specifies the number of days prior to certificate expiration when ACM starts generating `EventBridge` events.
	//
	// ACM sends one event per day per certificate until the certificate expires. By default, accounts receive events starting 45 days before certificate expiration.
	DaysBeforeExpiry *float64 `json:"daysBeforeExpiry" yaml:"daysBeforeExpiry"`
}

// Properties for defining a `CfnAccount`.
//
// TODO: EXAMPLE
//
type CfnAccountProps struct {
	// Object containing expiration events options associated with an AWS account .
	//
	// For more information, see [ExpiryEventsConfiguration](https://docs.aws.amazon.com/acm/latest/APIReference/API_ExpiryEventsConfiguration.html) in the API reference.
	ExpiryEventsConfiguration interface{} `json:"expiryEventsConfiguration" yaml:"expiryEventsConfiguration"`
}

// A CloudFormation `AWS::CertificateManager::Certificate`.
//
// The `AWS::CertificateManager::Certificate` resource requests an AWS Certificate Manager ( ACM ) certificate that you can use to enable secure connections. For example, you can deploy an ACM certificate to an Elastic Load Balancer to enable HTTPS support. For more information, see [RequestCertificate](https://docs.aws.amazon.com/acm/latest/APIReference/API_RequestCertificate.html) in the AWS Certificate Manager API Reference.
//
// > When you use the `AWS::CertificateManager::Certificate` resource in a CloudFormation stack, domain validation is handled automatically if all three of the following are true: The certificate domain is hosted in Amazon Route 53, the domain resides in your AWS account , and you are using DNS validation.
// >
// > However, if the certificate uses email validation, or if the domain is not hosted in Route 53, then the stack will remain in the `CREATE_IN_PROGRESS` state. Further stack operations are delayed until you validate the certificate request, either by acting upon the instructions in the validation email, or by adding a CNAME record to your DNS configuration. For more information, see [Option 1: DNS Validation](https://docs.aws.amazon.com/acm/latest/userguide/dns-validation.html) and [Option 2: Email Validation](https://docs.aws.amazon.com/acm/latest/userguide/email-validation.html) .
//
// TODO: EXAMPLE
//
type CfnCertificate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CertificateAuthorityArn() *string
	SetCertificateAuthorityArn(val *string)
	CertificateTransparencyLoggingPreference() *string
	SetCertificateTransparencyLoggingPreference(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DomainName() *string
	SetDomainName(val *string)
	DomainValidationOptions() interface{}
	SetDomainValidationOptions(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	SubjectAlternativeNames() *[]*string
	SetSubjectAlternativeNames(val *[]*string)
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	ValidationMethod() *string
	SetValidationMethod(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCertificate
type jsiiProxy_CfnCertificate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCertificate) CertificateAuthorityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CertificateTransparencyLoggingPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateTransparencyLoggingPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) DomainValidationOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domainValidationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) SubjectAlternativeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) ValidationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationMethod",
		&returns,
	)
	return returns
}


// Create a new `AWS::CertificateManager::Certificate`.
func NewCfnCertificate(scope constructs.Construct, id *string, props *CfnCertificateProps) CfnCertificate {
	_init_.Initialize()

	j := jsiiProxy_CfnCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_certificatemanager.CfnCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CertificateManager::Certificate`.
func NewCfnCertificate_Override(c CfnCertificate, scope constructs.Construct, id *string, props *CfnCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_certificatemanager.CfnCertificate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCertificate) SetCertificateAuthorityArn(val *string) {
	_jsii_.Set(
		j,
		"certificateAuthorityArn",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetCertificateTransparencyLoggingPreference(val *string) {
	_jsii_.Set(
		j,
		"certificateTransparencyLoggingPreference",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetDomainValidationOptions(val interface{}) {
	_jsii_.Set(
		j,
		"domainValidationOptions",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetSubjectAlternativeNames(val *[]*string) {
	_jsii_.Set(
		j,
		"subjectAlternativeNames",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetValidationMethod(val *string) {
	_jsii_.Set(
		j,
		"validationMethod",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCertificate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.CfnCertificate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCertificate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.CfnCertificate",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.CfnCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCertificate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_certificatemanager.CfnCertificate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnCertificate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnCertificate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnCertificate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnCertificate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnCertificate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnCertificate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnCertificate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnCertificate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCertificate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnCertificate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `DomainValidationOption` is a property of the [AWS::CertificateManager::Certificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html) resource that specifies the AWS Certificate Manager ( ACM ) certificate domain to validate. Depending on the chosen validation method, ACM checks the domain's DNS record for a validation CNAME, or it attempts to send a validation email message to the domain owner.
//
// TODO: EXAMPLE
//
type CfnCertificate_DomainValidationOptionProperty struct {
	// A fully qualified domain name (FQDN) in the certificate request.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// The `HostedZoneId` option, which is available if you are using Route 53 as your domain registrar, causes ACM to add your CNAME to the domain record.
	//
	// Your list of `DomainValidationOptions` must contain one and only one of the domain-validation options, and the `HostedZoneId` can be used only when `DNS` is specified as your validation method.
	//
	// Use the Route 53 `ListHostedZones` API to discover IDs for available hosted zones.
	//
	// This option is required for publicly trusted certificates.
	//
	// > The `ListHostedZones` API returns IDs in the format "/hostedzone/Z111111QQQQQQQ", but CloudFormation requires the IDs to be in the format "Z111111QQQQQQQ".
	//
	// When you change your `DomainValidationOptions` , a new resource is created.
	HostedZoneId *string `json:"hostedZoneId" yaml:"hostedZoneId"`
	// The domain name to which you want ACM to send validation emails.
	//
	// This domain name is the suffix of the email addresses that you want ACM to use. This must be the same as the `DomainName` value or a superdomain of the `DomainName` value. For example, if you request a certificate for `testing.example.com` , you can specify `example.com` as this value. In that case, ACM sends domain validation emails to the following five addresses:
	//
	// - admin@example.com
	// - administrator@example.com
	// - hostmaster@example.com
	// - postmaster@example.com
	// - webmaster@example.com
	ValidationDomain *string `json:"validationDomain" yaml:"validationDomain"`
}

// Properties for defining a `CfnCertificate`.
//
// TODO: EXAMPLE
//
type CfnCertificateProps struct {
	// The fully qualified domain name (FQDN), such as www.example.com, with which you want to secure an ACM certificate. Use an asterisk (*) to create a wildcard certificate that protects several sites in the same domain. For example, `*.example.com` protects `www.example.com` , `site.example.com` , and `images.example.com.`.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// The Amazon Resource Name (ARN) of the private certificate authority (CA) that will be used to issue the certificate.
	//
	// If you do not provide an ARN and you are trying to request a private certificate, ACM will attempt to issue a public certificate. For more information about private CAs, see the [AWS Certificate Manager Private Certificate Authority (PCA)](https://docs.aws.amazon.com/acm-pca/latest/userguide/PcaWelcome.html) user guide. The ARN must have the following form:
	//
	// `arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012`
	CertificateAuthorityArn *string `json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// You can opt out of certificate transparency logging by specifying the `DISABLED` option. Opt in by specifying `ENABLED` .
	//
	// If you do not specify a certificate transparency logging preference on a new CloudFormation template, or if you remove the logging preference from an existing template, this is the same as explicitly enabling the preference.
	//
	// Changing the certificate transparency logging preference will update the existing resource by calling `UpdateCertificateOptions` on the certificate. This action will not create a new resource.
	CertificateTransparencyLoggingPreference *string `json:"certificateTransparencyLoggingPreference" yaml:"certificateTransparencyLoggingPreference"`
	// Domain information that domain name registrars use to verify your identity.
	//
	// > In order for a AWS::CertificateManager::Certificate to be provisioned and validated in CloudFormation automatically, the `DomainName` property needs to be identical to one of the `DomainName` property supplied in DomainValidationOptions, if the ValidationMethod is **DNS**. Failing to keep them like-for-like will result in failure to create the domain validation records in Route53.
	DomainValidationOptions interface{} `json:"domainValidationOptions" yaml:"domainValidationOptions"`
	// Additional FQDNs to be included in the Subject Alternative Name extension of the ACM certificate.
	//
	// For example, you can add www.example.net to a certificate for which the `DomainName` field is www.example.com if users can reach your site by using either name.
	SubjectAlternativeNames *[]*string `json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// Key-value pairs that can identify the certificate.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The method you want to use to validate that you own or control the domain associated with a public certificate.
	//
	// You can [validate with DNS](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html) or [validate with email](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-email.html) . We recommend that you use DNS validation.
	//
	// If not specified, this property defaults to email validation.
	ValidationMethod *string `json:"validationMethod" yaml:"validationMethod"`
}

// A certificate managed by AWS Certificate Manager.
//
// Will be automatically
// validated using DNS validation against the specified Route 53 hosted zone.
//
// TODO: EXAMPLE
//
type DnsValidatedCertificate interface {
	awscdk.Resource
	ICertificate
	awscdk.ITaggable
	CertificateArn() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Region() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	MetricDaysToExpiry(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for DnsValidatedCertificate
type jsiiProxy_DnsValidatedCertificate struct {
	internal.Type__awscdkResource
	jsiiProxy_ICertificate
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_DnsValidatedCertificate) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsValidatedCertificate) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsValidatedCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsValidatedCertificate) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsValidatedCertificate) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsValidatedCertificate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsValidatedCertificate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}


func NewDnsValidatedCertificate(scope constructs.Construct, id *string, props *DnsValidatedCertificateProps) DnsValidatedCertificate {
	_init_.Initialize()

	j := jsiiProxy_DnsValidatedCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_certificatemanager.DnsValidatedCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDnsValidatedCertificate_Override(d DnsValidatedCertificate, scope constructs.Construct, id *string, props *DnsValidatedCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_certificatemanager.DnsValidatedCertificate",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DnsValidatedCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.DnsValidatedCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DnsValidatedCertificate_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.DnsValidatedCertificate",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (d *jsiiProxy_DnsValidatedCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DnsValidatedCertificate) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (d *jsiiProxy_DnsValidatedCertificate) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (d *jsiiProxy_DnsValidatedCertificate) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return the DaysToExpiry metric for this AWS Certificate Manager Certificate. By default, this is the minimum value over 1 day.
//
// This metric is no longer emitted once the certificate has effectively
// expired, so alarms configured on this metric should probably treat missing
// data as "breaching".
func (d *jsiiProxy_DnsValidatedCertificate) MetricDaysToExpiry(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDaysToExpiry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DnsValidatedCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to create a DNS validated certificate managed by AWS Certificate Manager.
//
// TODO: EXAMPLE
//
type DnsValidatedCertificateProps struct {
	// Fully-qualified domain name to request a certificate for.
	//
	// May contain wildcards, such as ``*.domain.com``.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Alternative domain names on your certificate.
	//
	// Use this to register alternative domain names that represent the same site.
	SubjectAlternativeNames *[]*string `json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// How to validate this certificate.
	Validation CertificateValidation `json:"validation" yaml:"validation"`
	// Route 53 Hosted Zone used to perform DNS validation of the request.
	//
	// The zone
	// must be authoritative for the domain name specified in the Certificate Request.
	HostedZone awsroute53.IHostedZone `json:"hostedZone" yaml:"hostedZone"`
	// Role to use for the custom resource that creates the validated certificate.
	CustomResourceRole awsiam.IRole `json:"customResourceRole" yaml:"customResourceRole"`
	// AWS region that will host the certificate.
	//
	// This is needed especially
	// for certificates used for CloudFront distributions, which require the region
	// to be us-east-1.
	Region *string `json:"region" yaml:"region"`
	// An endpoint of Route53 service, which is not necessary as AWS SDK could figure out the right endpoints for most regions, but for some regions such as those in aws-cn partition, the default endpoint is not working now, hence the right endpoint need to be specified through this prop.
	//
	// Route53 is not been officially launched in China, it is only available for AWS
	// internal accounts now. To make DnsValidatedCertificate work for internal accounts
	// now, a special endpoint needs to be provided.
	Route53Endpoint *string `json:"route53Endpoint" yaml:"route53Endpoint"`
}

// Represents a certificate in AWS Certificate Manager.
type ICertificate interface {
	awscdk.IResource
	// Return the DaysToExpiry metric for this AWS Certificate Manager Certificate. By default, this is the minimum value over 1 day.
	//
	// This metric is no longer emitted once the certificate has effectively
	// expired, so alarms configured on this metric should probably treat missing
	// data as "breaching".
	MetricDaysToExpiry(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The certificate's ARN.
	CertificateArn() *string
}

// The jsii proxy for ICertificate
type jsiiProxy_ICertificate struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICertificate) MetricDaysToExpiry(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDaysToExpiry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ICertificate) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

// A private certificate managed by AWS Certificate Manager.
//
// TODO: EXAMPLE
//
type PrivateCertificate interface {
	awscdk.Resource
	ICertificate
	CertificateArn() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Region() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	MetricDaysToExpiry(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for PrivateCertificate
type jsiiProxy_PrivateCertificate struct {
	internal.Type__awscdkResource
	jsiiProxy_ICertificate
}

func (j *jsiiProxy_PrivateCertificate) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateCertificate) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateCertificate) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateCertificate) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateCertificate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewPrivateCertificate(scope constructs.Construct, id *string, props *PrivateCertificateProps) PrivateCertificate {
	_init_.Initialize()

	j := jsiiProxy_PrivateCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_certificatemanager.PrivateCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPrivateCertificate_Override(p PrivateCertificate, scope constructs.Construct, id *string, props *PrivateCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_certificatemanager.PrivateCertificate",
		[]interface{}{scope, id, props},
		p,
	)
}

// Import a certificate.
func PrivateCertificate_FromCertificateArn(scope constructs.Construct, id *string, certificateArn *string) ICertificate {
	_init_.Initialize()

	var returns ICertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.PrivateCertificate",
		"fromCertificateArn",
		[]interface{}{scope, id, certificateArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PrivateCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.PrivateCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func PrivateCertificate_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.PrivateCertificate",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (p *jsiiProxy_PrivateCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_PrivateCertificate) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (p *jsiiProxy_PrivateCertificate) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (p *jsiiProxy_PrivateCertificate) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return the DaysToExpiry metric for this AWS Certificate Manager Certificate. By default, this is the minimum value over 1 day.
//
// This metric is no longer emitted once the certificate has effectively
// expired, so alarms configured on this metric should probably treat missing
// data as "breaching".
func (p *jsiiProxy_PrivateCertificate) MetricDaysToExpiry(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricDaysToExpiry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PrivateCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for your private certificate.
//
// TODO: EXAMPLE
//
type PrivateCertificateProps struct {
	// Private certificate authority (CA) that will be used to issue the certificate.
	CertificateAuthority awsacmpca.ICertificateAuthority `json:"certificateAuthority" yaml:"certificateAuthority"`
	// Fully-qualified domain name to request a private certificate for.
	//
	// May contain wildcards, such as ``*.domain.com``.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Alternative domain names on your private certificate.
	//
	// Use this to register alternative domain names that represent the same site.
	SubjectAlternativeNames *[]*string `json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

// Method used to assert ownership of the domain.
type ValidationMethod string

const (
	ValidationMethod_EMAIL ValidationMethod = "EMAIL"
	ValidationMethod_DNS ValidationMethod = "DNS"
)

