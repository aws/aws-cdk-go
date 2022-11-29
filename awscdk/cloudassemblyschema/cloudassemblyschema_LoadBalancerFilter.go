package cloudassemblyschema


// Filters for selecting load balancers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerFilter := &loadBalancerFilter{
//   	loadBalancerType: awscdk.Cloud_assembly_schema.loadBalancerType_NETWORK,
//
//   	// the properties below are optional
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerTags: []tag{
//   		&tag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
// Experimental.
type LoadBalancerFilter struct {
	// Filter load balancers by their type.
	// Experimental.
	LoadBalancerType LoadBalancerType `field:"required" json:"loadBalancerType" yaml:"loadBalancerType"`
	// Find by load balancer's ARN.
	// Experimental.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	// Experimental.
	LoadBalancerTags *[]*Tag `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

