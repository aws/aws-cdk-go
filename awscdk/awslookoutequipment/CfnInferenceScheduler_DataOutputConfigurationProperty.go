package awslookoutequipment


// Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataOutputConfigurationProperty := &DataOutputConfigurationProperty{
//   	S3OutputConfiguration: &S3OutputConfigurationProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		Prefix: jsii.String("prefix"),
//   	},
//
//   	// the properties below are optional
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-dataoutputconfiguration.html
//
type CfnInferenceScheduler_DataOutputConfigurationProperty struct {
	// Specifies configuration information for the output results from the inference, including output S3 location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-dataoutputconfiguration.html#cfn-lookoutequipment-inferencescheduler-dataoutputconfiguration-s3outputconfiguration
	//
	S3OutputConfiguration interface{} `field:"required" json:"s3OutputConfiguration" yaml:"s3OutputConfiguration"`
	// The ID number for the AWS KMS key used to encrypt the inference output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-dataoutputconfiguration.html#cfn-lookoutequipment-inferencescheduler-dataoutputconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

