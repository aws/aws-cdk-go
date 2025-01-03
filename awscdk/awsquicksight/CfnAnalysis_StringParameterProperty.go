package awsquicksight


// A string parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringParameterProperty := &StringParameterProperty{
//   	Name: jsii.String("name"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-stringparameter.html
//
type CfnAnalysis_StringParameterProperty struct {
	// A display name for a string parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-stringparameter.html#cfn-quicksight-analysis-stringparameter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values of a string parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-stringparameter.html#cfn-quicksight-analysis-stringparameter-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

