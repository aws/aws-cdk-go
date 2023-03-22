package awsredshift


// Properties for defining a `CfnEndpointAuthorization`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointAuthorizationProps := &CfnEndpointAuthorizationProps{
//   	Account: jsii.String("account"),
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//
//   	// the properties below are optional
//   	Force: jsii.Boolean(false),
//   	VpcIds: []*string{
//   		jsii.String("vpcIds"),
//   	},
//   }
//
type CfnEndpointAuthorizationProps struct {
	// The AWS account ID of either the cluster owner (grantor) or grantee.
	//
	// If `Grantee` parameter is true, then the `Account` value is of the grantor.
	Account *string `field:"required" json:"account" yaml:"account"`
	// The cluster identifier.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Indicates whether to force the revoke action.
	//
	// If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// The virtual private cloud (VPC) identifiers to grant access to.
	VpcIds *[]*string `field:"optional" json:"vpcIds" yaml:"vpcIds"`
}

