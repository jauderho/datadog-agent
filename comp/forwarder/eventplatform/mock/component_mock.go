// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

// Package mock provides the mock interafce for the event platform component.
package mock

import eventplatform "github.com/DataDog/datadog-agent/comp/forwarder/eventplatform/def"

// Mock implements mock-specific methods.
type Mock interface {
	eventplatform.Component
}