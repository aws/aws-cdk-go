package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnUseCase`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUseCaseProps := &CfnUseCaseProps{
//   	InstanceId: jsii.String("instanceId"),
//   	IntegrationAssociationId: jsii.String("integrationAssociationId"),
//   	UseCaseType: jsii.String("useCaseType"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-usecase.html
//
type CfnUseCaseProps struct {
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-usecase.html#cfn-connect-usecase-instanceid
	//
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The identifier for the integration association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-usecase.html#cfn-connect-usecase-integrationassociationid
	//
	IntegrationAssociationId *string `field:"required" json:"integrationAssociationId" yaml:"integrationAssociationId"`
	// The type of use case to associate to the integration association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-usecase.html#cfn-connect-usecase-usecasetype
	//
	UseCaseType *string `field:"required" json:"useCaseType" yaml:"useCaseType"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-usecase.html#cfn-connect-usecase-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

