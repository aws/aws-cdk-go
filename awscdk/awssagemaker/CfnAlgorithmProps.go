package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAlgorithm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlgorithmProps := &CfnAlgorithmProps{
//   	AlgorithmName: jsii.String("algorithmName"),
//   	TrainingSpecification: &TrainingSpecificationProperty{
//   		SupportedTrainingInstanceTypes: []*string{
//   			jsii.String("supportedTrainingInstanceTypes"),
//   		},
//   		TrainingChannels: []interface{}{
//   			&ChannelSpecificationProperty{
//   				Name: jsii.String("name"),
//   				SupportedContentTypes: []*string{
//   					jsii.String("supportedContentTypes"),
//   				},
//   				SupportedInputModes: []*string{
//   					jsii.String("supportedInputModes"),
//   				},
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   				IsRequired: jsii.Boolean(false),
//   				SupportedCompressionTypes: []*string{
//   					jsii.String("supportedCompressionTypes"),
//   				},
//   			},
//   		},
//   		TrainingImage: jsii.String("trainingImage"),
//
//   		// the properties below are optional
//   		MetricDefinitions: []interface{}{
//   			&MetricDefinitionProperty{
//   				Name: jsii.String("name"),
//   				Regex: jsii.String("regex"),
//   			},
//   		},
//   		SupportedHyperParameters: []interface{}{
//   			&HyperParameterSpecificationProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				DefaultValue: jsii.String("defaultValue"),
//   				Description: jsii.String("description"),
//   				IsRequired: jsii.Boolean(false),
//   				IsTunable: jsii.Boolean(false),
//   				Range: &ParameterRangeProperty{
//   					CategoricalParameterRangeSpecification: &CategoricalParameterRangeSpecificationProperty{
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					ContinuousParameterRangeSpecification: &ContinuousParameterRangeSpecificationProperty{
//   						MaxValue: jsii.String("maxValue"),
//   						MinValue: jsii.String("minValue"),
//   					},
//   					IntegerParameterRangeSpecification: &IntegerParameterRangeSpecificationProperty{
//   						MaxValue: jsii.String("maxValue"),
//   						MinValue: jsii.String("minValue"),
//   					},
//   				},
//   			},
//   		},
//   		SupportedTuningJobObjectiveMetrics: []interface{}{
//   			&HyperParameterTuningJobObjectiveProperty{
//   				MetricName: jsii.String("metricName"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		SupportsDistributedTraining: jsii.Boolean(false),
//   		TrainingImageDigest: jsii.String("trainingImageDigest"),
//   	},
//
//   	// the properties below are optional
//   	AlgorithmDescription: jsii.String("algorithmDescription"),
//   	CertifyForMarketplace: jsii.Boolean(false),
//   	InferenceSpecification: &InferenceSpecificationProperty{
//   		Containers: []interface{}{
//   			&ModelPackageContainerDefinitionProperty{
//   				Image: jsii.String("image"),
//
//   				// the properties below are optional
//   				ContainerHostname: jsii.String("containerHostname"),
//   				Environment: map[string]*string{
//   					"environmentKey": jsii.String("environment"),
//   				},
//   				Framework: jsii.String("framework"),
//   				FrameworkVersion: jsii.String("frameworkVersion"),
//   				ImageDigest: jsii.String("imageDigest"),
//   				IsCheckpoint: jsii.Boolean(false),
//   				ModelInput: &ModelInputProperty{
//   					DataInputConfig: jsii.String("dataInputConfig"),
//   				},
//   				NearestModelName: jsii.String("nearestModelName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		SupportedContentTypes: []*string{
//   			jsii.String("supportedContentTypes"),
//   		},
//   		SupportedRealtimeInferenceInstanceTypes: []*string{
//   			jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   		},
//   		SupportedResponseMimeTypes: []*string{
//   			jsii.String("supportedResponseMimeTypes"),
//   		},
//   		SupportedTransformInstanceTypes: []*string{
//   			jsii.String("supportedTransformInstanceTypes"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-algorithm.html
//
type CfnAlgorithmProps struct {
	// The name of the algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-algorithm.html#cfn-sagemaker-algorithm-algorithmname
	//
	AlgorithmName *string `field:"required" json:"algorithmName" yaml:"algorithmName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-algorithm.html#cfn-sagemaker-algorithm-trainingspecification
	//
	TrainingSpecification interface{} `field:"required" json:"trainingSpecification" yaml:"trainingSpecification"`
	// A description of the algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-algorithm.html#cfn-sagemaker-algorithm-algorithmdescription
	//
	AlgorithmDescription *string `field:"optional" json:"algorithmDescription" yaml:"algorithmDescription"`
	// Whether to certify the algorithm so that it can be listed in AWS Marketplace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-algorithm.html#cfn-sagemaker-algorithm-certifyformarketplace
	//
	CertifyForMarketplace interface{} `field:"optional" json:"certifyForMarketplace" yaml:"certifyForMarketplace"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-algorithm.html#cfn-sagemaker-algorithm-inferencespecification
	//
	InferenceSpecification interface{} `field:"optional" json:"inferenceSpecification" yaml:"inferenceSpecification"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-algorithm.html#cfn-sagemaker-algorithm-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

