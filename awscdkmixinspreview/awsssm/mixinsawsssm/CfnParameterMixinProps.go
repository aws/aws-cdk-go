package mixinsawsssm


// Properties for CfnParameterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnParameterMixinProps := &CfnParameterMixinProps{
//   	AllowedPattern: jsii.String("allowedPattern"),
//   	DataType: jsii.String("dataType"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Policies: jsii.String("policies"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Tier: jsii.String("tier"),
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html
//
type CfnParameterMixinProps struct {
	// A regular expression used to validate the parameter value.
	//
	// For example, for `String` types with values restricted to numbers, you can specify the following: `AllowedPattern=^\d+$`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-allowedpattern
	//
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// The data type of the parameter, such as `text` or `aws:ec2:image` .
	//
	// The default is `text` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-datatype
	//
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Information about the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter.
	//
	// > The reported maximum length of 2048 characters for a parameter name includes 1037 characters that are reserved for internal use by Systems Manager . The maximum length for a parameter name that you specify is 1011 characters.
	// >
	// > This count of 1011 characters includes the characters in the ARN that precede the name you specify. This ARN length will vary depending on your partition and Region. For example, the following 45 characters count toward the 1011 character maximum for a parameter created in the US East (Ohio) Region: `arn:aws:ssm:us-east-2:111122223333:parameter/` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about the policies assigned to a parameter.
	//
	// [Assigning parameter policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-policies.html) in the *AWS Systems Manager User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-policies
	//
	Policies *string `field:"optional" json:"policies" yaml:"policies"`
	// Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs).
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a Systems Manager parameter to identify the type of resource to which it applies, the environment, or the type of configuration data referenced by the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The parameter tier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-tier
	//
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// The type of parameter.
	//
	// > Parameters of type `SecureString` are not supported by AWS CloudFormation .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The parameter value.
	//
	// > If type is `StringList` , the system returns a comma-separated string with no spaces between commas in the `Value` field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

