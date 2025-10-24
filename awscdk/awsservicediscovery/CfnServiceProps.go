package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var serviceAttributes interface{}
//
//   cfnServiceProps := &CfnServiceProps{
//   	Description: jsii.String("description"),
//   	DnsConfig: &DnsConfigProperty{
//   		DnsRecords: []interface{}{
//   			&DnsRecordProperty{
//   				Ttl: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		NamespaceId: jsii.String("namespaceId"),
//   		RoutingPolicy: jsii.String("routingPolicy"),
//   	},
//   	HealthCheckConfig: &HealthCheckConfigProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		FailureThreshold: jsii.Number(123),
//   		ResourcePath: jsii.String("resourcePath"),
//   	},
//   	HealthCheckCustomConfig: &HealthCheckCustomConfigProperty{
//   		FailureThreshold: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//   	NamespaceId: jsii.String("namespaceId"),
//   	ServiceAttributes: serviceAttributes,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html
//
type CfnServiceProps struct {
	// The description of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html#cfn-servicediscovery-service-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A complex type that contains information about the Route 53 DNS records that you want AWS Cloud Map to create when you register an instance.
	//
	// > The record types of a service can only be changed by deleting the service and recreating it with a new `Dnsconfig` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html#cfn-servicediscovery-service-dnsconfig
	//
	DnsConfig interface{} `field:"optional" json:"dnsConfig" yaml:"dnsConfig"`
	// *Public DNS and HTTP namespaces only.* A complex type that contains settings for an optional health check. If you specify settings for a health check, AWS Cloud Map associates the health check with the records that you specify in `DnsConfig` .
	//
	// For information about the charges for health checks, see [Amazon Route 53 Pricing](https://docs.aws.amazon.com/route53/pricing/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html#cfn-servicediscovery-service-healthcheckconfig
	//
	HealthCheckConfig interface{} `field:"optional" json:"healthCheckConfig" yaml:"healthCheckConfig"`
	// A complex type that contains information about an optional custom health check.
	//
	// > If you specify a health check configuration, you can specify either `HealthCheckCustomConfig` or `HealthCheckConfig` but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html#cfn-servicediscovery-service-healthcheckcustomconfig
	//
	HealthCheckCustomConfig interface{} `field:"optional" json:"healthCheckCustomConfig" yaml:"healthCheckCustomConfig"`
	// The name of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html#cfn-servicediscovery-service-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID or Amazon Resource Name (ARN) of the namespace that you want to use to create the service.
	//
	// For namespaces shared with your AWS account, specify the namespace ARN. For more information about shared namespaces, see [Cross-account AWS Cloud Map namespace sharing](https://docs.aws.amazon.com/cloud-map/latest/dg/sharing-namespaces.html) in the *AWS Cloud Map Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html#cfn-servicediscovery-service-namespaceid
	//
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// A complex type that contains information about attributes associated with a specific service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html#cfn-servicediscovery-service-serviceattributes
	//
	ServiceAttributes interface{} `field:"optional" json:"serviceAttributes" yaml:"serviceAttributes"`
	// The tags for the service.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html#cfn-servicediscovery-service-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// If present, specifies that the service instances are only discoverable using the `DiscoverInstances` API operation.
	//
	// No DNS records is registered for the service instances. The only valid value is `HTTP` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html#cfn-servicediscovery-service-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

