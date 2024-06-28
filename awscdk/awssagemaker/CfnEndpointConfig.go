package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SageMaker::EndpointConfig` resource creates a configuration for an Amazon SageMaker endpoint.
//
// For more information, see [CreateEndpointConfig](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html) in the *SageMaker Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointConfig := awscdk.Aws_sagemaker.NewCfnEndpointConfig(this, jsii.String("MyCfnEndpointConfig"), &CfnEndpointConfigProps{
//   	ProductionVariants: []interface{}{
//   		&ProductionVariantProperty{
//   			VariantName: jsii.String("variantName"),
//
//   			// the properties below are optional
//   			AcceleratorType: jsii.String("acceleratorType"),
//   			ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   			EnableSsmAccess: jsii.Boolean(false),
//   			InitialInstanceCount: jsii.Number(123),
//   			InitialVariantWeight: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			ManagedInstanceScaling: &ManagedInstanceScalingProperty{
//   				MaxInstanceCount: jsii.Number(123),
//   				MinInstanceCount: jsii.Number(123),
//   				Status: jsii.String("status"),
//   			},
//   			ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   			ModelName: jsii.String("modelName"),
//   			RoutingConfig: &RoutingConfigProperty{
//   				RoutingStrategy: jsii.String("routingStrategy"),
//   			},
//   			ServerlessConfig: &ServerlessConfigProperty{
//   				MaxConcurrency: jsii.Number(123),
//   				MemorySizeInMb: jsii.Number(123),
//
//   				// the properties below are optional
//   				ProvisionedConcurrency: jsii.Number(123),
//   			},
//   			VolumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AsyncInferenceConfig: &AsyncInferenceConfigProperty{
//   		OutputConfig: &AsyncInferenceOutputConfigProperty{
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			NotificationConfig: &AsyncInferenceNotificationConfigProperty{
//   				ErrorTopic: jsii.String("errorTopic"),
//   				IncludeInferenceResponseIn: []*string{
//   					jsii.String("includeInferenceResponseIn"),
//   				},
//   				SuccessTopic: jsii.String("successTopic"),
//   			},
//   			S3FailurePath: jsii.String("s3FailurePath"),
//   			S3OutputPath: jsii.String("s3OutputPath"),
//   		},
//
//   		// the properties below are optional
//   		ClientConfig: &AsyncInferenceClientConfigProperty{
//   			MaxConcurrentInvocationsPerInstance: jsii.Number(123),
//   		},
//   	},
//   	DataCaptureConfig: &DataCaptureConfigProperty{
//   		CaptureOptions: []interface{}{
//   			&CaptureOptionProperty{
//   				CaptureMode: jsii.String("captureMode"),
//   			},
//   		},
//   		DestinationS3Uri: jsii.String("destinationS3Uri"),
//   		InitialSamplingPercentage: jsii.Number(123),
//
//   		// the properties below are optional
//   		CaptureContentTypeHeader: &CaptureContentTypeHeaderProperty{
//   			CsvContentTypes: []*string{
//   				jsii.String("csvContentTypes"),
//   			},
//   			JsonContentTypes: []*string{
//   				jsii.String("jsonContentTypes"),
//   			},
//   		},
//   		EnableCapture: jsii.Boolean(false),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	EnableNetworkIsolation: jsii.Boolean(false),
//   	EndpointConfigName: jsii.String("endpointConfigName"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	ExplainerConfig: &ExplainerConfigProperty{
//   		ClarifyExplainerConfig: &ClarifyExplainerConfigProperty{
//   			ShapConfig: &ClarifyShapConfigProperty{
//   				ShapBaselineConfig: &ClarifyShapBaselineConfigProperty{
//   					MimeType: jsii.String("mimeType"),
//   					ShapBaseline: jsii.String("shapBaseline"),
//   					ShapBaselineUri: jsii.String("shapBaselineUri"),
//   				},
//
//   				// the properties below are optional
//   				NumberOfSamples: jsii.Number(123),
//   				Seed: jsii.Number(123),
//   				TextConfig: &ClarifyTextConfigProperty{
//   					Granularity: jsii.String("granularity"),
//   					Language: jsii.String("language"),
//   				},
//   				UseLogit: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			EnableExplanations: jsii.String("enableExplanations"),
//   			InferenceConfig: &ClarifyInferenceConfigProperty{
//   				ContentTemplate: jsii.String("contentTemplate"),
//   				FeatureHeaders: []*string{
//   					jsii.String("featureHeaders"),
//   				},
//   				FeaturesAttribute: jsii.String("featuresAttribute"),
//   				FeatureTypes: []*string{
//   					jsii.String("featureTypes"),
//   				},
//   				LabelAttribute: jsii.String("labelAttribute"),
//   				LabelHeaders: []*string{
//   					jsii.String("labelHeaders"),
//   				},
//   				LabelIndex: jsii.Number(123),
//   				MaxPayloadInMb: jsii.Number(123),
//   				MaxRecordCount: jsii.Number(123),
//   				ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   				ProbabilityIndex: jsii.Number(123),
//   			},
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ShadowProductionVariants: []interface{}{
//   		&ProductionVariantProperty{
//   			VariantName: jsii.String("variantName"),
//
//   			// the properties below are optional
//   			AcceleratorType: jsii.String("acceleratorType"),
//   			ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   			EnableSsmAccess: jsii.Boolean(false),
//   			InitialInstanceCount: jsii.Number(123),
//   			InitialVariantWeight: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			ManagedInstanceScaling: &ManagedInstanceScalingProperty{
//   				MaxInstanceCount: jsii.Number(123),
//   				MinInstanceCount: jsii.Number(123),
//   				Status: jsii.String("status"),
//   			},
//   			ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   			ModelName: jsii.String("modelName"),
//   			RoutingConfig: &RoutingConfigProperty{
//   				RoutingStrategy: jsii.String("routingStrategy"),
//   			},
//   			ServerlessConfig: &ServerlessConfigProperty{
//   				MaxConcurrency: jsii.Number(123),
//   				MemorySizeInMb: jsii.Number(123),
//
//   				// the properties below are optional
//   				ProvisionedConcurrency: jsii.Number(123),
//   			},
//   			VolumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html
//
type CfnEndpointConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig() interface{}
	SetAsyncInferenceConfig(val interface{})
	// The name of the endpoint configuration, such as `MyEndpointConfiguration` .
	AttrEndpointConfigName() *string
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies how to capture endpoint data for model monitor.
	DataCaptureConfig() interface{}
	SetDataCaptureConfig(val interface{})
	EnableNetworkIsolation() interface{}
	SetEnableNetworkIsolation(val interface{})
	// The name of the endpoint configuration.
	EndpointConfigName() *string
	SetEndpointConfigName(val *string)
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// A parameter to activate explainers.
	ExplainerConfig() interface{}
	SetExplainerConfig(val interface{})
	// The Amazon Resource Name (ARN) of an AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	// A list of `ProductionVariant` objects, one for each model that you want to host at this endpoint.
	ProductionVariants() interface{}
	SetProductionVariants(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Array of `ProductionVariant` objects.
	ShadowProductionVariants() interface{}
	SetShadowProductionVariants(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// A list of key-value pairs to apply to this resource.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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
	VpcConfig() interface{}
	SetVpcConfig(val interface{})
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

// The jsii proxy struct for CfnEndpointConfig
type jsiiProxy_CfnEndpointConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnEndpointConfig) AsyncInferenceConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asyncInferenceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) AttrEndpointConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) DataCaptureConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCaptureConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) EnableNetworkIsolation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) EndpointConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) ExplainerConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"explainerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) ProductionVariants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productionVariants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) ShadowProductionVariants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shadowProductionVariants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) VpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}


func NewCfnEndpointConfig(scope constructs.Construct, id *string, props *CfnEndpointConfigProps) CfnEndpointConfig {
	_init_.Initialize()

	if err := validateNewCfnEndpointConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEndpointConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnEndpointConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnEndpointConfig_Override(c CfnEndpointConfig, scope constructs.Construct, id *string, props *CfnEndpointConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnEndpointConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEndpointConfig)SetAsyncInferenceConfig(val interface{}) {
	if err := j.validateSetAsyncInferenceConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asyncInferenceConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig)SetDataCaptureConfig(val interface{}) {
	if err := j.validateSetDataCaptureConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCaptureConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig)SetEnableNetworkIsolation(val interface{}) {
	if err := j.validateSetEnableNetworkIsolationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNetworkIsolation",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig)SetEndpointConfigName(val *string) {
	_jsii_.Set(
		j,
		"endpointConfigName",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig)SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig)SetExplainerConfig(val interface{}) {
	if err := j.validateSetExplainerConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"explainerConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig)SetProductionVariants(val interface{}) {
	if err := j.validateSetProductionVariantsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productionVariants",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig)SetShadowProductionVariants(val interface{}) {
	if err := j.validateSetShadowProductionVariantsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowProductionVariants",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig)SetVpcConfig(val interface{}) {
	if err := j.validateSetVpcConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEndpointConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEndpointConfig_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnEndpointConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnEndpointConfig_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEndpointConfig_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnEndpointConfig",
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
func CfnEndpointConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEndpointConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnEndpointConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpointConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sagemaker.CfnEndpointConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnEndpointConfig) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnEndpointConfig) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnEndpointConfig) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

