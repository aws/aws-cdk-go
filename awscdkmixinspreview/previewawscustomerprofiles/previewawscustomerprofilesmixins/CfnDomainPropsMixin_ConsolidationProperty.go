package previewawscustomerprofilesmixins


// A list of matching attributes that represent matching criteria.
//
// If two profiles meet at least one of the requirements in the matching attributes list, they will be merged.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   consolidationProperty := &ConsolidationProperty{
//   	MatchingAttributesList: []interface{}{
//   		[]*string{
//   			jsii.String("matchingAttributesList"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-consolidation.html
//
type CfnDomainPropsMixin_ConsolidationProperty struct {
	// A list of matching criteria.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-consolidation.html#cfn-customerprofiles-domain-consolidation-matchingattributeslist
	//
	MatchingAttributesList interface{} `field:"optional" json:"matchingAttributesList" yaml:"matchingAttributesList"`
}

