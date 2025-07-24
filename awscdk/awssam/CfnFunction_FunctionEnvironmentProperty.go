package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionEnvironmentProperty := &FunctionEnvironmentProperty{
//   	Variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-functionenvironment.html
//
type CfnFunction_FunctionEnvironmentProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-functionenvironment.html#cfn-serverless-function-functionenvironment-variables
	//
	Variables interface{} `field:"required" json:"variables" yaml:"variables"`
}

