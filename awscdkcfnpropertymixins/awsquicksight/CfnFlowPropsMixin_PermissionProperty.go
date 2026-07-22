package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   permissionProperty := &PermissionProperty{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	Principal: jsii.String("principal"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-flow-permission.html
//
type CfnFlowPropsMixin_PermissionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-flow-permission.html#cfn-quicksight-flow-permission-actions
	//
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-flow-permission.html#cfn-quicksight-flow-permission-principal
	//
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
}

