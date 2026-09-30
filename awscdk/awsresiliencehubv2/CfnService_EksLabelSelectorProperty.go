package awsresiliencehubv2


// Kubernetes label selector that scopes discovery to matching objects in the specified namespaces.
//
// An object must satisfy both MatchLabels and MatchExpressions. Specify at least one of them; a selector carrying neither is treated as though no selector were supplied, and all supported objects in the specified namespaces are discovered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksLabelSelectorProperty := &EksLabelSelectorProperty{
//   	MatchExpressions: []interface{}{
//   		&EksLabelSelectorRequirementProperty{
//   			Key: jsii.String("key"),
//   			Operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	MatchLabels: map[string]*string{
//   		"matchLabelsKey": jsii.String("matchLabels"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekslabelselector.html
//
type CfnService_EksLabelSelectorProperty struct {
	// Label selector requirements an object must satisfy to be discovered.
	//
	// Up to 20 requirements, all of which must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekslabelselector.html#cfn-resiliencehubv2-service-ekslabelselector-matchexpressions
	//
	MatchExpressions interface{} `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// Label key/value pairs an object must carry to be discovered.
	//
	// Up to 20 pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekslabelselector.html#cfn-resiliencehubv2-service-ekslabelselector-matchlabels
	//
	MatchLabels interface{} `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

