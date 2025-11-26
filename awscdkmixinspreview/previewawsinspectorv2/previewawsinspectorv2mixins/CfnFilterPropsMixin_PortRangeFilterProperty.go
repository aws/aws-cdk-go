package previewawsinspectorv2mixins


// An object that describes the details of a port range filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   portRangeFilterProperty := &PortRangeFilterProperty{
//   	BeginInclusive: jsii.Number(123),
//   	EndInclusive: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-portrangefilter.html
//
type CfnFilterPropsMixin_PortRangeFilterProperty struct {
	// The port number the port range begins at.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-portrangefilter.html#cfn-inspectorv2-filter-portrangefilter-begininclusive
	//
	BeginInclusive *float64 `field:"optional" json:"beginInclusive" yaml:"beginInclusive"`
	// The port number the port range ends at.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-portrangefilter.html#cfn-inspectorv2-filter-portrangefilter-endinclusive
	//
	EndInclusive *float64 `field:"optional" json:"endInclusive" yaml:"endInclusive"`
}

