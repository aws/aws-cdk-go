package awssagemaker


// Contains explainability metrics for a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   explainabilityProperty := &ExplainabilityProperty{
//   	Report: &MetricsSourceProperty{
//   		ContentType: jsii.String("contentType"),
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		ContentDigest: jsii.String("contentDigest"),
//   	},
//   }
//
type CfnModelPackage_ExplainabilityProperty struct {
	// The explainability report for a model.
	Report interface{} `field:"optional" json:"report" yaml:"report"`
}

