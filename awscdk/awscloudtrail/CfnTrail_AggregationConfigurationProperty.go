package awscloudtrail


// Configure to add aggregation rules to aggregate CloudTrail Events.
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
	// The category of events to be aggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-aggregationconfiguration.html#cfn-cloudtrail-trail-aggregationconfiguration-eventcategory
	//
	EventCategory *string `field:"required" json:"eventCategory" yaml:"eventCategory"`
	// Contains all templates in an aggregation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-aggregationconfiguration.html#cfn-cloudtrail-trail-aggregationconfiguration-templates
	//
	Templates *[]*string `field:"required" json:"templates" yaml:"templates"`
}

