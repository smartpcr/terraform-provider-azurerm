package schema_rules

import (
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/tools/schema-api/providerjson"
)

var propertyTypeBaseNode = providerjson.SchemaJSON{
	Type:        "TypeString",
	ConfigMode:  "",
	Optional:    false,
	Required:    false,
	Default:     nil,
	Description: "",
	Computed:    false,
	ForceNew:    false,
	Elem:        nil,
	MaxItems:    0,
	MinItems:    0,
}

var propertyTypePasses = providerjson.SchemaJSON{
	Type:        "TypeString",
	ConfigMode:  "",
	Optional:    false,
	Required:    false,
	Default:     nil,
	Description: "",
	Computed:    false,
	ForceNew:    false,
	Elem:        nil,
	MaxItems:    0,
	MinItems:    0,
}

var propertyTypeViolates = providerjson.SchemaJSON{
	Type:        "TypeInt",
	ConfigMode:  "",
	Optional:    false,
	Required:    false,
	Default:     nil,
	Description: "",
	Computed:    false,
	ForceNew:    false,
	Elem:        nil,
	MaxItems:    0,
	MinItems:    0,
}

var propertyTypeList = providerjson.SchemaJSON{
	Type:        "TypeList",
	ConfigMode:  "",
	Optional:    false,
	Required:    false,
	Default:     nil,
	Description: "",
	Computed:    false,
	ForceNew:    false,
	Elem:        nil,
	MaxItems:    0,
	MinItems:    0,
}

var propertyTypeSet = providerjson.SchemaJSON{
	Type:        "TypeSet",
	ConfigMode:  "",
	Optional:    false,
	Required:    false,
	Default:     nil,
	Description: "",
	Computed:    false,
	ForceNew:    false,
	Elem:        nil,
	MaxItems:    0,
	MinItems:    0,
}

func TestPropertyTypeMatches_Check(t *testing.T) {
	data := propertyType{}
	if res := data.Check(propertyTypeBaseNode, propertyTypePasses, ""); res != nil {
		t.Errorf("expected no violation, got %+v", res)
	}

	if res := data.Check(propertyTypeBaseNode, propertyTypeViolates, ""); res == nil {
		t.Errorf("expected violation, but didn't get one")
	}

	if res := data.Check(propertyTypeList, propertyTypeSet, ""); res == nil {
		t.Errorf("expected violation, but didn't get one")
	}

	if res := data.Check(propertyTypeSet, propertyTypeList, ""); res != nil {
		t.Errorf("expected no violation, got %+v", res)
	}
}
