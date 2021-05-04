package awsacmpca

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsacmpca/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Defines a Certificate for ACMPCA.
// Experimental.
type CertificateAuthority interface {
}

// The jsii proxy struct for CertificateAuthority
type jsiiProxy_CertificateAuthority struct {
	_ byte // padding
}

// Import an existing Certificate given an ARN.
// Experimental.
func CertificateAuthority_FromCertificateAuthorityArn(scope constructs.Construct, id *string, certificateAuthorityArn *string) ICertificateAuthority {
	_init_.Initialize()

	var returns ICertificateAuthority

	_jsii_.StaticInvoke(
		"monocdk.aws_acmpca.CertificateAuthority",
		"fromCertificateAuthorityArn",
		[]interface{}{scope, id, certificateAuthorityArn},
		&returns,
	)

	return returns
}

// A CloudFormation `AWS::ACMPCA::Certificate`.
type CfnCertificate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiPassthrough() interface{}
	SetApiPassthrough(val interface{})
	AttrArn() *string
	AttrCertificate() *string
	CertificateAuthorityArn() *string
	SetCertificateAuthorityArn(val *string)
	CertificateSigningRequest() *string
	SetCertificateSigningRequest(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	SigningAlgorithm() *string
	SetSigningAlgorithm(val *string)
	Stack() awscdk.Stack
	TemplateArn() *string
	SetTemplateArn(val *string)
	UpdatedProperites() *map[string]interface{}
	Validity() interface{}
	SetValidity(val interface{})
	ValidityNotBefore() interface{}
	SetValidityNotBefore(val interface{})
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCertificate
type jsiiProxy_CfnCertificate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCertificate) ApiPassthrough() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiPassthrough",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) AttrCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCertificate",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnCertificate) CertificateSigningRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSigningRequest",
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

func (j *jsiiProxy_CfnCertificate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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

func (j *jsiiProxy_CfnCertificate) SigningAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAlgorithm",
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

func (j *jsiiProxy_CfnCertificate) TemplateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateArn",
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

func (j *jsiiProxy_CfnCertificate) Validity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) ValidityNotBefore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validityNotBefore",
		&returns,
	)
	return returns
}


