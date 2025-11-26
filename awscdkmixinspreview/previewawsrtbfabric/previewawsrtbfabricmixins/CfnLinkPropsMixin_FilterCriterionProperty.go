package previewawsrtbfabricmixins


// Describes the criteria for a filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterCriterionProperty := &FilterCriterionProperty{
//   	Path: jsii.String("path"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-filtercriterion.html
//
type CfnLinkPropsMixin_FilterCriterionProperty struct {
	// The path to filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-filtercriterion.html#cfn-rtbfabric-link-filtercriterion-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The value to filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-filtercriterion.html#cfn-rtbfabric-link-filtercriterion-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

