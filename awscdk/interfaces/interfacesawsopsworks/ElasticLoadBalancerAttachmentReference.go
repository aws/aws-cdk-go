package interfacesawsopsworks


// A reference to a ElasticLoadBalancerAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticLoadBalancerAttachmentReference := &ElasticLoadBalancerAttachmentReference{
//   	ElasticLoadBalancerAttachmentId: jsii.String("elasticLoadBalancerAttachmentId"),
//   }
//
type ElasticLoadBalancerAttachmentReference struct {
	// The Id of the ElasticLoadBalancerAttachment resource.
	ElasticLoadBalancerAttachmentId *string `field:"required" json:"elasticLoadBalancerAttachmentId" yaml:"elasticLoadBalancerAttachmentId"`
}

