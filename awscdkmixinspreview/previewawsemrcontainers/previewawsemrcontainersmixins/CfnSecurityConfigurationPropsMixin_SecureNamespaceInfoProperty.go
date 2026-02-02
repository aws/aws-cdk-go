package previewawsemrcontainersmixins


// Secure namespace information for Lake Formation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   secureNamespaceInfoProperty := &SecureNamespaceInfoProperty{
//   	ClusterId: jsii.String("clusterId"),
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-securenamespaceinfo.html
//
type CfnSecurityConfigurationPropsMixin_SecureNamespaceInfoProperty struct {
	// The ID of the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-securenamespaceinfo.html#cfn-emrcontainers-securityconfiguration-securenamespaceinfo-clusterid
	//
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// The namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-securenamespaceinfo.html#cfn-emrcontainers-securityconfiguration-securenamespaceinfo-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

