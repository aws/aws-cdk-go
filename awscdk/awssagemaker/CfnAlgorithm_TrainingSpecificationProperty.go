package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainingSpecificationProperty := &TrainingSpecificationProperty{
//   	SupportedTrainingInstanceTypes: []*string{
//   		jsii.String("supportedTrainingInstanceTypes"),
//   	},
//   	TrainingChannels: []interface{}{
//   		&ChannelSpecificationProperty{
//   			Name: jsii.String("name"),
//   			SupportedContentTypes: []*string{
//   				jsii.String("supportedContentTypes"),
//   			},
//   			SupportedInputModes: []*string{
//   				jsii.String("supportedInputModes"),
//   			},
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			IsRequired: jsii.Boolean(false),
//   			SupportedCompressionTypes: []*string{
//   				jsii.String("supportedCompressionTypes"),
//   			},
//   		},
//   	},
//   	TrainingImage: jsii.String("trainingImage"),
//
//   	// the properties below are optional
//   	MetricDefinitions: []interface{}{
//   		&MetricDefinitionProperty{
//   			Name: jsii.String("name"),
//   			Regex: jsii.String("regex"),
//   		},
//   	},
//   	SupportedHyperParameters: []interface{}{
//   		&HyperParameterSpecificationProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			DefaultValue: jsii.String("defaultValue"),
//   			Description: jsii.String("description"),
//   			IsRequired: jsii.Boolean(false),
//   			IsTunable: jsii.Boolean(false),
//   			Range: &ParameterRangeProperty{
//   				CategoricalParameterRangeSpecification: &CategoricalParameterRangeSpecificationProperty{
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				ContinuousParameterRangeSpecification: &ContinuousParameterRangeSpecificationProperty{
//   					MaxValue: jsii.String("maxValue"),
//   					MinValue: jsii.String("minValue"),
//   				},
//   				IntegerParameterRangeSpecification: &IntegerParameterRangeSpecificationProperty{
//   					MaxValue: jsii.String("maxValue"),
//   					MinValue: jsii.String("minValue"),
//   				},
//   			},
//   		},
//   	},
//   	SupportedTuningJobObjectiveMetrics: []interface{}{
//   		&HyperParameterTuningJobObjectiveProperty{
//   			MetricName: jsii.String("metricName"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SupportsDistributedTraining: jsii.Boolean(false),
//   	TrainingImageDigest: jsii.String("trainingImageDigest"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-trainingspecification.html
//
type CfnAlgorithm_TrainingSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-trainingspecification.html#cfn-sagemaker-algorithm-trainingspecification-supportedtraininginstancetypes
	//
	SupportedTrainingInstanceTypes *[]*string `field:"required" json:"supportedTrainingInstanceTypes" yaml:"supportedTrainingInstanceTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-trainingspecification.html#cfn-sagemaker-algorithm-trainingspecification-trainingchannels
	//
	TrainingChannels interface{} `field:"required" json:"trainingChannels" yaml:"trainingChannels"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-trainingspecification.html#cfn-sagemaker-algorithm-trainingspecification-trainingimage
	//
	TrainingImage *string `field:"required" json:"trainingImage" yaml:"trainingImage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-trainingspecification.html#cfn-sagemaker-algorithm-trainingspecification-metricdefinitions
	//
	MetricDefinitions interface{} `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-trainingspecification.html#cfn-sagemaker-algorithm-trainingspecification-supportedhyperparameters
	//
	SupportedHyperParameters interface{} `field:"optional" json:"supportedHyperParameters" yaml:"supportedHyperParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-trainingspecification.html#cfn-sagemaker-algorithm-trainingspecification-supportedtuningjobobjectivemetrics
	//
	SupportedTuningJobObjectiveMetrics interface{} `field:"optional" json:"supportedTuningJobObjectiveMetrics" yaml:"supportedTuningJobObjectiveMetrics"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-trainingspecification.html#cfn-sagemaker-algorithm-trainingspecification-supportsdistributedtraining
	//
	SupportsDistributedTraining interface{} `field:"optional" json:"supportsDistributedTraining" yaml:"supportsDistributedTraining"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-trainingspecification.html#cfn-sagemaker-algorithm-trainingspecification-trainingimagedigest
	//
	TrainingImageDigest *string `field:"optional" json:"trainingImageDigest" yaml:"trainingImageDigest"`
}

