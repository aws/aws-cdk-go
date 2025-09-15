package awspcaconnectorad

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspcaconnectorad/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Active Directory compatible certificate template.
//
// The connectors issues certificates using these templates based on the requester’s Active Directory group membership.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTemplate := awscdk.Aws_pcaconnectorad.NewCfnTemplate(this, jsii.String("MyCfnTemplate"), &CfnTemplateProps{
//   	ConnectorArn: jsii.String("connectorArn"),
//   	Definition: &TemplateDefinitionProperty{
//   		TemplateV2: &TemplateV2Property{
//   			CertificateValidity: &CertificateValidityProperty{
//   				RenewalPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   				ValidityPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   			},
//   			EnrollmentFlags: &EnrollmentFlagsV2Property{
//   				EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   				IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   				NoSecurityExtension: jsii.Boolean(false),
//   				RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   				UserInteractionRequired: jsii.Boolean(false),
//   			},
//   			Extensions: &ExtensionsV2Property{
//   				KeyUsage: &KeyUsageProperty{
//   					UsageFlags: &KeyUsageFlagsProperty{
//   						DataEncipherment: jsii.Boolean(false),
//   						DigitalSignature: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						KeyEncipherment: jsii.Boolean(false),
//   						NonRepudiation: jsii.Boolean(false),
//   					},
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//
//   				// the properties below are optional
//   				ApplicationPolicies: &ApplicationPoliciesProperty{
//   					Policies: []interface{}{
//   						&ApplicationPolicyProperty{
//   							PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   							PolicyType: jsii.String("policyType"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//   			},
//   			GeneralFlags: &GeneralFlagsV2Property{
//   				AutoEnrollment: jsii.Boolean(false),
//   				MachineType: jsii.Boolean(false),
//   			},
//   			PrivateKeyAttributes: &PrivateKeyAttributesV2Property{
//   				KeySpec: jsii.String("keySpec"),
//   				MinimalKeyLength: jsii.Number(123),
//
//   				// the properties below are optional
//   				CryptoProviders: []*string{
//   					jsii.String("cryptoProviders"),
//   				},
//   			},
//   			PrivateKeyFlags: &PrivateKeyFlagsV2Property{
//   				ClientVersion: jsii.String("clientVersion"),
//
//   				// the properties below are optional
//   				ExportableKey: jsii.Boolean(false),
//   				StrongKeyProtectionRequired: jsii.Boolean(false),
//   			},
//   			SubjectNameFlags: &SubjectNameFlagsV2Property{
//   				RequireCommonName: jsii.Boolean(false),
//   				RequireDirectoryPath: jsii.Boolean(false),
//   				RequireDnsAsCn: jsii.Boolean(false),
//   				RequireEmail: jsii.Boolean(false),
//   				SanRequireDirectoryGuid: jsii.Boolean(false),
//   				SanRequireDns: jsii.Boolean(false),
//   				SanRequireDomainDns: jsii.Boolean(false),
//   				SanRequireEmail: jsii.Boolean(false),
//   				SanRequireSpn: jsii.Boolean(false),
//   				SanRequireUpn: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			SupersededTemplates: []*string{
//   				jsii.String("supersededTemplates"),
//   			},
//   		},
//   		TemplateV3: &TemplateV3Property{
//   			CertificateValidity: &CertificateValidityProperty{
//   				RenewalPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   				ValidityPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   			},
//   			EnrollmentFlags: &EnrollmentFlagsV3Property{
//   				EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   				IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   				NoSecurityExtension: jsii.Boolean(false),
//   				RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   				UserInteractionRequired: jsii.Boolean(false),
//   			},
//   			Extensions: &ExtensionsV3Property{
//   				KeyUsage: &KeyUsageProperty{
//   					UsageFlags: &KeyUsageFlagsProperty{
//   						DataEncipherment: jsii.Boolean(false),
//   						DigitalSignature: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						KeyEncipherment: jsii.Boolean(false),
//   						NonRepudiation: jsii.Boolean(false),
//   					},
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//
//   				// the properties below are optional
//   				ApplicationPolicies: &ApplicationPoliciesProperty{
//   					Policies: []interface{}{
//   						&ApplicationPolicyProperty{
//   							PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   							PolicyType: jsii.String("policyType"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//   			},
//   			GeneralFlags: &GeneralFlagsV3Property{
//   				AutoEnrollment: jsii.Boolean(false),
//   				MachineType: jsii.Boolean(false),
//   			},
//   			HashAlgorithm: jsii.String("hashAlgorithm"),
//   			PrivateKeyAttributes: &PrivateKeyAttributesV3Property{
//   				Algorithm: jsii.String("algorithm"),
//   				KeySpec: jsii.String("keySpec"),
//   				KeyUsageProperty: &KeyUsagePropertyProperty{
//   					PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   						Decrypt: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						Sign: jsii.Boolean(false),
//   					},
//   					PropertyType: jsii.String("propertyType"),
//   				},
//   				MinimalKeyLength: jsii.Number(123),
//
//   				// the properties below are optional
//   				CryptoProviders: []*string{
//   					jsii.String("cryptoProviders"),
//   				},
//   			},
//   			PrivateKeyFlags: &PrivateKeyFlagsV3Property{
//   				ClientVersion: jsii.String("clientVersion"),
//
//   				// the properties below are optional
//   				ExportableKey: jsii.Boolean(false),
//   				RequireAlternateSignatureAlgorithm: jsii.Boolean(false),
//   				StrongKeyProtectionRequired: jsii.Boolean(false),
//   			},
//   			SubjectNameFlags: &SubjectNameFlagsV3Property{
//   				RequireCommonName: jsii.Boolean(false),
//   				RequireDirectoryPath: jsii.Boolean(false),
//   				RequireDnsAsCn: jsii.Boolean(false),
//   				RequireEmail: jsii.Boolean(false),
//   				SanRequireDirectoryGuid: jsii.Boolean(false),
//   				SanRequireDns: jsii.Boolean(false),
//   				SanRequireDomainDns: jsii.Boolean(false),
//   				SanRequireEmail: jsii.Boolean(false),
//   				SanRequireSpn: jsii.Boolean(false),
//   				SanRequireUpn: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			SupersededTemplates: []*string{
//   				jsii.String("supersededTemplates"),
//   			},
//   		},
//   		TemplateV4: &TemplateV4Property{
//   			CertificateValidity: &CertificateValidityProperty{
//   				RenewalPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   				ValidityPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   			},
//   			EnrollmentFlags: &EnrollmentFlagsV4Property{
//   				EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   				IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   				NoSecurityExtension: jsii.Boolean(false),
//   				RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   				UserInteractionRequired: jsii.Boolean(false),
//   			},
//   			Extensions: &ExtensionsV4Property{
//   				KeyUsage: &KeyUsageProperty{
//   					UsageFlags: &KeyUsageFlagsProperty{
//   						DataEncipherment: jsii.Boolean(false),
//   						DigitalSignature: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						KeyEncipherment: jsii.Boolean(false),
//   						NonRepudiation: jsii.Boolean(false),
//   					},
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//
//   				// the properties below are optional
//   				ApplicationPolicies: &ApplicationPoliciesProperty{
//   					Policies: []interface{}{
//   						&ApplicationPolicyProperty{
//   							PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   							PolicyType: jsii.String("policyType"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//   			},
//   			GeneralFlags: &GeneralFlagsV4Property{
//   				AutoEnrollment: jsii.Boolean(false),
//   				MachineType: jsii.Boolean(false),
//   			},
//   			PrivateKeyAttributes: &PrivateKeyAttributesV4Property{
//   				KeySpec: jsii.String("keySpec"),
//   				MinimalKeyLength: jsii.Number(123),
//
//   				// the properties below are optional
//   				Algorithm: jsii.String("algorithm"),
//   				CryptoProviders: []*string{
//   					jsii.String("cryptoProviders"),
//   				},
//   				KeyUsageProperty: &KeyUsagePropertyProperty{
//   					PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   						Decrypt: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						Sign: jsii.Boolean(false),
//   					},
//   					PropertyType: jsii.String("propertyType"),
//   				},
//   			},
//   			PrivateKeyFlags: &PrivateKeyFlagsV4Property{
//   				ClientVersion: jsii.String("clientVersion"),
//
//   				// the properties below are optional
//   				ExportableKey: jsii.Boolean(false),
//   				RequireAlternateSignatureAlgorithm: jsii.Boolean(false),
//   				RequireSameKeyRenewal: jsii.Boolean(false),
//   				StrongKeyProtectionRequired: jsii.Boolean(false),
//   				UseLegacyProvider: jsii.Boolean(false),
//   			},
//   			SubjectNameFlags: &SubjectNameFlagsV4Property{
//   				RequireCommonName: jsii.Boolean(false),
//   				RequireDirectoryPath: jsii.Boolean(false),
//   				RequireDnsAsCn: jsii.Boolean(false),
//   				RequireEmail: jsii.Boolean(false),
//   				SanRequireDirectoryGuid: jsii.Boolean(false),
//   				SanRequireDns: jsii.Boolean(false),
//   				SanRequireDomainDns: jsii.Boolean(false),
//   				SanRequireEmail: jsii.Boolean(false),
//   				SanRequireSpn: jsii.Boolean(false),
//   				SanRequireUpn: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			HashAlgorithm: jsii.String("hashAlgorithm"),
//   			SupersededTemplates: []*string{
//   				jsii.String("supersededTemplates"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ReenrollAllCertificateHolders: jsii.Boolean(false),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-template.html
//
type CfnTemplate interface {
	awscdk.CfnResource
	ITemplateRef
	awscdk.IInspectable
	awscdk.ITaggableV2
	// The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html) .
	AttrTemplateArn() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html) .
	ConnectorArn() *string
	SetConnectorArn(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Template configuration to define the information included in certificates.
	Definition() interface{}
	SetDefinition(val interface{})
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
	// Name of the templates.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// This setting allows the major version of a template to be increased automatically.
	ReenrollAllCertificateHolders() interface{}
	SetReenrollAllCertificateHolders(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata assigned to a template consisting of a key-value pair.
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	// A reference to a Template resource.
	TemplateRef() *TemplateReference
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

// The jsii proxy struct for CfnTemplate
type jsiiProxy_CfnTemplate struct {
	internal.Type__awscdkCfnResource
	jsiiProxy_ITemplateRef
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnTemplate) AttrTemplateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTemplateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) ConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Definition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) ReenrollAllCertificateHolders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reenrollAllCertificateHolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) TemplateRef() *TemplateReference {
	var returns *TemplateReference
	_jsii_.Get(
		j,
		"templateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnTemplate(scope constructs.Construct, id *string, props *CfnTemplateProps) CfnTemplate {
	_init_.Initialize()

	if err := validateNewCfnTemplateParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pcaconnectorad.CfnTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnTemplate_Override(c CfnTemplate, scope constructs.Construct, id *string, props *CfnTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pcaconnectorad.CfnTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTemplate)SetConnectorArn(val *string) {
	if err := j.validateSetConnectorArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorArn",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate)SetDefinition(val interface{}) {
	if err := j.validateSetDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate)SetReenrollAllCertificateHolders(val interface{}) {
	if err := j.validateSetReenrollAllCertificateHoldersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reenrollAllCertificateHolders",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate)SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTemplate_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pcaconnectorad.CfnTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnTemplate_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTemplate_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pcaconnectorad.CfnTemplate",
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
func CfnTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pcaconnectorad.CfnTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pcaconnectorad.CfnTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTemplate) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTemplate) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTemplate) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTemplate) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTemplate) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTemplate) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnTemplate) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnTemplate) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTemplate) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTemplate) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTemplate) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnTemplate) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnTemplate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTemplate) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

