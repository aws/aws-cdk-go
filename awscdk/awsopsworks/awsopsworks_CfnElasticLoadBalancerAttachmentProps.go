package awsopsworks


// Properties for defining a `CfnElasticLoadBalancerAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnElasticLoadBalancerAttachmentProps := &cfnElasticLoadBalancerAttachmentProps{
//   	elasticLoadBalancerName: jsii.String("elasticLoadBalancerName"),
//   	layerId: jsii.String("layerId"),
//   }
//
type CfnElasticLoadBalancerAttachmentProps struct {
	// The Elastic Load Balancing instance name.
	ElasticLoadBalancerName *string `field:"required" json:"elasticLoadBalancerName" yaml:"elasticLoadBalancerName"`
	// The AWS OpsWorks layer ID to which the Elastic Load Balancing load balancer is attached.
	LayerId *string `field:"required" json:"layerId" yaml:"layerId"`
}

