package previewawssesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConfigurationSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationSetMixinProps := &CfnConfigurationSetMixinProps{
//   	DeliveryOptions: &DeliveryOptionsProperty{
//   		MaxDeliverySeconds: jsii.Number(123),
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
//   		ValidationOptions: &ValidationOptionsProperty{
//   			ConditionThreshold: &ConditionThresholdProperty{
//   				ConditionThresholdEnabled: jsii.String("conditionThresholdEnabled"),
//   				OverallConfidenceThreshold: &OverallConfidenceThresholdProperty{
//   					ConfidenceVerdictThreshold: jsii.String("confidenceVerdictThreshold"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrackingOptions: &TrackingOptionsProperty{
//   		CustomRedirectDomain: jsii.String("customRedirectDomain"),
//   		HttpsPolicy: jsii.String("httpsPolicy"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html
//
type CfnConfigurationSetMixinProps struct {
	// Specifies the name of the dedicated IP pool to associate with the configuration set and whether messages that use the configuration set are required to use Transport Layer Security (TLS).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html#cfn-ses-configurationset-deliveryoptions
	//
	DeliveryOptions interface{} `field:"optional" json:"deliveryOptions" yaml:"deliveryOptions"`
	// The name of the configuration set. The name must meet the following requirements:.
	//
	// - Contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-).
	// - Contain 64 characters or fewer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html#cfn-ses-configurationset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html#cfn-ses-configurationset-reputationoptions
	//
	ReputationOptions interface{} `field:"optional" json:"reputationOptions" yaml:"reputationOptions"`
	// An object that defines whether or not Amazon SES can send email that you send using the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html#cfn-ses-configurationset-sendingoptions
	//
	SendingOptions interface{} `field:"optional" json:"sendingOptions" yaml:"sendingOptions"`
	// An object that contains information about the suppression list preferences for your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html#cfn-ses-configurationset-suppressionoptions
	//
	SuppressionOptions interface{} `field:"optional" json:"suppressionOptions" yaml:"suppressionOptions"`
	// An array of objects that define the tags (keys and values) that are associated with the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html#cfn-ses-configurationset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An object that defines the open and click tracking options for emails that you send using the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html#cfn-ses-configurationset-trackingoptions
	//
	TrackingOptions interface{} `field:"optional" json:"trackingOptions" yaml:"trackingOptions"`
	// The Virtual Deliverability Manager (VDM) options that apply to the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html#cfn-ses-configurationset-vdmoptions
	//
	VdmOptions interface{} `field:"optional" json:"vdmOptions" yaml:"vdmOptions"`
}

