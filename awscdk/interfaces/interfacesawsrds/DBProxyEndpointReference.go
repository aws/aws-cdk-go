package interfacesawsrds


// A reference to a DBProxyEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBProxyEndpointReference := &DBProxyEndpointReference{
//   	DbProxyEndpointArn: jsii.String("dbProxyEndpointArn"),
//   	DbProxyEndpointName: jsii.String("dbProxyEndpointName"),
//   }
//
type DBProxyEndpointReference struct {
	// The ARN of the DBProxyEndpoint resource.
	DbProxyEndpointArn *string `field:"required" json:"dbProxyEndpointArn" yaml:"dbProxyEndpointArn"`
	// The DBProxyEndpointName of the DBProxyEndpoint resource.
	DbProxyEndpointName *string `field:"required" json:"dbProxyEndpointName" yaml:"dbProxyEndpointName"`
}

