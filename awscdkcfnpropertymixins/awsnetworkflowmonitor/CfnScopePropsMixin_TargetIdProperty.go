package awsnetworkflowmonitor


// A target ID is an internally-generated identifier for a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   targetIdProperty := &TargetIdProperty{
//   	AccountId: jsii.String("accountId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-scope-targetid.html
//
type CfnScopePropsMixin_TargetIdProperty struct {
	// The account ID for the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-scope-targetid.html#cfn-networkflowmonitor-scope-targetid-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
}

