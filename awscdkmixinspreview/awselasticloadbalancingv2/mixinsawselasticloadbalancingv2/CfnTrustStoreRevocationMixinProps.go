package mixinsawselasticloadbalancingv2


// Properties for CfnTrustStoreRevocationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTrustStoreRevocationMixinProps := &CfnTrustStoreRevocationMixinProps{
//   	RevocationContents: []interface{}{
//   		&RevocationContentProperty{
//   			RevocationType: jsii.String("revocationType"),
//   			S3Bucket: jsii.String("s3Bucket"),
//   			S3Key: jsii.String("s3Key"),
//   			S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   		},
//   	},
//   	TrustStoreArn: jsii.String("trustStoreArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-truststorerevocation.html
//
type CfnTrustStoreRevocationMixinProps struct {
	// The revocation file to add.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-truststorerevocation.html#cfn-elasticloadbalancingv2-truststorerevocation-revocationcontents
	//
	RevocationContents interface{} `field:"optional" json:"revocationContents" yaml:"revocationContents"`
	// The Amazon Resource Name (ARN) of the trust store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-truststorerevocation.html#cfn-elasticloadbalancingv2-truststorerevocation-truststorearn
	//
	TrustStoreArn *string `field:"optional" json:"trustStoreArn" yaml:"trustStoreArn"`
}

