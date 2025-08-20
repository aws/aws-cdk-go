package awslightsail


// Describes a domain recordset entry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainEntryProperty := &DomainEntryProperty{
//   	Name: jsii.String("name"),
//   	Target: jsii.String("target"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Id: jsii.String("id"),
//   	IsAlias: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-domain-domainentry.html
//
type CfnDomain_DomainEntryProperty struct {
	// The name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-domain-domainentry.html#cfn-lightsail-domain-domainentry-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The target IP address ( `192.0.2.0` ), or AWS name server ( `ns-111.awsdns-22.com.` ).
	//
	// For Lightsail load balancers, the value looks like `ab1234c56789c6b86aba6fb203d443bc-123456789.us-east-2.elb.amazonaws.com` . For Lightsail distributions, the value looks like `exampled1182ne.cloudfront.net` . For Lightsail container services, the value looks like `container-service-1.example23scljs.us-west-2.cs.amazonlightsail.com` . Be sure to also set `isAlias` to `true` when setting up an A record for a Lightsail load balancer, distribution, or container service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-domain-domainentry.html#cfn-lightsail-domain-domainentry-target
	//
	Target *string `field:"required" json:"target" yaml:"target"`
	// The type of domain entry, such as address for IPv4 (A), address for IPv6 (AAAA), canonical name (CNAME), mail exchanger (MX), name server (NS), start of authority (SOA), service locator (SRV), or text (TXT).
	//
	// The following domain entry types can be used:
	//
	// - `A`
	// - `AAAA`
	// - `CNAME`
	// - `MX`
	// - `NS`
	// - `SOA`
	// - `SRV`
	// - `TXT`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-domain-domainentry.html#cfn-lightsail-domain-domainentry-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The ID of the domain recordset entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-domain-domainentry.html#cfn-lightsail-domain-domainentry-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When `true` , specifies whether the domain entry is an alias used by the Lightsail load balancer, Lightsail container service, Lightsail content delivery network (CDN) distribution, or another AWS resource.
	//
	// You can include an alias (A type) record in your request, which points to the DNS name of a load balancer, container service, CDN distribution, or other AWS resource and routes traffic to that resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-domain-domainentry.html#cfn-lightsail-domain-domainentry-isalias
	//
	IsAlias interface{} `field:"optional" json:"isAlias" yaml:"isAlias"`
}

