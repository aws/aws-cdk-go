package awsbedrock


// Contains information about a version that the alias maps to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowAliasRoutingConfigurationListItemProperty := &FlowAliasRoutingConfigurationListItemProperty{
//   	FlowVersion: jsii.String("flowVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowalias-flowaliasroutingconfigurationlistitem.html
//
type CfnFlowAlias_FlowAliasRoutingConfigurationListItemProperty struct {
	// The version that the alias maps to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowalias-flowaliasroutingconfigurationlistitem.html#cfn-bedrock-flowalias-flowaliasroutingconfigurationlistitem-flowversion
	//
	FlowVersion *string `field:"optional" json:"flowVersion" yaml:"flowVersion"`
}

