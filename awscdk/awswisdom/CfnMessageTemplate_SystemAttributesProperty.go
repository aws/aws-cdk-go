package awswisdom


// The system attributes that are used with the message template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   systemAttributesProperty := &SystemAttributesProperty{
//   	CustomerEndpoint: &SystemEndpointAttributesProperty{
//   		Address: jsii.String("address"),
//   	},
//   	Name: jsii.String("name"),
//   	SystemEndpoint: &SystemEndpointAttributesProperty{
//   		Address: jsii.String("address"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-systemattributes.html
//
type CfnMessageTemplate_SystemAttributesProperty struct {
	// The CustomerEndpoint attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-systemattributes.html#cfn-wisdom-messagetemplate-systemattributes-customerendpoint
	//
	CustomerEndpoint interface{} `field:"optional" json:"customerEndpoint" yaml:"customerEndpoint"`
	// The name of the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-systemattributes.html#cfn-wisdom-messagetemplate-systemattributes-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The SystemEndpoint attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-systemattributes.html#cfn-wisdom-messagetemplate-systemattributes-systemendpoint
	//
	SystemEndpoint interface{} `field:"optional" json:"systemEndpoint" yaml:"systemEndpoint"`
}

