package awsses


// Properties for the Virtual Deliverablity Manager (VDM) attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vdmAttributesProps := &vdmAttributesProps{
//   	engagementMetrics: jsii.Boolean(false),
//   	optimizedSharedDelivery: jsii.Boolean(false),
//   }
//
type VdmAttributesProps struct {
	// Whether engagement metrics are enabled for your account.
	EngagementMetrics *bool `field:"optional" json:"engagementMetrics" yaml:"engagementMetrics"`
	// Whether optimized shared delivery is enabled for your account.
	OptimizedSharedDelivery *bool `field:"optional" json:"optimizedSharedDelivery" yaml:"optimizedSharedDelivery"`
}

