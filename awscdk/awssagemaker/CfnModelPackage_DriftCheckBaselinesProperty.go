package awssagemaker


// Represents the drift check baselines that can be used when the model monitor is set using the model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   driftCheckBaselinesProperty := &DriftCheckBaselinesProperty{
//   	Bias: &DriftCheckBiasProperty{
//   		ConfigFile: &FileSourceProperty{
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   			ContentType: jsii.String("contentType"),
//   		},
//   		PostTrainingConstraints: &MetricsSourceProperty{
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   		},
//   		PreTrainingConstraints: &MetricsSourceProperty{
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   		},
//   	},
//   	Explainability: &DriftCheckExplainabilityProperty{
//   		ConfigFile: &FileSourceProperty{
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   			ContentType: jsii.String("contentType"),
//   		},
//   		Constraints: &MetricsSourceProperty{
//   			ContentType: jsii.String("contentType"),
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ContentDigest: jsii.String("contentDigest"),
//   		},
//   	},
//   	ModelDataQuality: &DriftCheckModelDataQualityProperty{
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
//   	ModelQuality: &DriftCheckModelQualityProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckbaselines.html
//
type CfnModelPackage_DriftCheckBaselinesProperty struct {
	// Represents the drift check bias baselines that can be used when the model monitor is set using the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckbaselines.html#cfn-sagemaker-modelpackage-driftcheckbaselines-bias
	//
	Bias interface{} `field:"optional" json:"bias" yaml:"bias"`
	// Represents the drift check explainability baselines that can be used when the model monitor is set using the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckbaselines.html#cfn-sagemaker-modelpackage-driftcheckbaselines-explainability
	//
	Explainability interface{} `field:"optional" json:"explainability" yaml:"explainability"`
	// Represents the drift check model data quality baselines that can be used when the model monitor is set using the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckbaselines.html#cfn-sagemaker-modelpackage-driftcheckbaselines-modeldataquality
	//
	ModelDataQuality interface{} `field:"optional" json:"modelDataQuality" yaml:"modelDataQuality"`
	// Represents the drift check model quality baselines that can be used when the model monitor is set using the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckbaselines.html#cfn-sagemaker-modelpackage-driftcheckbaselines-modelquality
	//
	ModelQuality interface{} `field:"optional" json:"modelQuality" yaml:"modelQuality"`
}

