package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestIni_Read(t *testing.T) {
	ini := Ini{}
	src, dst, err := ini.Read("testutil/source.ini", "testutil/destination.ini")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, `map[[Server2]:Name = "Server1"]`, fmt.Sprint(src))
	assert.Equal(t, `map[[Server1]:Name = "Server 1" [Server2]:Name = "Server 2"]`, fmt.Sprint(dst.Digest))
}

func TestIni_Override(t *testing.T) {
	src := Source{
		"Server2": []string{
			"URL = 'http://example.com'",
		},
	}
	dest := Destination{
		Digest: map[string]interface{}{
			"Server1": []string{
				"Name = 'Server1'",
				"URL = 'http://server1.com'",
			},
			"Server2": []string{
				"Name = 'Server2'",
				"URL = 'http://server2.com'",
			},
		},
	}

	ini := Ini{}
	dst, _ := ini.Override(src, dest)
	assert.Equal(t, Destination{
		Digest: map[string]interface{}{
			"Server1": []string{
				"Name = 'Server1'",
				"URL = 'http://server1.com'",
			},
			"Server2": []string{
				"Name = 'Server2'",
				"URL = 'http://example.com'",
			},
		},
	}, dst)
}

func TestIni_Write(t *testing.T) {
	ini := Ini{}
	dest := Destination{
		Digest: map[string]interface{}{
			"[Server1]": []string{
				"Name = 'Server1'",
				"URL = 'http://server1.com'",
			},
			"[Server2]": []string{
				"Name = 'Server2'",
				"URL = 'http://server2.com'",
			},
		},
	}
	err := ini.Write(dest, "testutil/new.ini")
	log.Println(err)
}
