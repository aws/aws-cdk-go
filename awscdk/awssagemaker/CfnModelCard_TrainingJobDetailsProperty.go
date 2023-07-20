package awssagemaker


// The overview of a training job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainingJobDetailsProperty := &TrainingJobDetailsProperty{
//   	HyperParameters: []interface{}{
//   		&TrainingHyperParameterProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrainingArn: jsii.String("trainingArn"),
//   	TrainingDatasets: []*string{
//   		jsii.String("trainingDatasets"),
//   	},
//   	TrainingEnvironment: &TrainingEnvironmentProperty{
//   		ContainerImage: []*string{
//   			jsii.String("containerImage"),
//   		},
//   	},
//   	TrainingMetrics: []interface{}{
//   		&TrainingMetricProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.Number(123),
//
//   			// the properties below are optional
//   			Notes: jsii.String("notes"),
//   		},
//   	},
//   	UserProvidedHyperParameters: []interface{}{
//   		&TrainingHyperParameterProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserProvidedTrainingMetrics: []interface{}{
//   		&TrainingMetricProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.Number(123),
//
//   			// the properties below are optional
//   			Notes: jsii.String("notes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingjobdetails.html
//
type CfnModelCard_TrainingJobDetailsProperty struct {
	// The hyper parameters used in the training job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingjobdetails.html#cfn-sagemaker-modelcard-trainingjobdetails-hyperparameters
	//
	HyperParameters interface{} `field:"optional" json:"hyperParameters" yaml:"hyperParameters"`
	// The SageMaker training job Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingjobdetails.html#cfn-sagemaker-modelcard-trainingjobdetails-trainingarn
	//
	TrainingArn *string `field:"optional" json:"trainingArn" yaml:"trainingArn"`
	// The location of the datasets used to train the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingjobdetails.html#cfn-sagemaker-modelcard-trainingjobdetails-trainingdatasets
	//
	TrainingDatasets *[]*string `field:"optional" json:"trainingDatasets" yaml:"trainingDatasets"`
	// The SageMaker training job image URI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingjobdetails.html#cfn-sagemaker-modelcard-trainingjobdetails-trainingenvironment
	//
	TrainingEnvironment interface{} `field:"optional" json:"trainingEnvironment" yaml:"trainingEnvironment"`
	// The SageMaker training job results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingjobdetails.html#cfn-sagemaker-modelcard-trainingjobdetails-trainingmetrics
	//
	TrainingMetrics interface{} `field:"optional" json:"trainingMetrics" yaml:"trainingMetrics"`
	// Additional hyper parameters that you've specified when training the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingjobdetails.html#cfn-sagemaker-modelcard-trainingjobdetails-userprovidedhyperparameters
	//
	UserProvidedHyperParameters interface{} `field:"optional" json:"userProvidedHyperParameters" yaml:"userProvidedHyperParameters"`
	// Custom training job results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingjobdetails.html#cfn-sagemaker-modelcard-trainingjobdetails-userprovidedtrainingmetrics
	//
	UserProvidedTrainingMetrics interface{} `field:"optional" json:"userProvidedTrainingMetrics" yaml:"userProvidedTrainingMetrics"`
}

