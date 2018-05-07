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

	assert.Equal(t, `map[[Server "server2"]:[Name = "Server dua"]]`, fmt.Sprint(src))
	assert.Equal(t, `[# this is comment # baris kedua [Server "server1"] Name = "Server 1"  [Server "server2"] Name = "Server 2" ]`, fmt.Sprint(dst))
}

func TestIni_Override(t *testing.T) {
	type testcase struct {
		src    Source
		dest   Destination
		expect Destination
	}

	tests := []testcase{
		{
			src: Source{
				"[Server2]": []string{
					"URL = 'http://example.com'",
				},
			},
			dest: Destination{
				"[Server1]",
				"Name = 'Server1'",
				"URL = 'http://server1.com'",
				"[Server2]",
				"Name = 'Server2'",
				"URL = 'http://server2.com'",
			},
			expect: Destination{
				"[Server1]",
				"Name = 'Server1'",
				"URL = 'http://server1.com'",
				"[Server2]",
				"Name = 'Server2'",
				"URL = 'http://example.com'",
			},
		},
		{
			src: Source{
				"[Server1 \"ASDF\"]": []string{
					"URL = 'http://example1.com'",
				},
				"[Server2]": []string{
					"URL = 'http://example.com'",
				},
			},
			dest: Destination{
				"[Server2]   ",
				"Name = 'Server2'",
				"URL = 'http://server2.com'",
				"[Server1 \"ASDF\"]   ",
				"Name = 'Server1'",
				"URL = 'http://server1.com'",
			},
			expect: Destination{
				"[Server2]   ",
				"Name = 'Server2'",
				"URL = 'http://example.com'",
				"[Server1 \"ASDF\"]   ",
				"Name = 'Server1'",
				"URL = 'http://example1.com'",
			},
		},
	}

	for _, test := range tests {
		ini := Ini{}
		dst, err := ini.Override(test.src, test.dest)
		assert.NoError(t, err)
		assert.Equal(t, test.expect, dst)
	}

}

func TestIni_Write(t *testing.T) {
	ini := Ini{}
	dest := Destination{
		"[Server1]",
		"Name = 'Server1'",
		"URL = 'http://server1.com'",
		"[Server2]",
		"Name = 'Server2'",
		"URL = 'http://server2.com'",
	}
	err := ini.Write(dest, "testutil/new.ini")
	log.Println(err)
}
