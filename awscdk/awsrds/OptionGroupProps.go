package awsrds


// Construction properties for an OptionGroup.
//
// Example:
//   var vpc Vpc
//   var securityGroup SecurityGroup
//
//
//   rds.NewOptionGroup(this, jsii.String("Options"), &OptionGroupProps{
//   	Engine: rds.DatabaseInstanceEngine_OracleSe2(&OracleSe2InstanceEngineProps{
//   		Version: rds.OracleEngineVersion_VER_19(),
//   	}),
//   	Configurations: []OptionConfiguration{
//   		&OptionConfiguration{
//   			Name: jsii.String("OEM"),
//   			Port: jsii.Number(5500),
//   			Vpc: *Vpc,
//   			SecurityGroups: []ISecurityGroup{
//   				securityGroup,
//   			},
//   		},
//   	},
//   	OptionGroupName: jsii.String("MyOptionGroup"),
//   })
//
type OptionGroupProps struct {
	// The configurations for this option group.
	Configurations *[]*OptionConfiguration `field:"required" json:"configurations" yaml:"configurations"`
	// The database engine that this option group is associated with.
	Engine IInstanceEngine `field:"required" json:"engine" yaml:"engine"`
	// A description of the option group.
	// Default: a CDK generated description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the option group.
	// Default: - a CDK generated name.
	//
	OptionGroupName *string `field:"optional" json:"optionGroupName" yaml:"optionGroupName"`
}

