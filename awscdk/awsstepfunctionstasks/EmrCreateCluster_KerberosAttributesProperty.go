package awsstepfunctionstasks


// Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration.
//
// See the RunJobFlow API for complete documentation on input parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kerberosAttributesProperty := &KerberosAttributesProperty{
//   	Realm: jsii.String("realm"),
//
//   	// the properties below are optional
//   	AdDomainJoinPassword: jsii.String("adDomainJoinPassword"),
//   	AdDomainJoinUser: jsii.String("adDomainJoinUser"),
//   	CrossRealmTrustPrincipalPassword: jsii.String("crossRealmTrustPrincipalPassword"),
//   	KdcAdminPassword: jsii.String("kdcAdminPassword"),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_KerberosAttributes.html
//
type EmrCreateCluster_KerberosAttributesProperty struct {
	// The name of the Kerberos realm to which all nodes in a cluster belong.
	//
	// For example, EC2.INTERNAL.
	Realm *string `field:"required" json:"realm" yaml:"realm"`
	// The Active Directory password for ADDomainJoinUser.
	// Default: No adDomainJoinPassword.
	//
	AdDomainJoinPassword *string `field:"optional" json:"adDomainJoinPassword" yaml:"adDomainJoinPassword"`
	// Required only when establishing a cross-realm trust with an Active Directory domain.
	//
	// A user with sufficient privileges to join
	// resources to the domain.
	// Default: No adDomainJoinUser.
	//
	AdDomainJoinUser *string `field:"optional" json:"adDomainJoinUser" yaml:"adDomainJoinUser"`
	// Required only when establishing a cross-realm trust with a KDC in a different realm.
	//
	// The cross-realm principal password, which
	// must be identical across realms.
	// Default: No crossRealmTrustPrincipalPassword.
	//
	CrossRealmTrustPrincipalPassword *string `field:"optional" json:"crossRealmTrustPrincipalPassword" yaml:"crossRealmTrustPrincipalPassword"`
	// The password used within the cluster for the kadmin service on the cluster-dedicated KDC, which maintains Kerberos principals, password policies, and keytabs for the cluster.
	// Default: No kdcAdminPassword.
	//
	KdcAdminPassword *string `field:"optional" json:"kdcAdminPassword" yaml:"kdcAdminPassword"`
}

