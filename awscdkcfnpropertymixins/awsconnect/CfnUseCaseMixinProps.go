package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnUseCasePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnUseCaseMixinProps := &CfnUseCaseMixinProps{
//   	InstanceId: jsii.String("instanceId"),
//   	IntegrationAssociationId: jsii.String("integrationAssociationId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseCaseType: jsii.String("useCaseType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-usecase.html
//
type CfnUseCaseMixinProps struct {
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-usecase.html#cfn-connect-usecase-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The identifier for the integration association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-usecase.html#cfn-connect-usecase-integrationassociationid
	//
	IntegrationAssociationId *string `field:"optional" json:"integrationAssociationId" yaml:"integrationAssociationId"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-usecase.html#cfn-connect-usecase-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of use case to associate to the integration association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-usecase.html#cfn-connect-usecase-usecasetype
	//
	UseCaseType *string `field:"optional" json:"useCaseType" yaml:"useCaseType"`
}

