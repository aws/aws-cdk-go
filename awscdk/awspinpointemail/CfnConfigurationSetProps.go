package awspinpointemail


// Properties for defining a `CfnConfigurationSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationSetProps := &CfnConfigurationSetProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DeliveryOptions: &DeliveryOptionsProperty{
//   		SendingPoolName: jsii.String("sendingPoolName"),
//   	},
//   	ReputationOptions: &ReputationOptionsProperty{
//   		ReputationMetricsEnabled: jsii.Boolean(false),
//   	},
//   	SendingOptions: &SendingOptionsProperty{
//   		SendingEnabled: jsii.Boolean(false),
//   	},
//   	Tags: []tagsProperty{
//   		&tagsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrackingOptions: &TrackingOptionsProperty{
//   		CustomRedirectDomain: jsii.String("customRedirectDomain"),
//   	},
//   }
//
type CfnConfigurationSetProps struct {
	// The name of the configuration set.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.
	DeliveryOptions interface{} `field:"optional" json:"deliveryOptions" yaml:"deliveryOptions"`
	// An object that defines whether or not Amazon Pinpoint collects reputation metrics for the emails that you send that use the configuration set.
	ReputationOptions interface{} `field:"optional" json:"reputationOptions" yaml:"reputationOptions"`
	// An object that defines whether or not Amazon Pinpoint can send email that you send using the configuration set.
	SendingOptions interface{} `field:"optional" json:"sendingOptions" yaml:"sendingOptions"`
	// An object that defines the tags (keys and values) that you want to associate with the configuration set.
	Tags *[]*CfnConfigurationSet_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
	// An object that defines the open and click tracking options for emails that you send using the configuration set.
	TrackingOptions interface{} `field:"optional" json:"trackingOptions" yaml:"trackingOptions"`
}

