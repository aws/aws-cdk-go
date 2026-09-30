package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   staticFileS3SourceOptionsProperty := &StaticFileS3SourceOptionsProperty{
//   	BucketName: jsii.String("bucketName"),
//   	ObjectKey: jsii.String("objectKey"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-staticfiles3sourceoptions.html
//
type CfnTemplatePropsMixin_StaticFileS3SourceOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-staticfiles3sourceoptions.html#cfn-quicksight-template-staticfiles3sourceoptions-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-staticfiles3sourceoptions.html#cfn-quicksight-template-staticfiles3sourceoptions-objectkey
	//
	ObjectKey *string `field:"optional" json:"objectKey" yaml:"objectKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-staticfiles3sourceoptions.html#cfn-quicksight-template-staticfiles3sourceoptions-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

