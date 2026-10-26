package awscloudwatch


// Controls whether the alarm's tags are included in its action notifications and state-change events, and which tag keys are included.
//
// Specify IncludedTagKeys or ExcludedTagKeys, but not both. Omit both to include every tag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   alarmTagsPropagationProperty := &AlarmTagsPropagationProperty{
//   	Enabled: jsii.Boolean(false),
//   	ExcludedTagKeys: []*string{
//   		jsii.String("excludedTagKeys"),
//   	},
//   	IncludedTagKeys: []*string{
//   		jsii.String("includedTagKeys"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-alarmtagspropagation.html
//
type CfnLogAlarmPropsMixin_AlarmTagsPropagationProperty struct {
	// Whether the alarm's tags are included in its action notifications and state-change events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-alarmtagspropagation.html#cfn-cloudwatch-logalarm-alarmtagspropagation-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The tag keys to leave out.
	//
	// Cannot be specified together with IncludedTagKeys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-alarmtagspropagation.html#cfn-cloudwatch-logalarm-alarmtagspropagation-excludedtagkeys
	//
	ExcludedTagKeys *[]*string `field:"optional" json:"excludedTagKeys" yaml:"excludedTagKeys"`
	// The only tag keys to include.
	//
	// Cannot be specified together with ExcludedTagKeys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-alarmtagspropagation.html#cfn-cloudwatch-logalarm-alarmtagspropagation-includedtagkeys
	//
	IncludedTagKeys *[]*string `field:"optional" json:"includedTagKeys" yaml:"includedTagKeys"`
}

