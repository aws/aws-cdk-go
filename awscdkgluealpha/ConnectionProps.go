package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Construction properties for `Connection`.
//
// Example:
//   var securityGroup SecurityGroup
//   var subnet Subnet
//
//   glue.NewConnection(this, jsii.String("MyConnection"), &ConnectionProps{
//   	Type: glue.ConnectionType_NETWORK(),
//   	// The security groups granting AWS Glue inbound access to the data source within the VPC
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	// The VPC subnet which contains the data source
//   	Subnet: Subnet,
//   })
//
// Experimental.
type ConnectionProps struct {
	// The name of the connection.
	// Default: cloudformation generated name.
	//
	// Experimental.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The description of the connection.
	// Default: no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of criteria that can be used in selecting this connection.
	//
	// This is useful for filtering the results of https://awscli.amazonaws.com/v2/documentation/api/latest/reference/glue/get-connections.html
	// Default: no match criteria.
	//
	// Experimental.
	MatchCriteria *[]*string `field:"optional" json:"matchCriteria" yaml:"matchCriteria"`
	// Key-Value pairs that define parameters for the connection.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-connect.html
	//
	// Default: empty properties.
	//
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The list of security groups needed to successfully make this connection e.g. to successfully connect to VPC.
	// Default: no security group.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The VPC subnet to connect to resources within a VPC.
	//
	// See more at https://docs.aws.amazon.com/glue/latest/dg/start-connecting.html.
	// Default: no subnet.
	//
	// Experimental.
	Subnet awsec2.ISubnet `field:"optional" json:"subnet" yaml:"subnet"`
	// The type of the connection.
	// Experimental.
	Type ConnectionType `field:"required" json:"type" yaml:"type"`
}

