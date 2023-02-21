package awssagemaker


// Contains metrics captured from a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelMetricsProperty := &modelMetricsProperty{
//   	bias: &biasProperty{
//   		postTrainingReport: &metricsSourceProperty{
//   			contentType: jsii.String("contentType"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
//   		},
//   		preTrainingReport: &metricsSourceProperty{
//   			contentType: jsii.String("contentType"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
//   		},
//   		report: &metricsSourceProperty{
//   			contentType: jsii.String("contentType"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
//   		},
//   	},
//   	explainability: &explainabilityProperty{
//   		report: &metricsSourceProperty{
//   			contentType: jsii.String("contentType"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
//   		},
//   	},
//   	modelDataQuality: &modelDataQualityProperty{
//   		constraints: &metricsSourceProperty{
//   			contentType: jsii.String("contentType"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
//   		},
//   		statistics: &metricsSourceProperty{
//   			contentType: jsii.String("contentType"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
//   		},
//   	},
//   	modelQuality: &modelQualityProperty{
//   		constraints: &metricsSourceProperty{
//   			contentType: jsii.String("contentType"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
//   		},
//   		statistics: &metricsSourceProperty{
//   			contentType: jsii.String("contentType"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
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

