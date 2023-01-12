package awssesactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
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
type BounceProps struct {
	// The email address of the sender of the bounced email.
	//
	// This is the address
	// from which the bounce message will be sent.
	Sender *string `field:"required" json:"sender" yaml:"sender"`
	// The template containing the message, reply code and status code.
	Template BounceTemplate `field:"required" json:"template" yaml:"template"`
	// The SNS topic to notify when the bounce action is taken.
	Topic awssns.ITopic `field:"optional" json:"topic" yaml:"topic"`
}

