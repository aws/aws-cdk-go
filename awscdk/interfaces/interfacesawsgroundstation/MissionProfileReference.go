package interfacesawsgroundstation


// A reference to a MissionProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   missionProfileReference := &MissionProfileReference{
//   	MissionProfileArn: jsii.String("missionProfileArn"),
//   	MissionProfileId: jsii.String("missionProfileId"),
//   }
//
type MissionProfileReference struct {
	// The Arn of the MissionProfile resource.
	MissionProfileArn *string `field:"required" json:"missionProfileArn" yaml:"missionProfileArn"`
	// The Id of the MissionProfile resource.
	MissionProfileId *string `field:"required" json:"missionProfileId" yaml:"missionProfileId"`
}

