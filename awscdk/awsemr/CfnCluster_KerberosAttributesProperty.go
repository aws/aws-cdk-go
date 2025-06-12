package awsemr


// `KerberosAttributes` is a property of the `AWS::EMR::Cluster` resource.
//
// `KerberosAttributes` define the cluster-specific Kerberos configuration when Kerberos authentication is enabled using a security configuration. The cluster-specific configuration must be compatible with the security configuration. For more information see [Use Kerberos Authentication](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-kerberos.html) in the *EMR Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kerberosAttributesProperty := &KerberosAttributesProperty{
//   	KdcAdminPassword: jsii.String("kdcAdminPassword"),
//   	Realm: jsii.String("realm"),
//
//   	// the properties below are optional
//   	AdDomainJoinPassword: jsii.String("adDomainJoinPassword"),
//   	AdDomainJoinUser: jsii.String("adDomainJoinUser"),
//   	CrossRealmTrustPrincipalPassword: jsii.String("crossRealmTrustPrincipalPassword"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-kerberosattributes.html
//
type CfnCluster_KerberosAttributesProperty struct {
	// The password used within the cluster for the kadmin service on the cluster-dedicated KDC, which maintains Kerberos principals, password policies, and keytabs for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-kerberosattributes.html#cfn-emr-cluster-kerberosattributes-kdcadminpassword
	//
	KdcAdminPassword *string `field:"required" json:"kdcAdminPassword" yaml:"kdcAdminPassword"`
	// The name of the Kerberos realm to which all nodes in a cluster belong.
	//
	// For example, `EC2.INTERNAL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-kerberosattributes.html#cfn-emr-cluster-kerberosattributes-realm
	//
	Realm *string `field:"required" json:"realm" yaml:"realm"`
	// The Active Directory password for `ADDomainJoinUser` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-kerberosattributes.html#cfn-emr-cluster-kerberosattributes-addomainjoinpassword
	//
	AdDomainJoinPassword *string `field:"optional" json:"adDomainJoinPassword" yaml:"adDomainJoinPassword"`
	// Required only when establishing a cross-realm trust with an Active Directory domain.
	//
	// A user with sufficient privileges to join resources to the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-kerberosattributes.html#cfn-emr-cluster-kerberosattributes-addomainjoinuser
	//
	AdDomainJoinUser *string `field:"optional" json:"adDomainJoinUser" yaml:"adDomainJoinUser"`
	// Required only when establishing a cross-realm trust with a KDC in a different realm.
	//
	// The cross-realm principal password, which must be identical across realms.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-kerberosattributes.html#cfn-emr-cluster-kerberosattributes-crossrealmtrustprincipalpassword
	//
	CrossRealmTrustPrincipalPassword *string `field:"optional" json:"crossRealmTrustPrincipalPassword" yaml:"crossRealmTrustPrincipalPassword"`
}

