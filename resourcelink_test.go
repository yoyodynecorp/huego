package huego_test

import (
	"testing"
	"os"
	"github.com/amimof/huego"
)

func TestGetResourcelinks(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resourcelinks, err := hue.GetResourcelinks()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d resourcelinks", len(resourcelinks))
	for _, resourcelink := range resourcelinks {
		t.Log(resourcelink)
	}
}


func TestGetResourcelink(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resourcelinks, err := hue.GetResourcelinks()
	if err != nil {
		t.Error(err)
	}
	for _, resourcelink := range resourcelinks {
		l, err := hue.GetResourcelink(resourcelink.Id)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(l)
		}
		break
	}
}

func TestCreateResourcelink(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resourcelink := &huego.Resourcelink{
		Name: "Huego Test Resourcelink",
    Description: "Amir's wakeup experience",
    Type: "Link",
    Class: 1,
    Owner: "78H56B12BAABCDEF",
    Links: []string{"/schedules/2", "/schedules/3", "/scenes/ABCD"},
	}
	response, err := hue.CreateResourcelink(resourcelink)
	if err != nil {
		t.Error(err)
	}
	for _, r := range response {
		t.Logf("Address: %s Value: %s Interface: %s", r.Address, r.Value, r.Interface)
	}
}

func TestUpdateResourcelink(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resourcelinks, err := hue.GetSensors()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d resourcelinks, setting the first one", len(resourcelinks))
	for _, resourcelink := range resourcelinks {
		response, err := hue.UpdateSensor(resourcelink.Id, resourcelink)
		if err != nil {
			t.Error(err)
		}
		for _, r := range response {
			t.Logf("Address: %s Value: %s Interface: %s", r.Address, r.Value, r.Interface)
		}
		break
	}
}

func TestDeleteResourcelink(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	res, err := hue.DeleteResourcelink(1)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		for _, r := range res {
			t.Logf("Address: %s Value: %s Interface: %s", r.Address, r.Value, r.Interface)
		}
	}
}
