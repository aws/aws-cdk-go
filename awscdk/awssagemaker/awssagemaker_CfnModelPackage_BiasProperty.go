package awssagemaker


// Contains bias metrics for a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   biasProperty := &biasProperty{
//   	postTrainingReport: &metricsSourceProperty{
//   		contentType: jsii.String("contentType"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		contentDigest: jsii.String("contentDigest"),
//   	},
//   	preTrainingReport: &metricsSourceProperty{
//   		contentType: jsii.String("contentType"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		contentDigest: jsii.String("contentDigest"),
//   	},
//   	report: &metricsSourceProperty{
//   		contentType: jsii.String("contentType"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		contentDigest: jsii.String("contentDigest"),
//   	},
//   }
//
type CfnModelPackage_BiasProperty struct {
	// The post-training bias report for a model.
	PostTrainingReport interface{} `field:"optional" json:"postTrainingReport" yaml:"postTrainingReport"`
	// The pre-training bias report for a model.
	PreTrainingReport interface{} `field:"optional" json:"preTrainingReport" yaml:"preTrainingReport"`
	// The bias report for a model.
	Report interface{} `field:"optional" json:"report" yaml:"report"`
}

