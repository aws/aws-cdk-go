package awsroute53


// Properties for defining a `CfnTrafficPolicyInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficPolicyInstanceProps := &CfnTrafficPolicyInstanceProps{
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	Name: jsii.String("name"),
//   	TrafficPolicyId: jsii.String("trafficPolicyId"),
//   	TrafficPolicyVersion: jsii.Number(123),
//   	Ttl: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicyinstance.html
//
type CfnTrafficPolicyInstanceProps struct {
	// The ID of the hosted zone that Amazon Route 53 creates the resource record sets in.
	//
	// The bare ID as Route 53 returns it, without a /hostedzone/ prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicyinstance.html#cfn-route53-trafficpolicyinstance-hostedzoneid
	//
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// The domain name, or subdomain name, for which Amazon Route 53 answers DNS queries by using the resource record sets it creates for this traffic policy instance.
	//
	// Must be lower-case and end with a trailing dot, which is the form Route 53 returns: Route 53 normalizes DNS names, so admitting another form would neither round-trip through Read nor guarantee that a change to this property changes the resource's identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicyinstance.html#cfn-route53-trafficpolicyinstance-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the traffic policy that Amazon Route 53 uses to create resource record sets in the specified hosted zone.
	//
	// Lower-case, as Route 53 returns it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicyinstance.html#cfn-route53-trafficpolicyinstance-trafficpolicyid
	//
	TrafficPolicyId *string `field:"required" json:"trafficPolicyId" yaml:"trafficPolicyId"`
	// The version of the traffic policy that Amazon Route 53 uses to create resource record sets in the specified hosted zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicyinstance.html#cfn-route53-trafficpolicyinstance-trafficpolicyversion
	//
	TrafficPolicyVersion *float64 `field:"required" json:"trafficPolicyVersion" yaml:"trafficPolicyVersion"`
	// The TTL that Amazon Route 53 assigns to all of the resource record sets that it creates in the specified hosted zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicyinstance.html#cfn-route53-trafficpolicyinstance-ttl
	//
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}

