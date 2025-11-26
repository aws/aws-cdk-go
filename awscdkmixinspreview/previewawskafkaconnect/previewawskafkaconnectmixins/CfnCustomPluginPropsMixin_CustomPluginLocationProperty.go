package previewawskafkaconnectmixins


// Information about the location of a custom plugin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customPluginLocationProperty := &CustomPluginLocationProperty{
//   	S3Location: &S3LocationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		FileKey: jsii.String("fileKey"),
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-custompluginlocation.html
//
type CfnCustomPluginPropsMixin_CustomPluginLocationProperty struct {
	// The S3 bucket Amazon Resource Name (ARN), file key, and object version of the plugin file stored in Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-custompluginlocation.html#cfn-kafkaconnect-customplugin-custompluginlocation-s3location
	//
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}

