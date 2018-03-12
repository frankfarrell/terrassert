package terrassert

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"errors"
)

func terrassertEquals() *schema.Resource {
	return &schema.Resource{
		Read: terrassertEqualsRead,

		Schema: map[string]*schema.Schema{
			"actual": &schema.Schema{
				Type:     schema.TypeString,
			},
			"expected": &schema.Schema{
				Type:     schema.TypeString,
			},
			"message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func terrassertEqualsRead(d *schema.ResourceData, meta interface{}) error {

	actualValue := d.Get("actual")
	expectedValue := d.Get("expected")

	if actualValue != expectedValue {
		message := d.Get("message")
		if message != "" {
			errors.New(fmt.Sprintf("%s %s != %s", message, actualValue, expectedValue))
		} else {
			errors.New(fmt.Sprintf("%s != %s", actualValue, expectedValue))
		}
	}

	return nil
}