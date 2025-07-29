package awss3tables


// Properties for defining a `CfnNamespace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNamespaceProps := &CfnNamespaceProps{
//   	Namespace: jsii.String("namespace"),
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-namespace.html
//
type CfnNamespaceProps struct {
	// The name of the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-namespace.html#cfn-s3tables-namespace-namespace
	//
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The Amazon Resource Name (ARN) of the specified table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-namespace.html#cfn-s3tables-namespace-tablebucketarn
	//
	TableBucketArn *string `field:"required" json:"tableBucketArn" yaml:"tableBucketArn"`
}

