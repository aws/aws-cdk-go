package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   legacyS3OriginProperty := &legacyS3OriginProperty{
//   	dnsName: jsii.String("dnsName"),
//
//   	// the properties below are optional
//   	originAccessIdentity: jsii.String("originAccessIdentity"),
//   }
//
type CfnDistribution_LegacyS3OriginProperty struct {
	// `CfnDistribution.LegacyS3OriginProperty.DNSName`.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// `CfnDistribution.LegacyS3OriginProperty.OriginAccessIdentity`.
	OriginAccessIdentity *string `field:"optional" json:"originAccessIdentity" yaml:"originAccessIdentity"`
}

