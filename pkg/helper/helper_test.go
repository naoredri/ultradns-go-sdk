package helper_test

import (
	"testing"

	"github.com/naoredri/ultradns-go-sdk/pkg/helper"
)

const (
	testOwnerFQDN = "www.example.com."
)

func TestGetOwnerFQDNwithRelativeName(t *testing.T) {
	if found := helper.GetOwnerFQDN("WWW", "example.com"); testOwnerFQDN != found {
		t.Fatal("FQDN owner name not returned : owner name - www : zone name - example.com")
	}
}

func TestGetOwnerFQDNwithOwnerRelativeName(t *testing.T) {
	if found := helper.GetOwnerFQDN("www", "example.COM."); testOwnerFQDN != found {
		t.Fatal("FQDN owner name not returned : owner name - www : zone name - example.com.")
	}
}

func TestGetOwnerFQDNwithoutFQDN(t *testing.T) {
	if found := helper.GetOwnerFQDN("www.EXAMPLE.com", "example.com"); testOwnerFQDN != found {
		t.Fatal("FQDN owner name not returned : owner name - www.example.com : zone name - example.com")
	}
}

func TestGetOwnerFQDNwithZoneFQDN(t *testing.T) {
	if found := helper.GetOwnerFQDN("www.example.com", "EXAMPLE.com."); testOwnerFQDN != found {
		t.Fatal("FQDN owner name not returned : owner name - www.example.com : zone name - example.com.")
	}
}

func TestGetOwnerFQDNwithOwnerFQDN(t *testing.T) {
	if found := helper.GetOwnerFQDN("WWW.EXAMPLE.COM.", "example.com"); testOwnerFQDN != found {
		t.Fatal("FQDN owner name not returned : owner name - www.example.com. : zone name - example.com")
	}
}

func TestGetOwnerFQDNwithFQDN(t *testing.T) {
	if found := helper.GetOwnerFQDN("www.example.com.", "EXAMPLE.COM."); testOwnerFQDN != found {
		t.Fatal("FQDN owner name not returned : owner name - www.example.com. : zone name - example.com.")
	}
}

func TestGetRecordTypeString(t *testing.T) {
	if expected, found := "A", helper.GetRecordTypeString("A (1)"); expected != found {
		t.Fatal("record type mismatched")
	}
}

func TestGetRecordTypeNumber(t *testing.T) {
	if expected, found := "5", helper.GetRecordTypeNumber("CNAME (5)"); expected != found {
		t.Fatal("record type mismatched")
	}
}
