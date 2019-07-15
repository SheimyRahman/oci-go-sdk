// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts.
//

package budget

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Budget A budget.
type Budget struct {

	// The OCID of the budget
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the compartment on which budget is applied
	TargetCompartmentId *string `mandatory:"true" json:"targetCompartmentId"`

	// The display name of the budget.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The amount of the budget expressed in the currency of the customer's rate card.
	Amount *float32 `mandatory:"true" json:"amount"`

	// The reset period for the budget.
	ResetPeriod ResetPeriodEnum `mandatory:"true" json:"resetPeriod"`

	// The current state of the budget.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Total number of alert rules in the budget
	AlertRuleCount *int `mandatory:"true" json:"alertRuleCount"`

	// Time that budget was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time that budget was updated
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The description of the budget.
	Description *string `mandatory:"false" json:"description"`

	// Version of the budget. Starts from 1 and increments by 1.
	Version *int `mandatory:"false" json:"version"`

	// The actual spend in currency for the current budget cycle
	ActualSpend *float32 `mandatory:"false" json:"actualSpend"`

	// The forecasted spend in currency by the end of the current budget cycle
	ForecastedSpend *float32 `mandatory:"false" json:"forecastedSpend"`

	// The time that the budget spend was last computed
	TimeSpendComputed *common.SDKTime `mandatory:"false" json:"timeSpendComputed"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Budget) String() string {
	return common.PointerString(m)
}

// BudgetResetPeriodEnum is an alias to type: ResetPeriodEnum
// Consider using ResetPeriodEnum instead
// Deprecated
type BudgetResetPeriodEnum = ResetPeriodEnum

// Set of constants representing the allowable values for ResetPeriodEnum
// Deprecated
const (
	BudgetResetPeriodMonthly ResetPeriodEnum = "MONTHLY"
)

// GetBudgetResetPeriodEnumValues Enumerates the set of values for ResetPeriodEnum
// Consider using GetResetPeriodEnumValue
// Deprecated
var GetBudgetResetPeriodEnumValues = GetResetPeriodEnumValues

// BudgetLifecycleStateEnum is an alias to type: LifecycleStateEnum
// Consider using LifecycleStateEnum instead
// Deprecated
type BudgetLifecycleStateEnum = LifecycleStateEnum

// Set of constants representing the allowable values for LifecycleStateEnum
// Deprecated
const (
	BudgetLifecycleStateActive   LifecycleStateEnum = "ACTIVE"
	BudgetLifecycleStateInactive LifecycleStateEnum = "INACTIVE"
)

// GetBudgetLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
// Consider using GetLifecycleStateEnumValue
// Deprecated
var GetBudgetLifecycleStateEnumValues = GetLifecycleStateEnumValues
