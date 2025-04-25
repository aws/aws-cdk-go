package awsmacie


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   findingsFilterListItemProperty := &FindingsFilterListItemProperty{
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-findingsfilterlistitem.html
//
type CfnFindingsFilter_FindingsFilterListItemProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-findingsfilterlistitem.html#cfn-macie-findingsfilter-findingsfilterlistitem-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-findingsfilterlistitem.html#cfn-macie-findingsfilter-findingsfilterlistitem-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

