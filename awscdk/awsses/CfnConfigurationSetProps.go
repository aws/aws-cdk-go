package awsses


// Properties for defining a `CfnConfigurationSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationSetProps := &CfnConfigurationSetProps{
//   	DeliveryOptions: &DeliveryOptionsProperty{
//   		SendingPoolName: jsii.String("sendingPoolName"),
//   		TlsPolicy: jsii.String("tlsPolicy"),
//   	},
//   	Name: jsii.String("name"),
//   	ReputationOptions: &ReputationOptionsProperty{
//   		ReputationMetricsEnabled: jsii.Boolean(false),
//   	},
//   	SendingOptions: &SendingOptionsProperty{
//   		SendingEnabled: jsii.Boolean(false),
//   	},
//   	SuppressionOptions: &SuppressionOptionsProperty{
//   		SuppressedReasons: []*string{
//   			jsii.String("suppressedReasons"),
//   		},
//   	},
//   	TrackingOptions: &TrackingOptionsProperty{
//   		CustomRedirectDomain: jsii.String("customRedirectDomain"),
//   	},
//   	VdmOptions: &VdmOptionsProperty{
//   		DashboardOptions: &DashboardOptionsProperty{
//   			EngagementMetrics: jsii.String("engagementMetrics"),
//   		},
//   		GuardianOptions: &GuardianOptionsProperty{
//   			OptimizedSharedDelivery: jsii.String("optimizedSharedDelivery"),
//   		},
//   	},
//   }
//
type CfnConfigurationSetProps struct {
	// Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS).
	DeliveryOptions interface{} `field:"optional" json:"deliveryOptions" yaml:"deliveryOptions"`
	// The name of the configuration set. The name must meet the following requirements:.
	//
	// - Contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-).
	// - Contain 64 characters or fewer.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An object that represents the reputation settings for the configuration set.
	ReputationOptions interface{} `field:"optional" json:"reputationOptions" yaml:"reputationOptions"`
	// An object that defines whether or not Amazon SES can send email that you send using the configuration set.
	SendingOptions interface{} `field:"optional" json:"sendingOptions" yaml:"sendingOptions"`
	// An object that contains information about the suppression list preferences for your account.
	SuppressionOptions interface{} `field:"optional" json:"suppressionOptions" yaml:"suppressionOptions"`
	// The name of the custom open and click tracking domain associated with the configuration set.
	TrackingOptions interface{} `field:"optional" json:"trackingOptions" yaml:"trackingOptions"`
	// The Virtual Deliverability Manager (VDM) options that apply to the configuration set.
	VdmOptions interface{} `field:"optional" json:"vdmOptions" yaml:"vdmOptions"`
}

