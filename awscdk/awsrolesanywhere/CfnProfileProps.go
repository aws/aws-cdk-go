package awsrolesanywhere

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProfileProps := &CfnProfileProps{
//   	Name: jsii.String("name"),
//   	RoleArns: []*string{
//   		jsii.String("roleArns"),
//   	},
//
//   	// the properties below are optional
//   	AttributeMappings: []interface{}{
//   		&AttributeMappingProperty{
//   			CertificateField: jsii.String("certificateField"),
//   			MappingRules: []interface{}{
//   				&MappingRuleProperty{
//   					Specifier: jsii.String("specifier"),
//   				},
//   			},
//   		},
//   	},
//   	DurationSeconds: jsii.Number(123),
//   	Enabled: jsii.Boolean(false),
//   	ManagedPolicyArns: []*string{
//   		jsii.String("managedPolicyArns"),
//   	},
//   	RequireInstanceProperties: jsii.Boolean(false),
//   	SessionPolicy: jsii.String("sessionPolicy"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html
//
type CfnProfileProps struct {
	// The customer specified name of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of IAM role ARNs that can be assumed when this profile is specified in a CreateSession request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-rolearns
	//
	RoleArns *[]*string `field:"required" json:"roleArns" yaml:"roleArns"`
	// A mapping applied to the authenticating end-entity certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-attributemappings
	//
	AttributeMappings interface{} `field:"optional" json:"attributeMappings" yaml:"attributeMappings"`
	// The number of seconds vended session credentials will be valid for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-durationseconds
	//
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// The enabled status of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of managed policy ARNs.
	//
	// Managed policies identified by this list will be applied to the vended session credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-managedpolicyarns
	//
	ManagedPolicyArns *[]*string `field:"optional" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// Specifies whether instance properties are required in CreateSession requests with this profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-requireinstanceproperties
	//
	RequireInstanceProperties interface{} `field:"optional" json:"requireInstanceProperties" yaml:"requireInstanceProperties"`
	// A session policy that will applied to the trust boundary of the vended session credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-sessionpolicy
	//
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
	// A list of Tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

