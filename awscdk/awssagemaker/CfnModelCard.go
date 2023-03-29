package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::SageMaker::ModelCard`.
//
// Creates an Amazon SageMaker Model Card.
//
// For information about how to use model cards, see [Amazon SageMaker Model Card](https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   cfnModelCard := awscdk.Aws_sagemaker.NewCfnModelCard(this, jsii.String("MyCfnModelCard"), &CfnModelCardProps{
//   	Content: &ContentProperty{
//   		AdditionalInformation: &AdditionalInformationProperty{
//   			CaveatsAndRecommendations: jsii.String("caveatsAndRecommendations"),
//   			CustomDetails: map[string]*string{
//   				"customDetailsKey": jsii.String("customDetails"),
//   			},
//   			EthicalConsiderations: jsii.String("ethicalConsiderations"),
//   		},
//   		BusinessDetails: &BusinessDetailsProperty{
//   			BusinessProblem: jsii.String("businessProblem"),
//   			BusinessStakeholders: jsii.String("businessStakeholders"),
//   			LineOfBusiness: jsii.String("lineOfBusiness"),
//   		},
//   		EvaluationDetails: []interface{}{
//   			&EvaluationDetailProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Datasets: []*string{
//   					jsii.String("datasets"),
//   				},
//   				EvaluationJobArn: jsii.String("evaluationJobArn"),
//   				EvaluationObservation: jsii.String("evaluationObservation"),
//   				Metadata: map[string]*string{
//   					"metadataKey": jsii.String("metadata"),
//   				},
//   				MetricGroups: []interface{}{
//   					&MetricGroupProperty{
//   						MetricData: []interface{}{
//   							&MetricDataItemsProperty{
//   								Name: jsii.String("name"),
//   								Type: jsii.String("type"),
//   								Value: value,
//
//   								// the properties below are optional
//   								Notes: jsii.String("notes"),
//   								XAxisName: []*string{
//   									jsii.String("xAxisName"),
//   								},
//   								YAxisName: []*string{
//   									jsii.String("yAxisName"),
//   								},
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   		IntendedUses: &IntendedUsesProperty{
//   			ExplanationsForRiskRating: jsii.String("explanationsForRiskRating"),
//   			FactorsAffectingModelEfficiency: jsii.String("factorsAffectingModelEfficiency"),
//   			IntendedUses: jsii.String("intendedUses"),
//   			PurposeOfModel: jsii.String("purposeOfModel"),
//   			RiskRating: jsii.String("riskRating"),
//   		},
//   		ModelOverview: &ModelOverviewProperty{
//   			AlgorithmType: jsii.String("algorithmType"),
//   			InferenceEnvironment: &InferenceEnvironmentProperty{
//   				ContainerImage: []*string{
//   					jsii.String("containerImage"),
//   				},
//   			},
//   			ModelArtifact: []*string{
//   				jsii.String("modelArtifact"),
//   			},
//   			ModelCreator: jsii.String("modelCreator"),
//   			ModelDescription: jsii.String("modelDescription"),
//   			ModelId: jsii.String("modelId"),
//   			ModelName: jsii.String("modelName"),
//   			ModelOwner: jsii.String("modelOwner"),
//   			ModelVersion: jsii.Number(123),
//   			ProblemType: jsii.String("problemType"),
//   		},
//   		TrainingDetails: &TrainingDetailsProperty{
//   			ObjectiveFunction: &ObjectiveFunctionProperty{
//   				Function: &FunctionProperty{
//   					Condition: jsii.String("condition"),
//   					Facet: jsii.String("facet"),
//   					Function: jsii.String("function"),
//   				},
//   				Notes: jsii.String("notes"),
//   			},
//   			TrainingJobDetails: &TrainingJobDetailsProperty{
//   				HyperParameters: []interface{}{
//   					&TrainingHyperParameterProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				TrainingArn: jsii.String("trainingArn"),
//   				TrainingDatasets: []*string{
//   					jsii.String("trainingDatasets"),
//   				},
//   				TrainingEnvironment: &TrainingEnvironmentProperty{
//   					ContainerImage: []*string{
//   						jsii.String("containerImage"),
//   					},
//   				},
//   				TrainingMetrics: []interface{}{
//   					&TrainingMetricProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.Number(123),
//
//   						// the properties below are optional
//   						Notes: jsii.String("notes"),
//   					},
//   				},
//   				UserProvidedHyperParameters: []interface{}{
//   					&TrainingHyperParameterProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				UserProvidedTrainingMetrics: []interface{}{
//   					&TrainingMetricProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.Number(123),
//
//   						// the properties below are optional
//   						Notes: jsii.String("notes"),
//   					},
//   				},
//   			},
//   			TrainingObservations: jsii.String("trainingObservations"),
//   		},
//   	},
//   	ModelCardName: jsii.String("modelCardName"),
//   	ModelCardStatus: jsii.String("modelCardStatus"),
//
//   	// the properties below are optional
//   	CreatedBy: &UserContextProperty{
//   		DomainId: jsii.String("domainId"),
//   		UserProfileArn: jsii.String("userProfileArn"),
//   		UserProfileName: jsii.String("userProfileName"),
//   	},
//   	LastModifiedBy: &UserContextProperty{
//   		DomainId: jsii.String("domainId"),
//   		UserProfileArn: jsii.String("userProfileArn"),
//   		UserProfileName: jsii.String("userProfileName"),
//   	},
//   	SecurityConfig: &SecurityConfigProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnModelCard interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrCreatedByDomainId() *string
	AttrCreatedByUserProfileArn() *string
	AttrCreatedByUserProfileName() *string
	AttrCreationTime() *string
	AttrLastModifiedByDomainId() *string
	AttrLastModifiedByUserProfileArn() *string
	AttrLastModifiedByUserProfileName() *string
	AttrLastModifiedTime() *string
	// The Amazon Resource Number (ARN) of the model card.
	//
	// For example, `arn:aws:sagemaker:us-west-2:012345678901:modelcard/examplemodelcard` .
	AttrModelCardArn() *string
	AttrModelCardProcessingStatus() *string
	AttrModelCardVersion() *float64
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The content of the model card.
	//
	// Content uses the [model card JSON schema](https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html#model-cards-json-schema) .
	Content() interface{}
	SetContent(val interface{})
	// Information about the user who created or modified one or more of the following:.
	//
	// - Experiment
	// - Trial
	// - Trial component
	// - Lineage group
	// - Project
	// - Model Card.
	CreatedBy() interface{}
	SetCreatedBy(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::SageMaker::ModelCard.LastModifiedBy`.
	LastModifiedBy() interface{}
	SetLastModifiedBy(val interface{})
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
	// The unique name of the model card.
	ModelCardName() *string
	SetModelCardName(val *string)
	// The approval status of the model card within your organization.
	//
	// Different organizations might have different criteria for model card review and approval.
	//
	// - `Draft` : The model card is a work in progress.
	// - `PendingReview` : The model card is pending review.
	// - `Approved` : The model card is approved.
	// - `Archived` : The model card is archived. No more updates should be made to the model card, but it can still be exported.
	ModelCardStatus() *string
	SetModelCardStatus(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The security configuration used to protect model card data.
	SecurityConfig() interface{}
	SetSecurityConfig(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Key-value pairs used to manage metadata for the model card.
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnModelCard
type jsiiProxy_CfnModelCard struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnModelCard) AttrCreatedByDomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedByDomainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) AttrCreatedByUserProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedByUserProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) AttrCreatedByUserProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedByUserProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) AttrLastModifiedByDomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedByDomainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) AttrLastModifiedByUserProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedByUserProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) AttrLastModifiedByUserProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedByUserProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) AttrLastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) AttrModelCardArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModelCardArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) AttrModelCardProcessingStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModelCardProcessingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) AttrModelCardVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrModelCardVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) Content() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) CreatedBy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) LastModifiedBy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastModifiedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) ModelCardName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelCardName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) ModelCardStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelCardStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) SecurityConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCard) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::ModelCard`.
func NewCfnModelCard(scope constructs.Construct, id *string, props *CfnModelCardProps) CfnModelCard {
	_init_.Initialize()

	if err := validateNewCfnModelCardParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnModelCard{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnModelCard",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::ModelCard`.
