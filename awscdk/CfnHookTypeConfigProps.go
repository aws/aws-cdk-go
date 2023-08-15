package awscdk


// Properties for defining a `CfnHookTypeConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHookTypeConfigProps := &CfnHookTypeConfigProps{
//   	Configuration: jsii.String("configuration"),
//
//   	// the properties below are optional
//   	ConfigurationAlias: jsii.String("configurationAlias"),
//   	TypeArn: jsii.String("typeArn"),
//   	TypeName: jsii.String("typeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hooktypeconfig.html
//
type CfnHookTypeConfigProps struct {
	// Specifies the activated hook type configuration, in this AWS account and AWS Region .
	//
	// You must specify either `TypeName` and `Configuration` or `TypeARN` and `Configuration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hooktypeconfig.html#cfn-cloudformation-hooktypeconfig-configuration
	//
	Configuration *string `field:"required" json:"configuration" yaml:"configuration"`
	// Specifies the activated hook type configuration, in this AWS account and AWS Region .
	//
	// Defaults to `default` alias. Hook types currently support default configuration alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hooktypeconfig.html#cfn-cloudformation-hooktypeconfig-configurationalias
	//
	// Default: - "default".
	//
	ConfigurationAlias *string `field:"optional" json:"configurationAlias" yaml:"configurationAlias"`
	// The Amazon Resource Number (ARN) for the hook to set `Configuration` for.
	//
	// You must specify either `TypeName` and `Configuration` or `TypeARN` and `Configuration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hooktypeconfig.html#cfn-cloudformation-hooktypeconfig-typearn
	//
	TypeArn *string `field:"optional" json:"typeArn" yaml:"typeArn"`
	// The unique name for your hook.
	//
	// Specifies a three-part namespace for your hook, with a recommended pattern of `Organization::Service::Hook` .
	//
	// You must specify either `TypeName` and `Configuration` or `TypeARN` and `Configuration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hooktypeconfig.html#cfn-cloudformation-hooktypeconfig-typename
	//
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
}

