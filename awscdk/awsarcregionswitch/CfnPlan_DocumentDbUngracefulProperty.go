package awsarcregionswitch


// Configuration for handling failures when performing operations on DocumentDB global clusters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentDbUngracefulProperty := &DocumentDbUngracefulProperty{
//   	Ungraceful: jsii.String("ungraceful"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-documentdbungraceful.html
//
type CfnPlan_DocumentDbUngracefulProperty struct {
	// The settings for ungraceful execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-documentdbungraceful.html#cfn-arcregionswitch-plan-documentdbungraceful-ungraceful
	//
	Ungraceful *string `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

