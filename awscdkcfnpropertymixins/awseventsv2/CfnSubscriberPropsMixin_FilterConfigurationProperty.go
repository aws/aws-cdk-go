package awseventsv2


// Configuration for filtering which events are delivered to the target.
//
// An event must match every filter to be delivered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   filterConfigurationProperty := &FilterConfigurationProperty{
//   	Filters: []interface{}{
//   		&FilterProperty{
//   			Pattern: jsii.String("pattern"),
//   			Scope: jsii.String("scope"),
//   		},
//   	},
//   	Language: jsii.String("language"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-filterconfiguration.html
//
type CfnSubscriberPropsMixin_FilterConfigurationProperty struct {
	// The list of filters, 1-50 entries.
	//
	// An event must match every filter to be delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-filterconfiguration.html#cfn-eventsv2-subscriber-filterconfiguration-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// The filter language.
	//
	// The default is EVENT_BRIDGE_PATTERN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-filterconfiguration.html#cfn-eventsv2-subscriber-filterconfiguration-language
	//
	Language *string `field:"optional" json:"language" yaml:"language"`
}

