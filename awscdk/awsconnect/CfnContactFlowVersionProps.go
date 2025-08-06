package awsconnect


// Properties for defining a `CfnContactFlowVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContactFlowVersionProps := &CfnContactFlowVersionProps{
//   	ContactFlowId: jsii.String("contactFlowId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowversion.html
//
type CfnContactFlowVersionProps struct {
	// The identifier of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowversion.html#cfn-connect-contactflowversion-contactflowid
	//
	ContactFlowId *string `field:"required" json:"contactFlowId" yaml:"contactFlowId"`
	// The description of the flow version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowversion.html#cfn-connect-contactflowversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

