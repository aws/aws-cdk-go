package awsopsworks


// Properties for defining a `CfnElasticLoadBalancerAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnElasticLoadBalancerAttachmentProps := &CfnElasticLoadBalancerAttachmentProps{
//   	ElasticLoadBalancerName: jsii.String("elasticLoadBalancerName"),
//   	LayerId: jsii.String("layerId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elasticloadbalancerattachment.html
//
type CfnElasticLoadBalancerAttachmentProps struct {
	// The Elastic Load Balancing instance name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elasticloadbalancerattachment.html#cfn-opsworks-elasticloadbalancerattachment-elasticloadbalancername
	//
	ElasticLoadBalancerName *string `field:"required" json:"elasticLoadBalancerName" yaml:"elasticLoadBalancerName"`
	// The OpsWorks layer ID to which the Elastic Load Balancing load balancer is attached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elasticloadbalancerattachment.html#cfn-opsworks-elasticloadbalancerattachment-layerid
	//
	LayerId *string `field:"required" json:"layerId" yaml:"layerId"`
}

