package awsquicksight


// Properties for defining a `CfnFlow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var flowDefinition interface{}
//
//   cfnFlowProps := &CfnFlowProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	FlowDefinition: flowDefinition,
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
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
type CfnFlowProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-flow.html#cfn-quicksight-flow-awsaccountid
	//
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-flow.html#cfn-quicksight-flow-flowdefinition
	//
	FlowDefinition interface{} `field:"required" json:"flowDefinition" yaml:"flowDefinition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-flow.html#cfn-quicksight-flow-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-flow.html#cfn-quicksight-flow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-flow.html#cfn-quicksight-flow-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
}

