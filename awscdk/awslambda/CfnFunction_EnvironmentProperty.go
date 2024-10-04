package awslambda


// A function's environment variable settings.
//
// You can use environment variables to adjust your function's behavior without updating code. An environment variable is a pair of strings that are stored in a function's version-specific configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentProperty := &EnvironmentProperty{
//   	Variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html
//
type CfnFunction_EnvironmentProperty struct {
	// Environment variable key-value pairs.
	//
	// For more information, see [Using Lambda environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html#cfn-lambda-function-environment-variables
	//
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

