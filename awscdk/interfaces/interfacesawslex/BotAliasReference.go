package interfacesawslex


// A reference to a BotAlias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botAliasReference := &BotAliasReference{
//   	BotAliasArn: jsii.String("botAliasArn"),
//   	BotAliasId: jsii.String("botAliasId"),
//   	BotId: jsii.String("botId"),
//   }
//
type BotAliasReference struct {
	// The ARN of the BotAlias resource.
	BotAliasArn *string `field:"required" json:"botAliasArn" yaml:"botAliasArn"`
	// The BotAliasId of the BotAlias resource.
	BotAliasId *string `field:"required" json:"botAliasId" yaml:"botAliasId"`
	// The BotId of the BotAlias resource.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
}

