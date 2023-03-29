package awssagemaker


// The content of the model card.
//
// It follows the [model card json schema](https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html#model-cards-json-schema) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   contentProperty := &ContentProperty{
//   	AdditionalInformation: &AdditionalInformationProperty{
//   		CaveatsAndRecommendations: jsii.String("caveatsAndRecommendations"),
//   		CustomDetails: map[string]*string{
//   			"customDetailsKey": jsii.String("customDetails"),
//   		},
//   		EthicalConsiderations: jsii.String("ethicalConsiderations"),
//   	},
//   	BusinessDetails: &BusinessDetailsProperty{
//   		BusinessProblem: jsii.String("businessProblem"),
//   		BusinessStakeholders: jsii.String("businessStakeholders"),
//   		LineOfBusiness: jsii.String("lineOfBusiness"),
//   	},
//   	EvaluationDetails: []interface{}{
//   		&EvaluationDetailProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Datasets: []*string{
//   				jsii.String("datasets"),
//   			},
//   			EvaluationJobArn: jsii.String("evaluationJobArn"),
//   			EvaluationObservation: jsii.String("evaluationObservation"),
//   			Metadata: map[string]*string{
//   				"metadataKey": jsii.String("metadata"),
//   			},
//   			MetricGroups: []interface{}{
//   				&MetricGroupProperty{
//   					MetricData: []interface{}{
//   						&MetricDataItemsProperty{
//   							Name: jsii.String("name"),
//   							Type: jsii.String("type"),
//   							Value: value,
//
//   							// the properties below are optional
//   							Notes: jsii.String("notes"),
//   							XAxisName: []*string{
//   								jsii.String("xAxisName"),
//   							},
//   							YAxisName: []*string{
//   								jsii.String("yAxisName"),
//   							},
//   						},
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	IntendedUses: &IntendedUsesProperty{
//   		ExplanationsForRiskRating: jsii.String("explanationsForRiskRating"),
//   		FactorsAffectingModelEfficiency: jsii.String("factorsAffectingModelEfficiency"),
//   		IntendedUses: jsii.String("intendedUses"),
//   		PurposeOfModel: jsii.String("purposeOfModel"),
//   		RiskRating: jsii.String("riskRating"),
//   	},
//   	ModelOverview: &ModelOverviewProperty{
//   		AlgorithmType: jsii.String("algorithmType"),
//   		InferenceEnvironment: &InferenceEnvironmentProperty{
//   			ContainerImage: []*string{
//   				jsii.String("containerImage"),
//   			},
//   		},
//   		ModelArtifact: []*string{
//   			jsii.String("modelArtifact"),
//   		},
//   		ModelCreator: jsii.String("modelCreator"),
//   		ModelDescription: jsii.String("modelDescription"),
//   		ModelId: jsii.String("modelId"),
//   		ModelName: jsii.String("modelName"),
//   		ModelOwner: jsii.String("modelOwner"),
//   		ModelVersion: jsii.Number(123),
//   		ProblemType: jsii.String("problemType"),
//   	},
//   	TrainingDetails: &TrainingDetailsProperty{
//   		ObjectiveFunction: &ObjectiveFunctionProperty{
//   			Function: &FunctionProperty{
//   				Condition: jsii.String("condition"),
//   				Facet: jsii.String("facet"),
//   				Function: jsii.String("function"),
//   			},
//   			Notes: jsii.String("notes"),
//   		},
//   		TrainingJobDetails: &TrainingJobDetailsProperty{
//   			HyperParameters: []interface{}{
//   				&TrainingHyperParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			TrainingArn: jsii.String("trainingArn"),
//   			TrainingDatasets: []*string{
//   				jsii.String("trainingDatasets"),
//   			},
//   			TrainingEnvironment: &TrainingEnvironmentProperty{
//   				ContainerImage: []*string{
//   					jsii.String("containerImage"),
//   				},
//   			},
//   			TrainingMetrics: []interface{}{
//   				&TrainingMetricProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.Number(123),
//
//   					// the properties below are optional
//   					Notes: jsii.String("notes"),
//   				},
//   			},
//   			UserProvidedHyperParameters: []interface{}{
//   				&TrainingHyperParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			UserProvidedTrainingMetrics: []interface{}{
//   				&TrainingMetricProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.Number(123),
//
//   					// the properties below are optional
//   					Notes: jsii.String("notes"),
//   				},
//   			},
//   		},
//   		TrainingObservations: jsii.String("trainingObservations"),
//   	},
//   }
//
type CfnModelCard_ContentProperty struct {
	// Additional information about the model.
	AdditionalInformation interface{} `field:"optional" json:"additionalInformation" yaml:"additionalInformation"`
	// Information about how the model supports business goals.
	BusinessDetails interface{} `field:"optional" json:"businessDetails" yaml:"businessDetails"`
	// An overview about the model's evaluation.
	EvaluationDetails interface{} `field:"optional" json:"evaluationDetails" yaml:"evaluationDetails"`
	// The intended usage of the model.
	IntendedUses interface{} `field:"optional" json:"intendedUses" yaml:"intendedUses"`
	// An overview about the model.
	ModelOverview interface{} `field:"optional" json:"modelOverview" yaml:"modelOverview"`
	// An overview about model training.
	TrainingDetails interface{} `field:"optional" json:"trainingDetails" yaml:"trainingDetails"`
}

