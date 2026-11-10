package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiot"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an IoT OTA update on a target group of things or groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOTAUpdate := awscdk.Aws_iot.NewCfnOTAUpdate(this, jsii.String("MyCfnOTAUpdate"), &CfnOTAUpdateProps{
//   	Files: []interface{}{
//   		&OTAUpdateFileProperty{
//   			Attributes: map[string]*string{
//   				"attributesKey": jsii.String("attributes"),
//   			},
//   			CodeSigning: &CodeSigningProperty{
//   				AwsSignerJobId: jsii.String("awsSignerJobId"),
//   				CustomCodeSigning: &CustomCodeSigningProperty{
//   					CertificateChain: &CodeSigningCertificateChainProperty{
//   						CertificateName: jsii.String("certificateName"),
//   						InlineDocument: jsii.String("inlineDocument"),
//   					},
//   					HashAlgorithm: jsii.String("hashAlgorithm"),
//   					Signature: &CodeSigningSignatureProperty{
//   						InlineDocument: jsii.String("inlineDocument"),
//   					},
//   					SignatureAlgorithm: jsii.String("signatureAlgorithm"),
//   				},
//   				StartSigningJobParameter: &StartSigningJobParameterProperty{
//   					Destination: &DestinationProperty{
//   						S3Destination: &S3DestinationProperty{
//   							Bucket: jsii.String("bucket"),
//   							Prefix: jsii.String("prefix"),
//   						},
//   					},
//   					SigningProfileName: jsii.String("signingProfileName"),
//   					SigningProfileParameter: &SigningProfileParameterProperty{
//   						CertificateArn: jsii.String("certificateArn"),
//   						CertificatePathOnDevice: jsii.String("certificatePathOnDevice"),
//   						Platform: jsii.String("platform"),
//   					},
//   				},
//   			},
//   			FileLocation: &FileLocationProperty{
//   				S3Location: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Key: jsii.String("key"),
//   					Version: jsii.String("version"),
//   				},
//   				Stream: &StreamProperty{
//   					FileId: jsii.Number(123),
//   					StreamId: jsii.String("streamId"),
//   				},
//   			},
//   			FileName: jsii.String("fileName"),
//   			FileType: jsii.Number(123),
//   			FileVersion: jsii.String("fileVersion"),
//   		},
//   	},
//   	OtaUpdateId: jsii.String("otaUpdateId"),
//   	Targets: []*string{
//   		jsii.String("targets"),
//   	},
//
//   	// the properties below are optional
//   	AdditionalParameters: map[string]*string{
//   		"additionalParametersKey": jsii.String("additionalParameters"),
//   	},
//   	AwsJobAbortConfig: &AwsJobAbortConfigProperty{
//   		AbortCriteriaList: []interface{}{
//   			&AwsJobAbortCriteriaProperty{
//   				Action: jsii.String("action"),
//   				FailureType: jsii.String("failureType"),
//   				MinNumberOfExecutedThings: jsii.Number(123),
//   				ThresholdPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	AwsJobExecutionsRolloutConfig: &AwsJobExecutionsRolloutConfigProperty{
//   		ExponentialRate: &AwsJobExponentialRolloutRateProperty{
//   			BaseRatePerMinute: jsii.Number(123),
//   			IncrementFactor: jsii.Number(123),
//   			RateIncreaseCriteria: &AwsJobRateIncreaseCriteriaProperty{
//   				NumberOfNotifiedThings: jsii.Number(123),
//   				NumberOfSucceededThings: jsii.Number(123),
//   			},
//   		},
//   		MaximumPerMinute: jsii.Number(123),
//   	},
//   	AwsJobPresignedUrlConfig: &AwsJobPresignedUrlConfigProperty{
//   		ExpiresInSec: jsii.Number(123),
//   	},
//   	AwsJobTimeoutConfig: &AwsJobTimeoutConfigProperty{
//   		InProgressTimeoutInMinutes: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	Protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetSelection: jsii.String("targetSelection"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html
//
type CfnOTAUpdate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsiot.IOTAUpdateRef
	awscdk.ITaggableV2
	// A list of additional OTA update parameters, which are name-value pairs.
	AdditionalParameters() interface{}
	SetAdditionalParameters(val interface{})
	// The date when the OTA update was created.
	AttrCreationDate() *string
	// The date when the OTA update was last updated.
	AttrLastModifiedDate() *string
	// The OTA update ARN.
	AttrOtaUpdateArn() *string
	// The status of the OTA update.
	AttrOtaUpdateStatus() *string
	AwsJobAbortConfig() interface{}
	SetAwsJobAbortConfig(val interface{})
	AwsJobExecutionsRolloutConfig() interface{}
	SetAwsJobExecutionsRolloutConfig(val interface{})
	AwsJobPresignedUrlConfig() interface{}
	SetAwsJobPresignedUrlConfig(val interface{})
	AwsJobTimeoutConfig() interface{}
	SetAwsJobTimeoutConfig(val interface{})
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnPropertyNames() *map[string]*string
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the OTA update.
	Description() *string
	SetDescription(val *string)
	Env() *interfaces.ResourceEnvironment
	// The files to be streamed by the OTA update.
	Files() interface{}
	SetFiles(val interface{})
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
	// The ID of the OTA update to be created.
	OtaUpdateId() *string
	SetOtaUpdateId(val *string)
	// A reference to a OTAUpdate resource.
	OtaUpdateRef() *interfacesawsiot.OTAUpdateReference
	// The protocol used to transfer the OTA update image.
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The IAM role that grants AWS IoT Core access to the Amazon S3, IoT jobs and AWS Code Signing resources to create an OTA update job.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata which can be used to manage updates.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// The devices targeted to receive OTA updates.
	Targets() *[]*string
	SetTargets(val *[]*string)
	// Specifies whether the update will continue to run (CONTINUOUS), or will be complete after all things specified as targets have completed the update (SNAPSHOT).
	TargetSelection() *string
	SetTargetSelection(val *string)
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
	// This method has been renamed to `addResourceDependency` to more clearly
	// set it apart from `construct.node.addDependency`. See the documentation
	// of that function for more details.
	// Deprecated: Use `addResourceDependency` instead.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method has been renamed to `addResourceDependency`, which makes it
	// more clear that this method operates at a different level from the
	// construct-level `construct.node.addDependency()` mechanism.
	// Deprecated: Use `addResourceDependency` instead.
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
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method only adds dependencies between L1 resources. If you are
	// looking for a generic construct-to-construct dependency mechanism that works
	// for all constructs including L2s, use `construct.node.addDependency` instead.
	AddResourceDependency(target awscdk.CfnResource, reason *string)
	// Sets the cross-stack reference strength for this resource.
	//
	// When set, any cross-stack reference to this resource will use the specified
	// strength instead of the global default from the consuming stack's context.
	ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength)
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
	CfnPropertyName(cdkPropertyName *string) *string
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
	// Retrieves an array of resources and stacks this resource depends on.
	//
	// For resources depended on directly, returns the `CfnResource` object. For
	// dependencies on other stacks, returns the `Stack` object. The order of the
	// array is not guaranteed.
	ObtainDependencies() *[]interface{}
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	// Deprecated: Use `removeResourceDependency` instead.
	RemoveDependency(target awscdk.CfnResource)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveResourceDependency(target awscdk.CfnResource)
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnOTAUpdate
type jsiiProxy_CfnOTAUpdate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsiotIOTAUpdateRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnOTAUpdate) AdditionalParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) AttrCreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) AttrLastModifiedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) AttrOtaUpdateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOtaUpdateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) AttrOtaUpdateStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOtaUpdateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) AwsJobAbortConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsJobAbortConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) AwsJobExecutionsRolloutConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsJobExecutionsRolloutConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) AwsJobPresignedUrlConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsJobPresignedUrlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) AwsJobTimeoutConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsJobTimeoutConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) Files() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) OtaUpdateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"otaUpdateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) OtaUpdateRef() *interfacesawsiot.OTAUpdateReference {
	var returns *interfacesawsiot.OTAUpdateReference
	_jsii_.Get(
		j,
		"otaUpdateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) Targets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) TargetSelection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdate) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::OTAUpdate`.
func NewCfnOTAUpdate(scope constructs.Construct, id *string, props *CfnOTAUpdateProps) CfnOTAUpdate {
	_init_.Initialize()

	if err := validateNewCfnOTAUpdateParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOTAUpdate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnOTAUpdate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::OTAUpdate`.
func NewCfnOTAUpdate_Override(c CfnOTAUpdate, scope constructs.Construct, id *string, props *CfnOTAUpdateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnOTAUpdate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetAdditionalParameters(val interface{}) {
	if err := j.validateSetAdditionalParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalParameters",
		val,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetAwsJobAbortConfig(val interface{}) {
	if err := j.validateSetAwsJobAbortConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsJobAbortConfig",
		val,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetAwsJobExecutionsRolloutConfig(val interface{}) {
	if err := j.validateSetAwsJobExecutionsRolloutConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsJobExecutionsRolloutConfig",
		val,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetAwsJobPresignedUrlConfig(val interface{}) {
	if err := j.validateSetAwsJobPresignedUrlConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsJobPresignedUrlConfig",
		val,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetAwsJobTimeoutConfig(val interface{}) {
	if err := j.validateSetAwsJobTimeoutConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsJobTimeoutConfig",
		val,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetFiles(val interface{}) {
	if err := j.validateSetFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"files",
		val,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetOtaUpdateId(val *string) {
	if err := j.validateSetOtaUpdateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"otaUpdateId",
		val,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetProtocols(val *[]*string) {
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetTargets(val *[]*string) {
	if err := j.validateSetTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

func (j *jsiiProxy_CfnOTAUpdate)SetTargetSelection(val *string) {
	_jsii_.Set(
		j,
		"targetSelection",
		val,
	)
}

func CfnOTAUpdate_ArnForOTAUpdate(resource interfacesawsiot.IOTAUpdateRef) *string {
	_init_.Initialize()

	if err := validateCfnOTAUpdate_ArnForOTAUpdateParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnOTAUpdate",
		"arnForOTAUpdate",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Creates a new IOTAUpdateRef from a otaUpdateId.
func CfnOTAUpdate_FromOtaUpdateId(scope constructs.Construct, id *string, otaUpdateId *string) interfacesawsiot.IOTAUpdateRef {
	_init_.Initialize()

	if err := validateCfnOTAUpdate_FromOtaUpdateIdParameters(scope, id, otaUpdateId); err != nil {
		panic(err)
	}
	var returns interfacesawsiot.IOTAUpdateRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnOTAUpdate",
		"fromOtaUpdateId",
		[]interface{}{scope, id, otaUpdateId},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnOTAUpdate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOTAUpdate_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnOTAUpdate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnOTAUpdate.
func CfnOTAUpdate_IsCfnOTAUpdate(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOTAUpdate_IsCfnOTAUpdateParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnOTAUpdate",
		"isCfnOTAUpdate",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnOTAUpdate_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOTAUpdate_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnOTAUpdate",
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
func CfnOTAUpdate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOTAUpdate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnOTAUpdate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOTAUpdate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnOTAUpdate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOTAUpdate) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) AddResourceDependency(target awscdk.CfnResource, reason *string) {
	if err := c.validateAddResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addResourceDependency",
		[]interface{}{target, reason},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) CfnPropertyName(cdkPropertyName *string) *string {
	if err := c.validateCfnPropertyNameParameters(cdkPropertyName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"cfnPropertyName",
		[]interface{}{cdkPropertyName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOTAUpdate) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnOTAUpdate) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnOTAUpdate) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOTAUpdate) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) RemoveResourceDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeResourceDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnOTAUpdate) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOTAUpdate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOTAUpdate) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnOTAUpdate) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

