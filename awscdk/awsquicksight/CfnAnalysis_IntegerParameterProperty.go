package awsquicksight


// An integer parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerParameterProperty := &IntegerParameterProperty{
//   	Name: jsii.String("name"),
//   	Values: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-integerparameter.html
//
type CfnAnalysis_IntegerParameterProperty struct {
	// The name of the integer parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-integerparameter.html#cfn-quicksight-analysis-integerparameter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values for the integer parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-integerparameter.html#cfn-quicksight-analysis-integerparameter-values
	//
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

