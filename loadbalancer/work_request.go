// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// WorkRequest. Many of the API requests you use to create and configure load balancing do not take effect immediately.
// In these cases, the request spawns an asynchronous work flow to fulfill the request. WorkRequest objects provide visibility
// for in-progress work flows.
// For more information about work requests, see [Viewing the State of a Work Request]({{DOC_SERVER_URL}}/Content/Balance/Tasks/viewingworkrequest.htm).
type WorkRequest struct {
	ErrorDetails *[]WorkRequestError `mandatory:"true" json:"errorDetails,omitempty"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the work request.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the work request.
	LifecycleState WorkRequestLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the load balancer with which the work request
	// is associated.
	LoadBalancerID *string `mandatory:"true" json:"loadBalancerId,omitempty"`

	// A collection of data, related to the load balancer provisioning process, that helps with debugging in the event of failure.
	// Possible data elements include:
	// - workflow name
	// - event ID
	// - work request ID
	// - load balancer ID
	// - workflow completion message
	Message *string `mandatory:"true" json:"message,omitempty"`

	// The date and time the work request was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted,omitempty"`

	// The type of action the work request represents.
	Type_ *string `mandatory:"true" json:"type,omitempty"`

	// The date and time the work request was completed, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished,omitempty"`
}

func (model WorkRequest) String() string {
	return common.PointerString(model)
}

type WorkRequestLifecycleStateEnum string

const (
	WORK_REQUEST_LIFECYCLE_STATE_ACCEPTED    WorkRequestLifecycleStateEnum = "ACCEPTED"
	WORK_REQUEST_LIFECYCLE_STATE_IN_PROGRESS WorkRequestLifecycleStateEnum = "IN_PROGRESS"
	WORK_REQUEST_LIFECYCLE_STATE_FAILED      WorkRequestLifecycleStateEnum = "FAILED"
	WORK_REQUEST_LIFECYCLE_STATE_SUCCEEDED   WorkRequestLifecycleStateEnum = "SUCCEEDED"
	WORK_REQUEST_LIFECYCLE_STATE_UNKNOWN     WorkRequestLifecycleStateEnum = "UNKNOWN"
)

var mapping_workrequest_lifecycleState = map[string]WorkRequestLifecycleStateEnum{
	"ACCEPTED":    WORK_REQUEST_LIFECYCLE_STATE_ACCEPTED,
	"IN_PROGRESS": WORK_REQUEST_LIFECYCLE_STATE_IN_PROGRESS,
	"FAILED":      WORK_REQUEST_LIFECYCLE_STATE_FAILED,
	"SUCCEEDED":   WORK_REQUEST_LIFECYCLE_STATE_SUCCEEDED,
	"UNKNOWN":     WORK_REQUEST_LIFECYCLE_STATE_UNKNOWN,
}

func GetWorkRequestLifecycleStateEnumValues() []WorkRequestLifecycleStateEnum {
	values := make([]WorkRequestLifecycleStateEnum, 0)
	for _, v := range mapping_workrequest_lifecycleState {
		if v != WORK_REQUEST_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}