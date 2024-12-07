package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baselinePerformanceFactorsProperty := &BaselinePerformanceFactorsProperty{
//   	Cpu: &CpuProperty{
//   		References: []interface{}{
//   			&ReferenceProperty{
//   				InstanceFamily: jsii.String("instanceFamily"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-baselineperformancefactors.html
//
type CfnLaunchTemplate_BaselinePerformanceFactorsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-baselineperformancefactors.html#cfn-ec2-launchtemplate-baselineperformancefactors-cpu
	//
	Cpu interface{} `field:"optional" json:"cpu" yaml:"cpu"`
}