func NewCfnModelCard_Override(c CfnModelCard, scope constructs.Construct, id *string, props *CfnModelCardProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnModelCard",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnModelCard)SetContent(val interface{}) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_CfnModelCard)SetCreatedBy(val interface{}) {
	if err := j.validateSetCreatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBy",
		val,
	)
}

func (j *jsiiProxy_CfnModelCard)SetLastModifiedBy(val interface{}) {
	if err := j.validateSetLastModifiedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastModifiedBy",
		val,
	)
}

func (j *jsiiProxy_CfnModelCard)SetModelCardName(val *string) {
	if err := j.validateSetModelCardNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelCardName",
		val,
	)
}

func (j *jsiiProxy_CfnModelCard)SetModelCardStatus(val *string) {
	if err := j.validateSetModelCardStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelCardStatus",
		val,
	)
}

func (j *jsiiProxy_CfnModelCard)SetSecurityConfig(val interface{}) {
	if err := j.validateSetSecurityConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnModelCard_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelCard_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnModelCard",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnModelCard_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnModelCard_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnModelCard",
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
func CfnModelCard_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelCard_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnModelCard",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModelCard_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sagemaker.CfnModelCard",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModelCard) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnModelCard) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModelCard) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModelCard) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnModelCard) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnModelCard) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnModelCard) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnModelCard) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnModelCard) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnModelCard) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnModelCard) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnModelCard) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelCard) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelCard) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnModelCard) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModelCard) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnModelCard) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnModelCard) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelCard) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelCard) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

