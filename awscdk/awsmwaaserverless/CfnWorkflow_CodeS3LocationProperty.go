package awsmwaaserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeS3LocationProperty := &CodeS3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	ObjectKey: jsii.String("objectKey"),
//
//   	// the properties below are optional
//   	VersionId: jsii.String("versionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-codes3location.html
//
type CfnWorkflow_CodeS3LocationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-codes3location.html#cfn-mwaaserverless-workflow-codes3location-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-codes3location.html#cfn-mwaaserverless-workflow-codes3location-objectkey
	//
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-codes3location.html#cfn-mwaaserverless-workflow-codes3location-versionid
	//
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

