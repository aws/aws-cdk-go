package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecretsmanager"
)

// Base Connection Options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secretRef ISecretRef
//   var securityGroup SecurityGroup
//   var subnet Subnet
//   var subnetFilter SubnetFilter
//   var vpc Vpc
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
//   	Secret: secretRef,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	Subnet: subnet,
//   	Vpc: vpc,
//   	VpcSubnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []SubnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		Subnets: []ISubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
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
	// The VPC subnet to connect to resources within a VPC. See more at https://docs.aws.amazon.com/glue/latest/dg/start-connecting.html.
	//
	// Mutually exclusive with `vpc`: provide `subnet` to pin the connection to a
	// specific subnet, or provide `vpc` (optionally with `vpcSubnets`) to let the
	// CDK select one for you.
	// Default: - no subnet, unless `vpc` is provided.
	//
	// Experimental.
	Subnet awsec2.ISubnet `field:"optional" json:"subnet" yaml:"subnet"`
	// The VPC to connect to resources within.
	//
	// When provided, the CDK selects a
	// subnet from this VPC using `vpcSubnets`. A Glue connection targets a single
	// subnet, so the first subnet of the selection is used.
	//
	// Mutually exclusive with `subnet`.
	// Default: - no VPC, the subnet is taken from `subnet` if provided.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Which subnets of `vpc` to select the connection subnet from.
	//
	// Only used when
	// `vpc` is provided. Since a Glue connection targets a single subnet, the
	// first subnet of the selection is used.
	// Default: - private subnets.
	//
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

