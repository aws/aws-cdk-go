package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Base Connection Options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//   var subnet subnet
//
//   connectionOptions := &ConnectionOptions{
//   	ConnectionName: jsii.String("connectionName"),
//   	Description: jsii.String("description"),
//   	MatchCriteria: []*string{
//   		jsii.String("matchCriteria"),
//   	},
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	Subnet: subnet,
//   }
//
// Experimental.
type ConnectionOptions struct {
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
}

