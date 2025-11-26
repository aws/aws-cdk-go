package interfacesawsecr


// A reference to a PullTimeUpdateExclusion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pullTimeUpdateExclusionReference := &PullTimeUpdateExclusionReference{
//   	PrincipalArn: jsii.String("principalArn"),
//   }
//
type PullTimeUpdateExclusionReference struct {
	// The PrincipalArn of the PullTimeUpdateExclusion resource.
	PrincipalArn *string `field:"required" json:"principalArn" yaml:"principalArn"`
}

