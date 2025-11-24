package mixinsawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRuleGroupsNamespacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuleGroupsNamespaceMixinProps := &CfnRuleGroupsNamespaceMixinProps{
//   	Data: jsii.String("data"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Workspace: jsii.String("workspace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-rulegroupsnamespace.html
//
type CfnRuleGroupsNamespaceMixinProps struct {
	// The rules file used in the namespace.
	//
	// For more details about the rules file, see [Creating a rules file](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-ruler-rulesfile.html) in the *Amazon Managed Service for Prometheus User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-rulegroupsnamespace.html#cfn-aps-rulegroupsnamespace-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The name of the rule groups namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-rulegroupsnamespace.html#cfn-aps-rulegroupsnamespace-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The list of tag keys and values that are associated with the rule groups namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-rulegroupsnamespace.html#cfn-aps-rulegroupsnamespace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the workspace to add the rule groups namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-rulegroupsnamespace.html#cfn-aps-rulegroupsnamespace-workspace
	//
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}

