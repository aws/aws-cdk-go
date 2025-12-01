package previewawseksmixins


// The access configuration for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accessConfigProperty := &AccessConfigProperty{
//   	AuthenticationMode: jsii.String("authenticationMode"),
//   	BootstrapClusterCreatorAdminPermissions: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-accessconfig.html
//
type CfnClusterPropsMixin_AccessConfigProperty struct {
	// The desired authentication mode for the cluster.
	//
	// If you create a cluster by using the EKS API, AWS SDKs, or AWS CloudFormation , the default is `CONFIG_MAP` . If you create the cluster by using the AWS Management Console , the default value is `API_AND_CONFIG_MAP` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-accessconfig.html#cfn-eks-cluster-accessconfig-authenticationmode
	//
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Specifies whether or not the cluster creator IAM principal was set as a cluster admin access entry during cluster creation time.
	//
	// The default value is `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-accessconfig.html#cfn-eks-cluster-accessconfig-bootstrapclustercreatoradminpermissions
	//
	BootstrapClusterCreatorAdminPermissions interface{} `field:"optional" json:"bootstrapClusterCreatorAdminPermissions" yaml:"bootstrapClusterCreatorAdminPermissions"`
}

