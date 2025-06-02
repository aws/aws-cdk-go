package awsdatazone


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   designationConfigurationProperty := &DesignationConfigurationProperty{
//   	DesignationId: jsii.String("designationId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-designationconfiguration.html
//
type CfnProjectProfile_DesignationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-designationconfiguration.html#cfn-datazone-projectprofile-designationconfiguration-designationid
	//
	DesignationId *string `field:"required" json:"designationId" yaml:"designationId"`
}

