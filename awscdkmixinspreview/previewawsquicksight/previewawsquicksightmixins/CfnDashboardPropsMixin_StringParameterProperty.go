package previewawsquicksightmixins


// A string parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stringParameterProperty := &StringParameterProperty{
//   	Name: jsii.String("name"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-stringparameter.html
//
type CfnDashboardPropsMixin_StringParameterProperty struct {
	// A display name for a string parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-stringparameter.html#cfn-quicksight-dashboard-stringparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The values of a string parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-stringparameter.html#cfn-quicksight-dashboard-stringparameter-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

