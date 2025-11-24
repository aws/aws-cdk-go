package mixinsawssagemaker


// Contains metrics captured from a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   modelMetricsProperty := &ModelMetricsProperty{
//   	Bias: &BiasProperty{
//   		PostTrainingReport: &MetricsSourceProperty{
//   			ContentDigest: jsii.String("contentDigest"),
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   		PreTrainingReport: &MetricsSourceProperty{
//   			ContentDigest: jsii.String("contentDigest"),
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   		Report: &MetricsSourceProperty{
//   			ContentDigest: jsii.String("contentDigest"),
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	Explainability: &ExplainabilityProperty{
//   		Report: &MetricsSourceProperty{
//   			ContentDigest: jsii.String("contentDigest"),
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	ModelDataQuality: &ModelDataQualityProperty{
//   		Constraints: &MetricsSourceProperty{
//   			ContentDigest: jsii.String("contentDigest"),
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   		Statistics: &MetricsSourceProperty{
//   			ContentDigest: jsii.String("contentDigest"),
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	ModelQuality: &ModelQualityProperty{
//   		Constraints: &MetricsSourceProperty{
//   			ContentDigest: jsii.String("contentDigest"),
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   		Statistics: &MetricsSourceProperty{
//   			ContentDigest: jsii.String("contentDigest"),
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelmetrics.html
//
type CfnModelPackagePropsMixin_ModelMetricsProperty struct {
	// Metrics that measure bias in a model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelmetrics.html#cfn-sagemaker-modelpackage-modelmetrics-bias
	//
	Bias interface{} `field:"optional" json:"bias" yaml:"bias"`
	// Metrics that help explain a model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelmetrics.html#cfn-sagemaker-modelpackage-modelmetrics-explainability
	//
	Explainability interface{} `field:"optional" json:"explainability" yaml:"explainability"`
	// Metrics that measure the quality of the input data for a model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelmetrics.html#cfn-sagemaker-modelpackage-modelmetrics-modeldataquality
	//
	ModelDataQuality interface{} `field:"optional" json:"modelDataQuality" yaml:"modelDataQuality"`
	// Metrics that measure the quality of a model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelmetrics.html#cfn-sagemaker-modelpackage-modelmetrics-modelquality
	//
	ModelQuality interface{} `field:"optional" json:"modelQuality" yaml:"modelQuality"`
}

