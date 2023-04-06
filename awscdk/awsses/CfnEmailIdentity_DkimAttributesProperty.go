package awsses


// Used to enable or disable DKIM authentication for an email identity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dkimAttributesProperty := &DkimAttributesProperty{
//   	SigningEnabled: jsii.Boolean(false),
//   }
//
type CfnEmailIdentity_DkimAttributesProperty struct {
	// Sets the DKIM signing configuration for the identity.
	//
	// When you set this value `true` , then the messages that are sent from the identity are signed using DKIM. If you set this value to `false` , your messages are sent without DKIM signing.
	SigningEnabled interface{} `field:"optional" json:"signingEnabled" yaml:"signingEnabled"`
}

