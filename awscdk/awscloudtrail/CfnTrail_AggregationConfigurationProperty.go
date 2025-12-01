package awscloudtrail


// An object that contains configuration settings for aggregating events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregationConfigurationProperty := &AggregationConfigurationProperty{
//   	EventCategory: jsii.String("eventCategory"),
//   	Templates: []*string{
//   		jsii.String("templates"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-aggregationconfiguration.html
//
type CfnTrail_AggregationConfigurationProperty struct {
	// Specifies the event category for which aggregation should be performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-aggregationconfiguration.html#cfn-cloudtrail-trail-aggregationconfiguration-eventcategory
	//
	EventCategory *string `field:"required" json:"eventCategory" yaml:"eventCategory"`
	// A list of aggregation templates that can be used to configure event aggregation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-aggregationconfiguration.html#cfn-cloudtrail-trail-aggregationconfiguration-templates
	//
	Templates *[]*string `field:"required" json:"templates" yaml:"templates"`
}