// Create a new `AWS::ACMPCA::Certificate`.
func NewCfnCertificate(scope awscdk.Construct, id *string, props *CfnCertificateProps) CfnCertificate {
	_init_.Initialize()

	j := jsiiProxy_CfnCertificate{}

	_jsii_.Create(
		"monocdk.aws_acmpca.CfnCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ACMPCA::Certificate`.
func NewCfnCertificate_Override(c CfnCertificate, scope awscdk.Construct, id *string, props *CfnCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_acmpca.CfnCertificate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCertificate) SetApiPassthrough(val interface{}) {
	_jsii_.Set(
		j,
		"apiPassthrough",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetCertificateAuthorityArn(val *string) {
	_jsii_.Set(
		j,
		"certificateAuthorityArn",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetCertificateSigningRequest(val *string) {
	_jsii_.Set(
		j,
		"certificateSigningRequest",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetSigningAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"signingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetTemplateArn(val *string) {
	_jsii_.Set(
		j,
		"templateArn",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetValidity(val interface{}) {
	_jsii_.Set(
		j,
		"validity",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetValidityNotBefore(val interface{}) {
	_jsii_.Set(
		j,
		"validityNotBefore",
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
func CfnCertificate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_acmpca.CfnCertificate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCertificate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_acmpca.CfnCertificate",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_acmpca.CfnCertificate",
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
		"monocdk.aws_acmpca.CfnCertificate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnCertificate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCertificate) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCertificate) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCertificate) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCertificate) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCertificate) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCertificate) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnCertificate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnCertificate_ApiPassthroughProperty struct {
	// `CfnCertificate.ApiPassthroughProperty.Extensions`.
	Extensions interface{} `json:"extensions"`
	// `CfnCertificate.ApiPassthroughProperty.Subject`.
	Subject interface{} `json:"subject"`
}

type CfnCertificate_EdiPartyNameProperty struct {
	// `CfnCertificate.EdiPartyNameProperty.NameAssigner`.
	NameAssigner *string `json:"nameAssigner"`
	// `CfnCertificate.EdiPartyNameProperty.PartyName`.
	PartyName *string `json:"partyName"`
}

type CfnCertificate_ExtendedKeyUsageProperty struct {
	// `CfnCertificate.ExtendedKeyUsageProperty.ExtendedKeyUsageObjectIdentifier`.
	ExtendedKeyUsageObjectIdentifier *string `json:"extendedKeyUsageObjectIdentifier"`
	// `CfnCertificate.ExtendedKeyUsageProperty.ExtendedKeyUsageType`.
	ExtendedKeyUsageType *string `json:"extendedKeyUsageType"`
}

type CfnCertificate_ExtensionsProperty struct {
	// `CfnCertificate.ExtensionsProperty.CertificatePolicies`.
	CertificatePolicies interface{} `json:"certificatePolicies"`
	// `CfnCertificate.ExtensionsProperty.ExtendedKeyUsage`.
	ExtendedKeyUsage interface{} `json:"extendedKeyUsage"`
	// `CfnCertificate.ExtensionsProperty.KeyUsage`.
	KeyUsage interface{} `json:"keyUsage"`
	// `CfnCertificate.ExtensionsProperty.SubjectAlternativeNames`.
	SubjectAlternativeNames interface{} `json:"subjectAlternativeNames"`
}

type CfnCertificate_GeneralNameProperty struct {
	// `CfnCertificate.GeneralNameProperty.DirectoryName`.
	DirectoryName interface{} `json:"directoryName"`
	// `CfnCertificate.GeneralNameProperty.DnsName`.
	DnsName *string `json:"dnsName"`
	// `CfnCertificate.GeneralNameProperty.EdiPartyName`.
	EdiPartyName interface{} `json:"ediPartyName"`
	// `CfnCertificate.GeneralNameProperty.IpAddress`.
	IpAddress *string `json:"ipAddress"`
	// `CfnCertificate.GeneralNameProperty.OtherName`.
	OtherName interface{} `json:"otherName"`
	// `CfnCertificate.GeneralNameProperty.RegisteredId`.
	RegisteredId *string `json:"registeredId"`
	// `CfnCertificate.GeneralNameProperty.Rfc822Name`.
	Rfc822Name *string `json:"rfc822Name"`
	// `CfnCertificate.GeneralNameProperty.UniformResourceIdentifier`.
	UniformResourceIdentifier *string `json:"uniformResourceIdentifier"`
}

type CfnCertificate_KeyUsageProperty struct {
	// `CfnCertificate.KeyUsageProperty.CRLSign`.
	CrlSign interface{} `json:"crlSign"`
	// `CfnCertificate.KeyUsageProperty.DataEncipherment`.
	DataEncipherment interface{} `json:"dataEncipherment"`
	// `CfnCertificate.KeyUsageProperty.DecipherOnly`.
	DecipherOnly interface{} `json:"decipherOnly"`
	// `CfnCertificate.KeyUsageProperty.DigitalSignature`.
	DigitalSignature interface{} `json:"digitalSignature"`
	// `CfnCertificate.KeyUsageProperty.EncipherOnly`.
	EncipherOnly interface{} `json:"encipherOnly"`
	// `CfnCertificate.KeyUsageProperty.KeyAgreement`.
	KeyAgreement interface{} `json:"keyAgreement"`
	// `CfnCertificate.KeyUsageProperty.KeyCertSign`.
	KeyCertSign interface{} `json:"keyCertSign"`
	// `CfnCertificate.KeyUsageProperty.KeyEncipherment`.
	KeyEncipherment interface{} `json:"keyEncipherment"`
	// `CfnCertificate.KeyUsageProperty.NonRepudiation`.
	NonRepudiation interface{} `json:"nonRepudiation"`
}

type CfnCertificate_OtherNameProperty struct {
	// `CfnCertificate.OtherNameProperty.TypeId`.
	TypeId *string `json:"typeId"`
	// `CfnCertificate.OtherNameProperty.Value`.
	Value *string `json:"value"`
}

type CfnCertificate_PolicyInformationProperty struct {
	// `CfnCertificate.PolicyInformationProperty.CertPolicyId`.
	CertPolicyId *string `json:"certPolicyId"`
	// `CfnCertificate.PolicyInformationProperty.PolicyQualifiers`.
	PolicyQualifiers interface{} `json:"policyQualifiers"`
}

type CfnCertificate_PolicyQualifierInfoProperty struct {
	// `CfnCertificate.PolicyQualifierInfoProperty.PolicyQualifierId`.
	PolicyQualifierId *string `json:"policyQualifierId"`
	// `CfnCertificate.PolicyQualifierInfoProperty.Qualifier`.
	Qualifier interface{} `json:"qualifier"`
}

type CfnCertificate_QualifierProperty struct {
	// `CfnCertificate.QualifierProperty.CpsUri`.
	CpsUri *string `json:"cpsUri"`
}

type CfnCertificate_SubjectProperty struct {
	// `CfnCertificate.SubjectProperty.CommonName`.
	CommonName *string `json:"commonName"`
	// `CfnCertificate.SubjectProperty.Country`.
	Country *string `json:"country"`
	// `CfnCertificate.SubjectProperty.DistinguishedNameQualifier`.
	DistinguishedNameQualifier *string `json:"distinguishedNameQualifier"`
	// `CfnCertificate.SubjectProperty.GenerationQualifier`.
	GenerationQualifier *string `json:"generationQualifier"`
	// `CfnCertificate.SubjectProperty.GivenName`.
	GivenName *string `json:"givenName"`
	// `CfnCertificate.SubjectProperty.Initials`.
	Initials *string `json:"initials"`
	// `CfnCertificate.SubjectProperty.Locality`.
	Locality *string `json:"locality"`
	// `CfnCertificate.SubjectProperty.Organization`.
	Organization *string `json:"organization"`
	// `CfnCertificate.SubjectProperty.OrganizationalUnit`.
	OrganizationalUnit *string `json:"organizationalUnit"`
	// `CfnCertificate.SubjectProperty.Pseudonym`.
	Pseudonym *string `json:"pseudonym"`
	// `CfnCertificate.SubjectProperty.SerialNumber`.
	SerialNumber *string `json:"serialNumber"`
	// `CfnCertificate.SubjectProperty.State`.
	State *string `json:"state"`
	// `CfnCertificate.SubjectProperty.Surname`.
	Surname *string `json:"surname"`
	// `CfnCertificate.SubjectProperty.Title`.
	Title *string `json:"title"`
}

type CfnCertificate_ValidityProperty struct {
	// `CfnCertificate.ValidityProperty.Type`.
	Type *string `json:"type"`
	// `CfnCertificate.ValidityProperty.Value`.
	Value *float64 `json:"value"`
}

// A CloudFormation `AWS::ACMPCA::CertificateAuthority`.
type CfnCertificateAuthority interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrCertificateSigningRequest() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	CsrExtensions() interface{}
	SetCsrExtensions(val interface{})
	KeyAlgorithm() *string
	SetKeyAlgorithm(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	RevocationConfiguration() interface{}
	SetRevocationConfiguration(val interface{})
	SigningAlgorithm() *string
	SetSigningAlgorithm(val *string)
	Stack() awscdk.Stack
	Subject() interface{}
	SetSubject(val interface{})
	Tags() awscdk.TagManager
	Type() *string
	SetType(val *string)
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCertificateAuthority
type jsiiProxy_CfnCertificateAuthority struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCertificateAuthority) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) AttrCertificateSigningRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCertificateSigningRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) CsrExtensions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"csrExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) KeyAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) RevocationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revocationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) SigningAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) Subject() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthority) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ACMPCA::CertificateAuthority`.
func NewCfnCertificateAuthority(scope awscdk.Construct, id *string, props *CfnCertificateAuthorityProps) CfnCertificateAuthority {
	_init_.Initialize()

	j := jsiiProxy_CfnCertificateAuthority{}

	_jsii_.Create(
		"monocdk.aws_acmpca.CfnCertificateAuthority",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ACMPCA::CertificateAuthority`.
func NewCfnCertificateAuthority_Override(c CfnCertificateAuthority, scope awscdk.Construct, id *string, props *CfnCertificateAuthorityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_acmpca.CfnCertificateAuthority",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCertificateAuthority) SetCsrExtensions(val interface{}) {
	_jsii_.Set(
		j,
		"csrExtensions",
		val,
	)
}

func (j *jsiiProxy_CfnCertificateAuthority) SetKeyAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"keyAlgorithm",
		val,
	)
}

func (j *jsiiProxy_CfnCertificateAuthority) SetRevocationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"revocationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCertificateAuthority) SetSigningAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"signingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_CfnCertificateAuthority) SetSubject(val interface{}) {
	_jsii_.Set(
		j,
		"subject",
		val,
	)
}

func (j *jsiiProxy_CfnCertificateAuthority) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
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
func CfnCertificateAuthority_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_acmpca.CfnCertificateAuthority",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCertificateAuthority_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_acmpca.CfnCertificateAuthority",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCertificateAuthority_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_acmpca.CfnCertificateAuthority",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCertificateAuthority_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_acmpca.CfnCertificateAuthority",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnCertificateAuthority) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCertificateAuthority) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnCertificateAuthority) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnCertificateAuthority_AccessDescriptionProperty struct {
	// `CfnCertificateAuthority.AccessDescriptionProperty.AccessLocation`.
	AccessLocation interface{} `json:"accessLocation"`
	// `CfnCertificateAuthority.AccessDescriptionProperty.AccessMethod`.
	AccessMethod interface{} `json:"accessMethod"`
}

type CfnCertificateAuthority_AccessMethodProperty struct {
	// `CfnCertificateAuthority.AccessMethodProperty.AccessMethodType`.
	AccessMethodType *string `json:"accessMethodType"`
	// `CfnCertificateAuthority.AccessMethodProperty.CustomObjectIdentifier`.
	CustomObjectIdentifier *string `json:"customObjectIdentifier"`
}

type CfnCertificateAuthority_CrlConfigurationProperty struct {
	// `CfnCertificateAuthority.CrlConfigurationProperty.CustomCname`.
	CustomCname *string `json:"customCname"`
	// `CfnCertificateAuthority.CrlConfigurationProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnCertificateAuthority.CrlConfigurationProperty.ExpirationInDays`.
	ExpirationInDays *float64 `json:"expirationInDays"`
	// `CfnCertificateAuthority.CrlConfigurationProperty.S3BucketName`.
	S3BucketName *string `json:"s3BucketName"`
}

type CfnCertificateAuthority_CsrExtensionsProperty struct {
	// `CfnCertificateAuthority.CsrExtensionsProperty.KeyUsage`.
	KeyUsage interface{} `json:"keyUsage"`
	// `CfnCertificateAuthority.CsrExtensionsProperty.SubjectInformationAccess`.
	SubjectInformationAccess interface{} `json:"subjectInformationAccess"`
}

type CfnCertificateAuthority_EdiPartyNameProperty struct {
	// `CfnCertificateAuthority.EdiPartyNameProperty.NameAssigner`.
	NameAssigner *string `json:"nameAssigner"`
	// `CfnCertificateAuthority.EdiPartyNameProperty.PartyName`.
	PartyName *string `json:"partyName"`
}

type CfnCertificateAuthority_GeneralNameProperty struct {
	// `CfnCertificateAuthority.GeneralNameProperty.DirectoryName`.
	DirectoryName interface{} `json:"directoryName"`
	// `CfnCertificateAuthority.GeneralNameProperty.DnsName`.
	DnsName *string `json:"dnsName"`
	// `CfnCertificateAuthority.GeneralNameProperty.EdiPartyName`.
	EdiPartyName interface{} `json:"ediPartyName"`
	// `CfnCertificateAuthority.GeneralNameProperty.IpAddress`.
	IpAddress *string `json:"ipAddress"`
	// `CfnCertificateAuthority.GeneralNameProperty.OtherName`.
	OtherName interface{} `json:"otherName"`
	// `CfnCertificateAuthority.GeneralNameProperty.RegisteredId`.
	RegisteredId *string `json:"registeredId"`
	// `CfnCertificateAuthority.GeneralNameProperty.Rfc822Name`.
	Rfc822Name *string `json:"rfc822Name"`
	// `CfnCertificateAuthority.GeneralNameProperty.UniformResourceIdentifier`.
	UniformResourceIdentifier *string `json:"uniformResourceIdentifier"`
}

type CfnCertificateAuthority_KeyUsageProperty struct {
	// `CfnCertificateAuthority.KeyUsageProperty.CRLSign`.
	CrlSign interface{} `json:"crlSign"`
	// `CfnCertificateAuthority.KeyUsageProperty.DataEncipherment`.
	DataEncipherment interface{} `json:"dataEncipherment"`
	// `CfnCertificateAuthority.KeyUsageProperty.DecipherOnly`.
	DecipherOnly interface{} `json:"decipherOnly"`
	// `CfnCertificateAuthority.KeyUsageProperty.DigitalSignature`.
	DigitalSignature interface{} `json:"digitalSignature"`
	// `CfnCertificateAuthority.KeyUsageProperty.EncipherOnly`.
	EncipherOnly interface{} `json:"encipherOnly"`
	// `CfnCertificateAuthority.KeyUsageProperty.KeyAgreement`.
	KeyAgreement interface{} `json:"keyAgreement"`
	// `CfnCertificateAuthority.KeyUsageProperty.KeyCertSign`.
	KeyCertSign interface{} `json:"keyCertSign"`
	// `CfnCertificateAuthority.KeyUsageProperty.KeyEncipherment`.
	KeyEncipherment interface{} `json:"keyEncipherment"`
	// `CfnCertificateAuthority.KeyUsageProperty.NonRepudiation`.
	NonRepudiation interface{} `json:"nonRepudiation"`
}

type CfnCertificateAuthority_OtherNameProperty struct {
	// `CfnCertificateAuthority.OtherNameProperty.TypeId`.
	TypeId *string `json:"typeId"`
	// `CfnCertificateAuthority.OtherNameProperty.Value`.
	Value *string `json:"value"`
}

type CfnCertificateAuthority_RevocationConfigurationProperty struct {
	// `CfnCertificateAuthority.RevocationConfigurationProperty.CrlConfiguration`.
	CrlConfiguration interface{} `json:"crlConfiguration"`
}

type CfnCertificateAuthority_SubjectProperty struct {
	// `CfnCertificateAuthority.SubjectProperty.CommonName`.
	CommonName *string `json:"commonName"`
	// `CfnCertificateAuthority.SubjectProperty.Country`.
	Country *string `json:"country"`
	// `CfnCertificateAuthority.SubjectProperty.DistinguishedNameQualifier`.
	DistinguishedNameQualifier *string `json:"distinguishedNameQualifier"`
	// `CfnCertificateAuthority.SubjectProperty.GenerationQualifier`.
	GenerationQualifier *string `json:"generationQualifier"`
	// `CfnCertificateAuthority.SubjectProperty.GivenName`.
	GivenName *string `json:"givenName"`
	// `CfnCertificateAuthority.SubjectProperty.Initials`.
	Initials *string `json:"initials"`
	// `CfnCertificateAuthority.SubjectProperty.Locality`.
	Locality *string `json:"locality"`
	// `CfnCertificateAuthority.SubjectProperty.Organization`.
	Organization *string `json:"organization"`
	// `CfnCertificateAuthority.SubjectProperty.OrganizationalUnit`.
	OrganizationalUnit *string `json:"organizationalUnit"`
	// `CfnCertificateAuthority.SubjectProperty.Pseudonym`.
	Pseudonym *string `json:"pseudonym"`
	// `CfnCertificateAuthority.SubjectProperty.SerialNumber`.
	SerialNumber *string `json:"serialNumber"`
	// `CfnCertificateAuthority.SubjectProperty.State`.
	State *string `json:"state"`
	// `CfnCertificateAuthority.SubjectProperty.Surname`.
	Surname *string `json:"surname"`
	// `CfnCertificateAuthority.SubjectProperty.Title`.
	Title *string `json:"title"`
}

// A CloudFormation `AWS::ACMPCA::CertificateAuthorityActivation`.
type CfnCertificateAuthorityActivation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrCompleteCertificateChain() *string
	Certificate() *string
	SetCertificate(val *string)
	CertificateAuthorityArn() *string
	SetCertificateAuthorityArn(val *string)
	CertificateChain() *string
	SetCertificateChain(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	Stack() awscdk.Stack
	Status() *string
	SetStatus(val *string)
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCertificateAuthorityActivation
type jsiiProxy_CfnCertificateAuthorityActivation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) AttrCompleteCertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCompleteCertificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) CertificateAuthorityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) CertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ACMPCA::CertificateAuthorityActivation`.
func NewCfnCertificateAuthorityActivation(scope awscdk.Construct, id *string, props *CfnCertificateAuthorityActivationProps) CfnCertificateAuthorityActivation {
	_init_.Initialize()

	j := jsiiProxy_CfnCertificateAuthorityActivation{}

	_jsii_.Create(
		"monocdk.aws_acmpca.CfnCertificateAuthorityActivation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ACMPCA::CertificateAuthorityActivation`.
func NewCfnCertificateAuthorityActivation_Override(c CfnCertificateAuthorityActivation, scope awscdk.Construct, id *string, props *CfnCertificateAuthorityActivationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_acmpca.CfnCertificateAuthorityActivation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) SetCertificateAuthorityArn(val *string) {
	_jsii_.Set(
		j,
		"certificateAuthorityArn",
		val,
	)
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) SetCertificateChain(val *string) {
	_jsii_.Set(
		j,
		"certificateChain",
		val,
	)
}

func (j *jsiiProxy_CfnCertificateAuthorityActivation) SetStatus(val *string) {
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
func CfnCertificateAuthorityActivation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_acmpca.CfnCertificateAuthorityActivation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCertificateAuthorityActivation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_acmpca.CfnCertificateAuthorityActivation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCertificateAuthorityActivation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_acmpca.CfnCertificateAuthorityActivation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCertificateAuthorityActivation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_acmpca.CfnCertificateAuthorityActivation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnCertificateAuthorityActivation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCertificateAuthorityActivation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnCertificateAuthorityActivation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::ACMPCA::CertificateAuthorityActivation`.
type CfnCertificateAuthorityActivationProps struct {
	// `AWS::ACMPCA::CertificateAuthorityActivation.Certificate`.
	Certificate *string `json:"certificate"`
	// `AWS::ACMPCA::CertificateAuthorityActivation.CertificateAuthorityArn`.
	CertificateAuthorityArn *string `json:"certificateAuthorityArn"`
	// `AWS::ACMPCA::CertificateAuthorityActivation.CertificateChain`.
	CertificateChain *string `json:"certificateChain"`
	// `AWS::ACMPCA::CertificateAuthorityActivation.Status`.
	Status *string `json:"status"`
}

// Properties for defining a `AWS::ACMPCA::CertificateAuthority`.
type CfnCertificateAuthorityProps struct {
	// `AWS::ACMPCA::CertificateAuthority.KeyAlgorithm`.
	KeyAlgorithm *string `json:"keyAlgorithm"`
	// `AWS::ACMPCA::CertificateAuthority.SigningAlgorithm`.
	SigningAlgorithm *string `json:"signingAlgorithm"`
	// `AWS::ACMPCA::CertificateAuthority.Subject`.
	Subject interface{} `json:"subject"`
	// `AWS::ACMPCA::CertificateAuthority.Type`.
	Type *string `json:"type"`
	// `AWS::ACMPCA::CertificateAuthority.CsrExtensions`.
	CsrExtensions interface{} `json:"csrExtensions"`
	// `AWS::ACMPCA::CertificateAuthority.RevocationConfiguration`.
	RevocationConfiguration interface{} `json:"revocationConfiguration"`
	// `AWS::ACMPCA::CertificateAuthority.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// Properties for defining a `AWS::ACMPCA::Certificate`.
type CfnCertificateProps struct {
	// `AWS::ACMPCA::Certificate.CertificateAuthorityArn`.
	CertificateAuthorityArn *string `json:"certificateAuthorityArn"`
	// `AWS::ACMPCA::Certificate.CertificateSigningRequest`.
	CertificateSigningRequest *string `json:"certificateSigningRequest"`
	// `AWS::ACMPCA::Certificate.SigningAlgorithm`.
	SigningAlgorithm *string `json:"signingAlgorithm"`
	// `AWS::ACMPCA::Certificate.Validity`.
	Validity interface{} `json:"validity"`
	// `AWS::ACMPCA::Certificate.ApiPassthrough`.
	ApiPassthrough interface{} `json:"apiPassthrough"`
	// `AWS::ACMPCA::Certificate.TemplateArn`.
	TemplateArn *string `json:"templateArn"`
	// `AWS::ACMPCA::Certificate.ValidityNotBefore`.
	ValidityNotBefore interface{} `json:"validityNotBefore"`
}

// Interface which all CertificateAuthority based class must implement.
// Experimental.
type ICertificateAuthority interface {
	awscdk.IResource
	// The Amazon Resource Name of the Certificate.
	// Experimental.
	CertificateAuthorityArn() *string
}

// The jsii proxy for ICertificateAuthority
type jsiiProxy_ICertificateAuthority struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ICertificateAuthority) CertificateAuthorityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArn",
		&returns,
	)
	return returns
}

