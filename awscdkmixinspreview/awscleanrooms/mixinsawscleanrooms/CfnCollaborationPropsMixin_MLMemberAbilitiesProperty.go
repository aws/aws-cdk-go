package mixinsawscleanrooms


// The ML member abilities for a collaboration member.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mLMemberAbilitiesProperty := &MLMemberAbilitiesProperty{
//   	CustomMlMemberAbilities: []*string{
//   		jsii.String("customMlMemberAbilities"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-mlmemberabilities.html
//
type CfnCollaborationPropsMixin_MLMemberAbilitiesProperty struct {
	// The custom ML member abilities for a collaboration member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-mlmemberabilities.html#cfn-cleanrooms-collaboration-mlmemberabilities-custommlmemberabilities
	//
	CustomMlMemberAbilities *[]*string `field:"optional" json:"customMlMemberAbilities" yaml:"customMlMemberAbilities"`
}

