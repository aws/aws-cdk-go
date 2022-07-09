package awsacmpca

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines a Certificate for ACMPCA.
//
// Example:
//   import acmpca "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.tls(&tlsAuthProps{
//   		certificateAuthorities: []iCertificateAuthority{
//   			acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("CertificateAuthority"), jsii.String("arn:aws:acm-pca:us-west-2:1234567890:certificate-authority/11111111-1111-1111-1111-111111111111")),
//   		},
//   	}),
//   })
//
type CertificateAuthority interface {
}

// The jsii proxy struct for CertificateAuthority
type jsiiProxy_CertificateAuthority struct {
	_ byte // padding
}

// Import an existing Certificate given an ARN.
func CertificateAuthority_FromCertificateAuthorityArn(scope constructs.Construct, id *string, certificateAuthorityArn *string) ICertificateAuthority {
	_init_.Initialize()

	var returns ICertificateAuthority

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CertificateAuthority",
		"fromCertificateAuthorityArn",
		[]interface{}{scope, id, certificateAuthorityArn},
		&returns,
	)

	return returns
}

// A CloudFormation `AWS::ACMPCA::Certificate`.
//
// The `AWS::ACMPCA::Certificate` resource is used to issue a certificate using your private certificate authority. For more information, see the [IssueCertificate](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_IssueCertificate.html) action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificate := awscdk.Aws_acmpca.NewCfnCertificate(this, jsii.String("MyCfnCertificate"), &cfnCertificateProps{
//   	certificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	certificateSigningRequest: jsii.String("certificateSigningRequest"),
//   	signingAlgorithm: jsii.String("signingAlgorithm"),
//   	validity: &validityProperty{
//   		type: jsii.String("type"),
//   		value: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	apiPassthrough: &apiPassthroughProperty{
//   		extensions: &extensionsProperty{
//   			certificatePolicies: []interface{}{
//   				&policyInformationProperty{
//   					certPolicyId: jsii.String("certPolicyId"),
//
//   					// the properties below are optional
//   					policyQualifiers: []interface{}{
//   						&policyQualifierInfoProperty{
//   							policyQualifierId: jsii.String("policyQualifierId"),
//   							qualifier: &qualifierProperty{
//   								cpsUri: jsii.String("cpsUri"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			customExtensions: []interface{}{
//   				&customExtensionProperty{
//   					objectIdentifier: jsii.String("objectIdentifier"),
//   					value: jsii.String("value"),
//
//   					// the properties below are optional
//   					critical: jsii.Boolean(false),
//   				},
//   			},
//   			extendedKeyUsage: []interface{}{
//   				&extendedKeyUsageProperty{
//   					extendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   					extendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   				},
//   			},
//   			keyUsage: &keyUsageProperty{
//   				crlSign: jsii.Boolean(false),
//   				dataEncipherment: jsii.Boolean(false),
//   				decipherOnly: jsii.Boolean(false),
//   				digitalSignature: jsii.Boolean(false),
//   				encipherOnly: jsii.Boolean(false),
//   				keyAgreement: jsii.Boolean(false),
//   				keyCertSign: jsii.Boolean(false),
//   				keyEncipherment: jsii.Boolean(false),
//   				nonRepudiation: jsii.Boolean(false),
//   			},
//   			subjectAlternativeNames: []interface{}{
//   				&generalNameProperty{
//   					directoryName: &subjectProperty{
//   						commonName: jsii.String("commonName"),
//   						country: jsii.String("country"),
//   						customAttributes: []interface{}{
//   							&customAttributeProperty{
//   								objectIdentifier: jsii.String("objectIdentifier"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   						generationQualifier: jsii.String("generationQualifier"),
//   						givenName: jsii.String("givenName"),
//   						initials: jsii.String("initials"),
//   						locality: jsii.String("locality"),
//   						organization: jsii.String("organization"),
//   						organizationalUnit: jsii.String("organizationalUnit"),
//   						pseudonym: jsii.String("pseudonym"),
//   						serialNumber: jsii.String("serialNumber"),
//   						state: jsii.String("state"),
//   						surname: jsii.String("surname"),
//   						title: jsii.String("title"),
//   					},
//   					dnsName: jsii.String("dnsName"),
//   					ediPartyName: &ediPartyNameProperty{
//   						nameAssigner: jsii.String("nameAssigner"),
//   						partyName: jsii.String("partyName"),
//   					},
//   					ipAddress: jsii.String("ipAddress"),
//   					otherName: &otherNameProperty{
//   						typeId: jsii.String("typeId"),
//   						value: jsii.String("value"),
//   					},
//   					registeredId: jsii.String("registeredId"),
//   					rfc822Name: jsii.String("rfc822Name"),
//   					uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   				},
//   			},
//   		},
//   		subject: &subjectProperty{
//   			commonName: jsii.String("commonName"),
//   			country: jsii.String("country"),
//   			customAttributes: []interface{}{
//   				&customAttributeProperty{
//   					objectIdentifier: jsii.String("objectIdentifier"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   			generationQualifier: jsii.String("generationQualifier"),
//   			givenName: jsii.String("givenName"),
//   			initials: jsii.String("initials"),
//   			locality: jsii.String("locality"),
//   			organization: jsii.String("organization"),
//   			organizationalUnit: jsii.String("organizationalUnit"),
//   			pseudonym: jsii.String("pseudonym"),
//   			serialNumber: jsii.String("serialNumber"),
//   			state: jsii.String("state"),
//   			surname: jsii.String("surname"),
//   			title: jsii.String("title"),
//   		},
//   	},
//   	templateArn: jsii.String("templateArn"),
//   	validityNotBefore: &validityProperty{
//   		type: jsii.String("type"),
//   		value: jsii.Number(123),
//   	},
//   })
//
type CfnCertificate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Specifies X.509 certificate information to be included in the issued certificate. An `APIPassthrough` or `APICSRPassthrough` template variant must be selected, or else this parameter is ignored.
	ApiPassthrough() interface{}
	SetApiPassthrough(val interface{})
	// The Amazon Resource Name (ARN) of the issued certificate.
	AttrArn() *string
	// The issued Base64 PEM-encoded certificate.
	AttrCertificate() *string
	// The Amazon Resource Name (ARN) for the private CA issues the certificate.
	CertificateAuthorityArn() *string
	SetCertificateAuthorityArn(val *string)
	// The certificate signing request (CSR) for the certificate.
	CertificateSigningRequest() *string
	SetCertificateSigningRequest(val *string)
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
	// The name of the algorithm that will be used to sign the certificate to be issued.
	//
	// This parameter should not be confused with the `SigningAlgorithm` parameter used to sign a CSR in the `CreateCertificateAuthority` action.
	//
	// > The specified signing algorithm family (RSA or ECDSA) must match the algorithm family of the CA's secret key.
	SigningAlgorithm() *string
	SetSigningAlgorithm(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Specifies a custom configuration template to use when issuing a certificate.
	//
	// If this parameter is not provided, ACM Private CA defaults to the `EndEntityCertificate/V1` template. For more information about ACM Private CA templates, see [Using Templates](https://docs.aws.amazon.com/acm-pca/latest/userguide/UsingTemplates.html) .
	TemplateArn() *string
	SetTemplateArn(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The period of time during which the certificate will be valid.
	Validity() interface{}
	SetValidity(val interface{})
	// Information describing the start of the validity period of the certificate.
	//
	// This parameter sets the “Not Before" date for the certificate.
	//
	// By default, when issuing a certificate, ACM Private CA sets the "Not Before" date to the issuance time minus 60 minutes. This compensates for clock inconsistencies across computer systems. The `ValidityNotBefore` parameter can be used to customize the “Not Before” value.
	//
	// Unlike the `Validity` parameter, the `ValidityNotBefore` parameter is optional.
	//
	// The `ValidityNotBefore` value is expressed as an explicit date and time, using the `Validity` type value `ABSOLUTE` .
	ValidityNotBefore() interface{}
	SetValidityNotBefore(val interface{})
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
func NewCfnCertificate(scope constructs.Construct, id *string, props *CfnCertificateProps) CfnCertificate {
	_init_.Initialize()

	j := jsiiProxy_CfnCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_acmpca.CfnCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ACMPCA::Certificate`.
func NewCfnCertificate_Override(c CfnCertificate, scope constructs.Construct, id *string, props *CfnCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_acmpca.CfnCertificate",
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
func CfnCertificate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CfnCertificate",
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
		"aws-cdk-lib.aws_acmpca.CfnCertificate",
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
func CfnCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CfnCertificate",
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
		"aws-cdk-lib.aws_acmpca.CfnCertificate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCertificate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCertificate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCertificate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCertificate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCertificate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnCertificate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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

// Contains X.509 certificate information to be placed in an issued certificate. An `APIPassthrough` or `APICSRPassthrough` template variant must be selected, or else this parameter is ignored.
//
// If conflicting or duplicate certificate information is supplied from other sources, ACM Private CA applies [order of operation rules](https://docs.aws.amazon.com/acm-pca/latest/userguide/UsingTemplates.html#template-order-of-operations) to determine what information is used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiPassthroughProperty := &apiPassthroughProperty{
//   	extensions: &extensionsProperty{
//   		certificatePolicies: []interface{}{
//   			&policyInformationProperty{
//   				certPolicyId: jsii.String("certPolicyId"),
//
//   				// the properties below are optional
//   				policyQualifiers: []interface{}{
//   					&policyQualifierInfoProperty{
//   						policyQualifierId: jsii.String("policyQualifierId"),
//   						qualifier: &qualifierProperty{
//   							cpsUri: jsii.String("cpsUri"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		customExtensions: []interface{}{
//   			&customExtensionProperty{
//   				objectIdentifier: jsii.String("objectIdentifier"),
//   				value: jsii.String("value"),
//
//   				// the properties below are optional
//   				critical: jsii.Boolean(false),
//   			},
//   		},
//   		extendedKeyUsage: []interface{}{
//   			&extendedKeyUsageProperty{
//   				extendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   				extendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   			},
//   		},
//   		keyUsage: &keyUsageProperty{
//   			crlSign: jsii.Boolean(false),
//   			dataEncipherment: jsii.Boolean(false),
//   			decipherOnly: jsii.Boolean(false),
//   			digitalSignature: jsii.Boolean(false),
//   			encipherOnly: jsii.Boolean(false),
//   			keyAgreement: jsii.Boolean(false),
//   			keyCertSign: jsii.Boolean(false),
//   			keyEncipherment: jsii.Boolean(false),
//   			nonRepudiation: jsii.Boolean(false),
//   		},
//   		subjectAlternativeNames: []interface{}{
//   			&generalNameProperty{
//   				directoryName: &subjectProperty{
//   					commonName: jsii.String("commonName"),
//   					country: jsii.String("country"),
//   					customAttributes: []interface{}{
//   						&customAttributeProperty{
//   							objectIdentifier: jsii.String("objectIdentifier"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   					generationQualifier: jsii.String("generationQualifier"),
//   					givenName: jsii.String("givenName"),
//   					initials: jsii.String("initials"),
//   					locality: jsii.String("locality"),
//   					organization: jsii.String("organization"),
//   					organizationalUnit: jsii.String("organizationalUnit"),
//   					pseudonym: jsii.String("pseudonym"),
//   					serialNumber: jsii.String("serialNumber"),
//   					state: jsii.String("state"),
//   					surname: jsii.String("surname"),
//   					title: jsii.String("title"),
//   				},
//   				dnsName: jsii.String("dnsName"),
//   				ediPartyName: &ediPartyNameProperty{
//   					nameAssigner: jsii.String("nameAssigner"),
//   					partyName: jsii.String("partyName"),
//   				},
//   				ipAddress: jsii.String("ipAddress"),
//   				otherName: &otherNameProperty{
//   					typeId: jsii.String("typeId"),
//   					value: jsii.String("value"),
//   				},
//   				registeredId: jsii.String("registeredId"),
//   				rfc822Name: jsii.String("rfc822Name"),
//   				uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   			},
//   		},
//   	},
//   	subject: &subjectProperty{
//   		commonName: jsii.String("commonName"),
//   		country: jsii.String("country"),
//   		customAttributes: []interface{}{
//   			&customAttributeProperty{
//   				objectIdentifier: jsii.String("objectIdentifier"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   		generationQualifier: jsii.String("generationQualifier"),
//   		givenName: jsii.String("givenName"),
//   		initials: jsii.String("initials"),
//   		locality: jsii.String("locality"),
//   		organization: jsii.String("organization"),
//   		organizationalUnit: jsii.String("organizationalUnit"),
//   		pseudonym: jsii.String("pseudonym"),
//   		serialNumber: jsii.String("serialNumber"),
//   		state: jsii.String("state"),
//   		surname: jsii.String("surname"),
//   		title: jsii.String("title"),
//   	},
//   }
//
type CfnCertificate_ApiPassthroughProperty struct {
	// Specifies X.509 extension information for a certificate.
	Extensions interface{} `field:"optional" json:"extensions" yaml:"extensions"`
	// Contains information about the certificate subject.
	//
	// The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.
	Subject interface{} `field:"optional" json:"subject" yaml:"subject"`
}

// Defines the X.500 relative distinguished name (RDN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customAttributeProperty := &customAttributeProperty{
//   	objectIdentifier: jsii.String("objectIdentifier"),
//   	value: jsii.String("value"),
//   }
//
type CfnCertificate_CustomAttributeProperty struct {
	// Specifies the object identifier (OID) of the attribute type of the relative distinguished name (RDN).
	ObjectIdentifier *string `field:"required" json:"objectIdentifier" yaml:"objectIdentifier"`
	// Specifies the attribute value of relative distinguished name (RDN).
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Specifies the X.509 extension information for a certificate.
//
// Extensions present in `CustomExtensions` follow the `ApiPassthrough` [template rules](https://docs.aws.amazon.com/acm-pca/latest/userguide/UsingTemplates.html#template-order-of-operations) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customExtensionProperty := &customExtensionProperty{
//   	objectIdentifier: jsii.String("objectIdentifier"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	critical: jsii.Boolean(false),
//   }
//
type CfnCertificate_CustomExtensionProperty struct {
	// Specifies the object identifier (OID) of the X.509 extension. For more information, see the [Global OID reference database.](https://docs.aws.amazon.com/https://oidref.com/2.5.29).
	ObjectIdentifier *string `field:"required" json:"objectIdentifier" yaml:"objectIdentifier"`
	// Specifies the base64-encoded value of the X.509 extension.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Specifies the critical flag of the X.509 extension.
	Critical interface{} `field:"optional" json:"critical" yaml:"critical"`
}

// Describes an Electronic Data Interchange (EDI) entity as described in as defined in [Subject Alternative Name](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280) in RFC 5280.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ediPartyNameProperty := &ediPartyNameProperty{
//   	nameAssigner: jsii.String("nameAssigner"),
//   	partyName: jsii.String("partyName"),
//   }
//
type CfnCertificate_EdiPartyNameProperty struct {
	// Specifies the name assigner.
	NameAssigner *string `field:"required" json:"nameAssigner" yaml:"nameAssigner"`
	// Specifies the party name.
	PartyName *string `field:"required" json:"partyName" yaml:"partyName"`
}

// Specifies additional purposes for which the certified public key may be used other than basic purposes indicated in the `KeyUsage` extension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extendedKeyUsageProperty := &extendedKeyUsageProperty{
//   	extendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   	extendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   }
//
type CfnCertificate_ExtendedKeyUsageProperty struct {
	// Specifies a custom `ExtendedKeyUsage` with an object identifier (OID).
	ExtendedKeyUsageObjectIdentifier *string `field:"optional" json:"extendedKeyUsageObjectIdentifier" yaml:"extendedKeyUsageObjectIdentifier"`
	// Specifies a standard `ExtendedKeyUsage` as defined as in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12) .
	ExtendedKeyUsageType *string `field:"optional" json:"extendedKeyUsageType" yaml:"extendedKeyUsageType"`
}

// Contains X.509 extension information for a certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extensionsProperty := &extensionsProperty{
//   	certificatePolicies: []interface{}{
//   		&policyInformationProperty{
//   			certPolicyId: jsii.String("certPolicyId"),
//
//   			// the properties below are optional
//   			policyQualifiers: []interface{}{
//   				&policyQualifierInfoProperty{
//   					policyQualifierId: jsii.String("policyQualifierId"),
//   					qualifier: &qualifierProperty{
//   						cpsUri: jsii.String("cpsUri"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	customExtensions: []interface{}{
//   		&customExtensionProperty{
//   			objectIdentifier: jsii.String("objectIdentifier"),
//   			value: jsii.String("value"),
//
//   			// the properties below are optional
//   			critical: jsii.Boolean(false),
//   		},
//   	},
//   	extendedKeyUsage: []interface{}{
//   		&extendedKeyUsageProperty{
//   			extendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   			extendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   		},
//   	},
//   	keyUsage: &keyUsageProperty{
//   		crlSign: jsii.Boolean(false),
//   		dataEncipherment: jsii.Boolean(false),
//   		decipherOnly: jsii.Boolean(false),
//   		digitalSignature: jsii.Boolean(false),
//   		encipherOnly: jsii.Boolean(false),
//   		keyAgreement: jsii.Boolean(false),
//   		keyCertSign: jsii.Boolean(false),
//   		keyEncipherment: jsii.Boolean(false),
//   		nonRepudiation: jsii.Boolean(false),
//   	},
//   	subjectAlternativeNames: []interface{}{
//   		&generalNameProperty{
//   			directoryName: &subjectProperty{
//   				commonName: jsii.String("commonName"),
//   				country: jsii.String("country"),
//   				customAttributes: []interface{}{
//   					&customAttributeProperty{
//   						objectIdentifier: jsii.String("objectIdentifier"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   				generationQualifier: jsii.String("generationQualifier"),
//   				givenName: jsii.String("givenName"),
//   				initials: jsii.String("initials"),
//   				locality: jsii.String("locality"),
//   				organization: jsii.String("organization"),
//   				organizationalUnit: jsii.String("organizationalUnit"),
//   				pseudonym: jsii.String("pseudonym"),
//   				serialNumber: jsii.String("serialNumber"),
//   				state: jsii.String("state"),
//   				surname: jsii.String("surname"),
//   				title: jsii.String("title"),
//   			},
//   			dnsName: jsii.String("dnsName"),
//   			ediPartyName: &ediPartyNameProperty{
//   				nameAssigner: jsii.String("nameAssigner"),
//   				partyName: jsii.String("partyName"),
//   			},
//   			ipAddress: jsii.String("ipAddress"),
//   			otherName: &otherNameProperty{
//   				typeId: jsii.String("typeId"),
//   				value: jsii.String("value"),
//   			},
//   			registeredId: jsii.String("registeredId"),
//   			rfc822Name: jsii.String("rfc822Name"),
//   			uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   		},
//   	},
//   }
//
type CfnCertificate_ExtensionsProperty struct {
	// Contains a sequence of one or more policy information terms, each of which consists of an object identifier (OID) and optional qualifiers.
	//
	// For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
	//
	// In an end-entity certificate, these terms indicate the policy under which the certificate was issued and the purposes for which it may be used. In a CA certificate, these terms limit the set of policies for certification paths that include this certificate.
	CertificatePolicies interface{} `field:"optional" json:"certificatePolicies" yaml:"certificatePolicies"`
	// Contains a sequence of one or more X.509 extensions, each of which consists of an object identifier (OID), a base64-encoded value, and the critical flag. For more information, see the [Global OID reference database.](https://docs.aws.amazon.com/https://oidref.com/2.5.29).
	//
	// > The OID value of a [CustomExtension](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CustomExtension.html) must not match the OID of a predefined extension.
	CustomExtensions interface{} `field:"optional" json:"customExtensions" yaml:"customExtensions"`
	// Specifies additional purposes for which the certified public key may be used other than basic purposes indicated in the `KeyUsage` extension.
	ExtendedKeyUsage interface{} `field:"optional" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// Defines one or more purposes for which the key contained in the certificate can be used.
	//
	// Default value for each option is false.
	KeyUsage interface{} `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// The subject alternative name extension allows identities to be bound to the subject of the certificate.
	//
	// These identities may be included in addition to or in place of the identity in the subject field of the certificate.
	SubjectAlternativeNames interface{} `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

// Describes an ASN.1 X.400 `GeneralName` as defined in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280) . Only one of the following naming options should be provided. Providing more than one option results in an `InvalidArgsException` error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   generalNameProperty := &generalNameProperty{
//   	directoryName: &subjectProperty{
//   		commonName: jsii.String("commonName"),
//   		country: jsii.String("country"),
//   		customAttributes: []interface{}{
//   			&customAttributeProperty{
//   				objectIdentifier: jsii.String("objectIdentifier"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   		generationQualifier: jsii.String("generationQualifier"),
//   		givenName: jsii.String("givenName"),
//   		initials: jsii.String("initials"),
//   		locality: jsii.String("locality"),
//   		organization: jsii.String("organization"),
//   		organizationalUnit: jsii.String("organizationalUnit"),
//   		pseudonym: jsii.String("pseudonym"),
//   		serialNumber: jsii.String("serialNumber"),
//   		state: jsii.String("state"),
//   		surname: jsii.String("surname"),
//   		title: jsii.String("title"),
//   	},
//   	dnsName: jsii.String("dnsName"),
//   	ediPartyName: &ediPartyNameProperty{
//   		nameAssigner: jsii.String("nameAssigner"),
//   		partyName: jsii.String("partyName"),
//   	},
//   	ipAddress: jsii.String("ipAddress"),
//   	otherName: &otherNameProperty{
//   		typeId: jsii.String("typeId"),
//   		value: jsii.String("value"),
//   	},
//   	registeredId: jsii.String("registeredId"),
//   	rfc822Name: jsii.String("rfc822Name"),
//   	uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   }
//
type CfnCertificate_GeneralNameProperty struct {
	// Contains information about the certificate subject.
	//
	// The certificate can be one issued by your private certificate authority (CA) or it can be your private CA certificate. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate. The DN must be unique for each entity, but your private CA can issue more than one certificate with the same DN to the same entity.
	DirectoryName interface{} `field:"optional" json:"directoryName" yaml:"directoryName"`
	// Represents `GeneralName` as a DNS name.
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// Represents `GeneralName` as an `EdiPartyName` object.
	EdiPartyName interface{} `field:"optional" json:"ediPartyName" yaml:"ediPartyName"`
	// Represents `GeneralName` as an IPv4 or IPv6 address.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Represents `GeneralName` using an `OtherName` object.
	OtherName interface{} `field:"optional" json:"otherName" yaml:"otherName"`
	// Represents `GeneralName` as an object identifier (OID).
	RegisteredId *string `field:"optional" json:"registeredId" yaml:"registeredId"`
	// Represents `GeneralName` as an [RFC 822](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc822) email address.
	Rfc822Name *string `field:"optional" json:"rfc822Name" yaml:"rfc822Name"`
	// Represents `GeneralName` as a URI.
	UniformResourceIdentifier *string `field:"optional" json:"uniformResourceIdentifier" yaml:"uniformResourceIdentifier"`
}

// Defines one or more purposes for which the key contained in the certificate can be used.
//
// Default value for each option is false.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyUsageProperty := &keyUsageProperty{
//   	crlSign: jsii.Boolean(false),
//   	dataEncipherment: jsii.Boolean(false),
//   	decipherOnly: jsii.Boolean(false),
//   	digitalSignature: jsii.Boolean(false),
//   	encipherOnly: jsii.Boolean(false),
//   	keyAgreement: jsii.Boolean(false),
//   	keyCertSign: jsii.Boolean(false),
//   	keyEncipherment: jsii.Boolean(false),
//   	nonRepudiation: jsii.Boolean(false),
//   }
//
type CfnCertificate_KeyUsageProperty struct {
	// Key can be used to sign CRLs.
	CrlSign interface{} `field:"optional" json:"crlSign" yaml:"crlSign"`
	// Key can be used to decipher data.
	DataEncipherment interface{} `field:"optional" json:"dataEncipherment" yaml:"dataEncipherment"`
	// Key can be used only to decipher data.
	DecipherOnly interface{} `field:"optional" json:"decipherOnly" yaml:"decipherOnly"`
	// Key can be used for digital signing.
	DigitalSignature interface{} `field:"optional" json:"digitalSignature" yaml:"digitalSignature"`
	// Key can be used only to encipher data.
	EncipherOnly interface{} `field:"optional" json:"encipherOnly" yaml:"encipherOnly"`
	// Key can be used in a key-agreement protocol.
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// Key can be used to sign certificates.
	KeyCertSign interface{} `field:"optional" json:"keyCertSign" yaml:"keyCertSign"`
	// Key can be used to encipher data.
	KeyEncipherment interface{} `field:"optional" json:"keyEncipherment" yaml:"keyEncipherment"`
	// Key can be used for non-repudiation.
	NonRepudiation interface{} `field:"optional" json:"nonRepudiation" yaml:"nonRepudiation"`
}

// Defines a custom ASN.1 X.400 `GeneralName` using an object identifier (OID) and value. The OID must satisfy the regular expression shown below. For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   otherNameProperty := &otherNameProperty{
//   	typeId: jsii.String("typeId"),
//   	value: jsii.String("value"),
//   }
//
type CfnCertificate_OtherNameProperty struct {
	// Specifies an OID.
	TypeId *string `field:"required" json:"typeId" yaml:"typeId"`
	// Specifies an OID value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Defines the X.509 `CertificatePolicies` extension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyInformationProperty := &policyInformationProperty{
//   	certPolicyId: jsii.String("certPolicyId"),
//
//   	// the properties below are optional
//   	policyQualifiers: []interface{}{
//   		&policyQualifierInfoProperty{
//   			policyQualifierId: jsii.String("policyQualifierId"),
//   			qualifier: &qualifierProperty{
//   				cpsUri: jsii.String("cpsUri"),
//   			},
//   		},
//   	},
//   }
//
type CfnCertificate_PolicyInformationProperty struct {
	// Specifies the object identifier (OID) of the certificate policy under which the certificate was issued.
	//
	// For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
	CertPolicyId *string `field:"required" json:"certPolicyId" yaml:"certPolicyId"`
	// Modifies the given `CertPolicyId` with a qualifier.
	//
	// ACM Private CA supports the certification practice statement (CPS) qualifier.
	PolicyQualifiers interface{} `field:"optional" json:"policyQualifiers" yaml:"policyQualifiers"`
}

// Modifies the `CertPolicyId` of a `PolicyInformation` object with a qualifier.
//
// ACM Private CA supports the certification practice statement (CPS) qualifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyQualifierInfoProperty := &policyQualifierInfoProperty{
//   	policyQualifierId: jsii.String("policyQualifierId"),
//   	qualifier: &qualifierProperty{
//   		cpsUri: jsii.String("cpsUri"),
//   	},
//   }
//
type CfnCertificate_PolicyQualifierInfoProperty struct {
	// Identifies the qualifier modifying a `CertPolicyId` .
	PolicyQualifierId *string `field:"required" json:"policyQualifierId" yaml:"policyQualifierId"`
	// Defines the qualifier type.
	//
	// ACM Private CA supports the use of a URI for a CPS qualifier in this field.
	Qualifier interface{} `field:"required" json:"qualifier" yaml:"qualifier"`
}

// Defines a `PolicyInformation` qualifier.
//
// ACM Private CA supports the [certification practice statement (CPS) qualifier](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.4) defined in RFC 5280.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   qualifierProperty := &qualifierProperty{
//   	cpsUri: jsii.String("cpsUri"),
//   }
//
type CfnCertificate_QualifierProperty struct {
	// Contains a pointer to a certification practice statement (CPS) published by the CA.
	CpsUri *string `field:"required" json:"cpsUri" yaml:"cpsUri"`
}

// Contains information about the certificate subject.
//
// The `Subject` field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The `Subject` must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subjectProperty := &subjectProperty{
//   	commonName: jsii.String("commonName"),
//   	country: jsii.String("country"),
//   	customAttributes: []interface{}{
//   		&customAttributeProperty{
//   			objectIdentifier: jsii.String("objectIdentifier"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   	generationQualifier: jsii.String("generationQualifier"),
//   	givenName: jsii.String("givenName"),
//   	initials: jsii.String("initials"),
//   	locality: jsii.String("locality"),
//   	organization: jsii.String("organization"),
//   	organizationalUnit: jsii.String("organizationalUnit"),
//   	pseudonym: jsii.String("pseudonym"),
//   	serialNumber: jsii.String("serialNumber"),
//   	state: jsii.String("state"),
//   	surname: jsii.String("surname"),
//   	title: jsii.String("title"),
//   }
//
type CfnCertificate_SubjectProperty struct {
	// For CA and end-entity certificates in a private PKI, the common name (CN) can be any string within the length limit.
	//
	// Note: In publicly trusted certificates, the common name must be a fully qualified domain name (FQDN) associated with the certificate subject.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Two-digit code that specifies the country in which the certificate subject located.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Contains a sequence of one or more X.500 relative distinguished names (RDNs), each of which consists of an object identifier (OID) and a value. For more information, see NIST’s definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
	//
	// > Custom attributes cannot be used in combination with standard attributes.
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Disambiguating information for the certificate subject.
	DistinguishedNameQualifier *string `field:"optional" json:"distinguishedNameQualifier" yaml:"distinguishedNameQualifier"`
	// Typically a qualifier appended to the name of an individual.
	//
	// Examples include Jr. for junior, Sr. for senior, and III for third.
	GenerationQualifier *string `field:"optional" json:"generationQualifier" yaml:"generationQualifier"`
	// First name.
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
	// Concatenation that typically contains the first letter of the *GivenName* , the first letter of the middle name if one exists, and the first letter of the *Surname* .
	Initials *string `field:"optional" json:"initials" yaml:"initials"`
	// The locality (such as a city or town) in which the certificate subject is located.
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Legal name of the organization with which the certificate subject is affiliated.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Typically a shortened version of a longer *GivenName* .
	//
	// For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
	Pseudonym *string `field:"optional" json:"pseudonym" yaml:"pseudonym"`
	// The certificate serial number.
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// State in which the subject of the certificate is located.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Family name.
	//
	// In the US and the UK, for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.
	Surname *string `field:"optional" json:"surname" yaml:"surname"`
	// A title such as Mr.
	//
	// or Ms., which is pre-pended to the name to refer formally to the certificate subject.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

// Length of time for which the certificate issued by your private certificate authority (CA), or by the private CA itself, is valid in days, months, or years.
//
// You can issue a certificate by calling the `IssueCertificate` operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   validityProperty := &validityProperty{
//   	type: jsii.String("type"),
//   	value: jsii.Number(123),
//   }
//
type CfnCertificate_ValidityProperty struct {
	// Specifies whether the `Value` parameter represents days, months, or years.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A long integer interpreted according to the value of `Type` , below.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

// A CloudFormation `AWS::ACMPCA::CertificateAuthority`.
//
// Use the `AWS::ACMPCA::CertificateAuthority` resource to create a private CA. Once the CA exists, you can use the `AWS::ACMPCA::Certificate` resource to issue a new CA certificate. Alternatively, you can issue a CA certificate using an on-premises CA, and then use the `AWS::ACMPCA::CertificateAuthorityActivation` resource to import the new CA certificate and activate the CA.
//
// > Before removing a `AWS::ACMPCA::CertificateAuthority` resource from the CloudFormation stack, disable the affected CA. Otherwise, the action will fail. You can disable the CA by removing its associated `AWS::ACMPCA::CertificateAuthorityActivation` resource from CloudFormation.
//
// Example:
//   cfnCertificateAuthority := acmpca.NewCfnCertificateAuthority(this, jsii.String("CA"), &cfnCertificateAuthorityProps{
//   	type: jsii.String("ROOT"),
//   	keyAlgorithm: jsii.String("RSA_2048"),
//   	signingAlgorithm: jsii.String("SHA256WITHRSA"),
//   	subject: &subjectProperty{
//   		country: jsii.String("US"),
//   		organization: jsii.String("string"),
//   		organizationalUnit: jsii.String("string"),
//   		distinguishedNameQualifier: jsii.String("string"),
//   		state: jsii.String("string"),
//   		commonName: jsii.String("123"),
//   		serialNumber: jsii.String("string"),
//   		locality: jsii.String("string"),
//   		title: jsii.String("string"),
//   		surname: jsii.String("string"),
//   		givenName: jsii.String("string"),
//   		initials: jsii.String("DG"),
//   		pseudonym: jsii.String("string"),
//   		generationQualifier: jsii.String("DBG"),
//   	},
//   })
//
type CfnCertificateAuthority interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) for the private CA that issued the certificate.
	AttrArn() *string
	// The Base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.
	AttrCertificateSigningRequest() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies information to be added to the extension section of the certificate signing request (CSR).
	CsrExtensions() interface{}
	SetCsrExtensions(val interface{})
	// Type of the public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
	//
	// When you create a subordinate CA, you must use a key algorithm supported by the parent CA.
	KeyAlgorithm() *string
	SetKeyAlgorithm(val *string)
	// Specifies a cryptographic key management compliance standard used for handling CA keys.
	//
	// Default: FIPS_140_2_LEVEL_3_OR_HIGHER
	//
	// *Note:* `FIPS_140_2_LEVEL_3_OR_HIGHER` is not supported in the following Regions:
	//
	// - ap-northeast-3
	// - ap-southeast-3
	//
	// When creating a CA in these Regions, you must provide `FIPS_140_2_LEVEL_2_OR_HIGHER` as the argument for `KeyStorageSecurityStandard` . Failure to do this results in an `InvalidArgsException` with the message, "A certificate authority cannot be created in this region with the specified security standard."
	KeyStorageSecurityStandard() *string
	SetKeyStorageSecurityStandard(val *string)
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
	// Information about the certificate revocation list (CRL) created and maintained by your private CA.
	//
	// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions. Your certificate authority can create and maintain a certificate revocation list (CRL). A CRL contains information about certificates that have been revoked.
	RevocationConfiguration() interface{}
	SetRevocationConfiguration(val interface{})
	// Name of the algorithm your private CA uses to sign certificate requests.
	//
	// This parameter should not be confused with the `SigningAlgorithm` parameter used to sign certificates when they are issued.
	SigningAlgorithm() *string
	SetSigningAlgorithm(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Structure that contains X.500 distinguished name information for your private CA.
	Subject() interface{}
	SetSubject(val interface{})
	// Key-value pairs that will be attached to the new private CA.
	//
	// You can associate up to 50 tags with a private CA. For information using tags with IAM to manage permissions, see [Controlling Access Using IAM Tags](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html) .
	Tags() awscdk.TagManager
	// Type of your private CA.
	Type() *string
	SetType(val *string)
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

func (j *jsiiProxy_CfnCertificateAuthority) KeyStorageSecurityStandard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorageSecurityStandard",
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

func (j *jsiiProxy_CfnCertificateAuthority) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnCertificateAuthority(scope constructs.Construct, id *string, props *CfnCertificateAuthorityProps) CfnCertificateAuthority {
	_init_.Initialize()

	j := jsiiProxy_CfnCertificateAuthority{}

	_jsii_.Create(
		"aws-cdk-lib.aws_acmpca.CfnCertificateAuthority",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ACMPCA::CertificateAuthority`.
func NewCfnCertificateAuthority_Override(c CfnCertificateAuthority, scope constructs.Construct, id *string, props *CfnCertificateAuthorityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_acmpca.CfnCertificateAuthority",
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

func (j *jsiiProxy_CfnCertificateAuthority) SetKeyStorageSecurityStandard(val *string) {
	_jsii_.Set(
		j,
		"keyStorageSecurityStandard",
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
func CfnCertificateAuthority_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CfnCertificateAuthority",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCertificateAuthority_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CfnCertificateAuthority",
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
func CfnCertificateAuthority_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CfnCertificateAuthority",
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
		"aws-cdk-lib.aws_acmpca.CfnCertificateAuthority",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCertificateAuthority) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCertificateAuthority) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCertificateAuthority) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCertificateAuthority) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCertificateAuthority) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCertificateAuthority) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCertificateAuthority) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnCertificateAuthority) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCertificateAuthority) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

func (c *jsiiProxy_CfnCertificateAuthority) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Provides access information used by the `authorityInfoAccess` and `subjectInfoAccess` extensions described in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessDescriptionProperty := &accessDescriptionProperty{
//   	accessLocation: &generalNameProperty{
//   		directoryName: &subjectProperty{
//   			commonName: jsii.String("commonName"),
//   			country: jsii.String("country"),
//   			customAttributes: []interface{}{
//   				&customAttributeProperty{
//   					objectIdentifier: jsii.String("objectIdentifier"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   			generationQualifier: jsii.String("generationQualifier"),
//   			givenName: jsii.String("givenName"),
//   			initials: jsii.String("initials"),
//   			locality: jsii.String("locality"),
//   			organization: jsii.String("organization"),
//   			organizationalUnit: jsii.String("organizationalUnit"),
//   			pseudonym: jsii.String("pseudonym"),
//   			serialNumber: jsii.String("serialNumber"),
//   			state: jsii.String("state"),
//   			surname: jsii.String("surname"),
//   			title: jsii.String("title"),
//   		},
//   		dnsName: jsii.String("dnsName"),
//   		ediPartyName: &ediPartyNameProperty{
//   			nameAssigner: jsii.String("nameAssigner"),
//   			partyName: jsii.String("partyName"),
//   		},
//   		ipAddress: jsii.String("ipAddress"),
//   		otherName: &otherNameProperty{
//   			typeId: jsii.String("typeId"),
//   			value: jsii.String("value"),
//   		},
//   		registeredId: jsii.String("registeredId"),
//   		rfc822Name: jsii.String("rfc822Name"),
//   		uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   	},
//   	accessMethod: &accessMethodProperty{
//   		accessMethodType: jsii.String("accessMethodType"),
//   		customObjectIdentifier: jsii.String("customObjectIdentifier"),
//   	},
//   }
//
type CfnCertificateAuthority_AccessDescriptionProperty struct {
	// The location of `AccessDescription` information.
	AccessLocation interface{} `field:"required" json:"accessLocation" yaml:"accessLocation"`
	// The type and format of `AccessDescription` information.
	AccessMethod interface{} `field:"required" json:"accessMethod" yaml:"accessMethod"`
}

// Describes the type and format of extension access.
//
// Only one of `CustomObjectIdentifier` or `AccessMethodType` may be provided. Providing both results in `InvalidArgsException` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessMethodProperty := &accessMethodProperty{
//   	accessMethodType: jsii.String("accessMethodType"),
//   	customObjectIdentifier: jsii.String("customObjectIdentifier"),
//   }
//
type CfnCertificateAuthority_AccessMethodProperty struct {
	// Specifies the `AccessMethod` .
	AccessMethodType *string `field:"optional" json:"accessMethodType" yaml:"accessMethodType"`
	// An object identifier (OID) specifying the `AccessMethod` .
	//
	// The OID must satisfy the regular expression shown below. For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
	CustomObjectIdentifier *string `field:"optional" json:"customObjectIdentifier" yaml:"customObjectIdentifier"`
}

// Contains configuration information for a certificate revocation list (CRL).
//
// Your private certificate authority (CA) creates base CRLs. Delta CRLs are not supported. You can enable CRLs for your new or an existing private CA by setting the *Enabled* parameter to `true` . Your private CA writes CRLs to an S3 bucket that you specify in the *S3BucketName* parameter. You can hide the name of your bucket by specifying a value for the *CustomCname* parameter. Your private CA copies the CNAME or the S3 bucket name to the *CRL Distribution Points* extension of each certificate it issues. Your S3 bucket policy must give write permission to ACM Private CA.
//
// ACM Private CA assets that are stored in Amazon S3 can be protected with encryption. For more information, see [Encrypting Your CRLs](https://docs.aws.amazon.com/acm-pca/latest/userguide/PcaCreateCa.html#crl-encryption) .
//
// Your private CA uses the value in the *ExpirationInDays* parameter to calculate the *nextUpdate* field in the CRL. The CRL is refreshed prior to a certificate's expiration date or when a certificate is revoked. When a certificate is revoked, it appears in the CRL until the certificate expires, and then in one additional CRL after expiration, and it always appears in the audit report.
//
// A CRL is typically updated approximately 30 minutes after a certificate is revoked. If for any reason a CRL update fails, ACM Private CA makes further attempts every 15 minutes.
//
// CRLs contain the following fields:
//
// - *Version* : The current version number defined in RFC 5280 is V2. The integer value is 0x1.
// - *Signature Algorithm* : The name of the algorithm used to sign the CRL.
// - *Issuer* : The X.500 distinguished name of your private CA that issued the CRL.
// - *Last Update* : The issue date and time of this CRL.
// - *Next Update* : The day and time by which the next CRL will be issued.
// - *Revoked Certificates* : List of revoked certificates. Each list item contains the following information.
//
// - *Serial Number* : The serial number, in hexadecimal format, of the revoked certificate.
// - *Revocation Date* : Date and time the certificate was revoked.
// - *CRL Entry Extensions* : Optional extensions for the CRL entry.
//
// - *X509v3 CRL Reason Code* : Reason the certificate was revoked.
// - *CRL Extensions* : Optional extensions for the CRL.
//
// - *X509v3 Authority Key Identifier* : Identifies the public key associated with the private key used to sign the certificate.
// - *X509v3 CRL Number:* : Decimal sequence number for the CRL.
// - *Signature Algorithm* : Algorithm used by your private CA to sign the CRL.
// - *Signature Value* : Signature computed over the CRL.
//
// Certificate revocation lists created by ACM Private CA are DER-encoded. You can use the following OpenSSL command to list a CRL.
//
// `openssl crl -inform DER -text -in *crl_path* -noout`
//
// For more information, see [Planning a certificate revocation list (CRL)](https://docs.aws.amazon.com/acm-pca/latest/userguide/crl-planning.html) in the *AWS Certificate Manager Private Certificate Authority (PCA) User Guide*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crlConfigurationProperty := &crlConfigurationProperty{
//   	customCname: jsii.String("customCname"),
//   	enabled: jsii.Boolean(false),
//   	expirationInDays: jsii.Number(123),
//   	s3BucketName: jsii.String("s3BucketName"),
//   	s3ObjectAcl: jsii.String("s3ObjectAcl"),
//   }
//
type CfnCertificateAuthority_CrlConfigurationProperty struct {
	// Name inserted into the certificate *CRL Distribution Points* extension that enables the use of an alias for the CRL distribution point.
	//
	// Use this value if you don't want the name of your S3 bucket to be public.
	CustomCname *string `field:"optional" json:"customCname" yaml:"customCname"`
	// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled.
	//
	// You can use this value to enable certificate revocation for a new CA when you call the `CreateCertificateAuthority` operation or for an existing CA when you call the `UpdateCertificateAuthority` operation.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Validity period of the CRL in days.
	ExpirationInDays *float64 `field:"optional" json:"expirationInDays" yaml:"expirationInDays"`
	// Name of the S3 bucket that contains the CRL.
	//
	// If you do not provide a value for the *CustomCname* argument, the name of your S3 bucket is placed into the *CRL Distribution Points* extension of the issued certificate. You can change the name of your bucket by calling the [UpdateCertificateAuthority](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_UpdateCertificateAuthority.html) operation. You must specify a [bucket policy](https://docs.aws.amazon.com/acm-pca/latest/userguide/PcaCreateCa.html#s3-policies) that allows ACM Private CA to write the CRL to your bucket.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// Determines whether the CRL will be publicly readable or privately held in the CRL Amazon S3 bucket.
	//
	// If you choose PUBLIC_READ, the CRL will be accessible over the public internet. If you choose BUCKET_OWNER_FULL_CONTROL, only the owner of the CRL S3 bucket can access the CRL, and your PKI clients may need an alternative method of access.
	//
	// If no value is specified, the default is PUBLIC_READ.
	//
	// > This default can cause CA creation to fail in some circumstances. If you have enabled the Block Public Access (BPA) feature in your S3 account, then you must specify the value of this parameter as `BUCKET_OWNER_FULL_CONTROL` , and not doing so results in an error. If you have disabled BPA in S3, then you can specify either `BUCKET_OWNER_FULL_CONTROL` or `PUBLIC_READ` as the value.
	//
	// For more information, see [Blocking public access to the S3 bucket](https://docs.aws.amazon.com/acm-pca/latest/userguide/PcaCreateCa.html#s3-bpa) .
	S3ObjectAcl *string `field:"optional" json:"s3ObjectAcl" yaml:"s3ObjectAcl"`
}

// Describes the certificate extensions to be added to the certificate signing request (CSR).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csrExtensionsProperty := &csrExtensionsProperty{
//   	keyUsage: &keyUsageProperty{
//   		crlSign: jsii.Boolean(false),
//   		dataEncipherment: jsii.Boolean(false),
//   		decipherOnly: jsii.Boolean(false),
//   		digitalSignature: jsii.Boolean(false),
//   		encipherOnly: jsii.Boolean(false),
//   		keyAgreement: jsii.Boolean(false),
//   		keyCertSign: jsii.Boolean(false),
//   		keyEncipherment: jsii.Boolean(false),
//   		nonRepudiation: jsii.Boolean(false),
//   	},
//   	subjectInformationAccess: []interface{}{
//   		&accessDescriptionProperty{
//   			accessLocation: &generalNameProperty{
//   				directoryName: &subjectProperty{
//   					commonName: jsii.String("commonName"),
//   					country: jsii.String("country"),
//   					customAttributes: []interface{}{
//   						&customAttributeProperty{
//   							objectIdentifier: jsii.String("objectIdentifier"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   					generationQualifier: jsii.String("generationQualifier"),
//   					givenName: jsii.String("givenName"),
//   					initials: jsii.String("initials"),
//   					locality: jsii.String("locality"),
//   					organization: jsii.String("organization"),
//   					organizationalUnit: jsii.String("organizationalUnit"),
//   					pseudonym: jsii.String("pseudonym"),
//   					serialNumber: jsii.String("serialNumber"),
//   					state: jsii.String("state"),
//   					surname: jsii.String("surname"),
//   					title: jsii.String("title"),
//   				},
//   				dnsName: jsii.String("dnsName"),
//   				ediPartyName: &ediPartyNameProperty{
//   					nameAssigner: jsii.String("nameAssigner"),
//   					partyName: jsii.String("partyName"),
//   				},
//   				ipAddress: jsii.String("ipAddress"),
//   				otherName: &otherNameProperty{
//   					typeId: jsii.String("typeId"),
//   					value: jsii.String("value"),
//   				},
//   				registeredId: jsii.String("registeredId"),
//   				rfc822Name: jsii.String("rfc822Name"),
//   				uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   			},
//   			accessMethod: &accessMethodProperty{
//   				accessMethodType: jsii.String("accessMethodType"),
//   				customObjectIdentifier: jsii.String("customObjectIdentifier"),
//   			},
//   		},
//   	},
//   }
//
type CfnCertificateAuthority_CsrExtensionsProperty struct {
	// Indicates the purpose of the certificate and of the key contained in the certificate.
	KeyUsage interface{} `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// For CA certificates, provides a path to additional information pertaining to the CA, such as revocation and policy.
	//
	// For more information, see [Subject Information Access](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.2.2) in RFC 5280.
	SubjectInformationAccess interface{} `field:"optional" json:"subjectInformationAccess" yaml:"subjectInformationAccess"`
}

// Defines the X.500 relative distinguished name (RDN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customAttributeProperty := &customAttributeProperty{
//   	objectIdentifier: jsii.String("objectIdentifier"),
//   	value: jsii.String("value"),
//   }
//
type CfnCertificateAuthority_CustomAttributeProperty struct {
	// Specifies the object identifier (OID) of the attribute type of the relative distinguished name (RDN).
	ObjectIdentifier *string `field:"required" json:"objectIdentifier" yaml:"objectIdentifier"`
	// Specifies the attribute value of relative distinguished name (RDN).
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Describes an Electronic Data Interchange (EDI) entity as described in as defined in [Subject Alternative Name](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280) in RFC 5280.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ediPartyNameProperty := &ediPartyNameProperty{
//   	nameAssigner: jsii.String("nameAssigner"),
//   	partyName: jsii.String("partyName"),
//   }
//
type CfnCertificateAuthority_EdiPartyNameProperty struct {
	// Specifies the name assigner.
	NameAssigner *string `field:"required" json:"nameAssigner" yaml:"nameAssigner"`
	// Specifies the party name.
	PartyName *string `field:"required" json:"partyName" yaml:"partyName"`
}

// Describes an ASN.1 X.400 `GeneralName` as defined in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280) . Only one of the following naming options should be provided. Providing more than one option results in an `InvalidArgsException` error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   generalNameProperty := &generalNameProperty{
//   	directoryName: &subjectProperty{
//   		commonName: jsii.String("commonName"),
//   		country: jsii.String("country"),
//   		customAttributes: []interface{}{
//   			&customAttributeProperty{
//   				objectIdentifier: jsii.String("objectIdentifier"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   		generationQualifier: jsii.String("generationQualifier"),
//   		givenName: jsii.String("givenName"),
//   		initials: jsii.String("initials"),
//   		locality: jsii.String("locality"),
//   		organization: jsii.String("organization"),
//   		organizationalUnit: jsii.String("organizationalUnit"),
//   		pseudonym: jsii.String("pseudonym"),
//   		serialNumber: jsii.String("serialNumber"),
//   		state: jsii.String("state"),
//   		surname: jsii.String("surname"),
//   		title: jsii.String("title"),
//   	},
//   	dnsName: jsii.String("dnsName"),
//   	ediPartyName: &ediPartyNameProperty{
//   		nameAssigner: jsii.String("nameAssigner"),
//   		partyName: jsii.String("partyName"),
//   	},
//   	ipAddress: jsii.String("ipAddress"),
//   	otherName: &otherNameProperty{
//   		typeId: jsii.String("typeId"),
//   		value: jsii.String("value"),
//   	},
//   	registeredId: jsii.String("registeredId"),
//   	rfc822Name: jsii.String("rfc822Name"),
//   	uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   }
//
type CfnCertificateAuthority_GeneralNameProperty struct {
	// Contains information about the certificate subject.
	//
	// The certificate can be one issued by your private certificate authority (CA) or it can be your private CA certificate. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate. The DN must be unique for each entity, but your private CA can issue more than one certificate with the same DN to the same entity.
	DirectoryName interface{} `field:"optional" json:"directoryName" yaml:"directoryName"`
	// Represents `GeneralName` as a DNS name.
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// Represents `GeneralName` as an `EdiPartyName` object.
	EdiPartyName interface{} `field:"optional" json:"ediPartyName" yaml:"ediPartyName"`
	// Represents `GeneralName` as an IPv4 or IPv6 address.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Represents `GeneralName` using an `OtherName` object.
	OtherName interface{} `field:"optional" json:"otherName" yaml:"otherName"`
	// Represents `GeneralName` as an object identifier (OID).
	RegisteredId *string `field:"optional" json:"registeredId" yaml:"registeredId"`
	// Represents `GeneralName` as an [RFC 822](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc822) email address.
	Rfc822Name *string `field:"optional" json:"rfc822Name" yaml:"rfc822Name"`
	// Represents `GeneralName` as a URI.
	UniformResourceIdentifier *string `field:"optional" json:"uniformResourceIdentifier" yaml:"uniformResourceIdentifier"`
}

// Defines one or more purposes for which the key contained in the certificate can be used.
//
// Default value for each option is false.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyUsageProperty := &keyUsageProperty{
//   	crlSign: jsii.Boolean(false),
//   	dataEncipherment: jsii.Boolean(false),
//   	decipherOnly: jsii.Boolean(false),
//   	digitalSignature: jsii.Boolean(false),
//   	encipherOnly: jsii.Boolean(false),
//   	keyAgreement: jsii.Boolean(false),
//   	keyCertSign: jsii.Boolean(false),
//   	keyEncipherment: jsii.Boolean(false),
//   	nonRepudiation: jsii.Boolean(false),
//   }
//
type CfnCertificateAuthority_KeyUsageProperty struct {
	// Key can be used to sign CRLs.
	CrlSign interface{} `field:"optional" json:"crlSign" yaml:"crlSign"`
	// Key can be used to decipher data.
	DataEncipherment interface{} `field:"optional" json:"dataEncipherment" yaml:"dataEncipherment"`
	// Key can be used only to decipher data.
	DecipherOnly interface{} `field:"optional" json:"decipherOnly" yaml:"decipherOnly"`
	// Key can be used for digital signing.
	DigitalSignature interface{} `field:"optional" json:"digitalSignature" yaml:"digitalSignature"`
	// Key can be used only to encipher data.
	EncipherOnly interface{} `field:"optional" json:"encipherOnly" yaml:"encipherOnly"`
	// Key can be used in a key-agreement protocol.
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// Key can be used to sign certificates.
	KeyCertSign interface{} `field:"optional" json:"keyCertSign" yaml:"keyCertSign"`
	// Key can be used to encipher data.
	KeyEncipherment interface{} `field:"optional" json:"keyEncipherment" yaml:"keyEncipherment"`
	// Key can be used for non-repudiation.
	NonRepudiation interface{} `field:"optional" json:"nonRepudiation" yaml:"nonRepudiation"`
}

// Contains information to enable and configure Online Certificate Status Protocol (OCSP) for validating certificate revocation status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ocspConfigurationProperty := &ocspConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   	ocspCustomCname: jsii.String("ocspCustomCname"),
//   }
//
type CfnCertificateAuthority_OcspConfigurationProperty struct {
	// Flag enabling use of the Online Certificate Status Protocol (OCSP) for validating certificate revocation status.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// By default, ACM Private CA injects an Amazon domain into certificates being validated by the Online Certificate Status Protocol (OCSP).
	//
	// A customer can alternatively use this object to define a CNAME specifying a customized OCSP domain.
	//
	// Note: The value of the CNAME must not include a protocol prefix such as "http://" or "https://".
	OcspCustomCname *string `field:"optional" json:"ocspCustomCname" yaml:"ocspCustomCname"`
}

// Defines a custom ASN.1 X.400 `GeneralName` using an object identifier (OID) and value. The OID must satisfy the regular expression shown below. For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   otherNameProperty := &otherNameProperty{
//   	typeId: jsii.String("typeId"),
//   	value: jsii.String("value"),
//   }
//
type CfnCertificateAuthority_OtherNameProperty struct {
	// Specifies an OID.
	TypeId *string `field:"required" json:"typeId" yaml:"typeId"`
	// Specifies an OID value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
//
// Your private certificate authority (CA) can configure Online Certificate Status Protocol (OCSP) support and/or maintain a certificate revocation list (CRL). OCSP returns validation information about certificates as requested by clients, and a CRL contains an updated list of certificates revoked by your CA. For more information, see [RevokeCertificate](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_RevokeCertificate.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   revocationConfigurationProperty := &revocationConfigurationProperty{
//   	crlConfiguration: &crlConfigurationProperty{
//   		customCname: jsii.String("customCname"),
//   		enabled: jsii.Boolean(false),
//   		expirationInDays: jsii.Number(123),
//   		s3BucketName: jsii.String("s3BucketName"),
//   		s3ObjectAcl: jsii.String("s3ObjectAcl"),
//   	},
//   	ocspConfiguration: &ocspConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   		ocspCustomCname: jsii.String("ocspCustomCname"),
//   	},
//   }
//
type CfnCertificateAuthority_RevocationConfigurationProperty struct {
	// Configuration of the certificate revocation list (CRL), if any, maintained by your private CA.
	CrlConfiguration interface{} `field:"optional" json:"crlConfiguration" yaml:"crlConfiguration"`
	// Configuration of Online Certificate Status Protocol (OCSP) support, if any, maintained by your private CA.
	OcspConfiguration interface{} `field:"optional" json:"ocspConfiguration" yaml:"ocspConfiguration"`
}

// ASN1 subject for the certificate authority.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subjectProperty := &subjectProperty{
//   	commonName: jsii.String("commonName"),
//   	country: jsii.String("country"),
//   	customAttributes: []interface{}{
//   		&customAttributeProperty{
//   			objectIdentifier: jsii.String("objectIdentifier"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   	generationQualifier: jsii.String("generationQualifier"),
//   	givenName: jsii.String("givenName"),
//   	initials: jsii.String("initials"),
//   	locality: jsii.String("locality"),
//   	organization: jsii.String("organization"),
//   	organizationalUnit: jsii.String("organizationalUnit"),
//   	pseudonym: jsii.String("pseudonym"),
//   	serialNumber: jsii.String("serialNumber"),
//   	state: jsii.String("state"),
//   	surname: jsii.String("surname"),
//   	title: jsii.String("title"),
//   }
//
type CfnCertificateAuthority_SubjectProperty struct {
	// Fully qualified domain name (FQDN) associated with the certificate subject.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Two-digit code that specifies the country in which the certificate subject located.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Contains a sequence of one or more X.500 relative distinguished names (RDNs), each of which consists of an object identifier (OID) and a value. For more information, see NIST’s definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
	//
	// > Custom attributes cannot be used in combination with standard attributes.
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Disambiguating information for the certificate subject.
	DistinguishedNameQualifier *string `field:"optional" json:"distinguishedNameQualifier" yaml:"distinguishedNameQualifier"`
	// Typically a qualifier appended to the name of an individual.
	//
	// Examples include Jr. for junior, Sr. for senior, and III for third.
	GenerationQualifier *string `field:"optional" json:"generationQualifier" yaml:"generationQualifier"`
	// First name.
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
	// Concatenation that typically contains the first letter of the GivenName, the first letter of the middle name if one exists, and the first letter of the SurName.
	Initials *string `field:"optional" json:"initials" yaml:"initials"`
	// The locality (such as a city or town) in which the certificate subject is located.
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Legal name of the organization with which the certificate subject is affiliated.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Typically a shortened version of a longer GivenName.
	//
	// For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
	Pseudonym *string `field:"optional" json:"pseudonym" yaml:"pseudonym"`
	// The certificate serial number.
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// State in which the subject of the certificate is located.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Family name.
	Surname *string `field:"optional" json:"surname" yaml:"surname"`
	// A personal title such as Mr.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

// A CloudFormation `AWS::ACMPCA::CertificateAuthorityActivation`.
//
// The `AWS::ACMPCA::CertificateAuthorityActivation` resource creates and installs a CA certificate on a CA. If no status is specified, the `AWS::ACMPCA::CertificateAuthorityActivation` resource status defaults to ACTIVE. Once the CA has a CA certificate installed, you can use the resource to toggle the CA status field between `ACTIVE` and `DISABLED` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateAuthorityActivation := awscdk.Aws_acmpca.NewCfnCertificateAuthorityActivation(this, jsii.String("MyCfnCertificateAuthorityActivation"), &cfnCertificateAuthorityActivationProps{
//   	certificate: jsii.String("certificate"),
//   	certificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//
//   	// the properties below are optional
//   	certificateChain: jsii.String("certificateChain"),
//   	status: jsii.String("status"),
//   })
//
type CfnCertificateAuthorityActivation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The complete Base64 PEM-encoded certificate chain, including the certificate authority certificate.
	AttrCompleteCertificateChain() *string
	// The Base64 PEM-encoded certificate authority certificate.
	Certificate() *string
	SetCertificate(val *string)
	// The Amazon Resource Name (ARN) of your private CA.
	CertificateAuthorityArn() *string
	SetCertificateAuthorityArn(val *string)
	// The Base64 PEM-encoded certificate chain that chains up to the root CA certificate that you used to sign your private CA certificate.
	CertificateChain() *string
	SetCertificateChain(val *string)
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
	// Status of your private CA.
	Status() *string
	SetStatus(val *string)
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

func (j *jsiiProxy_CfnCertificateAuthorityActivation) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnCertificateAuthorityActivation(scope constructs.Construct, id *string, props *CfnCertificateAuthorityActivationProps) CfnCertificateAuthorityActivation {
	_init_.Initialize()

	j := jsiiProxy_CfnCertificateAuthorityActivation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_acmpca.CfnCertificateAuthorityActivation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ACMPCA::CertificateAuthorityActivation`.
func NewCfnCertificateAuthorityActivation_Override(c CfnCertificateAuthorityActivation, scope constructs.Construct, id *string, props *CfnCertificateAuthorityActivationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_acmpca.CfnCertificateAuthorityActivation",
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
func CfnCertificateAuthorityActivation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CfnCertificateAuthorityActivation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCertificateAuthorityActivation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CfnCertificateAuthorityActivation",
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
func CfnCertificateAuthorityActivation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CfnCertificateAuthorityActivation",
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
		"aws-cdk-lib.aws_acmpca.CfnCertificateAuthorityActivation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCertificateAuthorityActivation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCertificateAuthorityActivation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCertificateAuthorityActivation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCertificateAuthorityActivation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCertificateAuthorityActivation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCertificateAuthorityActivation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCertificateAuthorityActivation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnCertificateAuthorityActivation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCertificateAuthorityActivation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

func (c *jsiiProxy_CfnCertificateAuthorityActivation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnCertificateAuthorityActivation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateAuthorityActivationProps := &cfnCertificateAuthorityActivationProps{
//   	certificate: jsii.String("certificate"),
//   	certificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//
//   	// the properties below are optional
//   	certificateChain: jsii.String("certificateChain"),
//   	status: jsii.String("status"),
//   }
//
type CfnCertificateAuthorityActivationProps struct {
	// The Base64 PEM-encoded certificate authority certificate.
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// The Amazon Resource Name (ARN) of your private CA.
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The Base64 PEM-encoded certificate chain that chains up to the root CA certificate that you used to sign your private CA certificate.
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// Status of your private CA.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

// Properties for defining a `CfnCertificateAuthority`.
//
// Example:
//   cfnCertificateAuthority := acmpca.NewCfnCertificateAuthority(this, jsii.String("CA"), &cfnCertificateAuthorityProps{
//   	type: jsii.String("ROOT"),
//   	keyAlgorithm: jsii.String("RSA_2048"),
//   	signingAlgorithm: jsii.String("SHA256WITHRSA"),
//   	subject: &subjectProperty{
//   		country: jsii.String("US"),
//   		organization: jsii.String("string"),
//   		organizationalUnit: jsii.String("string"),
//   		distinguishedNameQualifier: jsii.String("string"),
//   		state: jsii.String("string"),
//   		commonName: jsii.String("123"),
//   		serialNumber: jsii.String("string"),
//   		locality: jsii.String("string"),
//   		title: jsii.String("string"),
//   		surname: jsii.String("string"),
//   		givenName: jsii.String("string"),
//   		initials: jsii.String("DG"),
//   		pseudonym: jsii.String("string"),
//   		generationQualifier: jsii.String("DBG"),
//   	},
//   })
//
type CfnCertificateAuthorityProps struct {
	// Type of the public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
	//
	// When you create a subordinate CA, you must use a key algorithm supported by the parent CA.
	KeyAlgorithm *string `field:"required" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Name of the algorithm your private CA uses to sign certificate requests.
	//
	// This parameter should not be confused with the `SigningAlgorithm` parameter used to sign certificates when they are issued.
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// Structure that contains X.500 distinguished name information for your private CA.
	Subject interface{} `field:"required" json:"subject" yaml:"subject"`
	// Type of your private CA.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies information to be added to the extension section of the certificate signing request (CSR).
	CsrExtensions interface{} `field:"optional" json:"csrExtensions" yaml:"csrExtensions"`
	// Specifies a cryptographic key management compliance standard used for handling CA keys.
	//
	// Default: FIPS_140_2_LEVEL_3_OR_HIGHER
	//
	// *Note:* `FIPS_140_2_LEVEL_3_OR_HIGHER` is not supported in the following Regions:
	//
	// - ap-northeast-3
	// - ap-southeast-3
	//
	// When creating a CA in these Regions, you must provide `FIPS_140_2_LEVEL_2_OR_HIGHER` as the argument for `KeyStorageSecurityStandard` . Failure to do this results in an `InvalidArgsException` with the message, "A certificate authority cannot be created in this region with the specified security standard."
	KeyStorageSecurityStandard *string `field:"optional" json:"keyStorageSecurityStandard" yaml:"keyStorageSecurityStandard"`
	// Information about the certificate revocation list (CRL) created and maintained by your private CA.
	//
	// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions. Your certificate authority can create and maintain a certificate revocation list (CRL). A CRL contains information about certificates that have been revoked.
	RevocationConfiguration interface{} `field:"optional" json:"revocationConfiguration" yaml:"revocationConfiguration"`
	// Key-value pairs that will be attached to the new private CA.
	//
	// You can associate up to 50 tags with a private CA. For information using tags with IAM to manage permissions, see [Controlling Access Using IAM Tags](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProps := &cfnCertificateProps{
//   	certificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	certificateSigningRequest: jsii.String("certificateSigningRequest"),
//   	signingAlgorithm: jsii.String("signingAlgorithm"),
//   	validity: &validityProperty{
//   		type: jsii.String("type"),
//   		value: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	apiPassthrough: &apiPassthroughProperty{
//   		extensions: &extensionsProperty{
//   			certificatePolicies: []interface{}{
//   				&policyInformationProperty{
//   					certPolicyId: jsii.String("certPolicyId"),
//
//   					// the properties below are optional
//   					policyQualifiers: []interface{}{
//   						&policyQualifierInfoProperty{
//   							policyQualifierId: jsii.String("policyQualifierId"),
//   							qualifier: &qualifierProperty{
//   								cpsUri: jsii.String("cpsUri"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			customExtensions: []interface{}{
//   				&customExtensionProperty{
//   					objectIdentifier: jsii.String("objectIdentifier"),
//   					value: jsii.String("value"),
//
//   					// the properties below are optional
//   					critical: jsii.Boolean(false),
//   				},
//   			},
//   			extendedKeyUsage: []interface{}{
//   				&extendedKeyUsageProperty{
//   					extendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   					extendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   				},
//   			},
//   			keyUsage: &keyUsageProperty{
//   				crlSign: jsii.Boolean(false),
//   				dataEncipherment: jsii.Boolean(false),
//   				decipherOnly: jsii.Boolean(false),
//   				digitalSignature: jsii.Boolean(false),
//   				encipherOnly: jsii.Boolean(false),
//   				keyAgreement: jsii.Boolean(false),
//   				keyCertSign: jsii.Boolean(false),
//   				keyEncipherment: jsii.Boolean(false),
//   				nonRepudiation: jsii.Boolean(false),
//   			},
//   			subjectAlternativeNames: []interface{}{
//   				&generalNameProperty{
//   					directoryName: &subjectProperty{
//   						commonName: jsii.String("commonName"),
//   						country: jsii.String("country"),
//   						customAttributes: []interface{}{
//   							&customAttributeProperty{
//   								objectIdentifier: jsii.String("objectIdentifier"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   						generationQualifier: jsii.String("generationQualifier"),
//   						givenName: jsii.String("givenName"),
//   						initials: jsii.String("initials"),
//   						locality: jsii.String("locality"),
//   						organization: jsii.String("organization"),
//   						organizationalUnit: jsii.String("organizationalUnit"),
//   						pseudonym: jsii.String("pseudonym"),
//   						serialNumber: jsii.String("serialNumber"),
//   						state: jsii.String("state"),
//   						surname: jsii.String("surname"),
//   						title: jsii.String("title"),
//   					},
//   					dnsName: jsii.String("dnsName"),
//   					ediPartyName: &ediPartyNameProperty{
//   						nameAssigner: jsii.String("nameAssigner"),
//   						partyName: jsii.String("partyName"),
//   					},
//   					ipAddress: jsii.String("ipAddress"),
//   					otherName: &otherNameProperty{
//   						typeId: jsii.String("typeId"),
//   						value: jsii.String("value"),
//   					},
//   					registeredId: jsii.String("registeredId"),
//   					rfc822Name: jsii.String("rfc822Name"),
//   					uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   				},
//   			},
//   		},
//   		subject: &subjectProperty{
//   			commonName: jsii.String("commonName"),
//   			country: jsii.String("country"),
//   			customAttributes: []interface{}{
//   				&customAttributeProperty{
//   					objectIdentifier: jsii.String("objectIdentifier"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   			generationQualifier: jsii.String("generationQualifier"),
//   			givenName: jsii.String("givenName"),
//   			initials: jsii.String("initials"),
//   			locality: jsii.String("locality"),
//   			organization: jsii.String("organization"),
//   			organizationalUnit: jsii.String("organizationalUnit"),
//   			pseudonym: jsii.String("pseudonym"),
//   			serialNumber: jsii.String("serialNumber"),
//   			state: jsii.String("state"),
//   			surname: jsii.String("surname"),
//   			title: jsii.String("title"),
//   		},
//   	},
//   	templateArn: jsii.String("templateArn"),
//   	validityNotBefore: &validityProperty{
//   		type: jsii.String("type"),
//   		value: jsii.Number(123),
//   	},
//   }
//
type CfnCertificateProps struct {
	// The Amazon Resource Name (ARN) for the private CA issues the certificate.
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The certificate signing request (CSR) for the certificate.
	CertificateSigningRequest *string `field:"required" json:"certificateSigningRequest" yaml:"certificateSigningRequest"`
	// The name of the algorithm that will be used to sign the certificate to be issued.
	//
	// This parameter should not be confused with the `SigningAlgorithm` parameter used to sign a CSR in the `CreateCertificateAuthority` action.
	//
	// > The specified signing algorithm family (RSA or ECDSA) must match the algorithm family of the CA's secret key.
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// The period of time during which the certificate will be valid.
	Validity interface{} `field:"required" json:"validity" yaml:"validity"`
	// Specifies X.509 certificate information to be included in the issued certificate. An `APIPassthrough` or `APICSRPassthrough` template variant must be selected, or else this parameter is ignored.
	ApiPassthrough interface{} `field:"optional" json:"apiPassthrough" yaml:"apiPassthrough"`
	// Specifies a custom configuration template to use when issuing a certificate.
	//
	// If this parameter is not provided, ACM Private CA defaults to the `EndEntityCertificate/V1` template. For more information about ACM Private CA templates, see [Using Templates](https://docs.aws.amazon.com/acm-pca/latest/userguide/UsingTemplates.html) .
	TemplateArn *string `field:"optional" json:"templateArn" yaml:"templateArn"`
	// Information describing the start of the validity period of the certificate.
	//
	// This parameter sets the “Not Before" date for the certificate.
	//
	// By default, when issuing a certificate, ACM Private CA sets the "Not Before" date to the issuance time minus 60 minutes. This compensates for clock inconsistencies across computer systems. The `ValidityNotBefore` parameter can be used to customize the “Not Before” value.
	//
	// Unlike the `Validity` parameter, the `ValidityNotBefore` parameter is optional.
	//
	// The `ValidityNotBefore` value is expressed as an explicit date and time, using the `Validity` type value `ABSOLUTE` .
	ValidityNotBefore interface{} `field:"optional" json:"validityNotBefore" yaml:"validityNotBefore"`
}

// A CloudFormation `AWS::ACMPCA::Permission`.
//
// Grants permissions to the AWS Certificate Manager (ACM) service principal ( `acm.amazonaws.com` ) to perform [IssueCertificate](https://docs.aws.amazon.com/latest/APIReference/API_IssueCertificate.html) , [GetCertificate](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_GetCertificate.html) , and [ListPermissions](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_ListPermissions.html) actions on a CA. These actions are needed for the ACM principal to renew private PKI certificates requested through ACM and residing in the same AWS account as the CA.
//
// **About permissions** - If the private CA and the certificates it issues reside in the same account, you can use `AWS::ACMPCA::Permission` to grant permissions for ACM to carry out automatic certificate renewals.
// - For automatic certificate renewal to succeed, the ACM service principal needs permissions to create, retrieve, and list permissions.
// - If the private CA and the ACM certificates reside in different accounts, then permissions cannot be used to enable automatic renewals. Instead, the ACM certificate owner must set up a resource-based policy to enable cross-account issuance and renewals. For more information, see [Using a Resource Based Policy with ACM Private CA](https://docs.aws.amazon.com/acm-pca/latest/userguide/pca-rbp.html) .
//
// > To update an `AWS::ACMPCA::Permission` resource, you must first delete the existing permission resource from the CloudFormation stack and then create a new permission resource with updated properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPermission := awscdk.Aws_acmpca.NewCfnPermission(this, jsii.String("MyCfnPermission"), &cfnPermissionProps{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	certificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	principal: jsii.String("principal"),
//
//   	// the properties below are optional
//   	sourceAccount: jsii.String("sourceAccount"),
//   })
//
type CfnPermission interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The private CA actions that can be performed by the designated AWS service.
	//
	// Supported actions are `IssueCertificate` , `GetCertificate` , and `ListPermissions` .
	Actions() *[]*string
	SetActions(val *[]*string)
	// The Amazon Resource Number (ARN) of the private CA from which the permission was issued.
	CertificateAuthorityArn() *string
	SetCertificateAuthorityArn(val *string)
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
	// The AWS service or entity that holds the permission.
	//
	// At this time, the only valid principal is `acm.amazonaws.com` .
	Principal() *string
	SetPrincipal(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the account that assigned the permission.
	SourceAccount() *string
	SetSourceAccount(val *string)
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

// The jsii proxy struct for CfnPermission
type jsiiProxy_CfnPermission struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPermission) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) CertificateAuthorityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) SourceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ACMPCA::Permission`.
func NewCfnPermission(scope constructs.Construct, id *string, props *CfnPermissionProps) CfnPermission {
	_init_.Initialize()

	j := jsiiProxy_CfnPermission{}

	_jsii_.Create(
		"aws-cdk-lib.aws_acmpca.CfnPermission",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ACMPCA::Permission`.
func NewCfnPermission_Override(c CfnPermission, scope constructs.Construct, id *string, props *CfnPermissionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_acmpca.CfnPermission",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPermission) SetActions(val *[]*string) {
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_CfnPermission) SetCertificateAuthorityArn(val *string) {
	_jsii_.Set(
		j,
		"certificateAuthorityArn",
		val,
	)
}

func (j *jsiiProxy_CfnPermission) SetPrincipal(val *string) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_CfnPermission) SetSourceAccount(val *string) {
	_jsii_.Set(
		j,
		"sourceAccount",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPermission_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CfnPermission",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPermission_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CfnPermission",
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
func CfnPermission_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_acmpca.CfnPermission",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPermission_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_acmpca.CfnPermission",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPermission) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPermission) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPermission) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPermission) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPermission) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPermission) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPermission) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPermission) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermission) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermission) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPermission) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPermission) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermission) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermission) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermission) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPermission`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPermissionProps := &cfnPermissionProps{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	certificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	principal: jsii.String("principal"),
//
//   	// the properties below are optional
//   	sourceAccount: jsii.String("sourceAccount"),
//   }
//
type CfnPermissionProps struct {
	// The private CA actions that can be performed by the designated AWS service.
	//
	// Supported actions are `IssueCertificate` , `GetCertificate` , and `ListPermissions` .
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The Amazon Resource Number (ARN) of the private CA from which the permission was issued.
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The AWS service or entity that holds the permission.
	//
	// At this time, the only valid principal is `acm.amazonaws.com` .
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// The ID of the account that assigned the permission.
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
}

// Interface which all CertificateAuthority based class must implement.
type ICertificateAuthority interface {
	awscdk.IResource
	// The Amazon Resource Name of the Certificate.
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

