package previewawsbatchmixins


// The Environment property type specifies environment variables to use in a job definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   environmentProperty := &EnvironmentProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-environment.html
//
type CfnJobDefinitionPropsMixin_EnvironmentProperty struct {
	// The name of the environment variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-environment.html#cfn-batch-jobdefinition-environment-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the environment variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-environment.html#cfn-batch-jobdefinition-environment-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

