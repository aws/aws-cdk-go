package awselasticloadbalancingv2


// Information about a revocation file in use by a trust store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trustStoreRevocationProperty := &TrustStoreRevocationProperty{
//   	NumberOfRevokedEntries: jsii.Number(123),
//   	RevocationId: jsii.String("revocationId"),
//   	RevocationType: jsii.String("revocationType"),
//   	TrustStoreArn: jsii.String("trustStoreArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-truststorerevocation-truststorerevocation.html
//
type CfnTrustStoreRevocation_TrustStoreRevocationProperty struct {
	// The number of revoked certificates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-truststorerevocation-truststorerevocation.html#cfn-elasticloadbalancingv2-truststorerevocation-truststorerevocation-numberofrevokedentries
	//
	NumberOfRevokedEntries *float64 `field:"optional" json:"numberOfRevokedEntries" yaml:"numberOfRevokedEntries"`
	// The revocation ID of the revocation file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-truststorerevocation-truststorerevocation.html#cfn-elasticloadbalancingv2-truststorerevocation-truststorerevocation-revocationid
	//
	RevocationId *string `field:"optional" json:"revocationId" yaml:"revocationId"`
	// The type of revocation file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-truststorerevocation-truststorerevocation.html#cfn-elasticloadbalancingv2-truststorerevocation-truststorerevocation-revocationtype
	//
	RevocationType *string `field:"optional" json:"revocationType" yaml:"revocationType"`
	// The Amazon Resource Name (ARN) of the trust store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-truststorerevocation-truststorerevocation.html#cfn-elasticloadbalancingv2-truststorerevocation-truststorerevocation-truststorearn
	//
	TrustStoreArn *string `field:"optional" json:"trustStoreArn" yaml:"trustStoreArn"`
}

