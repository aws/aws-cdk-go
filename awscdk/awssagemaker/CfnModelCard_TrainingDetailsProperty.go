package awssagemaker


// The training details of the model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainingDetailsProperty := &TrainingDetailsProperty{
//   	ObjectiveFunction: &ObjectiveFunctionProperty{
//   		Function: &FunctionProperty{
//   			Condition: jsii.String("condition"),
//   			Facet: jsii.String("facet"),
//   			Function: jsii.String("function"),
//   		},
//   		Notes: jsii.String("notes"),
//   	},
//   	TrainingJobDetails: &TrainingJobDetailsProperty{
//   		HyperParameters: []interface{}{
//   			&TrainingHyperParameterProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		TrainingArn: jsii.String("trainingArn"),
//   		TrainingDatasets: []*string{
//   			jsii.String("trainingDatasets"),
//   		},
//   		TrainingEnvironment: &TrainingEnvironmentProperty{
//   			ContainerImage: []*string{
//   				jsii.String("containerImage"),
//   			},
//   		},
//   		TrainingMetrics: []interface{}{
//   			&TrainingMetricProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.Number(123),
//
//   				// the properties below are optional
//   				Notes: jsii.String("notes"),
//   			},
//   		},
//   		UserProvidedHyperParameters: []interface{}{
//   			&TrainingHyperParameterProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		UserProvidedTrainingMetrics: []interface{}{
//   			&TrainingMetricProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.Number(123),
//
//   				// the properties below are optional
//   				Notes: jsii.String("notes"),
//   			},
//   		},
//   	},
//   	TrainingObservations: jsii.String("trainingObservations"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingdetails.html
//
type CfnModelCard_TrainingDetailsProperty struct {
	// The function that is optimized during model training.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingdetails.html#cfn-sagemaker-modelcard-trainingdetails-objectivefunction
	//
	ObjectiveFunction interface{} `field:"optional" json:"objectiveFunction" yaml:"objectiveFunction"`
	// Details about any associated training jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingdetails.html#cfn-sagemaker-modelcard-trainingdetails-trainingjobdetails
	//
	TrainingJobDetails interface{} `field:"optional" json:"trainingJobDetails" yaml:"trainingJobDetails"`
	// Any observations about training.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingdetails.html#cfn-sagemaker-modelcard-trainingdetails-trainingobservations
	//
	TrainingObservations *string `field:"optional" json:"trainingObservations" yaml:"trainingObservations"`
}

