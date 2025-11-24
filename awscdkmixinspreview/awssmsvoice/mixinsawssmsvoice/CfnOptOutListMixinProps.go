package mixinsawssmsvoice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnOptOutListPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOptOutListMixinProps := &CfnOptOutListMixinProps{
//   	OptOutListName: jsii.String("optOutListName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-optoutlist.html
//
type CfnOptOutListMixinProps struct {
	// The name of the OptOutList.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-optoutlist.html#cfn-smsvoice-optoutlist-optoutlistname
	//
	OptOutListName *string `field:"optional" json:"optOutListName" yaml:"optOutListName"`
	// An array of tags (key and value pairs) to associate with the new OptOutList.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-optoutlist.html#cfn-smsvoice-optoutlist-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

