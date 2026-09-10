package mock

import "testing"

func TestValidate(t *testing.T) {
	definition := Definition{
		Routes: []Route{
			{
				Method: "GET",
				Path:   "/hello",
				Response: Response{
					Status: 200,
					Body:   "Hello",
				},
			},
		},
	}

	if err := definition.Validate(); err != nil {
		t.Fatalf("expected definition to be valid: %v", err)
	}
}

func TestValidateRequiresRoutes(t *testing.T) {
	definition := Definition{}

	if err := definition.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidateRequiresMethod(t *testing.T) {
	definition := Definition{
		Routes: []Route{
			{
				Path: "/hello",
				Response: Response{
					Status: 200,
				},
			},
		},
	}

	if err := definition.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidateRequiresPath(t *testing.T) {
	definition := Definition{
		Routes: []Route{
			{
				Method: "GET",
				Response: Response{
					Status: 200,
				},
			},
		},
	}

	if err := definition.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidatePathMustStartWithSlash(t *testing.T) {
	definition := Definition{
		Routes: []Route{
			{
				Method: "GET",
				Path:   "hello",
				Response: Response{
					Status: 200,
				},
			},
		},
	}

	if err := definition.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidateStatus(t *testing.T) {
	definition := Definition{
		Routes: []Route{
			{
				Method: "GET",
				Path:   "/hello",
				Response: Response{
					Status: 700,
				},
			},
		},
	}

	if err := definition.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}
