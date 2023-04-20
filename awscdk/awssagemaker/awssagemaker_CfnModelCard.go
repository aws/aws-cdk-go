package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v3"
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
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
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
	// Experimental.
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
	// Experimental.
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The security configuration used to protect model card data.
	SecurityConfig() interface{}
	SetSecurityConfig(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Key-value pairs used to manage metadata for the model card.
	Tags() awscdk.TagManager
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

func (j *jsiiProxy_CfnModelCard) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Create a new `AWS::SageMaker::ModelCard`.
func NewCfnModelCard(scope awscdk.Construct, id *string, props *CfnModelCardProps) CfnModelCard {
	_init_.Initialize()

	if err := validateNewCfnModelCardParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnModelCard{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnModelCard",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::ModelCard`.
func NewCfnModelCard_Override(c CfnModelCard, scope awscdk.Construct, id *string, props *CfnModelCardProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnModelCard",
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
// Experimental.
func CfnModelCard_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelCard_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelCard",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnModelCard_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnModelCard_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelCard",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnModelCard_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelCard_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelCard",
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
		"monocdk.aws_sagemaker.CfnModelCard",
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

func (c *jsiiProxy_CfnModelCard) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnModelCard) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnModelCard) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnModelCard) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
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

func (c *jsiiProxy_CfnModelCard) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnModelCard) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnModelCard) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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

