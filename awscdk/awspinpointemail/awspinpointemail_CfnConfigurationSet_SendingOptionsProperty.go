package awspinpointemail


// Used to enable or disable email sending for messages that use this configuration set in the current AWS Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sendingOptionsProperty := &sendingOptionsProperty{
//   	sendingEnabled: jsii.Boolean(false),
//   }
//
type CfnConfigurationSet_SendingOptionsProperty struct {
	// If `true` , email sending is enabled for the configuration set.
	//
	// If `false` , email sending is disabled for the configuration set.
	SendingEnabled interface{} `field:"optional" json:"sendingEnabled" yaml:"sendingEnabled"`
}

