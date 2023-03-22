package awsssm


// Properties for defining a `CfnParameter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnParameterProps := &cfnParameterProps{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	allowedPattern: jsii.String("allowedPattern"),
//   	dataType: jsii.String("dataType"),
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	policies: jsii.String("policies"),
//   	tags: tags,
//   	tier: jsii.String("tier"),
//   }
//
type CfnParameterProps struct {
	// The type of parameter.
	//
	// > AWS CloudFormation doesn't support creating a `SecureString` parameter type.
	//
	// *Allowed Values* : String | StringList.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The parameter value.
	//
	// > If type is `StringList` , the system returns a comma-separated string with no spaces between commas in the `Value` field.
	Value *string `field:"required" json:"value" yaml:"value"`
	// A regular expression used to validate the parameter value.
	//
	// For example, for String types with values restricted to numbers, you can specify the following: `AllowedPattern=^\d+$`.
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// The data type of the parameter, such as `text` or `aws:ec2:image` .
	//
	// The default is `text` .
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Information about the parameter.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter.
	//
	// > The maximum length constraint listed below includes capacity for additional system attributes that aren't part of the name. The maximum length for a parameter name, including the full length of the parameter ARN, is 1011 characters. For example, the length of the following parameter name is 65 characters, not 20 characters: `arn:aws:ssm:us-east-2:111222333444:parameter/ExampleParameterName`
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about the policies assigned to a parameter.
	//
	// [Assigning parameter policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-policies.html) in the *AWS Systems Manager User Guide* .
	Policies *string `field:"optional" json:"policies" yaml:"policies"`
	// Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs).
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a Systems Manager parameter to identify the type of resource to which it applies, the environment, or the type of configuration data referenced by the parameter.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The parameter tier.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

