package cloudassemblyschema


// Query input for looking up a load balancer listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assumeRoleAdditionalOptions interface{}
//
//   loadBalancerListenerContextQuery := &LoadBalancerListenerContextQuery{
//   	Account: jsii.String("account"),
//   	LoadBalancerType: awscdk.Cloud_assembly_schema.LoadBalancerType_NETWORK,
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	AssumeRoleAdditionalOptions: map[string]interface{}{
//   		"assumeRoleAdditionalOptionsKey": assumeRoleAdditionalOptions,
//   	},
//   	ListenerArn: jsii.String("listenerArn"),
//   	ListenerPort: jsii.Number(123),
//   	ListenerProtocol: awscdk.*Cloud_assembly_schema.LoadBalancerListenerProtocol_HTTP,
//   	LoadBalancerArn: jsii.String("loadBalancerArn"),
//   	LoadBalancerTags: []tag{
//   		&tag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	LookupRoleArn: jsii.String("lookupRoleArn"),
//   	LookupRoleExternalId: jsii.String("lookupRoleExternalId"),
//   }
//
type LoadBalancerListenerContextQuery struct {
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Additional options to pass to STS when assuming the lookup role.
	//
	// - `RoleArn` should not be used. Use the dedicated `lookupRoleArn` property instead.
	// - `ExternalId` should not be used. Use the dedicated `lookupRoleExternalId` instead.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/STS.html#assumeRole-property
	//
	// Default: - No additional options.
	//
	AssumeRoleAdditionalOptions *map[string]interface{} `field:"optional" json:"assumeRoleAdditionalOptions" yaml:"assumeRoleAdditionalOptions"`
	// The ARN of the role that should be used to look up the missing values.
	// Default: - None.
	//
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	// Default: - No ExternalId will be supplied.
	//
	LookupRoleExternalId *string `field:"optional" json:"lookupRoleExternalId" yaml:"lookupRoleExternalId"`
	// Filter load balancers by their type.
	LoadBalancerType LoadBalancerType `field:"required" json:"loadBalancerType" yaml:"loadBalancerType"`
	// Find by load balancer's ARN.
	// Default: - does not search by load balancer arn.
	//
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	// Default: - does not match load balancers by tags.
	//
	LoadBalancerTags *[]*Tag `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// Find by listener's arn.
	// Default: - does not find by listener arn.
	//
	ListenerArn *string `field:"optional" json:"listenerArn" yaml:"listenerArn"`
	// Filter listeners by listener port.
	// Default: - does not filter by a listener port.
	//
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// Filter by listener protocol.
	// Default: - does not filter by listener protocol.
	//
	ListenerProtocol LoadBalancerListenerProtocol `field:"optional" json:"listenerProtocol" yaml:"listenerProtocol"`
}

