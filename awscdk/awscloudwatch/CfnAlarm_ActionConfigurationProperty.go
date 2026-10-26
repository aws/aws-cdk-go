package awscloudwatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionConfigurationProperty := &ActionConfigurationProperty{
//   	AlarmTagsPropagation: &AlarmTagsPropagationProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
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
type CfnAlarm_ActionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-actionconfiguration.html#cfn-cloudwatch-alarm-actionconfiguration-alarmtagspropagation
	//
	AlarmTagsPropagation interface{} `field:"optional" json:"alarmTagsPropagation" yaml:"alarmTagsPropagation"`
}

