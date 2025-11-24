package mixinsawsredshift


// The security groups associated with the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcSecurityGroupProperty := &VpcSecurityGroupProperty{
//   	Status: jsii.String("status"),
//   	VpcSecurityGroupId: jsii.String("vpcSecurityGroupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcsecuritygroup.html
//
type CfnEndpointAccessPropsMixin_VpcSecurityGroupProperty struct {
	// The status of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcsecuritygroup.html#cfn-redshift-endpointaccess-vpcsecuritygroup-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The identifier of the VPC security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcsecuritygroup.html#cfn-redshift-endpointaccess-vpcsecuritygroup-vpcsecuritygroupid
	//
	VpcSecurityGroupId *string `field:"optional" json:"vpcSecurityGroupId" yaml:"vpcSecurityGroupId"`
}

