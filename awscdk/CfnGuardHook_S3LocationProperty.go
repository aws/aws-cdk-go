package awscdk


// S3 Source Location for the Guard files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Uri: jsii.String("uri"),
//
//   	// the properties below are optional
//   	VersionId: jsii.String("versionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-s3location.html
//
type CfnGuardHook_S3LocationProperty struct {
	// S3 uri of Guard files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-s3location.html#cfn-cloudformation-guardhook-s3location-uri
	//
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// S3 object version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-s3location.html#cfn-cloudformation-guardhook-s3location-versionid
	//
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

