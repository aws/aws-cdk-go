package awsemrcontainers


// EKS information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksInfoProperty := &EksInfoProperty{
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-eksinfo.html
//
type CfnSecurityConfiguration_EksInfoProperty struct {
	// The EKS namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-eksinfo.html#cfn-emrcontainers-securityconfiguration-eksinfo-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

