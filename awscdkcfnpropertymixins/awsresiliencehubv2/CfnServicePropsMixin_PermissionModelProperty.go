package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   permissionModelProperty := &PermissionModelProperty{
//   	CrossAccountRoleArns: []interface{}{
//   		&CrossAccountRoleConfigurationProperty{
//   			CrossAccountRoleArn: jsii.String("crossAccountRoleArn"),
//   			ExternalId: jsii.String("externalId"),
//   		},
//   	},
//   	InvokerRoleName: jsii.String("invokerRoleName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-permissionmodel.html
//
type CfnServicePropsMixin_PermissionModelProperty struct {
	// Cross-account role ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-permissionmodel.html#cfn-resiliencehubv2-service-permissionmodel-crossaccountrolearns
	//
	CrossAccountRoleArns interface{} `field:"optional" json:"crossAccountRoleArns" yaml:"crossAccountRoleArns"`
	// Name of the invoker IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-permissionmodel.html#cfn-resiliencehubv2-service-permissionmodel-invokerrolename
	//
	InvokerRoleName *string `field:"optional" json:"invokerRoleName" yaml:"invokerRoleName"`
}

