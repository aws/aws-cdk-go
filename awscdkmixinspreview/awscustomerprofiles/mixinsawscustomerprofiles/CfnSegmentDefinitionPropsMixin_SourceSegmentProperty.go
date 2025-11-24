package mixinsawscustomerprofiles


// The source segments to build off of.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceSegmentProperty := &SourceSegmentProperty{
//   	SegmentDefinitionName: jsii.String("segmentDefinitionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-sourcesegment.html
//
type CfnSegmentDefinitionPropsMixin_SourceSegmentProperty struct {
	// The name of the source segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-sourcesegment.html#cfn-customerprofiles-segmentdefinition-sourcesegment-segmentdefinitionname
	//
	SegmentDefinitionName *string `field:"optional" json:"segmentDefinitionName" yaml:"segmentDefinitionName"`
}

