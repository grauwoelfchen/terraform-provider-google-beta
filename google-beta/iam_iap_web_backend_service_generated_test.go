// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIapWebBackendServiceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/iap.httpsResourceAccessor",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor", getTestProjectFromEnv(), fmt.Sprintf("backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccIapWebBackendServiceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor", getTestProjectFromEnv(), fmt.Sprintf("backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/iap.httpsResourceAccessor",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccIapWebBackendServiceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", getTestProjectFromEnv(), fmt.Sprintf("backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/iap.httpsResourceAccessor",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s", getTestProjectFromEnv(), fmt.Sprintf("backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/iap.httpsResourceAccessor",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor %s", getTestProjectFromEnv(), fmt.Sprintf("backend-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/iap.httpsResourceAccessor",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor", getTestProjectFromEnv(), fmt.Sprintf("backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor %s", getTestProjectFromEnv(), fmt.Sprintf("backend-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/iap.httpsResourceAccessor",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", getTestProjectFromEnv(), fmt.Sprintf("backend-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/iap.httpsResourceAccessor",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", getTestProjectFromEnv(), fmt.Sprintf("backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", getTestProjectFromEnv(), fmt.Sprintf("backend-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(10),
		"role":            "roles/iap.httpsResourceAccessor",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamPolicy_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s", getTestProjectFromEnv(), fmt.Sprintf("backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIapWebBackendServiceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_http_health_check" "default" {
  name               = "health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_member" "foo" {
	project = "${google_compute_backend_service.default.project}"
	web_backend_service = "${google_compute_backend_service.default.name}"
	role = "%{role}"
	member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccIapWebBackendServiceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_http_health_check" "default" {
  name               = "health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

data "google_iam_policy" "foo" {
	binding {
		role = "%{role}"
		members = ["user:admin@hashicorptest.com"]
	}
}

resource "google_iap_web_backend_service_iam_policy" "foo" {
	project = "${google_compute_backend_service.default.project}"
	web_backend_service = "${google_compute_backend_service.default.name}"
	policy_data = "${data.google_iam_policy.foo.policy_data}"
}
`, context)
}

func testAccIapWebBackendServiceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_http_health_check" "default" {
  name               = "health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_binding" "foo" {
	project = "${google_compute_backend_service.default.project}"
	web_backend_service = "${google_compute_backend_service.default.name}"
	role = "%{role}"
	members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccIapWebBackendServiceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_http_health_check" "default" {
  name               = "health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_binding" "foo" {
	project = "${google_compute_backend_service.default.project}"
	web_backend_service = "${google_compute_backend_service.default.name}"
	role = "%{role}"
	members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}

func testAccIapWebBackendServiceIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_http_health_check" "default" {
  name               = "health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_binding" "foo" {
	project = "${google_compute_backend_service.default.project}"
	web_backend_service = "${google_compute_backend_service.default.name}"
	role = "%{role}"
	members = ["user:admin@hashicorptest.com"]
	condition {
		title       = "%{condition_title}"
		description = "Expiring at midnight of 2019-12-31"
		expression  = "%{condition_expr}"
	}
}
`, context)
}

func testAccIapWebBackendServiceIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_http_health_check" "default" {
  name               = "health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_binding" "foo" {
	project = "${google_compute_backend_service.default.project}"
	web_backend_service = "${google_compute_backend_service.default.name}"
	role = "%{role}"
	members = ["user:admin@hashicorptest.com"]
}

resource "google_iap_web_backend_service_iam_binding" "foo2" {
	project = "${google_compute_backend_service.default.project}"
	web_backend_service = "${google_compute_backend_service.default.name}"
	role = "%{role}"
	members = ["user:admin@hashicorptest.com"]
	condition {
		title       = "%{condition_title}"
		description = "Expiring at midnight of 2019-12-31"
		expression  = "%{condition_expr}"
	}
}
`, context)
}

func testAccIapWebBackendServiceIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_http_health_check" "default" {
  name               = "health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_member" "foo" {
	project = "${google_compute_backend_service.default.project}"
	web_backend_service = "${google_compute_backend_service.default.name}"
	role = "%{role}"
	member = "user:admin@hashicorptest.com"
	condition {
		title       = "%{condition_title}"
		description = "Expiring at midnight of 2019-12-31"
		expression  = "%{condition_expr}"
	}
}
`, context)
}

func testAccIapWebBackendServiceIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_http_health_check" "default" {
  name               = "health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_member" "foo" {
	project = "${google_compute_backend_service.default.project}"
	web_backend_service = "${google_compute_backend_service.default.name}"
	role = "%{role}"
	member = "user:admin@hashicorptest.com"
}

resource "google_iap_web_backend_service_iam_member" "foo2" {
	project = "${google_compute_backend_service.default.project}"
	web_backend_service = "${google_compute_backend_service.default.name}"
	role = "%{role}"
	member = "user:admin@hashicorptest.com"
	condition {
		title       = "%{condition_title}"
		description = "Expiring at midnight of 2019-12-31"
		expression  = "%{condition_expr}"
	}
}
`, context)
}

func testAccIapWebBackendServiceIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_http_health_check" "default" {
  name               = "health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

data "google_iam_policy" "foo" {
	binding {
		role = "%{role}"
		members = ["user:admin@hashicorptest.com"]
		condition {
			title       = "%{condition_title}"
			description = "Expiring at midnight of 2019-12-31"
			expression  = "%{condition_expr}"
		}
	}
}

resource "google_iap_web_backend_service_iam_policy" "foo" {
	project = "${google_compute_backend_service.default.project}"
	web_backend_service = "${google_compute_backend_service.default.name}"
	policy_data = "${data.google_iam_policy.foo.policy_data}"
}
`, context)
}
