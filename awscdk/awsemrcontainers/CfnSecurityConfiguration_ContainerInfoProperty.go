package awsemrcontainers


// Container information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerInfoProperty := &ContainerInfoProperty{
//   	EksInfo: &EksInfoProperty{
//   		Namespace: jsii.String("namespace"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-containerinfo.html
//
type CfnSecurityConfiguration_ContainerInfoProperty struct {
	// EKS information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-containerinfo.html#cfn-emrcontainers-securityconfiguration-containerinfo-eksinfo
	//
	EksInfo interface{} `field:"optional" json:"eksInfo" yaml:"eksInfo"`
}

