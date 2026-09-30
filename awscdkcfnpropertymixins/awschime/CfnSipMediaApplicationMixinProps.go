package awschime

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSipMediaApplicationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSipMediaApplicationMixinProps := &CfnSipMediaApplicationMixinProps{
//   	AwsRegion: jsii.String("awsRegion"),
//   	Endpoints: []interface{}{
//   		&SipMediaApplicationEndpointProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-sipmediaapplication.html
//
type CfnSipMediaApplicationMixinProps struct {
	// The AWS Region in which the SIP media application is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-sipmediaapplication.html#cfn-chime-sipmediaapplication-awsregion
	//
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// List of endpoints (Lambda ARNs) specified for the SIP media application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-sipmediaapplication.html#cfn-chime-sipmediaapplication-endpoints
	//
	Endpoints interface{} `field:"optional" json:"endpoints" yaml:"endpoints"`
	// The name of the SIP media application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-sipmediaapplication.html#cfn-chime-sipmediaapplication-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags assigned to the SIP media application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-sipmediaapplication.html#cfn-chime-sipmediaapplication-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

