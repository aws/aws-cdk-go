package previewawsdmsmixins


// The property describes schema conversion application attributes for the migration project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schemaConversionApplicationAttributesProperty := &SchemaConversionApplicationAttributesProperty{
//   	S3BucketPath: jsii.String("s3BucketPath"),
//   	S3BucketRoleArn: jsii.String("s3BucketRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-migrationproject-schemaconversionapplicationattributes.html
//
type CfnMigrationProjectPropsMixin_SchemaConversionApplicationAttributesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-migrationproject-schemaconversionapplicationattributes.html#cfn-dms-migrationproject-schemaconversionapplicationattributes-s3bucketpath
	//
	S3BucketPath *string `field:"optional" json:"s3BucketPath" yaml:"s3BucketPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-migrationproject-schemaconversionapplicationattributes.html#cfn-dms-migrationproject-schemaconversionapplicationattributes-s3bucketrolearn
	//
	S3BucketRoleArn *string `field:"optional" json:"s3BucketRoleArn" yaml:"s3BucketRoleArn"`
}

