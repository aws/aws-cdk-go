package awssagemaker


// Contains metrics captured from a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelMetricsProperty := &ModelMetricsProperty{
//   	Bias: &BiasProperty{
//   		PostTrainingReport: &MetricsSourceProperty{
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   		},
//   		PreTrainingReport: &MetricsSourceProperty{
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   		},
//   		Report: &MetricsSourceProperty{
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   		},
//   	},
//   	Explainability: &ExplainabilityProperty{
//   		Report: &MetricsSourceProperty{
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   		},
//   	},
//   	ModelDataQuality: &ModelDataQualityProperty{
//   		Constraints: &MetricsSourceProperty{
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   		},
//   		Statistics: &MetricsSourceProperty{
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   		},
//   	},
//   	ModelQuality: &ModelQualityProperty{
//   		Constraints: &MetricsSourceProperty{
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   		},
//   		Statistics: &MetricsSourceProperty{
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   		},
//   	},
//   }
//
type CfnModelPackage_ModelMetricsProperty struct {
	// Metrics that measure bais in a model.
	Bias interface{} `field:"optional" json:"bias" yaml:"bias"`
	// Metrics that help explain a model.
	Explainability interface{} `field:"optional" json:"explainability" yaml:"explainability"`
	// Metrics that measure the quality of the input data for a model.
	ModelDataQuality interface{} `field:"optional" json:"modelDataQuality" yaml:"modelDataQuality"`
	// Metrics that measure the quality of a model.
	ModelQuality interface{} `field:"optional" json:"modelQuality" yaml:"modelQuality"`
}

