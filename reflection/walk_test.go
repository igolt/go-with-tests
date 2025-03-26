package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"John Doe"},
			[]string{"John Doe"},
		},
		{
			"struct with two string fields",
			struct {
				Name     string
				LastName string
			}{"John", "Doe"},
			[]string{"John", "Doe"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"John Doe", 33},
			[]string{"John Doe"},
		},
		{
			"nested fields",
			Person{"John Doe", Profile{33, "London"}},
			[]string{"John Doe", "London"},
		},
		{
			"pointers to things",
			&Person{
				"John Doe",
				Profile{33, "London"},
			},
			[]string{"John Doe", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"}, {44, "Liverpool"},
			},
			[]string{"London", "Liverpool"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"}, {44, "Liverpool"},
			},
			[]string{"London", "Liverpool"},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			var got []string

			walk(testCase.Input, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(testCase.ExpectedCalls, got) {
				t.Errorf("expected %q but got %q", testCase.ExpectedCalls, got)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		var got []string

		walk(map[string]string{
			"Cow": "Moo", "Sheep": "Baa",
		}, func(s string) {
			got = append(got, s)
		},
		)
		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")

		if len(got) != 2 {
			t.Errorf("expected %v to have only two values", got)
		}
	})

	t.Run("with channels", func(t *testing.T) {
		channel := make(chan Profile)

		go func() {
			channel <- Profile{33, "Berlin"}
			channel <- Profile{18, "Paris"}
			close(channel)
		}()

		var got []string
		expected := []string{"Berlin", "Paris"}

		walk(channel, func(s string) { got = append(got, s) })

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{18, "Paris"}
		}

		var got []string
		expected := []string{"Berlin", "Paris"}

		walk(aFunction, func(s string) { got = append(got, s) })

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	for _, x := range haystack {
		if x == needle {
			return
		}
	}
	t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
