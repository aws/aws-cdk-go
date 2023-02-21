package awsredshift


// Properties for defining a `CfnEndpointAccess`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpcEndpoint interface{}
//
//   cfnEndpointAccessProps := &cfnEndpointAccessProps{
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   	endpointName: jsii.String("endpointName"),
//   	subnetGroupName: jsii.String("subnetGroupName"),
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//
//   	// the properties below are optional
//   	resourceOwner: jsii.String("resourceOwner"),
//   	vpcEndpoint: vpcEndpoint,
//   	vpcSecurityGroups: []interface{}{
//   		&vpcSecurityGroupProperty{
//   			status: jsii.String("status"),
//   			vpcSecurityGroupId: jsii.String("vpcSecurityGroupId"),
//   		},
//   	},
//   }
//
type CfnEndpointAccessProps struct {
	// The cluster identifier of the cluster associated with the endpoint.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The name of the endpoint.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The subnet group name where Amazon Redshift chooses to deploy the endpoint.
	SubnetGroupName *string `field:"required" json:"subnetGroupName" yaml:"subnetGroupName"`
	// The security group that defines the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.
	VpcSecurityGroupIds *[]*string `field:"required" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// The AWS account ID of the owner of the cluster.
	ResourceOwner *string `field:"optional" json:"resourceOwner" yaml:"resourceOwner"`
	// `AWS::Redshift::EndpointAccess.VpcEndpoint`.
	VpcEndpoint interface{} `field:"optional" json:"vpcEndpoint" yaml:"vpcEndpoint"`
	// `AWS::Redshift::EndpointAccess.VpcSecurityGroups`.
	VpcSecurityGroups interface{} `field:"optional" json:"vpcSecurityGroups" yaml:"vpcSecurityGroups"`
}

