package previewawstransfermixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   endpointDetailsProperty := &EndpointDetailsProperty{
//   	Vpc: &VpcProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-webapp-endpointdetails.html
//
type CfnWebAppPropsMixin_EndpointDetailsProperty struct {
	// You can provide a structure that contains the details for the VPC endpoint to use with your web app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-webapp-endpointdetails.html#cfn-transfer-webapp-endpointdetails-vpc
	//
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

