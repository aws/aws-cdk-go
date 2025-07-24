package awsdatazone


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-environmentparameter.html
//
type CfnProject_EnvironmentParameterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-environmentparameter.html#cfn-datazone-project-environmentparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-environmentparameter.html#cfn-datazone-project-environmentparameter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

