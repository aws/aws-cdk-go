package awscdkioteventsalpha


// Specifies the actions to be performed when the condition evaluates to `true`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iotevents_alpha "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//
//   var action IAction
//   var expression Expression
//
//   event := &Event{
//   	EventName: jsii.String("eventName"),
//
//   	// the properties below are optional
//   	Actions: []IAction{
//   		action,
//   	},
//   	Condition: expression,
//   }
//
// Experimental.
type Event struct {
	// The name of the event.
	// Experimental.
	EventName *string `field:"required" json:"eventName" yaml:"eventName"`
	// The actions to be performed.
	// Default: - no actions will be performed.
	//
	// Experimental.
	Actions *[]IAction `field:"optional" json:"actions" yaml:"actions"`
	// The Boolean expression that, when `true`, causes the actions to be performed.
	// Default: - none (the actions are always executed).
	//
	// Experimental.
	Condition Expression `field:"optional" json:"condition" yaml:"condition"`
}

