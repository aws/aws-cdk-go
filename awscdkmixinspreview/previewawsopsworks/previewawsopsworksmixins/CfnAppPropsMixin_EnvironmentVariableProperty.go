package previewawsopsworksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   environmentVariableProperty := &EnvironmentVariableProperty{
//   	Key: jsii.String("key"),
//   	Secure: jsii.Boolean(false),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environmentvariable.html
//
type CfnAppPropsMixin_EnvironmentVariableProperty struct {
	// (Required) The environment variable's name, which can consist of up to 64 characters and must be specified.
	//
	// The name can contain upper- and lowercase letters, numbers, and underscores (_), but it must start with a letter or underscore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environmentvariable.html#cfn-opsworks-app-environmentvariable-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// (Optional) Whether the variable's value is returned by the `DescribeApps` action.
	//
	// To hide an environment variable's value, set `Secure` to `true` . `DescribeApps` returns `*****FILTERED*****` instead of the actual value. The default value for `Secure` is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environmentvariable.html#cfn-opsworks-app-environmentvariable-secure
	//
	Secure interface{} `field:"optional" json:"secure" yaml:"secure"`
	// (Optional) The environment variable's value, which can be left empty.
	//
	// If you specify a value, it can contain up to 256 characters, which must all be printable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environmentvariable.html#cfn-opsworks-app-environmentvariable-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

