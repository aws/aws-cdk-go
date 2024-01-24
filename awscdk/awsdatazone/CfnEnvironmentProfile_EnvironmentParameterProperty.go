package awsdatazone


// The parameter details of an environment profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentParameterProperty := &EnvironmentParameterProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentprofile-environmentparameter.html
//
type CfnEnvironmentProfile_EnvironmentParameterProperty struct {
	// The name of an environment profile parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentprofile-environmentparameter.html#cfn-datazone-environmentprofile-environmentparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of an environment profile parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentprofile-environmentparameter.html#cfn-datazone-environmentprofile-environmentparameter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

