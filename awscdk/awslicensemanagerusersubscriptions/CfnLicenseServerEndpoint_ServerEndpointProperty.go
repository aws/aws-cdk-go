package awslicensemanagerusersubscriptions


// The network address of a license server endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverEndpointProperty := &ServerEndpointProperty{
//   	Endpoint: jsii.String("endpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-serverendpoint.html
//
type CfnLicenseServerEndpoint_ServerEndpointProperty struct {
	// The network address of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-serverendpoint.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-serverendpoint-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

