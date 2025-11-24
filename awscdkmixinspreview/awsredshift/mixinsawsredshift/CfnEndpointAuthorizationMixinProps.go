package mixinsawsredshift


// Properties for CfnEndpointAuthorizationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEndpointAuthorizationMixinProps := &CfnEndpointAuthorizationMixinProps{
//   	Account: jsii.String("account"),
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	Force: jsii.Boolean(false),
//   	VpcIds: []*string{
//   		jsii.String("vpcIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointauthorization.html
//
type CfnEndpointAuthorizationMixinProps struct {
	// The AWS account ID of either the cluster owner (grantor) or grantee.
	//
	// If `Grantee` parameter is true, then the `Account` value is of the grantor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointauthorization.html#cfn-redshift-endpointauthorization-account
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The cluster identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointauthorization.html#cfn-redshift-endpointauthorization-clusteridentifier
	//
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Indicates whether to force the revoke action.
	//
	// If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointauthorization.html#cfn-redshift-endpointauthorization-force
	//
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// The virtual private cloud (VPC) identifiers to grant access to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointauthorization.html#cfn-redshift-endpointauthorization-vpcids
	//
	VpcIds *[]*string `field:"optional" json:"vpcIds" yaml:"vpcIds"`
}

