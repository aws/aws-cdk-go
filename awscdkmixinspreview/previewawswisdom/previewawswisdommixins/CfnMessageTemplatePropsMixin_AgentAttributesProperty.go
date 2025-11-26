package previewawswisdommixins


// Information about an agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   agentAttributesProperty := &AgentAttributesProperty{
//   	FirstName: jsii.String("firstName"),
//   	LastName: jsii.String("lastName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-agentattributes.html
//
type CfnMessageTemplatePropsMixin_AgentAttributesProperty struct {
	// The agent’s first name as entered in their Amazon Connect user account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-agentattributes.html#cfn-wisdom-messagetemplate-agentattributes-firstname
	//
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// The agent’s last name as entered in their Amazon Connect user account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-agentattributes.html#cfn-wisdom-messagetemplate-agentattributes-lastname
	//
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
}

