package awsmsk


// Properties for defining a `CfnServerlessCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServerlessClusterProps := &cfnServerlessClusterProps{
//   	clientAuthentication: &clientAuthenticationProperty{
//   		sasl: &saslProperty{
//   			iam: &iamProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	clusterName: jsii.String("clusterName"),
//   	vpcConfigs: []interface{}{
//   		&vpcConfigProperty{
//   			subnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//
//   			// the properties below are optional
//   			securityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnServerlessClusterProps struct {
	// `AWS::MSK::ServerlessCluster.ClientAuthentication`.
	ClientAuthentication interface{} `field:"required" json:"clientAuthentication" yaml:"clientAuthentication"`
	// `AWS::MSK::ServerlessCluster.ClusterName`.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// `AWS::MSK::ServerlessCluster.VpcConfigs`.
	VpcConfigs interface{} `field:"required" json:"vpcConfigs" yaml:"vpcConfigs"`
	// `AWS::MSK::ServerlessCluster.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

