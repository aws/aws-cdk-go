package interfacesawslex


// A reference to a BotVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botVersionReference := &BotVersionReference{
//   	BotId: jsii.String("botId"),
//   	BotVersion: jsii.String("botVersion"),
//   }
//
type BotVersionReference struct {
	// The BotId of the BotVersion resource.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// The BotVersion of the BotVersion resource.
	BotVersion *string `field:"required" json:"botVersion" yaml:"botVersion"`
}

