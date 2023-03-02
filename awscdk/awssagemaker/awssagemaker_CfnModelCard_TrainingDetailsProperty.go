package awssagemaker


// The training details of the model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainingDetailsProperty := &trainingDetailsProperty{
//   	objectiveFunction: &objectiveFunctionProperty{
//   		function: &functionProperty{
//   			condition: jsii.String("condition"),
//   			facet: jsii.String("facet"),
//   			function: jsii.String("function"),
//   		},
//   		notes: jsii.String("notes"),
//   	},
//   	trainingJobDetails: &trainingJobDetailsProperty{
//   		hyperParameters: []interface{}{
//   			&trainingHyperParameterProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		trainingArn: jsii.String("trainingArn"),
//   		trainingDatasets: []*string{
//   			jsii.String("trainingDatasets"),
//   		},
//   		trainingEnvironment: &trainingEnvironmentProperty{
//   			containerImage: []*string{
//   				jsii.String("containerImage"),
//   			},
//   		},
//   		trainingMetrics: []interface{}{
//   			&trainingMetricProperty{
//   				name: jsii.String("name"),
//   				value: jsii.Number(123),
//
//   				// the properties below are optional
//   				notes: jsii.String("notes"),
//   			},
//   		},
//   		userProvidedHyperParameters: []interface{}{
//   			&trainingHyperParameterProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		userProvidedTrainingMetrics: []interface{}{
//   			&trainingMetricProperty{
//   				name: jsii.String("name"),
//   				value: jsii.Number(123),
//
//   				// the properties below are optional
//   				notes: jsii.String("notes"),
//   			},
//   		},
//   	},
//   	trainingObservations: jsii.String("trainingObservations"),
//   }
//
type CfnModelCard_TrainingDetailsProperty struct {
	// The function that is optimized during model training.
	ObjectiveFunction interface{} `field:"optional" json:"objectiveFunction" yaml:"objectiveFunction"`
	// Details about any associated training jobs.
	TrainingJobDetails interface{} `field:"optional" json:"trainingJobDetails" yaml:"trainingJobDetails"`
	// Any observations about training.
	TrainingObservations *string `field:"optional" json:"trainingObservations" yaml:"trainingObservations"`
}

