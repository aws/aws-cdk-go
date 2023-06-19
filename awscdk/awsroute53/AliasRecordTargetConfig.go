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
// Experimental.
type AliasRecordTargetConfig struct {
	// DNS name of the target.
	// Experimental.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// Hosted zone ID of the target.
	// Experimental.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
}

