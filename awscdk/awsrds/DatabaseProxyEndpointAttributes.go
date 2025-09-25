package awsrds


// Properties that describe an existing DB Proxy Endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseProxyEndpointAttributes := &DatabaseProxyEndpointAttributes{
//   	DbProxyEndpointArn: jsii.String("dbProxyEndpointArn"),
//   	DbProxyEndpointName: jsii.String("dbProxyEndpointName"),
//   	Endpoint: jsii.String("endpoint"),
//   }
//
type DatabaseProxyEndpointAttributes struct {
	// DB Proxy Endpoint ARN.
	DbProxyEndpointArn *string `field:"required" json:"dbProxyEndpointArn" yaml:"dbProxyEndpointArn"`
	// DB Proxy Endpoint Name.
	DbProxyEndpointName *string `field:"required" json:"dbProxyEndpointName" yaml:"dbProxyEndpointName"`
	// The endpoint that you can use to connect to the DB proxy.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
}

