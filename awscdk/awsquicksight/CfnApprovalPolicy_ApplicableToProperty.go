package awsquicksight


// Scoping: who the policy applies to.
//
// GROUP: `groupArns` required (one or more group ARNs).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var type interface{}
//
//   applicableToProperty := &ApplicableToProperty{
//   	Type: type,
//
//   	// the properties below are optional
//   	GroupArns: []*string{
//   		jsii.String("groupArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-approvalpolicy-applicableto.html
//
type CfnApprovalPolicy_ApplicableToProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-approvalpolicy-applicableto.html#cfn-quicksight-approvalpolicy-applicableto-type
	//
	Type interface{} `field:"required" json:"type" yaml:"type"`
	// Required when type = GROUP.
	//
	// One or more group ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-approvalpolicy-applicableto.html#cfn-quicksight-approvalpolicy-applicableto-grouparns
	//
	GroupArns *[]*string `field:"optional" json:"groupArns" yaml:"groupArns"`
}

