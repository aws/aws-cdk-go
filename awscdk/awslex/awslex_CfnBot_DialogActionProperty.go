package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dialogActionProperty := &dialogActionProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	slotToElicit: jsii.String("slotToElicit"),
//   	suppressNextMessage: jsii.Boolean(false),
//   }
//
type CfnBot_DialogActionProperty struct {
	// `CfnBot.DialogActionProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnBot.DialogActionProperty.SlotToElicit`.
	SlotToElicit *string `field:"optional" json:"slotToElicit" yaml:"slotToElicit"`
	// `CfnBot.DialogActionProperty.SuppressNextMessage`.
	SuppressNextMessage interface{} `field:"optional" json:"suppressNextMessage" yaml:"suppressNextMessage"`
}

