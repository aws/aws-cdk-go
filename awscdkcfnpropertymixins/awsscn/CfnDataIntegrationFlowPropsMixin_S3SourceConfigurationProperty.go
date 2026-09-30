package awsscn


// The S3 source configuration parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3SourceConfigurationProperty := &S3SourceConfigurationProperty{
//   	BucketName: jsii.String("bucketName"),
//   	Options: &S3OptionsProperty{
//   		FileType: jsii.String("fileType"),
//   	},
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-s3sourceconfiguration.html
//
type CfnDataIntegrationFlowPropsMixin_S3SourceConfigurationProperty struct {
	// The S3 bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-s3sourceconfiguration.html#cfn-scn-dataintegrationflow-s3sourceconfiguration-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The Amazon S3 options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-s3sourceconfiguration.html#cfn-scn-dataintegrationflow-s3sourceconfiguration-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The S3 prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-s3sourceconfiguration.html#cfn-scn-dataintegrationflow-s3sourceconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

