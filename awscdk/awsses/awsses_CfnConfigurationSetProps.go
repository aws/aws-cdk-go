package awsses


// Properties for defining a `CfnConfigurationSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationSetProps := &cfnConfigurationSetProps{
//   	deliveryOptions: &deliveryOptionsProperty{
//   		sendingPoolName: jsii.String("sendingPoolName"),
//   		tlsPolicy: jsii.String("tlsPolicy"),
//   	},
//   	name: jsii.String("name"),
//   	reputationOptions: &reputationOptionsProperty{
//   		reputationMetricsEnabled: jsii.Boolean(false),
//   	},
//   	sendingOptions: &sendingOptionsProperty{
//   		sendingEnabled: jsii.Boolean(false),
//   	},
//   	suppressionOptions: &suppressionOptionsProperty{
//   		suppressedReasons: []*string{
//   			jsii.String("suppressedReasons"),
//   		},
//   	},
//   	trackingOptions: &trackingOptionsProperty{
//   		customRedirectDomain: jsii.String("customRedirectDomain"),
//   	},
//   	vdmOptions: &vdmOptionsProperty{
//   		dashboardOptions: &dashboardOptionsProperty{
//   			engagementMetrics: jsii.String("engagementMetrics"),
//   		},
//   		guardianOptions: &guardianOptionsProperty{
//   			optimizedSharedDelivery: jsii.String("optimizedSharedDelivery"),
//   		},
//   	},
//   }
//
type CfnConfigurationSetProps struct {
	// `AWS::SES::ConfigurationSet.DeliveryOptions`.
	DeliveryOptions interface{} `field:"optional" json:"deliveryOptions" yaml:"deliveryOptions"`
	// The name of the configuration set. The name must meet the following requirements:.
	//
	// - Contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-).
	// - Contain 64 characters or fewer.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::SES::ConfigurationSet.ReputationOptions`.
	ReputationOptions interface{} `field:"optional" json:"reputationOptions" yaml:"reputationOptions"`
	// `AWS::SES::ConfigurationSet.SendingOptions`.
	SendingOptions interface{} `field:"optional" json:"sendingOptions" yaml:"sendingOptions"`
	// `AWS::SES::ConfigurationSet.SuppressionOptions`.
	SuppressionOptions interface{} `field:"optional" json:"suppressionOptions" yaml:"suppressionOptions"`
	// `AWS::SES::ConfigurationSet.TrackingOptions`.
	TrackingOptions interface{} `field:"optional" json:"trackingOptions" yaml:"trackingOptions"`
	// `AWS::SES::ConfigurationSet.VdmOptions`.
	VdmOptions interface{} `field:"optional" json:"vdmOptions" yaml:"vdmOptions"`
}

