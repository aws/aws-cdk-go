package awscontroltower


// A key-value parameter to an `EnabledBaseline` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   parameterProperty := &ParameterProperty{
//   	Key: jsii.String("key"),
//   	Value: value,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controltower-enabledbaseline-parameter.html
//
type CfnEnabledBaseline_ParameterProperty struct {
	// A string denoting the parameter key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controltower-enabledbaseline-parameter.html#cfn-controltower-enabledbaseline-parameter-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A low-level `Document` object of any type (for example, a Java Object).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controltower-enabledbaseline-parameter.html#cfn-controltower-enabledbaseline-parameter-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

