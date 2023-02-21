package awssagemaker


// The overview of a training job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainingJobDetailsProperty := &trainingJobDetailsProperty{
//   	hyperParameters: []interface{}{
//   		&trainingHyperParameterProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	trainingArn: jsii.String("trainingArn"),
//   	trainingDatasets: []*string{
//   		jsii.String("trainingDatasets"),
//   	},
//   	trainingEnvironment: &trainingEnvironmentProperty{
//   		containerImage: []*string{
//   			jsii.String("containerImage"),
//   		},
//   	},
//   	trainingMetrics: []interface{}{
//   		&trainingMetricProperty{
//   			name: jsii.String("name"),
//   			value: jsii.Number(123),
//
//   			// the properties below are optional
//   			notes: jsii.String("notes"),
//   		},
//   	},
//   	userProvidedHyperParameters: []interface{}{
//   		&trainingHyperParameterProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userProvidedTrainingMetrics: []interface{}{
//   		&trainingMetricProperty{
//   			name: jsii.String("name"),
//   			value: jsii.Number(123),
//
//   			// the properties below are optional
//   			notes: jsii.String("notes"),
//   		},
//   	},
//   }
//
type CfnModelCard_TrainingJobDetailsProperty struct {
	// The hyper parameters used in the training job.
	HyperParameters interface{} `field:"optional" json:"hyperParameters" yaml:"hyperParameters"`
	// The SageMaker training job Amazon Resource Name (ARN).
	TrainingArn *string `field:"optional" json:"trainingArn" yaml:"trainingArn"`
	// The location of the datasets used to train the model.
	TrainingDatasets *[]*string `field:"optional" json:"trainingDatasets" yaml:"trainingDatasets"`
	// The SageMaker training job image URI.
	TrainingEnvironment interface{} `field:"optional" json:"trainingEnvironment" yaml:"trainingEnvironment"`
	// The SageMaker training job results.
	TrainingMetrics interface{} `field:"optional" json:"trainingMetrics" yaml:"trainingMetrics"`
	// Additional hyper parameters that you've specified when training the model.
	UserProvidedHyperParameters interface{} `field:"optional" json:"userProvidedHyperParameters" yaml:"userProvidedHyperParameters"`
	// Custom training job results.
	UserProvidedTrainingMetrics interface{} `field:"optional" json:"userProvidedTrainingMetrics" yaml:"userProvidedTrainingMetrics"`
}

