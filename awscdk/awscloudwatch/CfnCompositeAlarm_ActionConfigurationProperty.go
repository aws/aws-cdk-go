package awscloudwatch


// Configuration that applies to the notifications this alarm sends when it performs its actions.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-compositealarm-actionconfiguration.html
//
type CfnCompositeAlarm_ActionConfigurationProperty struct {
	// Controls whether the alarm's tags are included in its action notifications and state-change events, and which tag keys are included.
	//
	// Specify IncludedTagKeys or ExcludedTagKeys, but not both. Omit both to include every tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-compositealarm-actionconfiguration.html#cfn-cloudwatch-compositealarm-actionconfiguration-alarmtagspropagation
	//
	AlarmTagsPropagation interface{} `field:"optional" json:"alarmTagsPropagation" yaml:"alarmTagsPropagation"`
}

