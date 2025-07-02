package awsdatazone


// Properties for defining a `CfnSubscriptionTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubscriptionTargetProps := &CfnSubscriptionTargetProps{
//   	ApplicableAssetTypes: []*string{
//   		jsii.String("applicableAssetTypes"),
//   	},
//   	AuthorizedPrincipals: []*string{
//   		jsii.String("authorizedPrincipals"),
//   	},
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	Name: jsii.String("name"),
//   	SubscriptionTargetConfig: []interface{}{
//   		&SubscriptionTargetFormProperty{
//   			Content: jsii.String("content"),
//   			FormName: jsii.String("formName"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ManageAccessRole: jsii.String("manageAccessRole"),
//   	Provider: jsii.String("provider"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html
//
type CfnSubscriptionTargetProps struct {
	// The asset types included in the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-applicableassettypes
	//
	ApplicableAssetTypes *[]*string `field:"required" json:"applicableAssetTypes" yaml:"applicableAssetTypes"`
	// The authorized principals included in the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-authorizedprincipals
	//
	AuthorizedPrincipals *[]*string `field:"required" json:"authorizedPrincipals" yaml:"authorizedPrincipals"`
	// The ID of the Amazon DataZone domain in which subscription target is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The ID of the environment in which subscription target is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-environmentidentifier
	//
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The name of the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The configuration of the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-subscriptiontargetconfig
	//
	SubscriptionTargetConfig interface{} `field:"required" json:"subscriptionTargetConfig" yaml:"subscriptionTargetConfig"`
	// The type of the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The manage access role that is used to create the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-manageaccessrole
	//
	ManageAccessRole *string `field:"optional" json:"manageAccessRole" yaml:"manageAccessRole"`
	// The provider of the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-provider
	//
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
}

