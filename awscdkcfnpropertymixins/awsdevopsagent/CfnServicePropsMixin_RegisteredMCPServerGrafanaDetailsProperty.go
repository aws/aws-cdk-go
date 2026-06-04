package awsdevopsagent


// Grafana MCP server details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   registeredMCPServerGrafanaDetailsProperty := &RegisteredMCPServerGrafanaDetailsProperty{
//   	AuthorizationMethod: jsii.String("authorizationMethod"),
//   	Description: jsii.String("description"),
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpservergrafanadetails.html
//
type CfnServicePropsMixin_RegisteredMCPServerGrafanaDetailsProperty struct {
	// MCP server authorization method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpservergrafanadetails.html#cfn-devopsagent-service-registeredmcpservergrafanadetails-authorizationmethod
	//
	AuthorizationMethod *string `field:"optional" json:"authorizationMethod" yaml:"authorizationMethod"`
	// Optional description for the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpservergrafanadetails.html#cfn-devopsagent-service-registeredmcpservergrafanadetails-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpservergrafanadetails.html#cfn-devopsagent-service-registeredmcpservergrafanadetails-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// MCP server name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpservergrafanadetails.html#cfn-devopsagent-service-registeredmcpservergrafanadetails-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

