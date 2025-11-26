package previewawssesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMailManagerRelayPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var noAuthentication interface{}
//
//   cfnMailManagerRelayMixinProps := &CfnMailManagerRelayMixinProps{
//   	Authentication: &RelayAuthenticationProperty{
//   		NoAuthentication: noAuthentication,
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	RelayName: jsii.String("relayName"),
//   	ServerName: jsii.String("serverName"),
//   	ServerPort: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerrelay.html
//
type CfnMailManagerRelayMixinProps struct {
	// Authentication for the relay destination server—specify the secretARN where the SMTP credentials are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerrelay.html#cfn-ses-mailmanagerrelay-authentication
	//
	Authentication interface{} `field:"optional" json:"authentication" yaml:"authentication"`
	// The unique relay name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerrelay.html#cfn-ses-mailmanagerrelay-relayname
	//
	RelayName *string `field:"optional" json:"relayName" yaml:"relayName"`
	// The destination relay server address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerrelay.html#cfn-ses-mailmanagerrelay-servername
	//
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The destination relay server port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerrelay.html#cfn-ses-mailmanagerrelay-serverport
	//
	ServerPort *float64 `field:"optional" json:"serverPort" yaml:"serverPort"`
	// The tags used to organize, track, or control access for the resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerrelay.html#cfn-ses-mailmanagerrelay-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

