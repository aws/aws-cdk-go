package awssagemaker


// Contains explainability metrics for a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   explainabilityProperty := &explainabilityProperty{
//   	report: &metricsSourceProperty{
//   		contentType: jsii.String("contentType"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		contentDigest: jsii.String("contentDigest"),
//   	},
//   }
//
type CfnModelPackage_ExplainabilityProperty struct {
	// The explainability report for a model.
	Report interface{} `field:"optional" json:"report" yaml:"report"`
}

