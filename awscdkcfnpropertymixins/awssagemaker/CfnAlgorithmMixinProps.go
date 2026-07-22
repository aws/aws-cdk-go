package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAlgorithmPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAlgorithmMixinProps := &CfnAlgorithmMixinProps{
//   	AlgorithmDescription: jsii.String("algorithmDescription"),
//   	AlgorithmName: jsii.String("algorithmName"),
//   	CertifyForMarketplace: jsii.Boolean(false),
//   	InferenceSpecification: &InferenceSpecificationProperty{
//   		Containers: []interface{}{
//   			&ModelPackageContainerDefinitionProperty{
//   				ContainerHostname: jsii.String("containerHostname"),
//   				Environment: map[string]*string{
//   					"environmentKey": jsii.String("environment"),
//   				},
//   				Framework: jsii.String("framework"),
//   				FrameworkVersion: jsii.String("frameworkVersion"),
//   				Image: jsii.String("image"),
//   				ImageDigest: jsii.String("imageDigest"),
//   				IsCheckpoint: jsii.Boolean(false),
//   				ModelInput: &ModelInputProperty{
//   					DataInputConfig: jsii.String("dataInputConfig"),
//   				},
//   				NearestModelName: jsii.String("nearestModelName"),
//   			},
//   		},
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
//   	TrainingSpecification: &TrainingSpecificationProperty{
//   		MetricDefinitions: []interface{}{
//   			&MetricDefinitionProperty{
//   				Name: jsii.String("name"),
//   				Regex: jsii.String("regex"),
//   			},
//   		},
//   		SupportedHyperParameters: []interface{}{
//   			&HyperParameterSpecificationProperty{
//   				DefaultValue: jsii.String("defaultValue"),
//   				Description: jsii.String("description"),
//   				IsRequired: jsii.Boolean(false),
//   				IsTunable: jsii.Boolean(false),
//   				Name: jsii.String("name"),
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
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		SupportedTrainingInstanceTypes: []*string{
//   			jsii.String("supportedTrainingInstanceTypes"),
//   		},
//   		SupportedTuningJobObjectiveMetrics: []interface{}{
//   			&HyperParameterTuningJobObjectiveProperty{
//   				MetricName: jsii.String("metricName"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		SupportsDistributedTraining: jsii.Boolean(false),
//   		TrainingChannels: []interface{}{
//   			&ChannelSpecificationProperty{
//   				Description: jsii.String("description"),
//   				IsRequired: jsii.Boolean(false),
//   				Name: jsii.String("name"),
//   				SupportedCompressionTypes: []*string{
//   					jsii.String("supportedCompressionTypes"),
//   				},
//   				SupportedContentTypes: []*string{
//   					jsii.String("supportedContentTypes"),
//   				},
//   				SupportedInputModes: []*string{
//   					jsii.String("supportedInputModes"),
//   				},
//   			},
//   		},
//   		TrainingImage: jsii.String("trainingImage"),
//   		TrainingImageDigest: jsii.String("trainingImageDigest"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-algorithm.html
//
type CfnAlgorithmMixinProps struct {
	// A description of the algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-algorithm.html#cfn-sagemaker-algorithm-algorithmdescription
	//
	AlgorithmDescription *string `field:"optional" json:"algorithmDescription" yaml:"algorithmDescription"`
	// The name of the algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-algorithm.html#cfn-sagemaker-algorithm-algorithmname
	//
	AlgorithmName *string `field:"optional" json:"algorithmName" yaml:"algorithmName"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-algorithm.html#cfn-sagemaker-algorithm-trainingspecification
	//
	TrainingSpecification interface{} `field:"optional" json:"trainingSpecification" yaml:"trainingSpecification"`
}

