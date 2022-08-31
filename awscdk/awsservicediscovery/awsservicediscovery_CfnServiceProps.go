package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceProps := &cfnServiceProps{
//   	description: jsii.String("description"),
//   	dnsConfig: &dnsConfigProperty{
//   		dnsRecords: []interface{}{
//   			&dnsRecordProperty{
//   				ttl: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		namespaceId: jsii.String("namespaceId"),
//   		routingPolicy: jsii.String("routingPolicy"),
//   	},
//   	healthCheckConfig: &healthCheckConfigProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		failureThreshold: jsii.Number(123),
//   		resourcePath: jsii.String("resourcePath"),
//   	},
//   	healthCheckCustomConfig: &healthCheckCustomConfigProperty{
//   		failureThreshold: jsii.Number(123),
//   	},
//   	name: jsii.String("name"),
//   	namespaceId: jsii.String("namespaceId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnServiceProps struct {
	// The description of the service.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A complex type that contains information about the Route 53 DNS records that you want AWS Cloud Map to create when you register an instance.
	DnsConfig interface{} `field:"optional" json:"dnsConfig" yaml:"dnsConfig"`
	// *Public DNS and HTTP namespaces only.* A complex type that contains settings for an optional health check. If you specify settings for a health check, AWS Cloud Map associates the health check with the records that you specify in `DnsConfig` .
	//
	// For information about the charges for health checks, see [Amazon Route 53 Pricing](https://docs.aws.amazon.com/route53/pricing/) .
	HealthCheckConfig interface{} `field:"optional" json:"healthCheckConfig" yaml:"healthCheckConfig"`
	// A complex type that contains information about an optional custom health check.
	//
	// > If you specify a health check configuration, you can specify either `HealthCheckCustomConfig` or `HealthCheckConfig` but not both.
	HealthCheckCustomConfig interface{} `field:"optional" json:"healthCheckCustomConfig" yaml:"healthCheckCustomConfig"`
	// The name of the service.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the namespace that was used to create the service.
	//
	// > You must specify a value for `NamespaceId` either for the service properties or for [DnsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html) . Don't specify a value in both places.
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// The tags for the service.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// If present, specifies that the service instances are only discoverable using the `DiscoverInstances` API operation.
	//
	// No DNS records is registered for the service instances. The only valid value is `HTTP` .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

