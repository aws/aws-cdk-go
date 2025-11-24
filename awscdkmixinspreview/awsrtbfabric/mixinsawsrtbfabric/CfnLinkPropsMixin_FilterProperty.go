package mixinsawsrtbfabric


// Describes the configuration of a filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterProperty := &FilterProperty{
//   	Criteria: []interface{}{
//   		&FilterCriterionProperty{
//   			Path: jsii.String("path"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-filter.html
//
type CfnLinkPropsMixin_FilterProperty struct {
	// Describes the criteria for a filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-filter.html#cfn-rtbfabric-link-filter-criteria
	//
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
}

