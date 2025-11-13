package interfacesalexaask


// A reference to a Skill resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   skillReference := &SkillReference{
//   	SkillId: jsii.String("skillId"),
//   }
//
type SkillReference struct {
	// The Id of the Skill resource.
	SkillId *string `field:"required" json:"skillId" yaml:"skillId"`
}

