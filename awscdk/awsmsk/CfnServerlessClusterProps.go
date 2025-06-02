package awsmsk


// Properties for defining a `CfnServerlessCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServerlessClusterProps := &CfnServerlessClusterProps{
//   	ClientAuthentication: &ClientAuthenticationProperty{
//   		Sasl: &SaslProperty{
//   			Iam: &IamProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	VpcConfigs: []interface{}{
//   		&VpcConfigProperty{
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//
//   			// the properties below are optional
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-serverlesscluster.html
//
type CfnServerlessClusterProps struct {
	// Includes all client authentication related information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-serverlesscluster.html#cfn-msk-serverlesscluster-clientauthentication
	//
	ClientAuthentication interface{} `field:"required" json:"clientAuthentication" yaml:"clientAuthentication"`
	// The name of the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-serverlesscluster.html#cfn-msk-serverlesscluster-clustername
	//
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// VPC configuration information for the serverless cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-serverlesscluster.html#cfn-msk-serverlesscluster-vpcconfigs
	//
	VpcConfigs interface{} `field:"required" json:"vpcConfigs" yaml:"vpcConfigs"`
	// An arbitrary set of tags (key-value pairs) for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-serverlesscluster.html#cfn-msk-serverlesscluster-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

