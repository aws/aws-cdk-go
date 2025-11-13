package interfacesawsiotevents


// A reference to a AlarmModel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmModelReference := &AlarmModelReference{
//   	AlarmModelName: jsii.String("alarmModelName"),
//   }
//
type AlarmModelReference struct {
	// The AlarmModelName of the AlarmModel resource.
	AlarmModelName *string `field:"required" json:"alarmModelName" yaml:"alarmModelName"`
}

