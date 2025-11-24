package mixinsawsrtbfabric


// Describes the configuration of a managed endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedEndpointConfigurationProperty := &ManagedEndpointConfigurationProperty{
//   	AutoScalingGroupsConfiguration: &AutoScalingGroupsConfigurationProperty{
//   		AutoScalingGroupNameList: []*string{
//   			jsii.String("autoScalingGroupNameList"),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	EksEndpointsConfiguration: &EksEndpointsConfigurationProperty{
//   		ClusterApiServerCaCertificateChain: jsii.String("clusterApiServerCaCertificateChain"),
//   		ClusterApiServerEndpointUri: jsii.String("clusterApiServerEndpointUri"),
//   		ClusterName: jsii.String("clusterName"),
//   		EndpointsResourceName: jsii.String("endpointsResourceName"),
//   		EndpointsResourceNamespace: jsii.String("endpointsResourceNamespace"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-managedendpointconfiguration.html
//
type CfnResponderGatewayPropsMixin_ManagedEndpointConfigurationProperty struct {
	// Describes the configuration of an auto scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-managedendpointconfiguration.html#cfn-rtbfabric-respondergateway-managedendpointconfiguration-autoscalinggroupsconfiguration
	//
	AutoScalingGroupsConfiguration interface{} `field:"optional" json:"autoScalingGroupsConfiguration" yaml:"autoScalingGroupsConfiguration"`
	// Describes the configuration of an Amazon Elastic Kubernetes Service endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-managedendpointconfiguration.html#cfn-rtbfabric-respondergateway-managedendpointconfiguration-eksendpointsconfiguration
	//
	EksEndpointsConfiguration interface{} `field:"optional" json:"eksEndpointsConfiguration" yaml:"eksEndpointsConfiguration"`
}

