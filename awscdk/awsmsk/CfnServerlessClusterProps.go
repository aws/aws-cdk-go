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
type CfnServerlessClusterProps struct {
	// Specifies client authentication information for the serverless cluster.
	ClientAuthentication interface{} `field:"required" json:"clientAuthentication" yaml:"clientAuthentication"`
	// The name of the serverless cluster.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// VPC configuration information.
	VpcConfigs interface{} `field:"required" json:"vpcConfigs" yaml:"vpcConfigs"`
	// A map of key:value pairs to apply to this serverless cluster.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

