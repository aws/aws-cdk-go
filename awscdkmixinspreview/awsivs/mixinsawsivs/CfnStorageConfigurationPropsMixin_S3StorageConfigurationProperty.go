package mixinsawsivs


// The S3StorageConfiguration property type describes an S3 location where recorded videos will be stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3StorageConfigurationProperty := &S3StorageConfigurationProperty{
//   	BucketName: jsii.String("bucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-storageconfiguration-s3storageconfiguration.html
//
type CfnStorageConfigurationPropsMixin_S3StorageConfigurationProperty struct {
	// Name of the S3 bucket where recorded video will be stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-storageconfiguration-s3storageconfiguration.html#cfn-ivs-storageconfiguration-s3storageconfiguration-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
}

