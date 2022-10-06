package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFargateProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFargateProfileProps := &cfnFargateProfileProps{
//   	clusterName: jsii.String("clusterName"),
//   	podExecutionRoleArn: jsii.String("podExecutionRoleArn"),
//   	selectors: []interface{}{
//   		&selectorProperty{
//   			namespace: jsii.String("namespace"),
//
//   			// the properties below are optional
//   			labels: []interface{}{
//   				&labelProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	fargateProfileName: jsii.String("fargateProfileName"),
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFargateProfileProps struct {
	// The name of the Amazon EKS cluster to apply the Fargate profile to.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The Amazon Resource Name (ARN) of the pod execution role to use for pods that match the selectors in the Fargate profile.
	//
	// The pod execution role allows Fargate infrastructure to register with your cluster as a node, and it provides read access to Amazon ECR image repositories. For more information, see [Pod Execution Role](https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html) in the *Amazon EKS User Guide* .
	PodExecutionRoleArn *string `field:"required" json:"podExecutionRoleArn" yaml:"podExecutionRoleArn"`
	// The selectors to match for pods to use this Fargate profile.
	//
	// Each selector must have an associated namespace. Optionally, you can also specify labels for a namespace. You may specify up to five selectors in a Fargate profile.
	Selectors interface{} `field:"required" json:"selectors" yaml:"selectors"`
	// The name of the Fargate profile.
	FargateProfileName *string `field:"optional" json:"fargateProfileName" yaml:"fargateProfileName"`
	// The IDs of subnets to launch your pods into.
	//
	// At this time, pods running on Fargate are not assigned public IP addresses, so only private subnets (with no direct route to an Internet Gateway) are accepted for this parameter.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// The metadata to apply to the Fargate profile to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Fargate profile tags do not propagate to any other resources associated with the Fargate profile, such as the pods that are scheduled with it.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

