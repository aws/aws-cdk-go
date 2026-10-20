package awsworkspaces


// Describes the access type and endpoint for a WorkSpace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessEndpointProperty := &AccessEndpointProperty{
//   	AccessEndpointType: jsii.String("accessEndpointType"),
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-accessendpoint.html
//
type CfnDirectory_AccessEndpointProperty struct {
	// Indicates the type of access endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-accessendpoint.html#cfn-workspaces-directory-accessendpoint-accessendpointtype
	//
	AccessEndpointType *string `field:"optional" json:"accessEndpointType" yaml:"accessEndpointType"`
	// Indicates the VPC endpoint to use for access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-accessendpoint.html#cfn-workspaces-directory-accessendpoint-vpcendpointid
	//
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

