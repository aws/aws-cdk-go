package previewawsmwaaserverlessmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	ObjectKey: jsii.String("objectKey"),
//   	VersionId: jsii.String("versionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-s3location.html
//
type CfnWorkflowPropsMixin_S3LocationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-s3location.html#cfn-mwaaserverless-workflow-s3location-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-s3location.html#cfn-mwaaserverless-workflow-s3location-objectkey
	//
	ObjectKey *string `field:"optional" json:"objectKey" yaml:"objectKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-s3location.html#cfn-mwaaserverless-workflow-s3location-versionid
	//
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

