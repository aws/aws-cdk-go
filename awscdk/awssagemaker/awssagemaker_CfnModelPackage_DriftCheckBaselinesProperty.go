package awssagemaker


// Represents the drift check baselines that can be used when the model monitor is set using the model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   driftCheckBaselinesProperty := &driftCheckBaselinesProperty{
//   	bias: &driftCheckBiasProperty{
//   		configFile: &fileSourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
//   			contentType: jsii.String("contentType"),
//   		},
//   		postTrainingConstraints: &metricsSourceProperty{
//   			contentType: jsii.String("contentType"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
//   		},
//   		preTrainingConstraints: &metricsSourceProperty{
//   			contentType: jsii.String("contentType"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
//   		},
//   	},
//   	explainability: &driftCheckExplainabilityProperty{
//   		configFile: &fileSourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
//   			contentType: jsii.String("contentType"),
//   		},
//   		constraints: &metricsSourceProperty{
//   			contentType: jsii.String("contentType"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			contentDigest: jsii.String("contentDigest"),
//   		},
//   	},
//   	modelDataQuality: &driftCheckModelDataQualityProperty{
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
//   	modelQuality: &driftCheckModelQualityProperty{
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
type CfnModelPackage_DriftCheckBaselinesProperty struct {
	// Represents the drift check bias baselines that can be used when the model monitor is set using the model package.
	Bias interface{} `field:"optional" json:"bias" yaml:"bias"`
	// Represents the drift check explainability baselines that can be used when the model monitor is set using the model package.
	Explainability interface{} `field:"optional" json:"explainability" yaml:"explainability"`
	// Represents the drift check model data quality baselines that can be used when the model monitor is set using the model package.
	ModelDataQuality interface{} `field:"optional" json:"modelDataQuality" yaml:"modelDataQuality"`
	// Represents the drift check model quality baselines that can be used when the model monitor is set using the model package.
	ModelQuality interface{} `field:"optional" json:"modelQuality" yaml:"modelQuality"`
}

