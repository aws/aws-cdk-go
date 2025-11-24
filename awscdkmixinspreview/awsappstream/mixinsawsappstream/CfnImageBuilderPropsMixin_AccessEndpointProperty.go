package mixinsawsappstream


// Describes an interface VPC endpoint (interface endpoint) that lets you create a private connection between the virtual private cloud (VPC) that you specify and WorkSpaces Applications.
//
// When you specify an interface endpoint for a stack, users of the stack can connect to WorkSpaces Applications only through that endpoint. When you specify an interface endpoint for an image builder, administrators can connect to the image builder only through that endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accessEndpointProperty := &AccessEndpointProperty{
//   	EndpointType: jsii.String("endpointType"),
//   	VpceId: jsii.String("vpceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-imagebuilder-accessendpoint.html
//
type CfnImageBuilderPropsMixin_AccessEndpointProperty struct {
	// The type of interface endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-imagebuilder-accessendpoint.html#cfn-appstream-imagebuilder-accessendpoint-endpointtype
	//
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// The identifier (ID) of the VPC in which the interface endpoint is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-imagebuilder-accessendpoint.html#cfn-appstream-imagebuilder-accessendpoint-vpceid
	//
	VpceId *string `field:"optional" json:"vpceId" yaml:"vpceId"`
}

