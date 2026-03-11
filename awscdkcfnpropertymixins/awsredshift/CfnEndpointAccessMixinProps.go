package awsredshift


// Properties for CfnEndpointAccessPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnEndpointAccessMixinProps := &CfnEndpointAccessMixinProps{
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	EndpointName: jsii.String("endpointName"),
//   	ResourceOwner: jsii.String("resourceOwner"),
//   	SubnetGroupName: jsii.String("subnetGroupName"),
//   	VpcSecurityGroupIds: []interface{}{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html
//
type CfnEndpointAccessMixinProps struct {
	// The cluster identifier of the cluster associated with the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html#cfn-redshift-endpointaccess-clusteridentifier
	//
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The name of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html#cfn-redshift-endpointaccess-endpointname
	//
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// The AWS account ID of the owner of the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html#cfn-redshift-endpointaccess-resourceowner
	//
	ResourceOwner *string `field:"optional" json:"resourceOwner" yaml:"resourceOwner"`
	// The subnet group name where Amazon Redshift chooses to deploy the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html#cfn-redshift-endpointaccess-subnetgroupname
	//
	SubnetGroupName interface{} `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// The security group that defines the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html#cfn-redshift-endpointaccess-vpcsecuritygroupids
	//
	VpcSecurityGroupIds *[]interface{} `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

