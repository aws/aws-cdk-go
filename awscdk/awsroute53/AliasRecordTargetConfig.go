package awsroute53


// Represents the properties of an alias target destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aliasRecordTargetConfig := &AliasRecordTargetConfig{
//   	DnsName: jsii.String("dnsName"),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   }
//
type AliasRecordTargetConfig struct {
	// DNS name of the target.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// Hosted zone ID of the target.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
}

