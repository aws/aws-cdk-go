package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnModelCard`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   cfnModelCardProps := &cfnModelCardProps{
//   	content: &contentProperty{
//   		additionalInformation: &additionalInformationProperty{
//   			caveatsAndRecommendations: jsii.String("caveatsAndRecommendations"),
//   			customDetails: map[string]*string{
//   				"customDetailsKey": jsii.String("customDetails"),
//   			},
//   			ethicalConsiderations: jsii.String("ethicalConsiderations"),
//   		},
//   		businessDetails: &businessDetailsProperty{
//   			businessProblem: jsii.String("businessProblem"),
//   			businessStakeholders: jsii.String("businessStakeholders"),
//   			lineOfBusiness: jsii.String("lineOfBusiness"),
//   		},
//   		evaluationDetails: []interface{}{
//   			&evaluationDetailProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				datasets: []*string{
//   					jsii.String("datasets"),
//   				},
//   				evaluationJobArn: jsii.String("evaluationJobArn"),
//   				evaluationObservation: jsii.String("evaluationObservation"),
//   				metadata: map[string]*string{
//   					"metadataKey": jsii.String("metadata"),
//   				},
//   				metricGroups: []interface{}{
//   					&metricGroupProperty{
//   						metricData: []interface{}{
//   							&metricDataItemsProperty{
//   								name: jsii.String("name"),
//   								type: jsii.String("type"),
//   								value: value,
//
//   								// the properties below are optional
//   								notes: jsii.String("notes"),
//   								xAxisName: []*string{
//   									jsii.String("xAxisName"),
//   								},
//   								yAxisName: []*string{
//   									jsii.String("yAxisName"),
//   								},
//   							},
//   						},
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   		intendedUses: &intendedUsesProperty{
//   			explanationsForRiskRating: jsii.String("explanationsForRiskRating"),
//   			factorsAffectingModelEfficiency: jsii.String("factorsAffectingModelEfficiency"),
//   			intendedUses: jsii.String("intendedUses"),
//   			purposeOfModel: jsii.String("purposeOfModel"),
//   			riskRating: jsii.String("riskRating"),
//   		},
//   		modelOverview: &modelOverviewProperty{
//   			algorithmType: jsii.String("algorithmType"),
//   			inferenceEnvironment: &inferenceEnvironmentProperty{
//   				containerImage: []*string{
//   					jsii.String("containerImage"),
//   				},
//   			},
//   			modelArtifact: []*string{
//   				jsii.String("modelArtifact"),
//   			},
//   			modelCreator: jsii.String("modelCreator"),
//   			modelDescription: jsii.String("modelDescription"),
//   			modelId: jsii.String("modelId"),
//   			modelName: jsii.String("modelName"),
//   			modelOwner: jsii.String("modelOwner"),
//   			modelVersion: jsii.Number(123),
//   			problemType: jsii.String("problemType"),
//   		},
//   		trainingDetails: &trainingDetailsProperty{
//   			objectiveFunction: &objectiveFunctionProperty{
//   				function: &functionProperty{
//   					condition: jsii.String("condition"),
//   					facet: jsii.String("facet"),
//   					function: jsii.String("function"),
//   				},
//   				notes: jsii.String("notes"),
//   			},
//   			trainingJobDetails: &trainingJobDetailsProperty{
//   				hyperParameters: []interface{}{
//   					&trainingHyperParameterProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				trainingArn: jsii.String("trainingArn"),
//   				trainingDatasets: []*string{
//   					jsii.String("trainingDatasets"),
//   				},
//   				trainingEnvironment: &trainingEnvironmentProperty{
//   					containerImage: []*string{
//   						jsii.String("containerImage"),
//   					},
//   				},
//   				trainingMetrics: []interface{}{
//   					&trainingMetricProperty{
//   						name: jsii.String("name"),
//   						value: jsii.Number(123),
//
//   						// the properties below are optional
//   						notes: jsii.String("notes"),
//   					},
//   				},
//   				userProvidedHyperParameters: []interface{}{
//   					&trainingHyperParameterProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				userProvidedTrainingMetrics: []interface{}{
//   					&trainingMetricProperty{
//   						name: jsii.String("name"),
//   						value: jsii.Number(123),
//
//   						// the properties below are optional
//   						notes: jsii.String("notes"),
//   					},
//   				},
//   			},
//   			trainingObservations: jsii.String("trainingObservations"),
//   		},
//   	},
//   	modelCardName: jsii.String("modelCardName"),
//   	modelCardStatus: jsii.String("modelCardStatus"),
//
//   	// the properties below are optional
//   	createdBy: &userContextProperty{
//   		domainId: jsii.String("domainId"),
//   		userProfileArn: jsii.String("userProfileArn"),
//   		userProfileName: jsii.String("userProfileName"),
//   	},
//   	lastModifiedBy: &userContextProperty{
//   		domainId: jsii.String("domainId"),
//   		userProfileArn: jsii.String("userProfileArn"),
//   		userProfileName: jsii.String("userProfileName"),
//   	},
//   	securityConfig: &securityConfigProperty{
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnModelCardProps struct {
	// The content of the model card.
	//
	// Content uses the [model card JSON schema](https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html#model-cards-json-schema) .
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// The unique name of the model card.
	ModelCardName *string `field:"required" json:"modelCardName" yaml:"modelCardName"`
	// The approval status of the model card within your organization.
	//
	// Different organizations might have different criteria for model card review and approval.
	//
	// - `Draft` : The model card is a work in progress.
	// - `PendingReview` : The model card is pending review.
	// - `Approved` : The model card is approved.
	// - `Archived` : The model card is archived. No more updates should be made to the model card, but it can still be exported.
	ModelCardStatus *string `field:"required" json:"modelCardStatus" yaml:"modelCardStatus"`
	// `AWS::SageMaker::ModelCard.CreatedBy`.
	CreatedBy interface{} `field:"optional" json:"createdBy" yaml:"createdBy"`
	// `AWS::SageMaker::ModelCard.LastModifiedBy`.
	LastModifiedBy interface{} `field:"optional" json:"lastModifiedBy" yaml:"lastModifiedBy"`
	// The security configuration used to protect model card data.
	SecurityConfig interface{} `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// Key-value pairs used to manage metadata for the model card.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

