package previewawscloudtrailmixins


// Configure to add aggregation rules to aggregate CloudTrail Events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnTrailPropsMixin_AggregationConfigurationProperty struct {
	// The category of events to be aggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-aggregationconfiguration.html#cfn-cloudtrail-trail-aggregationconfiguration-eventcategory
	//
	EventCategory *string `field:"optional" json:"eventCategory" yaml:"eventCategory"`
	// Contains all templates in an aggregation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-aggregationconfiguration.html#cfn-cloudtrail-trail-aggregationconfiguration-templates
	//
	Templates *[]*string `field:"optional" json:"templates" yaml:"templates"`
}

