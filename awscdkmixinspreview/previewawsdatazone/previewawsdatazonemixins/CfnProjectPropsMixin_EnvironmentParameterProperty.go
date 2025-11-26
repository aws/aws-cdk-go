package previewawsdatazonemixins


// The parameter details of an evironment profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   environmentParameterProperty := &EnvironmentParameterProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-environmentparameter.html
//
type CfnProjectPropsMixin_EnvironmentParameterProperty struct {
	// The name of an environment profile parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-environmentparameter.html#cfn-datazone-project-environmentparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of an environment profile parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-environmentparameter.html#cfn-datazone-project-environmentparameter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

