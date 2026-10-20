package awsscn

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnInstancePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnInstanceMixinProps := &CfnInstanceMixinProps{
//   	InstanceDescription: jsii.String("instanceDescription"),
//   	InstanceName: jsii.String("instanceName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WebAppDnsDomain: jsii.String("webAppDnsDomain"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-instance.html
//
type CfnInstanceMixinProps struct {
	// The instance description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-instance.html#cfn-scn-instance-instancedescription
	//
	InstanceDescription *string `field:"optional" json:"instanceDescription" yaml:"instanceDescription"`
	// The instance name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-instance.html#cfn-scn-instance-instancename
	//
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// Tags for the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-instance.html#cfn-scn-instance-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The WebApp DNS domain name of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-instance.html#cfn-scn-instance-webappdnsdomain
	//
	WebAppDnsDomain *string `field:"optional" json:"webAppDnsDomain" yaml:"webAppDnsDomain"`
}

