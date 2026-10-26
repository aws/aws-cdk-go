package awscloudwatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   actionConfigurationProperty := &ActionConfigurationProperty{
//   	AlarmTagsPropagation: &AlarmTagsPropagationProperty{
//   		Enabled: jsii.Boolean(false),
//   		ExcludedTagKeys: []*string{
//   			jsii.String("excludedTagKeys"),
//   		},
//   		IncludedTagKeys: []*string{
//   			jsii.String("includedTagKeys"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-actionconfiguration.html
//
type CfnAlarmPropsMixin_ActionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-actionconfiguration.html#cfn-cloudwatch-alarm-actionconfiguration-alarmtagspropagation
	//
	AlarmTagsPropagation interface{} `field:"optional" json:"alarmTagsPropagation" yaml:"alarmTagsPropagation"`
}

