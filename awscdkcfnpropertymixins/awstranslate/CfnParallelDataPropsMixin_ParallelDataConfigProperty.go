package awstranslate


// Specifies the format and S3 location of the parallel data input file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   parallelDataConfigProperty := &ParallelDataConfigProperty{
//   	Format: jsii.String("format"),
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-paralleldata-paralleldataconfig.html
//
type CfnParallelDataPropsMixin_ParallelDataConfigProperty struct {
	// The format of the parallel data input file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-paralleldata-paralleldataconfig.html#cfn-translate-paralleldata-paralleldataconfig-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The URI of the Amazon S3 folder that contains the parallel data input file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-paralleldata-paralleldataconfig.html#cfn-translate-paralleldata-paralleldataconfig-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

