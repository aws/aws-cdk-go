package cloudassemblyschema


// Query input for looking up a load balancer listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerListenerContextQuery := &loadBalancerListenerContextQuery{
//   	account: jsii.String("account"),
//   	loadBalancerType: awscdk.Cloud_assembly_schema.loadBalancerType_NETWORK,
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	listenerArn: jsii.String("listenerArn"),
//   	listenerPort: jsii.Number(123),
//   	listenerProtocol: awscdk.*Cloud_assembly_schema.loadBalancerListenerProtocol_HTTP,
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerTags: []tag{
//   		&tag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
type LoadBalancerListenerContextQuery struct {
	// Filter load balancers by their type.
	LoadBalancerType LoadBalancerType `field:"required" json:"loadBalancerType" yaml:"loadBalancerType"`
	// Find by load balancer's ARN.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	LoadBalancerTags *[]*Tag `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Find by listener's arn.
	ListenerArn *string `field:"optional" json:"listenerArn" yaml:"listenerArn"`
	// Filter listeners by listener port.
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// Filter by listener protocol.
	ListenerProtocol LoadBalancerListenerProtocol `field:"optional" json:"listenerProtocol" yaml:"listenerProtocol"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

