package main_test

import (
	"testing"
	"time"

	"github.com/xoob/goa-testcase-payloads/app"
)

// These tests fail statically, we don't actually assert anything.

func Test_UserTypes(t *testing.T) {
	{
		// Expected fields in parent user type
		var _ int = (&app.Greeting{}).ID
		var _ string = (&app.Greeting{}).Message
		var _ *bool = (&app.Greeting{}).ParentOptional
	}

	{
		// Expected inherited fields in child user type
		var _ int = (&app.ExtendedGreeting{}).ID
		var _ string = (&app.ExtendedGreeting{}).Message
		// Expected fields in child user type
		var _ bool = (&app.ExtendedGreeting{}).Active
		var _ time.Time = (&app.ExtendedGreeting{}).CreatedAt
		var _ *int = (&app.ExtendedGreeting{}).Optional
	}
}

func Test_MediaTypes(t *testing.T) {
	{
		// Expected fields in parent media type
		var _ int = (&app.GreetingMedia{}).ID
		var _ string = (&app.GreetingMedia{}).Message
		var _ *bool = (&app.GreetingMedia{}).ParentOptional
		var _ string = (&app.GreetingMedia{}).Href
	}

	{
		// Expected inherited fields in child media type
		var _ int = (&app.ExtendedGreetingMedia{}).ID
		var _ string = (&app.ExtendedGreetingMedia{}).Message
		var _ *bool = (&app.Greeting{}).ParentOptional
		// Expected fields in child media type
		var _ bool = (&app.ExtendedGreetingMedia{}).Active
		var _ time.Time = (&app.ExtendedGreetingMedia{}).CreatedAt
		var _ *int = (&app.ExtendedGreetingMedia{}).Optional
		var _ string = (&app.ExtendedGreetingMedia{}).Href
	}
}

func Test_Payloads(t *testing.T) {
	{
		// Expected fields in parent media type
		var _ string = (&app.CreateGreetingPayload{}).Message
		var _ *bool = (&app.CreateGreetingPayload{}).ParentOptional
	}

	{
		// Expected inherited fields in child media type
		var _ string = (&app.CreateExtendedGreetingPayload{}).Message
		var _ *bool = (&app.CreateExtendedGreetingPayload{}).ParentOptional
		// Expected fields in child media type
		var _ bool = (&app.CreateExtendedGreetingPayload{}).Active
		var _ *int = (&app.CreateExtendedGreetingPayload{}).Optional
	}

	{
		// Expected inherited fields in child media type
		var _ string = (&app.CreateOptionalExtendedGreetingPayload{}).Message
		var _ *bool = (&app.CreateOptionalExtendedGreetingPayload{}).ParentOptional
		// Expected fields in child media type
		var _ bool = (&app.CreateOptionalExtendedGreetingPayload{}).Active
		var _ *int = (&app.CreateOptionalExtendedGreetingPayload{}).Optional
	}
}
