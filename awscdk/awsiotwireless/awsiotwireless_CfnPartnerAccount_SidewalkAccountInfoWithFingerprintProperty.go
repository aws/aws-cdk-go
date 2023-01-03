package awsiotwireless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sidewalkAccountInfoWithFingerprintProperty := &sidewalkAccountInfoWithFingerprintProperty{
//   	amazonId: jsii.String("amazonId"),
//   	arn: jsii.String("arn"),
//   	fingerprint: jsii.String("fingerprint"),
//   }
//
type CfnPartnerAccount_SidewalkAccountInfoWithFingerprintProperty struct {
	// `CfnPartnerAccount.SidewalkAccountInfoWithFingerprintProperty.AmazonId`.
	AmazonId *string `field:"optional" json:"amazonId" yaml:"amazonId"`
	// `CfnPartnerAccount.SidewalkAccountInfoWithFingerprintProperty.Arn`.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// `CfnPartnerAccount.SidewalkAccountInfoWithFingerprintProperty.Fingerprint`.
	Fingerprint *string `field:"optional" json:"fingerprint" yaml:"fingerprint"`
}

