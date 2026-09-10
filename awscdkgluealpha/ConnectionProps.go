package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecretsmanager"
)

// Construction properties for `Connection`.
//
// Example:
//   var securityGroup SecurityGroup
//   var vpc Vpc
//
//   glue.NewConnection(this, jsii.String("MyConnection"), &ConnectionProps{
//   	Type: glue.ConnectionType_NETWORK(),
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	// vpcSubnets is optional - defaults to private subnets
//   	Network: glue.ConnectionNetwork_Vpc(vpc, &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	}),
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
	// The VPC network placement for this connection, so it can reach resources inside a VPC. See more at https://docs.aws.amazon.com/glue/latest/dg/start-connecting.html.
	//
	// Build it with `ConnectionNetwork.subnet(subnet)` to pin a specific subnet,
	// or `ConnectionNetwork.vpc(vpc, vpcSubnets?)` to let the CDK select one.
	// Default: - no VPC network placement.
	//
	// Experimental.
	Network ConnectionNetwork `field:"optional" json:"network" yaml:"network"`
	// Key-Value pairs that define parameters for the connection.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-connect.html
	//
	// Default: empty properties.
	//
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// A reference to a Secrets Manager secret holding the credentials for this connection.
	//
	// The secret is referenced through the connection's `SECRET_ID` property, so
	// Glue reads the credentials at runtime and the secret value never appears in
	// the synthesized template. Prefer this over placing credentials directly in
	// `properties`. Accepts any `secretsmanager.ISecret`.
	// Default: - no secret; any credentials must be supplied via `properties`.
	//
	// Experimental.
	Secret interfacesawssecretsmanager.ISecretRef `field:"optional" json:"secret" yaml:"secret"`
	// The list of security groups needed to successfully make this connection e.g. to successfully connect to VPC.
	// Default: no security group.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The type of the connection.
	// Experimental.
	Type ConnectionType `field:"required" json:"type" yaml:"type"`
}

