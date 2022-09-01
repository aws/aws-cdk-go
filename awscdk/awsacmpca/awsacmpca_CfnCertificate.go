package awsacmpca

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

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

func (j *jsiiProxy_CfnCertificate) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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

	if err := validateNewCfnCertificateParameters(scope, id, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_CfnCertificate)SetApiPassthrough(val interface{}) {
	if err := j.validateSetApiPassthroughParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiPassthrough",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate)SetCertificateAuthorityArn(val *string) {
	if err := j.validateSetCertificateAuthorityArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateAuthorityArn",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate)SetCertificateSigningRequest(val *string) {
	if err := j.validateSetCertificateSigningRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateSigningRequest",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate)SetSigningAlgorithm(val *string) {
	if err := j.validateSetSigningAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate)SetTemplateArn(val *string) {
	_jsii_.Set(
		j,
		"templateArn",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate)SetValidity(val interface{}) {
	if err := j.validateSetValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validity",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate)SetValidityNotBefore(val interface{}) {
	if err := j.validateSetValidityNotBeforeParameters(val); err != nil {
		panic(err)
	}
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

	if err := validateCfnCertificate_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateCfnCertificate_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
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

	if err := validateCfnCertificate_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCertificate) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCertificate) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCertificate) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCertificate) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCertificate) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCertificate) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnCertificate) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCertificate) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCertificate) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCertificate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

