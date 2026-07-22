package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   hyperParameterTuningJobObjectiveProperty := &HyperParameterTuningJobObjectiveProperty{
//   	MetricName: jsii.String("metricName"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-hyperparametertuningjobobjective.html
//
type CfnAlgorithmPropsMixin_HyperParameterTuningJobObjectiveProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-hyperparametertuningjobobjective.html#cfn-sagemaker-algorithm-hyperparametertuningjobobjective-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-hyperparametertuningjobobjective.html#cfn-sagemaker-algorithm-hyperparametertuningjobobjective-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

