package mixinsawscontroltower

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLandingZonePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var manifest interface{}
//
//   cfnLandingZoneMixinProps := &CfnLandingZoneMixinProps{
//   	Manifest: manifest,
//   	RemediationTypes: []*string{
//   		jsii.String("remediationTypes"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-landingzone.html
//
type CfnLandingZoneMixinProps struct {
	// The landing zone manifest JSON text file that specifies the landing zone configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-landingzone.html#cfn-controltower-landingzone-manifest
	//
	Manifest interface{} `field:"optional" json:"manifest" yaml:"manifest"`
	// The types of remediation actions configured for the landing zone, such as automatic drift correction or compliance enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-landingzone.html#cfn-controltower-landingzone-remediationtypes
	//
	RemediationTypes *[]*string `field:"optional" json:"remediationTypes" yaml:"remediationTypes"`
	// Tags to be applied to the landing zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-landingzone.html#cfn-controltower-landingzone-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The landing zone's current deployed version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-landingzone.html#cfn-controltower-landingzone-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

