package previewawsamplifymixins


// Environment variables are key-value pairs that are available at build time.
//
// Set environment variables for all branches in your app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   environmentVariableProperty := &EnvironmentVariableProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-environmentvariable.html
//
type CfnAppPropsMixin_EnvironmentVariableProperty struct {
	// The environment variable name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-environmentvariable.html#cfn-amplify-app-environmentvariable-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The environment variable value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-environmentvariable.html#cfn-amplify-app-environmentvariable-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

