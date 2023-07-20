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
//   evaluationDetailProperty := &EvaluationDetailProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Datasets: []*string{
//   		jsii.String("datasets"),
//   	},
//   	EvaluationJobArn: jsii.String("evaluationJobArn"),
//   	EvaluationObservation: jsii.String("evaluationObservation"),
//   	Metadata: map[string]*string{
//   		"metadataKey": jsii.String("metadata"),
//   	},
//   	MetricGroups: []interface{}{
//   		&MetricGroupProperty{
//   			MetricData: []interface{}{
//   				&MetricDataItemsProperty{
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//   					Value: value,
//
//   					// the properties below are optional
//   					Notes: jsii.String("notes"),
//   					XAxisName: []*string{
//   						jsii.String("xAxisName"),
//   					},
//   					YAxisName: []*string{
//   						jsii.String("yAxisName"),
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-evaluationdetail.html
//
type CfnModelCard_EvaluationDetailProperty struct {
	// The evaluation job name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-evaluationdetail.html#cfn-sagemaker-modelcard-evaluationdetail-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The location of the datasets used to evaluate the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-evaluationdetail.html#cfn-sagemaker-modelcard-evaluationdetail-datasets
	//
	Datasets *[]*string `field:"optional" json:"datasets" yaml:"datasets"`
	// The Amazon Resource Name (ARN) of the evaluation job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-evaluationdetail.html#cfn-sagemaker-modelcard-evaluationdetail-evaluationjobarn
	//
	EvaluationJobArn *string `field:"optional" json:"evaluationJobArn" yaml:"evaluationJobArn"`
	// Any observations made during the model evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-evaluationdetail.html#cfn-sagemaker-modelcard-evaluationdetail-evaluationobservation
	//
	EvaluationObservation *string `field:"optional" json:"evaluationObservation" yaml:"evaluationObservation"`
	// Additional attributes associated with the evaluation results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-evaluationdetail.html#cfn-sagemaker-modelcard-evaluationdetail-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// An evaluation Metric Group object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-evaluationdetail.html#cfn-sagemaker-modelcard-evaluationdetail-metricgroups
	//
	MetricGroups interface{} `field:"optional" json:"metricGroups" yaml:"metricGroups"`
}

