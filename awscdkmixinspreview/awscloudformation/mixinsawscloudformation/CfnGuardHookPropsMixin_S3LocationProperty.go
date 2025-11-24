package mixinsawscloudformation


// Specifies the S3 location where your Guard rules or input parameters are located.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Uri: jsii.String("uri"),
//   	VersionId: jsii.String("versionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-s3location.html
//
type CfnGuardHookPropsMixin_S3LocationProperty struct {
	// Specifies the S3 path to the file that contains your Guard rules or input parameters (in the form `s3://<bucket name>/<file name>` ).
	//
	// For Guard rules, the object stored in S3 must have one of the following file extensions: `.guard` , `.zip` , or `.tar.gz` .
	//
	// For input parameters, the object stored in S3 must have one of the following file extensions: `.yaml` , `.json` , `.zip` , or `.tar.gz` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-s3location.html#cfn-cloudformation-guardhook-s3location-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
	// For S3 buckets with versioning enabled, specifies the unique ID of the S3 object version to download your Guard rules or input parameters from.
	//
	// The Guard Hook downloads files from S3 every time the Hook is invoked. To prevent accidental changes or deletions, we recommend using a version when configuring your Guard Hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-s3location.html#cfn-cloudformation-guardhook-s3location-versionid
	//
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

