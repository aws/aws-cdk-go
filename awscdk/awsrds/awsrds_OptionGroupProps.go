package awsrds


// Construction properties for an OptionGroup.
//
// Example:
//   // Set open cursors with parameter group
//   parameterGroup := rds.NewParameterGroup(this, jsii.String("ParameterGroup"), &parameterGroupProps{
//   	engine: rds.databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
//   		version: rds.oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
//   	}),
//   	parameters: map[string]*string{
//   		"open_cursors": jsii.String("2500"),
//   	},
//   })
//
//   optionGroup := rds.NewOptionGroup(this, jsii.String("OptionGroup"), &optionGroupProps{
//   	engine: rds.*databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
//   		version: rds.*oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
//   	}),
//   	configurations: []optionConfiguration{
//   		&optionConfiguration{
//   			name: jsii.String("LOCATOR"),
//   		},
//   		&optionConfiguration{
//   			name: jsii.String("OEM"),
//   			port: jsii.Number(1158),
//   			vpc: vpc,
//   		},
//   	},
//   })
//
//   // Allow connections to OEM
//   optionGroup.optionConnections.oEM.connections.allowDefaultPortFromAnyIpv4()
//
//   // Database instance with production values
//   instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
//   	engine: rds.*databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
//   		version: rds.*oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
//   	}),
//   	licenseModel: rds.licenseModel_BRING_YOUR_OWN_LICENSE,
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE3, ec2.instanceSize_MEDIUM),
//   	multiAz: jsii.Boolean(true),
//   	storageType: rds.storageType_IO1,
//   	credentials: rds.credentials.fromUsername(jsii.String("syscdk")),
//   	vpc: vpc,
//   	databaseName: jsii.String("ORCL"),
//   	storageEncrypted: jsii.Boolean(true),
//   	backupRetention: cdk.duration.days(jsii.Number(7)),
//   	monitoringInterval: cdk.*duration.seconds(jsii.Number(60)),
//   	enablePerformanceInsights: jsii.Boolean(true),
//   	cloudwatchLogsExports: []*string{
//   		jsii.String("trace"),
//   		jsii.String("audit"),
//   		jsii.String("alert"),
//   		jsii.String("listener"),
//   	},
//   	cloudwatchLogsRetention: logs.retentionDays_ONE_MONTH,
//   	autoMinorVersionUpgrade: jsii.Boolean(true),
//   	 // required to be true if LOCATOR is used in the option group
//   	optionGroup: optionGroup,
//   	parameterGroup: parameterGroup,
//   	removalPolicy: awscdk.RemovalPolicy_DESTROY,
//   })
//
//   // Allow connections on default port from any IPV4
//   instance.connections.allowDefaultPortFromAnyIpv4()
//
//   // Rotate the master user password every 30 days
//   instance.addRotationSingleUser()
//
//   // Add alarm for high CPU
//   // Add alarm for high CPU
//   cloudwatch.NewAlarm(this, jsii.String("HighCPU"), &alarmProps{
//   	metric: instance.metricCPUUtilization(),
//   	threshold: jsii.Number(90),
//   	evaluationPeriods: jsii.Number(1),
//   })
//
//   // Trigger Lambda function on instance availability events
//   fn := lambda.NewFunction(this, jsii.String("Function"), &functionProps{
//   	code: lambda.code.fromInline(jsii.String("exports.handler = (event) => console.log(event);")),
//   	handler: jsii.String("index.handler"),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   })
//
//   availabilityRule := instance.onEvent(jsii.String("Availability"), &onEventOptions{
//   	target: targets.NewLambdaFunction(fn),
//   })
//   availabilityRule.addEventPattern(&eventPattern{
//   	detail: map[string]interface{}{
//   		"EventCategories": []interface{}{
//   			jsii.String("availability"),
//   		},
//   	},
//   })
//
type OptionGroupProps struct {
	// The configurations for this option group.
	Configurations *[]*OptionConfiguration `field:"required" json:"configurations" yaml:"configurations"`
	// The database engine that this option group is associated with.
	Engine IInstanceEngine `field:"required" json:"engine" yaml:"engine"`
	// A description of the option group.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

