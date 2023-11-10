package awseks


// An object representing the VPC configuration to use for an Amazon EKS cluster.
//
// > When updating a resource, you must include these properties if the previous CloudFormation template of the resource had them:
// >
// > - `EndpointPublicAccess`
// > - `EndpointPrivateAccess`
// > - `PublicAccessCidrs`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcesVpcConfigProperty := &ResourcesVpcConfigProperty{
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	EndpointPrivateAccess: jsii.Boolean(false),
//   	EndpointPublicAccess: jsii.Boolean(false),
//   	PublicAccessCidrs: []*string{
//   		jsii.String("publicAccessCidrs"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html
//
type CfnCluster_ResourcesVpcConfigProperty struct {
	// Specify subnets for your Amazon EKS nodes.
	//
	// Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.
	//
	// > All subnets that you add must be in the same set of AZs as originally provided when you created the cluster. New subnets must satisfy all of the other requirements, for example they must have sufficient IP addresses.
	// >
	// > For example, assume that you made a cluster and specified four subnets. In the order that you specified them, the first subnet is in the `us-west-2a` Availability Zone, the second and third subnets are in `us-west-2b` Availability Zone, and the fourth subnet is in `us-west-2c` Availability Zone. If you want to change the subnets, you must provide at least one subnet in each of the three Availability Zones, and the subnets must be in the same VPC as the original subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html#cfn-eks-cluster-resourcesvpcconfig-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Set this value to `true` to enable private access for your cluster's Kubernetes API server endpoint.
	//
	// If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is `false` , which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that `publicAccessCidrs` includes the necessary CIDR blocks for communication with the nodes or Fargate pods. For more information, see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the **Amazon EKS User Guide** .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html#cfn-eks-cluster-resourcesvpcconfig-endpointprivateaccess
	//
	EndpointPrivateAccess interface{} `field:"optional" json:"endpointPrivateAccess" yaml:"endpointPrivateAccess"`
	// Set this value to `false` to disable public access to your cluster's Kubernetes API server endpoint.
	//
	// If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is `true` , which enables public access for your Kubernetes API server. For more information, see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the **Amazon EKS User Guide** .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html#cfn-eks-cluster-resourcesvpcconfig-endpointpublicaccess
	//
	EndpointPublicAccess interface{} `field:"optional" json:"endpointPublicAccess" yaml:"endpointPublicAccess"`
	// The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint.
	//
	// Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is `0.0.0.0/0` . If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks. For more information, see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the **Amazon EKS User Guide** .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html#cfn-eks-cluster-resourcesvpcconfig-publicaccesscidrs
	//
	PublicAccessCidrs *[]*string `field:"optional" json:"publicAccessCidrs" yaml:"publicAccessCidrs"`
	// Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use that allow communication between your nodes and the Kubernetes control plane.
	//
	// If you don't specify any security groups, then familiarize yourself with the difference between Amazon EKS defaults for clusters deployed with Kubernetes. For more information, see [Amazon EKS security group considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the **Amazon EKS User Guide** .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html#cfn-eks-cluster-resourcesvpcconfig-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

