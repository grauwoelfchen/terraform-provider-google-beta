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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudRunDomainMapping_cloudRunDomainMappingBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"namespace":       getTestProjectFromEnv(),
		"verified_domain": "tftest-domainmapping.com",
		"random_suffix":   acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunDomainMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunDomainMapping_cloudRunDomainMappingBasicExample(context),
			},
			{
				ResourceName:      "google_cloud_run_domain_mapping.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudRunDomainMapping_cloudRunDomainMappingBasicExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_cloud_run_service" "default" {
  name     = "tftest-cloudrun%{random_suffix}"
  location = "us-central1"

  metadata {
    namespace = "%{namespace}"
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
      }
    }
  }
}

resource "google_cloud_run_domain_mapping" "default" {
  location = "us-central1"
  name     = "%{verified_domain}"

  metadata {
    namespace = "%{namespace}"
  }

  spec {
    route_name = google_cloud_run_service.default.name
  }
}
`, context)
}

func testAccCheckCloudRunDomainMappingDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_cloud_run_domain_mapping" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{CloudRunBasePath}}domains.cloudrun.com/v1/namespaces/{{project}}/domainmappings/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("CloudRunDomainMapping still exists at %s", url)
		}
	}

	return nil
}
