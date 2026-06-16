package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   permissionModelProperty := &PermissionModelProperty{
//   	InvokerRoleName: jsii.String("invokerRoleName"),
//
//   	// the properties below are optional
//   	CrossAccountRoleArns: []interface{}{
//   		&CrossAccountRoleConfigurationProperty{
//   			CrossAccountRoleArn: jsii.String("crossAccountRoleArn"),
//
//   			// the properties below are optional
//   			ExternalId: jsii.String("externalId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-permissionmodel.html
//
type CfnService_PermissionModelProperty struct {
	// Name of the invoker IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-permissionmodel.html#cfn-resiliencehubv2-service-permissionmodel-invokerrolename
	//
	InvokerRoleName *string `field:"required" json:"invokerRoleName" yaml:"invokerRoleName"`
	// Cross-account role ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-permissionmodel.html#cfn-resiliencehubv2-service-permissionmodel-crossaccountrolearns
	//
	CrossAccountRoleArns interface{} `field:"optional" json:"crossAccountRoleArns" yaml:"crossAccountRoleArns"`
}

