package awsnetworkflowmonitor


// A target identifier is a pair of identifying information for a scope target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetIdentifierProperty := &TargetIdentifierProperty{
//   	TargetId: &TargetIdProperty{
//   		AccountId: jsii.String("accountId"),
//   	},
//   	TargetType: jsii.String("targetType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-scope-targetidentifier.html
//
type CfnScope_TargetIdentifierProperty struct {
	// A target ID is an internally-generated identifier for a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-scope-targetidentifier.html#cfn-networkflowmonitor-scope-targetidentifier-targetid
	//
	TargetId interface{} `field:"required" json:"targetId" yaml:"targetId"`
	// The type of the target.
	//
	// Currently always ACCOUNT.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-scope-targetidentifier.html#cfn-networkflowmonitor-scope-targetidentifier-targettype
	//
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

