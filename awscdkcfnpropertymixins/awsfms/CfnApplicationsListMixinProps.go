package awsfms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApplicationsListPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnApplicationsListMixinProps := &CfnApplicationsListMixinProps{
//   	AppsList: []interface{}{
//   		&AppProperty{
//   			AppName: jsii.String("appName"),
//   			Port: jsii.Number(123),
//   			Protocol: jsii.String("protocol"),
//   		},
//   	},
//   	ListName: jsii.String("listName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-applicationslist.html
//
type CfnApplicationsListMixinProps struct {
	// An array of applications in the Firewall Manager applications list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-applicationslist.html#cfn-fms-applicationslist-appslist
	//
	AppsList interface{} `field:"optional" json:"appsList" yaml:"appsList"`
	// The name of the Firewall Manager applications list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-applicationslist.html#cfn-fms-applicationslist-listname
	//
	ListName *string `field:"optional" json:"listName" yaml:"listName"`
	// An array of key-value pairs to apply to the applications list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-applicationslist.html#cfn-fms-applicationslist-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

