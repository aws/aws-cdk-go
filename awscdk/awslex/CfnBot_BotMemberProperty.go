package awslex


// A bot that is a member of a bot network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botMemberProperty := &BotMemberProperty{
//   	BotMemberAliasId: jsii.String("botMemberAliasId"),
//   	BotMemberAliasName: jsii.String("botMemberAliasName"),
//   	BotMemberId: jsii.String("botMemberId"),
//   	BotMemberName: jsii.String("botMemberName"),
//   	BotMemberVersion: jsii.String("botMemberVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botmember.html
//
type CfnBot_BotMemberProperty struct {
	// The alias ID of a bot that is a member of this network of bots.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botmember.html#cfn-lex-bot-botmember-botmemberaliasid
	//
	BotMemberAliasId *string `field:"required" json:"botMemberAliasId" yaml:"botMemberAliasId"`
	// The alias name of a bot that is a member of this network of bots.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botmember.html#cfn-lex-bot-botmember-botmemberaliasname
	//
	BotMemberAliasName *string `field:"required" json:"botMemberAliasName" yaml:"botMemberAliasName"`
	// The unique ID of a bot that is a member of this network of bots.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botmember.html#cfn-lex-bot-botmember-botmemberid
	//
	BotMemberId *string `field:"required" json:"botMemberId" yaml:"botMemberId"`
	// The unique name of a bot that is a member of this network of bots.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botmember.html#cfn-lex-bot-botmember-botmembername
	//
	BotMemberName *string `field:"required" json:"botMemberName" yaml:"botMemberName"`
	// The version of a bot that is a member of this network of bots.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botmember.html#cfn-lex-bot-botmember-botmemberversion
	//
	BotMemberVersion *string `field:"required" json:"botMemberVersion" yaml:"botMemberVersion"`
}

