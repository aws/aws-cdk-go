package awscustomerprofiles


// Defines how segments should be sorted and ordered in the results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   segmentSortProperty := &SegmentSortProperty{
//   	Attributes: []interface{}{
//   		&SortAttributeProperty{
//   			DataType: jsii.String("dataType"),
//   			Name: jsii.String("name"),
//   			Order: jsii.String("order"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-segmentsort.html
//
type CfnSegmentDefinitionPropsMixin_SegmentSortProperty struct {
	// A list of attributes used to sort the segments and their ordering preferences.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-segmentsort.html#cfn-customerprofiles-segmentdefinition-segmentsort-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
}

