package mixinsawscustomerprofiles


// The S3 location where Identity Resolution Jobs write result files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3ExportingConfigProperty := &S3ExportingConfigProperty{
//   	S3BucketName: jsii.String("s3BucketName"),
//   	S3KeyName: jsii.String("s3KeyName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-s3exportingconfig.html
//
type CfnDomainPropsMixin_S3ExportingConfigProperty struct {
	// The name of the S3 bucket where Identity Resolution Jobs write result files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-s3exportingconfig.html#cfn-customerprofiles-domain-s3exportingconfig-s3bucketname
	//
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// The S3 key name of the location where Identity Resolution Jobs write result files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-s3exportingconfig.html#cfn-customerprofiles-domain-s3exportingconfig-s3keyname
	//
	S3KeyName *string `field:"optional" json:"s3KeyName" yaml:"s3KeyName"`
}

