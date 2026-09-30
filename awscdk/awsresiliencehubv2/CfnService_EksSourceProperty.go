package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksSourceProperty := &EksSourceProperty{
//   	ClusterArn: jsii.String("clusterArn"),
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//
//   	// the properties below are optional
//   	LabelSelector: &EksLabelSelectorProperty{
//   		MatchExpressions: []interface{}{
//   			&EksLabelSelectorRequirementProperty{
//   				Key: jsii.String("key"),
//   				Operator: jsii.String("operator"),
//
//   				// the properties below are optional
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		MatchLabels: map[string]*string{
//   			"matchLabelsKey": jsii.String("matchLabels"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekssource.html
//
type CfnService_EksSourceProperty struct {
	// ARN of the EKS cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekssource.html#cfn-resiliencehubv2-service-ekssource-clusterarn
	//
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// EKS namespaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekssource.html#cfn-resiliencehubv2-service-ekssource-namespaces
	//
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
	// Kubernetes label selector that scopes discovery to matching objects in the specified namespaces.
	//
	// An object must satisfy both MatchLabels and MatchExpressions. Specify at least one of them; a selector carrying neither is treated as though no selector were supplied, and all supported objects in the specified namespaces are discovered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekssource.html#cfn-resiliencehubv2-service-ekssource-labelselector
	//
	LabelSelector interface{} `field:"optional" json:"labelSelector" yaml:"labelSelector"`
}

