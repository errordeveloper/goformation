package cloudformation

// AWSEMRCluster_JobFlowInstancesConfig AWS CloudFormation Resource (AWS::EMR::Cluster.JobFlowInstancesConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html
type AWSEMRCluster_JobFlowInstancesConfig struct {

	// AdditionalMasterSecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-additionalmastersecuritygroups
	AdditionalMasterSecurityGroups []*stringIntrinsic `json:"AdditionalMasterSecurityGroups,omitempty"`

	// AdditionalSlaveSecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-additionalslavesecuritygroups
	AdditionalSlaveSecurityGroups []*stringIntrinsic `json:"AdditionalSlaveSecurityGroups,omitempty"`

	// CoreInstanceFleet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-coreinstancefleet
	CoreInstanceFleet *AWSEMRCluster_InstanceFleetConfig `json:"CoreInstanceFleet,omitempty"`

	// CoreInstanceGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-coreinstancegroup
	CoreInstanceGroup *AWSEMRCluster_InstanceGroupConfig `json:"CoreInstanceGroup,omitempty"`

	// Ec2KeyName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-ec2keyname
	Ec2KeyName *stringIntrinsic `json:"Ec2KeyName,omitempty"`

	// Ec2SubnetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-ec2subnetid
	Ec2SubnetId *stringIntrinsic `json:"Ec2SubnetId,omitempty"`

	// EmrManagedMasterSecurityGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-emrmanagedmastersecuritygroup
	EmrManagedMasterSecurityGroup *stringIntrinsic `json:"EmrManagedMasterSecurityGroup,omitempty"`

	// EmrManagedSlaveSecurityGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-emrmanagedslavesecuritygroup
	EmrManagedSlaveSecurityGroup *stringIntrinsic `json:"EmrManagedSlaveSecurityGroup,omitempty"`

	// HadoopVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-hadoopversion
	HadoopVersion *stringIntrinsic `json:"HadoopVersion,omitempty"`

	// MasterInstanceFleet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-masterinstancefleet
	MasterInstanceFleet *AWSEMRCluster_InstanceFleetConfig `json:"MasterInstanceFleet,omitempty"`

	// MasterInstanceGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-masterinstancegroup
	MasterInstanceGroup *AWSEMRCluster_InstanceGroupConfig `json:"MasterInstanceGroup,omitempty"`

	// Placement AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-placement
	Placement *AWSEMRCluster_PlacementType `json:"Placement,omitempty"`

	// ServiceAccessSecurityGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-serviceaccesssecuritygroup
	ServiceAccessSecurityGroup *stringIntrinsic `json:"ServiceAccessSecurityGroup,omitempty"`

	// TerminationProtected AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-terminationprotected
	TerminationProtected bool `json:"TerminationProtected,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_JobFlowInstancesConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.JobFlowInstancesConfig"
}
