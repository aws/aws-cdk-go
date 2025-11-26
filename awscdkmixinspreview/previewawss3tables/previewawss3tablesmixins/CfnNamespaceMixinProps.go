package previewawss3tablesmixins


// Properties for CfnNamespacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNamespaceMixinProps := &CfnNamespaceMixinProps{
//   	Namespace: jsii.String("namespace"),
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-namespace.html
//
type CfnNamespaceMixinProps struct {
	// The name of the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-namespace.html#cfn-s3tables-namespace-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The Amazon Resource Name (ARN) of the table bucket to create the namespace in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-namespace.html#cfn-s3tables-namespace-tablebucketarn
	//
	TableBucketArn *string `field:"optional" json:"tableBucketArn" yaml:"tableBucketArn"`
}

