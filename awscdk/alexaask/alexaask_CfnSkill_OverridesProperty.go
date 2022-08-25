package alexaask


// The `Overrides` property type provides overrides to the skill package to apply when creating or updating the skill.
//
// Values provided here do not modify the contents of the original skill package. Currently, only overriding values inside of the skill manifest component of the package is supported.
//
// `Overrides` is a property of the `Alexa::ASK::Skill SkillPackage` property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var manifest interface{}
//
//   overridesProperty := &overridesProperty{
//   	manifest: manifest,
//   }
//
type CfnSkill_OverridesProperty struct {
	// Overrides to apply to the skill manifest inside of the skill package.
	//
	// The skill manifest contains metadata about the skill. For more information, see  .
	Manifest interface{} `field:"optional" json:"manifest" yaml:"manifest"`
}

