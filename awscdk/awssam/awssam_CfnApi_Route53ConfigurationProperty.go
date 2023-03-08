package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   route53ConfigurationProperty := &route53ConfigurationProperty{
//   	distributedDomainName: jsii.String("distributedDomainName"),
//   	evaluateTargetHealth: jsii.Boolean(false),
//   	hostedZoneId: jsii.String("hostedZoneId"),
//   	hostedZoneName: jsii.String("hostedZoneName"),
//   	ipV6: jsii.Boolean(false),
//   }
//
type CfnApi_Route53ConfigurationProperty struct {
	// `CfnApi.Route53ConfigurationProperty.DistributedDomainName`.
	DistributedDomainName *string `field:"optional" json:"distributedDomainName" yaml:"distributedDomainName"`
	// `CfnApi.Route53ConfigurationProperty.EvaluateTargetHealth`.
	EvaluateTargetHealth interface{} `field:"optional" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// `CfnApi.Route53ConfigurationProperty.HostedZoneId`.
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// `CfnApi.Route53ConfigurationProperty.HostedZoneName`.
	HostedZoneName *string `field:"optional" json:"hostedZoneName" yaml:"hostedZoneName"`
	// `CfnApi.Route53ConfigurationProperty.IpV6`.
	IpV6 interface{} `field:"optional" json:"ipV6" yaml:"ipV6"`
}

