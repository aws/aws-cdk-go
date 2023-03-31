package awsopsworks


// Represents an app's environment variable.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentVariableProperty := &environmentVariableProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	secure: jsii.Boolean(false),
//   }
//
type CfnApp_EnvironmentVariableProperty struct {
	// (Required) The environment variable's name, which can consist of up to 64 characters and must be specified.
	//
	// The name can contain upper- and lowercase letters, numbers, and underscores (_), but it must start with a letter or underscore.
	Key *string `field:"required" json:"key" yaml:"key"`
	// (Optional) The environment variable's value, which can be left empty.
	//
	// If you specify a value, it can contain up to 256 characters, which must all be printable.
	Value *string `field:"required" json:"value" yaml:"value"`
	// (Optional) Whether the variable's value is returned by the [DescribeApps](https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeApps) action. To hide an environment variable's value, set `Secure` to `true` . `DescribeApps` returns `*****FILTERED*****` instead of the actual value. The default value for `Secure` is `false` .
	Secure interface{} `field:"optional" json:"secure" yaml:"secure"`
}

