package awssagemaker


// Represents the drift check bias baselines that can be used when the model monitor is set using the model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   driftCheckBiasProperty := &driftCheckBiasProperty{
//   	configFile: &fileSourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		contentDigest: jsii.String("contentDigest"),
//   		contentType: jsii.String("contentType"),
//   	},
//   	postTrainingConstraints: &metricsSourceProperty{
//   		contentType: jsii.String("contentType"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		contentDigest: jsii.String("contentDigest"),
//   	},
//   	preTrainingConstraints: &metricsSourceProperty{
//   		contentType: jsii.String("contentType"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		contentDigest: jsii.String("contentDigest"),
//   	},
//   }
//
type CfnModelPackage_DriftCheckBiasProperty struct {
	// The bias config file for a model.
	ConfigFile interface{} `field:"optional" json:"configFile" yaml:"configFile"`
	// `CfnModelPackage.DriftCheckBiasProperty.PostTrainingConstraints`.
	PostTrainingConstraints interface{} `field:"optional" json:"postTrainingConstraints" yaml:"postTrainingConstraints"`
	// `CfnModelPackage.DriftCheckBiasProperty.PreTrainingConstraints`.
	PreTrainingConstraints interface{} `field:"optional" json:"preTrainingConstraints" yaml:"preTrainingConstraints"`
}

