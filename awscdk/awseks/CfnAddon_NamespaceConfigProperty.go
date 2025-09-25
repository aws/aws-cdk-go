package awseks


// The custom namespace configuration to use with the add-on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   namespaceConfigProperty := &NamespaceConfigProperty{
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-addon-namespaceconfig.html
//
type CfnAddon_NamespaceConfigProperty struct {
	// The custom namespace for creating the add-on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-addon-namespaceconfig.html#cfn-eks-addon-namespaceconfig-namespace
	//
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

