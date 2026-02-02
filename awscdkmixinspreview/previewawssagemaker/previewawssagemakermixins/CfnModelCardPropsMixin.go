package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon SageMaker Model Card.
//
// For information about how to use model cards, see [Amazon SageMaker Model Card](https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var value interface{}
//
//   cfnModelCardPropsMixin := awscdkmixinspreview.Mixins.NewCfnModelCardPropsMixin(&CfnModelCardMixinProps{
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
//   								Notes: jsii.String("notes"),
//   								Type: jsii.String("type"),
//   								Value: value,
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
//   				Name: jsii.String("name"),
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
//   		ModelPackageDetails: &ModelPackageDetailsProperty{
//   			ApprovalDescription: jsii.String("approvalDescription"),
//   			CreatedBy: &ModelPackageCreatorProperty{
//   				UserProfileName: jsii.String("userProfileName"),
//   			},
//   			Domain: jsii.String("domain"),
//   			InferenceSpecification: &InferenceSpecificationProperty{
//   				Containers: []interface{}{
//   					&ContainerProperty{
//   						Image: jsii.String("image"),
//   						ModelDataUrl: jsii.String("modelDataUrl"),
//   						NearestModelName: jsii.String("nearestModelName"),
//   					},
//   				},
//   			},
//   			ModelApprovalStatus: jsii.String("modelApprovalStatus"),
//   			ModelPackageArn: jsii.String("modelPackageArn"),
//   			ModelPackageDescription: jsii.String("modelPackageDescription"),
//   			ModelPackageGroupName: jsii.String("modelPackageGroupName"),
//   			ModelPackageName: jsii.String("modelPackageName"),
//   			ModelPackageStatus: jsii.String("modelPackageStatus"),
//   			ModelPackageVersion: jsii.Number(123),
//   			SourceAlgorithms: []interface{}{
//   				&SourceAlgorithmProperty{
//   					AlgorithmName: jsii.String("algorithmName"),
//   					ModelDataUrl: jsii.String("modelDataUrl"),
//   				},
//   			},
//   			Task: jsii.String("task"),
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
//   						Notes: jsii.String("notes"),
//   						Value: jsii.Number(123),
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
//   						Notes: jsii.String("notes"),
//   						Value: jsii.Number(123),
//   					},
//   				},
//   			},
//   			TrainingObservations: jsii.String("trainingObservations"),
//   		},
//   	},
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
//   	ModelCardName: jsii.String("modelCardName"),
//   	ModelCardStatus: jsii.String("modelCardStatus"),
//   	SecurityConfig: &SecurityConfigProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelcard.html
//
type CfnModelCardPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnModelCardMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnModelCardPropsMixin
type jsiiProxy_CfnModelCardPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnModelCardPropsMixin) Props() *CfnModelCardMixinProps {
	var returns *CfnModelCardMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelCardPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::ModelCard`.
func NewCfnModelCardPropsMixin(props *CfnModelCardMixinProps, options *mixins.CfnPropertyMixinOptions) CfnModelCardPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnModelCardPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnModelCardPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::ModelCard`.
func NewCfnModelCardPropsMixin_Override(c CfnModelCardPropsMixin, props *CfnModelCardMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnModelCardPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelCardPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModelCardPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModelCardPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnModelCardPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

