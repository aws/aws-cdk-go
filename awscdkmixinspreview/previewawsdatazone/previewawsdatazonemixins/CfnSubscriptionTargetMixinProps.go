package previewawsdatazonemixins


// Properties for CfnSubscriptionTargetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSubscriptionTargetMixinProps := &CfnSubscriptionTargetMixinProps{
//   	ApplicableAssetTypes: []*string{
//   		jsii.String("applicableAssetTypes"),
//   	},
//   	AuthorizedPrincipals: []*string{
//   		jsii.String("authorizedPrincipals"),
//   	},
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	ManageAccessRole: jsii.String("manageAccessRole"),
//   	Name: jsii.String("name"),
//   	Provider: jsii.String("provider"),
//   	SubscriptionTargetConfig: []interface{}{
//   		&SubscriptionTargetFormProperty{
//   			Content: jsii.String("content"),
//   			FormName: jsii.String("formName"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html
//
type CfnSubscriptionTargetMixinProps struct {
	// The asset types included in the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-applicableassettypes
	//
	ApplicableAssetTypes *[]*string `field:"optional" json:"applicableAssetTypes" yaml:"applicableAssetTypes"`
	// The authorized principals included in the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-authorizedprincipals
	//
	AuthorizedPrincipals *[]*string `field:"optional" json:"authorizedPrincipals" yaml:"authorizedPrincipals"`
	// The ID of the Amazon DataZone domain in which subscription target is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-domainidentifier
	//
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The ID of the environment in which subscription target is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-environmentidentifier
	//
	EnvironmentIdentifier *string `field:"optional" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The manage access role that is used to create the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-manageaccessrole
	//
	ManageAccessRole *string `field:"optional" json:"manageAccessRole" yaml:"manageAccessRole"`
	// The name of the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The provider of the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-provider
	//
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// The configuration of the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-subscriptiontargetconfig
	//
	SubscriptionTargetConfig interface{} `field:"optional" json:"subscriptionTargetConfig" yaml:"subscriptionTargetConfig"`
	// The type of the subscription target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-subscriptiontarget.html#cfn-datazone-subscriptiontarget-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

