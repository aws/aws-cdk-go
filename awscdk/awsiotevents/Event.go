package awsiotevents


// Specifies the actions to be performed when the condition evaluates to `true`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var action iAction
//   var expression expression
//
//   event := &Event{
//   	EventName: jsii.String("eventName"),
//
//   	// the properties below are optional
//   	Actions: []*iAction{
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
	// Experimental.
	Actions *[]IAction `field:"optional" json:"actions" yaml:"actions"`
	// The Boolean expression that, when `true`, causes the actions to be performed.
	// Experimental.
	Condition Expression `field:"optional" json:"condition" yaml:"condition"`
}

