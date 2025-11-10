package awsrtbfabric


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnResponderGateway_EksEndpointsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html#cfn-rtbfabric-respondergateway-eksendpointsconfiguration-clusterapiservercacertificatechain
	//
	ClusterApiServerCaCertificateChain *string `field:"required" json:"clusterApiServerCaCertificateChain" yaml:"clusterApiServerCaCertificateChain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html#cfn-rtbfabric-respondergateway-eksendpointsconfiguration-clusterapiserverendpointuri
	//
	ClusterApiServerEndpointUri *string `field:"required" json:"clusterApiServerEndpointUri" yaml:"clusterApiServerEndpointUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html#cfn-rtbfabric-respondergateway-eksendpointsconfiguration-clustername
	//
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html#cfn-rtbfabric-respondergateway-eksendpointsconfiguration-endpointsresourcename
	//
	EndpointsResourceName *string `field:"required" json:"endpointsResourceName" yaml:"endpointsResourceName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html#cfn-rtbfabric-respondergateway-eksendpointsconfiguration-endpointsresourcenamespace
	//
	EndpointsResourceNamespace *string `field:"required" json:"endpointsResourceNamespace" yaml:"endpointsResourceNamespace"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html#cfn-rtbfabric-respondergateway-eksendpointsconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

