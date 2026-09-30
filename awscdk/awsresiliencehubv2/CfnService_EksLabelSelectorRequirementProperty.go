package awsresiliencehubv2


// A single label selector requirement.
//
// Specify Values when Operator is IN or NOT_IN, and omit Values when Operator is EXISTS or DOES_NOT_EXIST.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksLabelSelectorRequirementProperty := &EksLabelSelectorRequirementProperty{
//   	Key: jsii.String("key"),
//   	Operator: jsii.String("operator"),
//
//   	// the properties below are optional
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekslabelselectorrequirement.html
//
type CfnService_EksLabelSelectorRequirementProperty struct {
	// Label key the requirement applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekslabelselectorrequirement.html#cfn-resiliencehubv2-service-ekslabelselectorrequirement-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// Operator applied to the label key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekslabelselectorrequirement.html#cfn-resiliencehubv2-service-ekslabelselectorrequirement-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Label values the requirement compares against.
	//
	// Up to 20 values. Required for IN and NOT_IN; omit for EXISTS and DOES_NOT_EXIST.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekslabelselectorrequirement.html#cfn-resiliencehubv2-service-ekslabelselectorrequirement-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

