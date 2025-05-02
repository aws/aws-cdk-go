package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMailManagerIngressPoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMailManagerIngressPointProps := &CfnMailManagerIngressPointProps{
//   	RuleSetId: jsii.String("ruleSetId"),
//   	TrafficPolicyId: jsii.String("trafficPolicyId"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	IngressPointConfiguration: &IngressPointConfigurationProperty{
//   		SecretArn: jsii.String("secretArn"),
//   		SmtpPassword: jsii.String("smtpPassword"),
//   	},
//   	IngressPointName: jsii.String("ingressPointName"),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		PrivateNetworkConfiguration: &PrivateNetworkConfigurationProperty{
//   			VpcEndpointId: jsii.String("vpcEndpointId"),
//   		},
//   		PublicNetworkConfiguration: &PublicNetworkConfigurationProperty{
//   			IpType: jsii.String("ipType"),
//   		},
//   	},
//   	StatusToUpdate: jsii.String("statusToUpdate"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageringresspoint.html
//
type CfnMailManagerIngressPointProps struct {
	// The identifier of an existing rule set that you attach to an ingress endpoint resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageringresspoint.html#cfn-ses-mailmanageringresspoint-rulesetid
	//
	RuleSetId *string `field:"required" json:"ruleSetId" yaml:"ruleSetId"`
	// The identifier of an existing traffic policy that you attach to an ingress endpoint resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageringresspoint.html#cfn-ses-mailmanageringresspoint-trafficpolicyid
	//
	TrafficPolicyId *string `field:"required" json:"trafficPolicyId" yaml:"trafficPolicyId"`
	// The type of the ingress endpoint to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageringresspoint.html#cfn-ses-mailmanageringresspoint-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration of the ingress endpoint resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageringresspoint.html#cfn-ses-mailmanageringresspoint-ingresspointconfiguration
	//
	IngressPointConfiguration interface{} `field:"optional" json:"ingressPointConfiguration" yaml:"ingressPointConfiguration"`
	// A user friendly name for an ingress endpoint resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageringresspoint.html#cfn-ses-mailmanageringresspoint-ingresspointname
	//
	IngressPointName *string `field:"optional" json:"ingressPointName" yaml:"ingressPointName"`
	// The network type (IPv4-only, Dual-Stack, PrivateLink) of the ingress endpoint resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageringresspoint.html#cfn-ses-mailmanageringresspoint-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The update status of an ingress endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageringresspoint.html#cfn-ses-mailmanageringresspoint-statustoupdate
	//
	StatusToUpdate *string `field:"optional" json:"statusToUpdate" yaml:"statusToUpdate"`
	// The tags used to organize, track, or control access for the resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageringresspoint.html#cfn-ses-mailmanageringresspoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

