package awsdevopsagent


// Grafana MCP server details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registeredMCPServerGrafanaDetailsProperty := &RegisteredMCPServerGrafanaDetailsProperty{
//   	AuthorizationMethod: jsii.String("authorizationMethod"),
//   	Endpoint: jsii.String("endpoint"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpservergrafanadetails.html
//
type CfnService_RegisteredMCPServerGrafanaDetailsProperty struct {
	// MCP server authorization method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpservergrafanadetails.html#cfn-devopsagent-service-registeredmcpservergrafanadetails-authorizationmethod
	//
	AuthorizationMethod *string `field:"required" json:"authorizationMethod" yaml:"authorizationMethod"`
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpservergrafanadetails.html#cfn-devopsagent-service-registeredmcpservergrafanadetails-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Optional description for the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpservergrafanadetails.html#cfn-devopsagent-service-registeredmcpservergrafanadetails-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// MCP server name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpservergrafanadetails.html#cfn-devopsagent-service-registeredmcpservergrafanadetails-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

