package previewawsrtbfabricmixins


// Describes the configuration of an Amazon Elastic Kubernetes Service endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eksEndpointsConfigurationProperty := &EksEndpointsConfigurationProperty{
//   	ClusterApiServerCaCertificateChain: jsii.String("clusterApiServerCaCertificateChain"),
//   	ClusterApiServerEndpointUri: jsii.String("clusterApiServerEndpointUri"),
//   	ClusterName: jsii.String("clusterName"),
//   	EndpointsResourceName: jsii.String("endpointsResourceName"),
//   	EndpointsResourceNamespace: jsii.String("endpointsResourceNamespace"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html
//
type CfnResponderGatewayPropsMixin_EksEndpointsConfigurationProperty struct {
	// The CA certificate chain of the cluster API server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html#cfn-rtbfabric-respondergateway-eksendpointsconfiguration-clusterapiservercacertificatechain
	//
	ClusterApiServerCaCertificateChain *string `field:"optional" json:"clusterApiServerCaCertificateChain" yaml:"clusterApiServerCaCertificateChain"`
	// The URI of the cluster API server endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html#cfn-rtbfabric-respondergateway-eksendpointsconfiguration-clusterapiserverendpointuri
	//
	ClusterApiServerEndpointUri *string `field:"optional" json:"clusterApiServerEndpointUri" yaml:"clusterApiServerEndpointUri"`
	// The name of the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html#cfn-rtbfabric-respondergateway-eksendpointsconfiguration-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The name of the endpoint resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html#cfn-rtbfabric-respondergateway-eksendpointsconfiguration-endpointsresourcename
	//
	EndpointsResourceName *string `field:"optional" json:"endpointsResourceName" yaml:"endpointsResourceName"`
	// The namespace of the endpoint resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html#cfn-rtbfabric-respondergateway-eksendpointsconfiguration-endpointsresourcenamespace
	//
	EndpointsResourceNamespace *string `field:"optional" json:"endpointsResourceNamespace" yaml:"endpointsResourceNamespace"`
	// The role ARN for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html#cfn-rtbfabric-respondergateway-eksendpointsconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

