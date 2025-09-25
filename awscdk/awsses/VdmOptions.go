package awsses


// Properties for the Virtual Deliverability Manager (VDM) options that apply to the configuration set.
//
// Example:
//   ses.NewConfigurationSet(this, jsii.String("ConfigurationSetWithVdmOptions"), &ConfigurationSetProps{
//   	VdmOptions: &VdmOptions{
//   		EngagementMetrics: jsii.Boolean(true),
//   		OptimizedSharedDelivery: jsii.Boolean(true),
//   	},
//   })
//
type VdmOptions struct {
	// If true, engagement metrics are enabled for the configuration set.
	// Default: - Engagement metrics not configured at the configuration set level. In this case, use account level settings.
	//
	EngagementMetrics *bool `field:"optional" json:"engagementMetrics" yaml:"engagementMetrics"`
	// If true, optimized shared delivery is enabled for the configuration set.
	// Default: - Optimized shared delivery not configured at the configuration set level. In this case, use account level settings.
	//
	OptimizedSharedDelivery *bool `field:"optional" json:"optimizedSharedDelivery" yaml:"optimizedSharedDelivery"`
}

