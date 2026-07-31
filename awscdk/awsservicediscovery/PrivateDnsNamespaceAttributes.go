package awsservicediscovery


// Example:
//   var vpc Vpc
//
//
//   importedNamespace := cloudmap.PrivateDnsNamespace_FromPrivateDnsNamespaceAttributes(this, jsii.String("ImportedNamespace"), &PrivateDnsNamespaceAttributes{
//   	NamespaceId: jsii.String("ns-xxxxxxxxxxxxx"),
//   	NamespaceArn: jsii.String("arn:aws:servicediscovery:us-east-1:123456789012:namespace/ns-xxxxxxxxxxxxx"),
//   	NamespaceName: jsii.String("example.local"),
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   cluster.AddExistingDefaultCloudMapNamespace(&ExistingCloudMapNamespaceOptions{
//   	Namespace: importedNamespace,
//   })
//
type PrivateDnsNamespaceAttributes struct {
	// Namespace ARN for the Namespace.
	NamespaceArn *string `field:"required" json:"namespaceArn" yaml:"namespaceArn"`
	// Namespace Id for the Namespace.
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
	// A name for the Namespace.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
}

