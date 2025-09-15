package awsdatazone


// The environment configuration parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentConfigurationParameterProperty := &EnvironmentConfigurationParameterProperty{
//   	IsEditable: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfigurationparameter.html
//
type CfnProjectProfile_EnvironmentConfigurationParameterProperty struct {
	// Specifies whether the environment parameter is editable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfigurationparameter.html#cfn-datazone-projectprofile-environmentconfigurationparameter-iseditable
	//
	IsEditable interface{} `field:"optional" json:"isEditable" yaml:"isEditable"`
	// The name of the environment configuration parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfigurationparameter.html#cfn-datazone-projectprofile-environmentconfigurationparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the environment configuration parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfigurationparameter.html#cfn-datazone-projectprofile-environmentconfigurationparameter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

