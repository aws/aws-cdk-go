package awsiotwireless


// Information about a Sidewalk account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sidewalkAccountInfoWithFingerprintProperty := &SidewalkAccountInfoWithFingerprintProperty{
//   	AmazonId: jsii.String("amazonId"),
//   	Arn: jsii.String("arn"),
//   	Fingerprint: jsii.String("fingerprint"),
//   }
//
type CfnPartnerAccount_SidewalkAccountInfoWithFingerprintProperty struct {
	// The Sidewalk Amazon ID.
	AmazonId *string `field:"optional" json:"amazonId" yaml:"amazonId"`
	// The Amazon Resource Name (ARN) of the resource.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The fingerprint of the Sidewalk application server private key.
	Fingerprint *string `field:"optional" json:"fingerprint" yaml:"fingerprint"`
}

