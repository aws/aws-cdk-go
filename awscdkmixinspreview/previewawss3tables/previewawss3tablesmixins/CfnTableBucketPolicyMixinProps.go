package previewawss3tablesmixins


// Properties for CfnTableBucketPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var resourcePolicy interface{}
//
//   cfnTableBucketPolicyMixinProps := &CfnTableBucketPolicyMixinProps{
//   	ResourcePolicy: resourcePolicy,
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablebucketpolicy.html
//
type CfnTableBucketPolicyMixinProps struct {
	// The bucket policy JSON for the table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablebucketpolicy.html#cfn-s3tables-tablebucketpolicy-resourcepolicy
	//
	ResourcePolicy interface{} `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// The Amazon Resource Name (ARN) of the table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablebucketpolicy.html#cfn-s3tables-tablebucketpolicy-tablebucketarn
	//
	TableBucketArn *string `field:"optional" json:"tableBucketArn" yaml:"tableBucketArn"`
}

