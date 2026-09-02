package awsquicksight


// Scoping: who the policy applies to.
//
// GROUP: `groupArns` required (one or more group ARNs).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var type interface{}
//
//   applicableToProperty := &ApplicableToProperty{
//   	GroupArns: []*string{
//   		jsii.String("groupArns"),
//   	},
//   	Type: type,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-approvalpolicy-applicableto.html
//
type CfnApprovalPolicyPropsMixin_ApplicableToProperty struct {
	// Required when type = GROUP.
	//
	// One or more group ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-approvalpolicy-applicableto.html#cfn-quicksight-approvalpolicy-applicableto-grouparns
	//
	GroupArns *[]*string `field:"optional" json:"groupArns" yaml:"groupArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-approvalpolicy-applicableto.html#cfn-quicksight-approvalpolicy-applicableto-type
	//
	Type interface{} `field:"optional" json:"type" yaml:"type"`
}

