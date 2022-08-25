package awsredshift


// Properties for defining a `CfnEndpointAuthorization`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointAuthorizationProps := &cfnEndpointAuthorizationProps{
//   	account: jsii.String("account"),
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//
//   	// the properties below are optional
//   	force: jsii.Boolean(false),
//   	vpcIds: []*string{
//   		jsii.String("vpcIds"),
//   	},
//   }
//
type CfnEndpointAuthorizationProps struct {
	// The A AWS account ID of either the cluster owner (grantor) or grantee.
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

