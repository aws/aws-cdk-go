package previewawsec2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnNetworkInsightsAccessScopeAnalysisPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNetworkInsightsAccessScopeAnalysisMixinProps := &CfnNetworkInsightsAccessScopeAnalysisMixinProps{
//   	NetworkInsightsAccessScopeId: jsii.String("networkInsightsAccessScopeId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightsaccessscopeanalysis.html
//
type CfnNetworkInsightsAccessScopeAnalysisMixinProps struct {
	// The ID of the Network Access Scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightsaccessscopeanalysis.html#cfn-ec2-networkinsightsaccessscopeanalysis-networkinsightsaccessscopeid
	//
	NetworkInsightsAccessScopeId *string `field:"optional" json:"networkInsightsAccessScopeId" yaml:"networkInsightsAccessScopeId"`
	// The tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightsaccessscopeanalysis.html#cfn-ec2-networkinsightsaccessscopeanalysis-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

