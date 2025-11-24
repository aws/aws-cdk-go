package mixinsawscloudformation


// Properties for CfnHookTypeConfigPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHookTypeConfigMixinProps := &CfnHookTypeConfigMixinProps{
//   	Configuration: jsii.String("configuration"),
//   	ConfigurationAlias: jsii.String("configurationAlias"),
//   	TypeArn: jsii.String("typeArn"),
//   	TypeName: jsii.String("typeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hooktypeconfig.html
//
type CfnHookTypeConfigMixinProps struct {
	// Specifies the activated Hook type configuration, in this AWS account and AWS Region .
	//
	// You must specify either `TypeName` and `Configuration` or `TypeArn` and `Configuration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hooktypeconfig.html#cfn-cloudformation-hooktypeconfig-configuration
	//
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// An alias by which to refer to this configuration data.
	//
	// Defaults to `default` alias. Hook types currently support default configuration alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hooktypeconfig.html#cfn-cloudformation-hooktypeconfig-configurationalias
	//
	// Default: - "default".
	//
	ConfigurationAlias *string `field:"optional" json:"configurationAlias" yaml:"configurationAlias"`
	// The Amazon Resource Number (ARN) for the Hook to set `Configuration` for.
	//
	// You must specify either `TypeName` and `Configuration` or `TypeArn` and `Configuration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hooktypeconfig.html#cfn-cloudformation-hooktypeconfig-typearn
	//
	TypeArn *string `field:"optional" json:"typeArn" yaml:"typeArn"`
	// The unique name for your Hook.
	//
	// Specifies a three-part namespace for your Hook, with a recommended pattern of `Organization::Service::Hook` .
	//
	// You must specify either `TypeName` and `Configuration` or `TypeArn` and `Configuration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hooktypeconfig.html#cfn-cloudformation-hooktypeconfig-typename
	//
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
}

