package mixinsawssagemaker


// Contains bias metrics for a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   biasProperty := &BiasProperty{
//   	PostTrainingReport: &MetricsSourceProperty{
//   		ContentDigest: jsii.String("contentDigest"),
//   		ContentType: jsii.String("contentType"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	PreTrainingReport: &MetricsSourceProperty{
//   		ContentDigest: jsii.String("contentDigest"),
//   		ContentType: jsii.String("contentType"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	Report: &MetricsSourceProperty{
//   		ContentDigest: jsii.String("contentDigest"),
//   		ContentType: jsii.String("contentType"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-bias.html
//
type CfnModelPackagePropsMixin_BiasProperty struct {
	// The post-training bias report for a model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-bias.html#cfn-sagemaker-modelpackage-bias-posttrainingreport
	//
	PostTrainingReport interface{} `field:"optional" json:"postTrainingReport" yaml:"postTrainingReport"`
	// The pre-training bias report for a model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-bias.html#cfn-sagemaker-modelpackage-bias-pretrainingreport
	//
	PreTrainingReport interface{} `field:"optional" json:"preTrainingReport" yaml:"preTrainingReport"`
	// The bias report for a model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-bias.html#cfn-sagemaker-modelpackage-bias-report
	//
	Report interface{} `field:"optional" json:"report" yaml:"report"`
}

