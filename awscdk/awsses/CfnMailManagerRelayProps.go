package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMailManagerRelay`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var noAuthentication interface{}
//
//   cfnMailManagerRelayProps := &CfnMailManagerRelayProps{
//   	Authentication: &RelayAuthenticationProperty{
//   		NoAuthentication: noAuthentication,
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	ServerName: jsii.String("serverName"),
//   	ServerPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	RelayName: jsii.String("relayName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerrelay.html
//
type CfnMailManagerRelayProps struct {
	// Authentication for the relay destination server—specify the secretARN where the SMTP credentials are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerrelay.html#cfn-ses-mailmanagerrelay-authentication
	//
	Authentication interface{} `field:"required" json:"authentication" yaml:"authentication"`
	// The destination relay server address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerrelay.html#cfn-ses-mailmanagerrelay-servername
	//
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
	// The destination relay server port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerrelay.html#cfn-ses-mailmanagerrelay-serverport
	//
	ServerPort *float64 `field:"required" json:"serverPort" yaml:"serverPort"`
	// The unique relay name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerrelay.html#cfn-ses-mailmanagerrelay-relayname
	//
	RelayName *string `field:"optional" json:"relayName" yaml:"relayName"`
	// The tags used to organize, track, or control access for the resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerrelay.html#cfn-ses-mailmanagerrelay-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

