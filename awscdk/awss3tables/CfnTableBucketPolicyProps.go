package awss3tables


// Properties for defining a `CfnTableBucketPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var resourcePolicy interface{}
//
//   cfnTableBucketPolicyProps := &CfnTableBucketPolicyProps{
//   	ResourcePolicy: resourcePolicy,
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablebucketpolicy.html
//
type CfnTableBucketPolicyProps struct {
	// The bucket policy JSON for the table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablebucketpolicy.html#cfn-s3tables-tablebucketpolicy-resourcepolicy
	//
	ResourcePolicy interface{} `field:"required" json:"resourcePolicy" yaml:"resourcePolicy"`
	// The Amazon Resource Name (ARN) of the table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablebucketpolicy.html#cfn-s3tables-tablebucketpolicy-tablebucketarn
	//
	TableBucketArn *string `field:"required" json:"tableBucketArn" yaml:"tableBucketArn"`
}

