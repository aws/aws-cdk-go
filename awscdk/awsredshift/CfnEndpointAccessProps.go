package awsredshift


// Properties for defining a `CfnEndpointAccess`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointAccessProps := &CfnEndpointAccessProps{
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	EndpointName: jsii.String("endpointName"),
//   	SubnetGroupName: jsii.String("subnetGroupName"),
//   	VpcSecurityGroupIds: []interface{}{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//
//   	// the properties below are optional
//   	ResourceOwner: jsii.String("resourceOwner"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html
//
type CfnEndpointAccessProps struct {
	// The cluster identifier of the cluster associated with the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html#cfn-redshift-endpointaccess-clusteridentifier
	//
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The name of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html#cfn-redshift-endpointaccess-endpointname
	//
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The subnet group name where Amazon Redshift chooses to deploy the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html#cfn-redshift-endpointaccess-subnetgroupname
	//
	SubnetGroupName interface{} `field:"required" json:"subnetGroupName" yaml:"subnetGroupName"`
	// The security group that defines the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html#cfn-redshift-endpointaccess-vpcsecuritygroupids
	//
	VpcSecurityGroupIds *[]interface{} `field:"required" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// The AWS account ID of the owner of the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html#cfn-redshift-endpointaccess-resourceowner
	//
	ResourceOwner *string `field:"optional" json:"resourceOwner" yaml:"resourceOwner"`
}

