package mixinsawstimestream


// The location to write error reports for records rejected, asynchronously, during magnetic store writes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   magneticStoreRejectedDataLocationProperty := &MagneticStoreRejectedDataLocationProperty{
//   	S3Configuration: &S3ConfigurationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		EncryptionOption: jsii.String("encryptionOption"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-magneticstorerejecteddatalocation.html
//
type CfnTablePropsMixin_MagneticStoreRejectedDataLocationProperty struct {
	// Configuration of an S3 location to write error reports for records rejected, asynchronously, during magnetic store writes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-magneticstorerejecteddatalocation.html#cfn-timestream-table-magneticstorerejecteddatalocation-s3configuration
	//
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

