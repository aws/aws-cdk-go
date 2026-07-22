package awsquicksight


// Properties for CfnFlowPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var flowDefinition interface{}
//
//   cfnFlowMixinProps := &CfnFlowMixinProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	Description: jsii.String("description"),
//   	FlowDefinition: flowDefinition,
//   	Name: jsii.String("name"),
//   	Permissions: []interface{}{
//   		&PermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-flow.html
//
type CfnFlowMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-flow.html#cfn-quicksight-flow-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-flow.html#cfn-quicksight-flow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-flow.html#cfn-quicksight-flow-flowdefinition
	//
	FlowDefinition interface{} `field:"optional" json:"flowDefinition" yaml:"flowDefinition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-flow.html#cfn-quicksight-flow-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-flow.html#cfn-quicksight-flow-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
}

