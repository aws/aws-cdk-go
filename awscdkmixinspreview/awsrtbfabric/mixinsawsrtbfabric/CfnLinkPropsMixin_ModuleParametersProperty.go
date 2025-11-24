package mixinsawsrtbfabric


// Describes the parameters of a module.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   moduleParametersProperty := &ModuleParametersProperty{
//   	NoBid: &NoBidModuleParametersProperty{
//   		PassThroughPercentage: jsii.Number(123),
//   		Reason: jsii.String("reason"),
//   		ReasonCode: jsii.Number(123),
//   	},
//   	OpenRtbAttribute: &OpenRtbAttributeModuleParametersProperty{
//   		Action: &ActionProperty{
//   			HeaderTag: &HeaderTagActionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   			NoBid: &NoBidActionProperty{
//   				NoBidReasonCode: jsii.Number(123),
//   			},
//   		},
//   		FilterConfiguration: []interface{}{
//   			&FilterProperty{
//   				Criteria: []interface{}{
//   					&FilterCriterionProperty{
//   						Path: jsii.String("path"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		FilterType: jsii.String("filterType"),
//   		HoldbackPercentage: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-moduleparameters.html
//
type CfnLinkPropsMixin_ModuleParametersProperty struct {
	// Describes the parameters of a no bid module.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-moduleparameters.html#cfn-rtbfabric-link-moduleparameters-nobid
	//
	NoBid interface{} `field:"optional" json:"noBid" yaml:"noBid"`
	// Describes the parameters of an open RTB attribute module.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-moduleparameters.html#cfn-rtbfabric-link-moduleparameters-openrtbattribute
	//
	OpenRtbAttribute interface{} `field:"optional" json:"openRtbAttribute" yaml:"openRtbAttribute"`
}

