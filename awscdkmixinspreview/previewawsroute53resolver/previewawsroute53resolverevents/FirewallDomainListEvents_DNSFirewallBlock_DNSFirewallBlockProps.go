package previewawsroute53resolverevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for FirewallDomainList aws.route53resolver@DNSFirewallBlock event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dNSFirewallBlockProps := &DNSFirewallBlockProps{
//   	AccountId: []*string{
//   		jsii.String("accountId"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	FirewallDomainListId: []*string{
//   		jsii.String("firewallDomainListId"),
//   	},
//   	FirewallRuleAction: []*string{
//   		jsii.String("firewallRuleAction"),
//   	},
//   	FirewallRuleGroupId: []*string{
//   		jsii.String("firewallRuleGroupId"),
//   	},
//   	LastObservedAt: []*string{
//   		jsii.String("lastObservedAt"),
//   	},
//   	QueryClass: []*string{
//   		jsii.String("queryClass"),
//   	},
//   	QueryName: []*string{
//   		jsii.String("queryName"),
//   	},
//   	QueryType: []*string{
//   		jsii.String("queryType"),
//   	},
//   	Resources: []DnsFirewallBlockItem{
//   		&DnsFirewallBlockItem{
//   			ResolverEndpointDetails: &ResolverEndpointDetails{
//   				Id: []*string{
//   					jsii.String("id"),
//   				},
//   			},
//   			ResolverNetworkInterfaceDetails: &ResolverNetworkInterfaceDetails{
//   				Id: []*string{
//   					jsii.String("id"),
//   				},
//   			},
//   			ResourceType: []*string{
//   				jsii.String("resourceType"),
//   			},
//   		},
//   	},
//   	SrcAddr: []*string{
//   		jsii.String("srcAddr"),
//   	},
//   	SrcPort: []*string{
//   		jsii.String("srcPort"),
//   	},
//   	Transport: []*string{
//   		jsii.String("transport"),
//   	},
//   	VpcId: []*string{
//   		jsii.String("vpcId"),
//   	},
//   }
//
// Experimental.
type FirewallDomainListEvents_DNSFirewallBlock_DNSFirewallBlockProps struct {
	// account-id property.
	//
	// Specify an array of string values to match this event if the actual value of account-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// firewall-domain-list-id property.
	//
	// Specify an array of string values to match this event if the actual value of firewall-domain-list-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the FirewallDomainList reference.
	//
	// Experimental.
	FirewallDomainListId *[]*string `field:"optional" json:"firewallDomainListId" yaml:"firewallDomainListId"`
	// firewall-rule-action property.
	//
	// Specify an array of string values to match this event if the actual value of firewall-rule-action is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FirewallRuleAction *[]*string `field:"optional" json:"firewallRuleAction" yaml:"firewallRuleAction"`
	// firewall-rule-group-id property.
	//
	// Specify an array of string values to match this event if the actual value of firewall-rule-group-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FirewallRuleGroupId *[]*string `field:"optional" json:"firewallRuleGroupId" yaml:"firewallRuleGroupId"`
	// last-observed-at property.
	//
	// Specify an array of string values to match this event if the actual value of last-observed-at is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastObservedAt *[]*string `field:"optional" json:"lastObservedAt" yaml:"lastObservedAt"`
	// query-class property.
	//
	// Specify an array of string values to match this event if the actual value of query-class is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	QueryClass *[]*string `field:"optional" json:"queryClass" yaml:"queryClass"`
	// query-name property.
	//
	// Specify an array of string values to match this event if the actual value of query-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	QueryName *[]*string `field:"optional" json:"queryName" yaml:"queryName"`
	// query-type property.
	//
	// Specify an array of string values to match this event if the actual value of query-type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	QueryType *[]*string `field:"optional" json:"queryType" yaml:"queryType"`
	// resources property.
	//
	// Specify an array of string values to match this event if the actual value of resources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Resources *[]*FirewallDomainListEvents_DNSFirewallBlock_DnsFirewallBlockItem `field:"optional" json:"resources" yaml:"resources"`
	// src-addr property.
	//
	// Specify an array of string values to match this event if the actual value of src-addr is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SrcAddr *[]*string `field:"optional" json:"srcAddr" yaml:"srcAddr"`
	// src-port property.
	//
	// Specify an array of string values to match this event if the actual value of src-port is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SrcPort *[]*string `field:"optional" json:"srcPort" yaml:"srcPort"`
	// transport property.
	//
	// Specify an array of string values to match this event if the actual value of transport is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Transport *[]*string `field:"optional" json:"transport" yaml:"transport"`
	// vpc-id property.
	//
	// Specify an array of string values to match this event if the actual value of vpc-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VpcId *[]*string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

