package interfacesawscloudwatch


// A reference to a AlarmMuteRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmMuteRuleReference := &AlarmMuteRuleReference{
//   	AlarmMuteRuleArn: jsii.String("alarmMuteRuleArn"),
//   }
//
type AlarmMuteRuleReference struct {
	// The Arn of the AlarmMuteRule resource.
	AlarmMuteRuleArn *string `field:"required" json:"alarmMuteRuleArn" yaml:"alarmMuteRuleArn"`
}

