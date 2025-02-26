package awselasticloadbalancingv2


// The mutual authentication configuration information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mutualAuthenticationProperty := &MutualAuthenticationProperty{
//   	AdvertiseTrustStoreCaNames: jsii.String("advertiseTrustStoreCaNames"),
//   	IgnoreClientCertificateExpiry: jsii.Boolean(false),
//   	Mode: jsii.String("mode"),
//   	TrustStoreArn: jsii.String("trustStoreArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-mutualauthentication.html
//
type CfnListener_MutualAuthenticationProperty struct {
	// Indicates whether trust store CA certificate names are advertised.
	//
	// The default value is `off` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-mutualauthentication.html#cfn-elasticloadbalancingv2-listener-mutualauthentication-advertisetruststorecanames
	//
	AdvertiseTrustStoreCaNames *string `field:"optional" json:"advertiseTrustStoreCaNames" yaml:"advertiseTrustStoreCaNames"`
	// Indicates whether expired client certificates are ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-mutualauthentication.html#cfn-elasticloadbalancingv2-listener-mutualauthentication-ignoreclientcertificateexpiry
	//
	IgnoreClientCertificateExpiry interface{} `field:"optional" json:"ignoreClientCertificateExpiry" yaml:"ignoreClientCertificateExpiry"`
	// The client certificate handling method.
	//
	// The possible values are `off` , `passthrough` , and `verify` . The default value is `off` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-mutualauthentication.html#cfn-elasticloadbalancingv2-listener-mutualauthentication-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The Amazon Resource Name (ARN) of the trust store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-mutualauthentication.html#cfn-elasticloadbalancingv2-listener-mutualauthentication-truststorearn
	//
	TrustStoreArn *string `field:"optional" json:"trustStoreArn" yaml:"trustStoreArn"`
}

