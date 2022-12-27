package awsce


// Properties for defining a `CfnAnomalySubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnomalySubscriptionProps := &cfnAnomalySubscriptionProps{
//   	frequency: jsii.String("frequency"),
//   	monitorArnList: []*string{
//   		jsii.String("monitorArnList"),
//   	},
//   	subscribers: []interface{}{
//   		&subscriberProperty{
//   			address: jsii.String("address"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			status: jsii.String("status"),
//   		},
//   	},
//   	subscriptionName: jsii.String("subscriptionName"),
//   	threshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	resourceTags: []interface{}{
//   		&resourceTagProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAnomalySubscriptionProps struct {
	// The frequency that anomaly reports are sent over email.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// A list of cost anomaly monitors.
	MonitorArnList *[]*string `field:"required" json:"monitorArnList" yaml:"monitorArnList"`
	// A list of subscribers to notify.
	Subscribers interface{} `field:"required" json:"subscribers" yaml:"subscribers"`
	// The name for the subscription.
	SubscriptionName *string `field:"required" json:"subscriptionName" yaml:"subscriptionName"`
	// The dollar value that triggers a notification if the threshold is exceeded.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// `AWS::CE::AnomalySubscription.ResourceTags`.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

