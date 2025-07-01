package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMatchmakingRuleSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMatchmakingRuleSetProps := &CfnMatchmakingRuleSetProps{
//   	Name: jsii.String("name"),
//   	RuleSetBody: jsii.String("ruleSetBody"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingruleset.html
//
type CfnMatchmakingRuleSetProps struct {
	// A unique identifier for the matchmaking rule set.
	//
	// A matchmaking configuration identifies the rule set it uses by this name value. Note that the rule set name is different from the optional `name` field in the rule set body.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingruleset.html#cfn-gamelift-matchmakingruleset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A collection of matchmaking rules, formatted as a JSON string.
	//
	// Comments are not allowed in JSON, but most elements support a description field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingruleset.html#cfn-gamelift-matchmakingruleset-rulesetbody
	//
	RuleSetBody *string `field:"required" json:"ruleSetBody" yaml:"ruleSetBody"`
	// A list of labels to assign to the new matchmaking rule set resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources are useful for resource management, access management and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* . Once the resource is created, you can use TagResource, UntagResource, and ListTagsForResource to add, remove, and view tags. The maximum tag limit may be lower than stated. See the AWS General Reference for actual tagging limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingruleset.html#cfn-gamelift-matchmakingruleset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

