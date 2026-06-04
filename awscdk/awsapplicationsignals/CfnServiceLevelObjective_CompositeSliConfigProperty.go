package awsapplicationsignals


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   compositeSliConfigProperty := &CompositeSliConfigProperty{
//   	SelectionConfig: &SelectionConfigProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Pattern: jsii.String("pattern"),
//   	},
//
//   	// the properties below are optional
//   	CompositeSliComponents: []interface{}{
//   		&CompositeSliComponentProperty{
//   			OperationName: jsii.String("operationName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-compositesliconfig.html
//
type CfnServiceLevelObjective_CompositeSliConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-compositesliconfig.html#cfn-applicationsignals-servicelevelobjective-compositesliconfig-selectionconfig
	//
	SelectionConfig interface{} `field:"required" json:"selectionConfig" yaml:"selectionConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-compositesliconfig.html#cfn-applicationsignals-servicelevelobjective-compositesliconfig-compositeslicomponents
	//
	CompositeSliComponents interface{} `field:"optional" json:"compositeSliComponents" yaml:"compositeSliComponents"`
}

