package awssesactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Construction properties for a bounce action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bounceTemplate bounceTemplate
//   var topic topic
//
//   bounceProps := &bounceProps{
//   	sender: jsii.String("sender"),
//   	template: bounceTemplate,
//
//   	// the properties below are optional
//   	topic: topic,
//   }
//
// Experimental.
type BounceProps struct {
	// The email address of the sender of the bounced email.
	//
	// This is the address
	// from which the bounce message will be sent.
	// Experimental.
	Sender *string `field:"required" json:"sender" yaml:"sender"`
	// The template containing the message, reply code and status code.
	// Experimental.
	Template BounceTemplate `field:"required" json:"template" yaml:"template"`
	// The SNS topic to notify when the bounce action is taken.
	// Experimental.
	Topic awssns.ITopic `field:"optional" json:"topic" yaml:"topic"`
}

