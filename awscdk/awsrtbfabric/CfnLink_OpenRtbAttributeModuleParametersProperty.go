package awsrtbfabric


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openRtbAttributeModuleParametersProperty := &OpenRtbAttributeModuleParametersProperty{
//   	Action: &ActionProperty{
//   		HeaderTag: &HeaderTagActionProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   		NoBid: &NoBidActionProperty{
//   			NoBidReasonCode: jsii.Number(123),
//   		},
//   	},
//   	FilterConfiguration: []interface{}{
//   		&FilterProperty{
//   			Criteria: []interface{}{
//   				&FilterCriterionProperty{
//   					Path: jsii.String("path"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	FilterType: jsii.String("filterType"),
//   	HoldbackPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-openrtbattributemoduleparameters.html
//
type CfnLink_OpenRtbAttributeModuleParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-openrtbattributemoduleparameters.html#cfn-rtbfabric-link-openrtbattributemoduleparameters-action
	//
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-openrtbattributemoduleparameters.html#cfn-rtbfabric-link-openrtbattributemoduleparameters-filterconfiguration
	//
	FilterConfiguration interface{} `field:"required" json:"filterConfiguration" yaml:"filterConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-openrtbattributemoduleparameters.html#cfn-rtbfabric-link-openrtbattributemoduleparameters-filtertype
	//
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-openrtbattributemoduleparameters.html#cfn-rtbfabric-link-openrtbattributemoduleparameters-holdbackpercentage
	//
	HoldbackPercentage *float64 `field:"required" json:"holdbackPercentage" yaml:"holdbackPercentage"`
}

