package previewawslookoutequipmentmixins


// Specifies configuration information for the output results from the inference, including output S3 location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3OutputConfigurationProperty := &S3OutputConfigurationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-s3outputconfiguration.html
//
type CfnInferenceSchedulerPropsMixin_S3OutputConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-s3outputconfiguration.html#cfn-lookoutequipment-inferencescheduler-s3outputconfiguration-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-s3outputconfiguration.html#cfn-lookoutequipment-inferencescheduler-s3outputconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

