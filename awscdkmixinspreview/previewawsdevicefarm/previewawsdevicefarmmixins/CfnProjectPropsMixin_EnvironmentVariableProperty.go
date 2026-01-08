package previewawsdevicefarmmixins


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devicefarm-project-environmentvariable.html
//
type CfnProjectPropsMixin_EnvironmentVariableProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devicefarm-project-environmentvariable.html#cfn-devicefarm-project-environmentvariable-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devicefarm-project-environmentvariable.html#cfn-devicefarm-project-environmentvariable-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

