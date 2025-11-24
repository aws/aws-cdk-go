package mixinsawssagemaker


// The evaluation details of the model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var value interface{}
//
//   evaluationDetailProperty := &EvaluationDetailProperty{
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
//   					Notes: jsii.String("notes"),
//   					Type: jsii.String("type"),
//   					Value: value,
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
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-evaluationdetail.html
//
type CfnModelCardPropsMixin_EvaluationDetailProperty struct {
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
	// The evaluation job name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-evaluationdetail.html#cfn-sagemaker-modelcard-evaluationdetail-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

