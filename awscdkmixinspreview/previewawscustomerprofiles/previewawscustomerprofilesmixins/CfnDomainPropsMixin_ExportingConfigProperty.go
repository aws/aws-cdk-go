package previewawscustomerprofilesmixins


// Configuration information for exporting Identity Resolution results, for example, to an S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   exportingConfigProperty := &ExportingConfigProperty{
//   	S3Exporting: &S3ExportingConfigProperty{
//   		S3BucketName: jsii.String("s3BucketName"),
//   		S3KeyName: jsii.String("s3KeyName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-exportingconfig.html
//
type CfnDomainPropsMixin_ExportingConfigProperty struct {
	// The S3 location where Identity Resolution Jobs write result files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-exportingconfig.html#cfn-customerprofiles-domain-exportingconfig-s3exporting
	//
	S3Exporting interface{} `field:"optional" json:"s3Exporting" yaml:"s3Exporting"`
}

