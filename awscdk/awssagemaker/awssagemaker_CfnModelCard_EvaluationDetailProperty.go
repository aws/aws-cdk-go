package awssagemaker


// The evaluation details of the model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   evaluationDetailProperty := &evaluationDetailProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	datasets: []*string{
//   		jsii.String("datasets"),
//   	},
//   	evaluationJobArn: jsii.String("evaluationJobArn"),
//   	evaluationObservation: jsii.String("evaluationObservation"),
//   	metadata: map[string]*string{
//   		"metadataKey": jsii.String("metadata"),
//   	},
//   	metricGroups: []interface{}{
//   		&metricGroupProperty{
//   			metricData: []interface{}{
//   				&metricDataItemsProperty{
//   					name: jsii.String("name"),
//   					type: jsii.String("type"),
//   					value: value,
//
//   					// the properties below are optional
//   					notes: jsii.String("notes"),
//   					xAxisName: []*string{
//   						jsii.String("xAxisName"),
//   					},
//   					yAxisName: []*string{
//   						jsii.String("yAxisName"),
//   					},
//   				},
//   			},
//   			name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnModelCard_EvaluationDetailProperty struct {
	// The evaluation job name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The location of the datasets used to evaluate the model.
	Datasets *[]*string `field:"optional" json:"datasets" yaml:"datasets"`
	// The Amazon Resource Name (ARN) of the evaluation job.
	EvaluationJobArn *string `field:"optional" json:"evaluationJobArn" yaml:"evaluationJobArn"`
	// Any observations made during the model evaluation.
	EvaluationObservation *string `field:"optional" json:"evaluationObservation" yaml:"evaluationObservation"`
	// Additional attributes associated with the evaluation results.
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// An evaluation Metric Group object.
	MetricGroups interface{} `field:"optional" json:"metricGroups" yaml:"metricGroups"`
}

