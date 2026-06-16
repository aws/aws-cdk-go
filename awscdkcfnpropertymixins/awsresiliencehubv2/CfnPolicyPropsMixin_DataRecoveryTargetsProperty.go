package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dataRecoveryTargetsProperty := &DataRecoveryTargetsProperty{
//   	TimeBetweenBackupsInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-policy-datarecoverytargets.html
//
type CfnPolicyPropsMixin_DataRecoveryTargetsProperty struct {
	// Time between backups in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-policy-datarecoverytargets.html#cfn-resiliencehubv2-policy-datarecoverytargets-timebetweenbackupsinminutes
	//
	TimeBetweenBackupsInMinutes *float64 `field:"optional" json:"timeBetweenBackupsInMinutes" yaml:"timeBetweenBackupsInMinutes"`
}

