package cloudformation

// AWSBudgetsBudget_Notification AWS CloudFormation Resource (AWS::Budgets::Budget.Notification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html
type AWSBudgetsBudget_Notification struct {

	// ComparisonOperator AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-comparisonoperator
	ComparisonOperator *stringIntrinsic `json:"ComparisonOperator,omitempty"`

	// NotificationType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-notificationtype
	NotificationType *stringIntrinsic `json:"NotificationType,omitempty"`

	// Threshold AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-threshold
	Threshold float64 `json:"Threshold,omitempty"`

	// ThresholdType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-thresholdtype
	ThresholdType *stringIntrinsic `json:"ThresholdType,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBudgetsBudget_Notification) AWSCloudFormationType() string {
	return "AWS::Budgets::Budget.Notification"
}
